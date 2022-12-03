package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ts1106/favomic-api/ent"
	"github.com/ts1106/favomic-api/ent/migrate"
	"github.com/ts1106/favomic-api/internal/server"

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

	srv := server.NewServer(client)
	srv.RouteRegister()
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}
