package author_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ts1106/favomic-api/ent/enttest"
	"github.com/ts1106/favomic-api/internal/domain/author"
	"github.com/ts1106/favomic-api/internal/util/response"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateHandler(t *testing.T) {
	for name, f := range map[string]func(t *testing.T, h *author.Handler){
		"normal":       testCreateHandler,
		"invalid name": testCreateHandlerInvalidName,
	} {
		t.Run(name, func(t *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer client.Close()

			ctx := context.Background()
			client.Schema.Create(ctx)

			srv := author.NewServer(client)
			h := author.NewHandler(srv)

			f(t, h)
		})
	}
}

func testCreateHandler(t *testing.T, h *author.Handler) {
	param := author.CreateRequest{Name: "name"}
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(param)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRequest(http.MethodPost, "http://localhost/author", &body)
	w := httptest.NewRecorder()
	h.Create(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("got %d, want %d", w.Code, http.StatusCreated)
	}
	if w.Result().ContentLength == 0 {
		t.Errorf("response Content length is 0")
	}
}

func testCreateHandlerInvalidName(t *testing.T, h *author.Handler) {
	param := author.CreateRequest{Name: ""}
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(param)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRequest(http.MethodPost, "http://localhost/author", &body)
	w := httptest.NewRecorder()
	h.Create(w, r)

	var want bytes.Buffer
	err = json.NewEncoder(&want).Encode(response.Error{Code: http.StatusBadRequest, Errors: []response.Errors{{Message: "Nameは必須フィールドです"}}})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want.String(), w.Body.String())
}
