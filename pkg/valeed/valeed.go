package valeed

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func Input(dto interface{}, r *http.Request) error {
	body, errIo := io.ReadAll(r.Body)
	if errIo != nil {
		return fmt.Errorf("error read body: %v", errIo)
	}
	errUnMar := json.Unmarshal(body, &dto)
	return errUnMar
}

func Validate(dto interface{}) error {
	validates := validator.New(validator.WithRequiredStructEnabled())
	err := validates.Struct(dto)
	var invalidValidationError *validator.InvalidValidationError
	if err != nil {
		if !errors.As(err, &invalidValidationError) {
			return err
		}
		return err
	}
	return nil
}
