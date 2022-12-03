package author_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ts1106/favomic-api/internal/domain/author"
)

func TestCreateRequestValidation(t *testing.T) {
	tests := []struct {
		name string
		obj  author.CreateRequest
		want []error
	}{
		{"normal", author.CreateRequest{Name: "name"}, nil},
		{"invalid: name is empty", author.CreateRequest{Name: ""}, []error{errors.New("Nameは必須フィールドです")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.obj.Validate()
			assert.Equal(t, got, tt.want)
		})
	}
}
