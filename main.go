package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ts1106/favomic-api/ent"
	"github.com/ts1106/favomic-api/ent/migrate"
	"github.com/ts1106/favomic-api/internal/router"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "user:password@tcp(db:3306)/database?parseTime=True")
	if err != nil {
		log.Fatalf("failed to opening connection to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed to creating schema resources: %v", err)
	}

	r := router.NewRouter(client)

	http.ListenAndServe(
		"0.0.0.0:8080",
		h2c.NewHandler(r, &http2.Server{}),
	)
}
