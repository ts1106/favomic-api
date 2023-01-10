package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ts1106/favomic-api/ent"
)

func main() {
	ctx := context.Background()
	client, err := ent.Open("mysql", "user:password@tcp(db:3306)/database?parseTime=True")
	if err != nil {
		log.Fatalf("failed to opening connection to mysql: %v", err)
	}
	defer client.Close()

	q := client.Debug().Author.Query().
		WithComics()
	c, _ := q.All(ctx)
	fmt.Printf("%v", c[0].Edges)
}
