package controllers

import (
	"github.com/ahsansaif47/cdc-app/http/dto"
	"github.com/ahsansaif47/cdc-app/repository/postgres"
	"github.com/ahsansaif47/cdc-app/repository/redis"
)

type IUserService interface {
	CreateUser(user *dto.UserSignupRequest) (int, string, error)
	SignIn(*dto.UserLoginRequest) (int, string, error)
	GenerateOTP(email string) error
	VerifyOTP(email, otp string) (bool, error)
	// FindAll() ([]models.User, error)
	// FindByID(id string) (*models.User, error)
	// GetAllVendors() ([]models.User, error)
	// GetAllUsers() ([]models.User, error)
	// SetNewPassword(email, newPassword string) (bool, error)
	// FindUserByEmail(string) (int, *models.User, error)
}

type UserService struct {
	repo      postgres.IUserRepository
	cacheRepo redis.ICacheRepository
}

func NewUserService(repo postgres.IUserRepository, cache redis.ICacheRepository) IUserService {
	return &UserService{
		repo:      repo,
		cacheRepo: cache,
	}
}
