package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"
	"wynn-member-api/pkg/errs"
)

// NewValidator func for create a new validator for models fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book models.
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, er := range err.(validator.ValidationErrors) {
		fields[er.Field()] = err.Error()
	}

	return fields
}

// ValidateBody : for fiber
func ValidateBody(c *fiber.Ctx, reqBody interface{}) error {
	err := c.BodyParser(&reqBody)
	if err != nil {
		return errs.NewBadRequestError(errs.INVALID_VALIDATION)
	}

	validate := NewValidator()
	err = validate.Struct(reqBody)
	if err != nil {
		return errs.NewBadRequestError(errs.INVALID_VALIDATION)
	}

	return nil
}

func ValidateQueryStr(c *fiber.Ctx, reqBody interface{}) error {
	err := c.QueryParser(&reqBody)
	if err != nil {
		fmt.Println(err.Error())
		return errs.NewBadRequestError(errs.INVALID_VALIDATION)
	}

	validate := NewValidator()
	err = validate.Struct(reqBody)
	if err != nil {
		return errs.NewBadRequestError(errs.INVALID_VALIDATION)
	}

	return nil
}
