package storage

import (
	"fmt"

	sq "github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

var _ Queryer = new(sqlx.DB)

// Queryer allows sqlx.DB binding
type Queryer interface {
	Select(destination interface{}, query string, args ...interface{}) error
	Get(destination interface{}, query string, args ...interface{}) error
}

// NewProjection generates a SQL string based on the given SQLBuilderOption
func NewProjection(options ...SQLBuilderOption) (string, error) {
	wrapper := projectionBuilder{}

	for _, opt := range options {
		opt(&wrapper)
	}

	statement, _, err := wrapper.builder.ToSql()
	return statement, err
}

type projectionBuilder struct {
	builder *sq.SelectBuilder
}

// SQLBuilderOption functional option wrapper for SQL builder
type SQLBuilderOption func(b *projectionBuilder)

func Select(columns ...string) SQLBuilderOption {
	return func(b *projectionBuilder) {
		// FIXME: This assignation is dangerous and may lead to "panic" due nil references
		b.builder = sq.Select(columns...)
	}
}

func From(tables ...string) SQLBuilderOption {
	return func(b *projectionBuilder) {
		b.builder.From(tables...)
	}
}

func Where(expression string) SQLBuilderOption {
	return func(b *projectionBuilder) {
		b.builder.Where(expression, "")
	}
}

func Join(fmtExpresion string, tables ...interface{}) SQLBuilderOption {
	return func(b *projectionBuilder) {
		b.builder.Join(fmt.Sprintf(fmtExpresion, tables...))
	}
}

func LeftJoin(fmtExpresion string, tables ...interface{}) SQLBuilderOption {
	return func(b *projectionBuilder) {
		b.builder.LeftJoin(fmt.Sprintf(fmtExpresion, tables...))
	}
}

func OrderBy(columns ...string) SQLBuilderOption {
	return func(b *projectionBuilder) {
		b.builder.OrderBy(columns...)
	}
}

func Limit(limit int) SQLBuilderOption {
	return func(b *projectionBuilder) {
		b.builder.Limit(uint64(limit))
	}
}

func Offset(offset int) SQLBuilderOption {
	return func(b *projectionBuilder) {
		b.builder.Offset(uint64(offset))
	}
}
