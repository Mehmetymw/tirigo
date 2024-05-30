package service_test

import (
	"testing"
	"time"
	"tirigo/internal/user-management/domain"
	"tirigo/internal/user-management/dtos"
	"tirigo/internal/user-management/repository"
	"tirigo/internal/user-management/service"
	util "tirigo/pkg"
	"tirigo/pkg/jwt"
	"tirigo/pkg/redis"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB setups the in-memory database for testing.
func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&domain.User{})

	return db
}

// setupTestRedis setups the in-memory Redis server for testing.
func setupTestRedis() *miniredis.Miniredis {

	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	redis.Init(s.Addr())
	return s
}

func TestRegister(t *testing.T) {
	db := setupTestDB()
	userRepo := repository.NewGormUserRepository(db)
	userService := service.NewUserService(userRepo)

	userParam := dtos.UserRegisterParameter{
		Username: util.RandomUsername(),
		Password: util.RandomPassword(),
		Email:    util.RandomEmail(),
	}

	user, err := userService.Register(userParam)
	assert.NoError(t, err)
	assert.Equal(t, userParam.Username, user.Username)
	assert.Equal(t, userParam.Email, user.Email)
	assert.NotEmpty(t, user.ID)
	assert.WithinDuration(t, time.Now(), user.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), user.UpdatedAt, time.Second)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userParam.Password))
	assert.NoError(t, err)
}

func TestAuthenticate(t *testing.T) {
	db := setupTestDB()
	s := setupTestRedis()
	defer s.Close()

	jwt.Init("my_secret_key")

	userRepo := repository.NewGormUserRepository(db)
	userService := service.NewUserService(userRepo)

	userParam := dtos.UserRegisterParameter{
		Username: util.RandomUsername(),
		Password: util.RandomPassword(),
		Email:    util.RandomEmail(),
	}

	_, err := userService.Register(userParam)
	assert.NoError(t, err)

	authParam := dtos.UserAuthParameter{
		Username: userParam.Username,
		Password: userParam.Password,
	}

	token, err := userService.Authenticate(authParam)
	assert.NoError(t, err)

	username, err := redis.Client.Get(token).Result()
	assert.NoError(t, err)
	assert.Equal(t, userParam.Username, username)
}

func TestLogout(t *testing.T) {
	db := setupTestDB()
	s := setupTestRedis()
	defer s.Close()

	jwt.Init("my_secret_key")

	userRepo := repository.NewGormUserRepository(db)
	userService := service.NewUserService(userRepo)

	userParam := dtos.UserRegisterParameter{
		Username: util.RandomUsername(),
		Password: util.RandomPassword(),
		Email:    util.RandomEmail(),
	}

	_, err := userService.Register(userParam)
	assert.NoError(t, err)

	authParam := dtos.UserAuthParameter{
		Username: userParam.Username,
		Password: userParam.Password,
	}

	token, err := userService.Authenticate(authParam)
	assert.NoError(t, err)

	err = userService.Logout(token)
	assert.NoError(t, err)

	_, err = redis.Client.Get(token).Result()
	assert.Error(t, err)
}
