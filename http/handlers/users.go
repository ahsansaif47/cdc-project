package handlers

import (
	"log"

	"github.com/ahsansaif47/cdc-app/constants"
	"github.com/ahsansaif47/cdc-app/http/controllers"
	"github.com/ahsansaif47/cdc-app/http/dto"
	"github.com/ahsansaif47/cdc-app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service   controllers.IUserService
	validator *validator.Validate
}

func NewAuthHandler(service controllers.IUserService) *AuthHandler {

	val := validator.New()

	// Register all auth-related custom validators here
	if err := val.RegisterValidation("password", utils.PasswordValidator); err != nil {
		log.Fatalf("failed to register password validation: %v", err)
	}

	return &AuthHandler{
		service:   service,
		validator: val,
	}
}

// SignUp
//
//	@Summary		Create User
//	@Description	Create a New User
//	@Tags			User
//	@Accept			json
//	@Param			body	body	dto.UserSignupRequest	true	"Signup User Request"
//
//	@Produce		json
//	@Body			user  dto.UserSignupRequest true "User Signup Request"
//	@Success		200	{object}	dto.UserSignupResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/users/signup [post]
func (h *AuthHandler) CreateUser(ctx *fiber.Ctx) error {
	userReq := &dto.UserSignupRequest{}

	err := ctx.BodyParser(userReq)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if val_errors := utils.ValidateRequest(h.validator, userReq, constants.CustomValidationErrors); val_errors != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": val_errors,
		})
	}

	status_code, tokenStr, err := h.service.CreateUser(userReq)
	if err != nil {
		return ctx.Status(status_code).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	} else {
		return ctx.Status(status_code).JSON(dto.UserLoginResponse{
			AccessToken: tokenStr,
		})
	}
}

// SignIn
//
//	@Summary		Authenticate User
//	@Description	Login User into the System
//	@Tags			User
//	@Accept			json
//	@Param			body	body	dto.UserLoginRequest	true	"Signin User Request"
//
//	@Produce		json
//	@Success		200	{object}	dto.UserLoginResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/users/signin [post]
func (h *AuthHandler) Signin(ctx *fiber.Ctx) error {
	signinReq := &dto.UserLoginRequest{}

	err := ctx.BodyParser(signinReq)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if val_errors := utils.ValidateRequest(h.validator, signinReq, constants.CustomValidationErrors); val_errors != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": val_errors,
		})
	}

	status_code, token, err := h.service.SignIn(signinReq)
	if err != nil {
		return ctx.Status(status_code).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	} else {
		return ctx.Status(status_code).JSON(dto.UserLoginResponse{
			AccessToken: token,
		})
	}
}

// GenerateOTP
//
//	@Summary		Generate OTP for User
//	@Description	Generate OTP for User Authentication
//	@Tags			User
//	@Accept			json
//	@Param			email	query	string	true	"Email"
//
//	@Produce		json
//	@Success		200	{object}	dto.GenerateOTPResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/users/generate-otp [post]
func (h *AuthHandler) GenerateOTP(ctx *fiber.Ctx) error {
	email := ctx.Query("email")

	if email == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Email is required",
		})
	}

	if err := h.service.GenerateOTP(email); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(dto.GenerateOTPResponse{
			Success: true,
		})
	}

}

// VerifyOTP
//
//	@Summary		Verify OTP for User
//	@Description	Verify OTP for User Authentication
//	@Tags			User
//	@Accept			json
//	@Param			body	body	dto.VerifyOTPRequest	true	"Verify OTP Request"
//
//	@Produce		json
//	@Success		200	{object}	dto.VerifyOTPResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/users/verify-otp [post]
func (h *AuthHandler) VerifyOTP(ctx *fiber.Ctx) error {
	var otpReq dto.VerifyOTPRequest

	err := ctx.BodyParser(&otpReq)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Cannot parse JSON",
		})
	}

	if val_errors := utils.ValidateRequest(h.validator, otpReq, constants.CustomValidationErrors); val_errors != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": val_errors,
		})
	}

	valid, err := h.service.VerifyOTP(otpReq.Email, otpReq.OTP)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.VerifyOTPResponse{
		Valid: valid,
	})
}
