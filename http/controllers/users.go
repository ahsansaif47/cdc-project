package controllers

import (
	"errors"
	"time"

	"github.com/ahsansaif47/cdc-app/constants"
	"github.com/ahsansaif47/cdc-app/http/dto"
	sqlcgenerated "github.com/ahsansaif47/cdc-app/repository/postgres/schema/sqlc/generated"
	"github.com/ahsansaif47/cdc-app/utils"
	"github.com/ahsansaif47/cdc-app/utils/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
	pgUUID := pgtype.UUID{}
	err = pgUUID.Scan(uuid)

	pgPassword := utils.ToPgText(&passwordHash)
	dbTIme := utils.ToPgTime(now)
	dbPhone := utils.ToPgText(user.PhoneNumber)

	newUser := sqlcgenerated.CreateUserParams{
		ID:               pgUUID,
		Username:         user.UserName,
		Email:            user.Email,
		PasswordHash:     pgPassword,
		CreatedAt:        dbTIme,
		UpdatedAt:        dbTIme,
		AuthProviderType: user.AuthProvider,
		PhoneNumber:      dbPhone,
		RoleID:           int32(user.RoleID),
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

	passwordHash := utils.PgTextToString(user.PasswordHash)

	switch user.AuthProviderType {
	case "local":
		// Check User Hash
		ok := utils.CheckPasswordHash(signinReq.Password, passwordHash)
		if !ok {
			return fiber.StatusInternalServerError, "", errors.New("password does not match")
		}

		uuid, err := utils.PgUUIDToUUID(user.ID)
		if err != nil {
			return fiber.StatusBadRequest, "", err
		}

		token, err := jwt.GenerateJWT(uuid, signinReq.Email, user.Username, uint(user.RoleID))
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
