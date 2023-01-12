package helpers

import (
	"api-go-gin/models"

	"gopkg.in/validator.v2"
)

func ModelValidator(foobar *models.Foobar) error {
	if err := validator.Validate(foobar); err != nil {
		return err
	}
	return nil
}
