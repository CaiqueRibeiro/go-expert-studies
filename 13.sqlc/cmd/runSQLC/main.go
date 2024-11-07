package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CaiqueRibeiro/go-expert-studies/sqlc/internal/db"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("postgresql", "root:root@tcp(localhost:5432)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	// instantiates the "repository" with the db connect
	queries := db.New(dbConn)

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend",
		Description: sql.NullString{String: "Courses for backend development"},
	})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)
	}
}
