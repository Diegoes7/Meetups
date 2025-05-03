package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type DBLogger struct{}

func (d *DBLogger) BeforeQuery(cxt context.Context, q *pg.QueryEvent) (context.Context, error) {
	return cxt, nil
}

func (d *DBLogger) AfterQuery(cxt context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)
	return db
}
