package author

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRequestValidation(t *testing.T) {
	tests := []struct {
		name string
		obj  CreateRequest
		want []error
	}{
		{"normal", CreateRequest{Name: "name"}, nil},
		{"invalid: name is empty", CreateRequest{Name: ""}, []error{errors.New("Nameは必須フィールドです")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.obj.Validate()
			assert.Equal(t, got, tt.want)
		})
	}
}
