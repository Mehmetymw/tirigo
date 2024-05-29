package service_test

import (
	"testing"
	"time"
	"tirigo/internal/user-management/domain"
	"tirigo/internal/user-management/dtos"
	"tirigo/internal/user-management/repository"
	"tirigo/internal/user-management/service"
	util "tirigo/pkg"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&domain.User{})

	return db
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

	user, err := userService.Authenticate(authParam)
	assert.NoError(t, err)
	assert.Equal(t, userParam.Username, user.Username)
	assert.Equal(t, userParam.Email, user.Email)
}
