package author

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
	"github.com/ts1106/favomic-api/internal/util/validator"
)

type CreateRequest struct {
	Name   string      `json:"name" validate:"required"`
	Comics []uuid.UUID `json:"comics"`
}

func (r *CreateRequest) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *CreateRequest) Validate() []error {
	return validator.Validate(r)
}
