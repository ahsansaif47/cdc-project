package controllers

import (
	"errors"
	"time"

	"github.com/ahsansaif47/cdc-app/constants"
	"github.com/ahsansaif47/cdc-app/http/dto"
	"github.com/ahsansaif47/cdc-app/models"
	"github.com/ahsansaif47/cdc-app/utils"
	"github.com/ahsansaif47/cdc-app/utils/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *UserService) CreateUser(user *dto.UserSignupRequest) (int, string, error) {
	passwordHash, err := utils.GeneratePasswordHash(user.Password)
	if passwordHash == "" {
		return fiber.StatusBadRequest, "", constants.ErrFailedToHashPassword
	} else if err != nil {
		return fiber.StatusInternalServerError, "", err
	}

	now := time.Now().UTC()
	uuid := uuid.New()
	newUser := &models.User{
		ID:               uuid,
		UserName:         user.UserName,
		Email:            user.Email,
		PasswordHash:     &passwordHash,
		CreatedAt:        &now,
		UpdatedAt:        &now,
		AuthProviderType: user.AuthProvider,
		PhoneNumber:      user.PhoneNumber,
		RoleID:           user.RoleID,
	}
	if err := s.repo.CreateUser(newUser); err != nil {
		if errors.Is(err, constants.ErrUserAlreadyExists) {
			return fiber.StatusBadRequest, "", err
		}
	}

	tokenStr, err := jwt.GenerateJWT(uuid.String(), user.Email, user.UserName, user.RoleID)
	if err != nil {
		return fiber.StatusInternalServerError, "", err
	}

	return fiber.StatusOK, tokenStr, nil

}

func (s *UserService) SignIn(signinReq *dto.UserLoginRequest) (int, string, error) {

	// Get the user details from DB
	user, err := s.repo.FindUserByEmail(signinReq.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.StatusNotFound, "", err
		} else {
			return fiber.StatusInternalServerError, "", err
		}
	}

	switch user.AuthProviderType {
	case "local":
		// Check User Hash
		ok := utils.CheckPasswordHash(signinReq.Password, *user.PasswordHash)
		if !ok {
			return fiber.StatusInternalServerError, "", errors.New("password does not match")
		}

		token, err := jwt.GenerateJWT(user.ID.String(), signinReq.Email, user.UserName, user.RoleID)
		if err != nil {
			return fiber.StatusInternalServerError, "", err
		}

		return fiber.StatusOK, token, nil

	case "google":
		return fiber.StatusOK, "token", nil

	case "apple":
		return fiber.StatusOK, "token", nil

		// call the google apis api to verify the token, then send the token back as verification...

	default:
		// This should never happen
		return fiber.StatusInternalServerError, "", errors.New("unknown auth type")
	}
}

func (s *UserService) GenerateOTP(email string) error {
	otp := utils.GenerateOTP()

	otpHash := utils.HashOTP(otp)
	err := s.cacheRepo.StoreOTP(email, otpHash, 1*time.Minute)
	if err != nil {
		return err
	}
	// Email this otp to the client
	return nil
}

func (s *UserService) VerifyOTP(email, otp string) (bool, error) {
	storedOTPHash, err := s.cacheRepo.RetrieveOTP(email)
	if err != nil {
		return false, err
	}

	isValid := utils.VerifyOTPHash(otp, storedOTPHash)
	return isValid, nil
}
