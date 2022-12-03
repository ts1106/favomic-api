package author

import (
	"encoding/json"
	"io"

	"github.com/ts1106/favomic-api/internal/util/validator"
)

type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

func (r *CreateRequest) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *CreateRequest) Validate() []error {
	return validator.Validate(r)
}
