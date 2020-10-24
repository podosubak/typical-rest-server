package service

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/typical-go/typical-rest-server/internal/app/data_access/postgresdb"
	"github.com/typical-go/typical-rest-server/internal/generated/postgresdb_repo"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"github.com/typical-go/typical-rest-server/pkg/typrest"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// BookSvc contain logic for Book Controller
	// @mock
	BookSvc interface {
		FindOne(context.Context, string) (*postgresdb.Book, error)
		Find(context.Context, *FindReq) ([]*postgresdb.Book, error)
		Create(context.Context, *postgresdb.Book) (*postgresdb.Book, error)
		Delete(context.Context, string) error
		Update(context.Context, string, *postgresdb.Book) (*postgresdb.Book, error)
		Patch(context.Context, string, *postgresdb.Book) (*postgresdb.Book, error)
	}
	// BookSvcImpl is implementation of BookSvc
	BookSvcImpl struct {
		dig.In
		Repo postgresdb_repo.BookRepo
	}
	// FindReq find request
	FindReq struct {
		Limit  uint64 `query:"limit"`
		Offset uint64 `query:"offset"`
	}
)

// NewBookSvc return new instance of BookSvc
// @ctor
func NewBookSvc(impl BookSvcImpl) BookSvc {
	return &impl
}

// Create Book
func (b *BookSvcImpl) Create(ctx context.Context, book *postgresdb.Book) (*postgresdb.Book, error) {
	if err := validator.New().Struct(book); err != nil {
		return nil, typrest.NewValidErr(err.Error())
	}
	id, err := b.Repo.Create(ctx, book)
	if err != nil {
		return nil, err
	}
	return b.findOne(ctx, id)
}

// Find books
func (b *BookSvcImpl) Find(ctx context.Context, req *FindReq) ([]*postgresdb.Book, error) {
	return b.Repo.Find(ctx, b.findSelectOpt(req)...)
}

func (b *BookSvcImpl) findSelectOpt(req *FindReq) []dbkit.SelectOption {
	return []dbkit.SelectOption{
		&dbkit.OffsetPagination{Offset: req.Offset, Limit: req.Limit},
	}
}

// FindOne book
func (b *BookSvcImpl) FindOne(ctx context.Context, paramID string) (*postgresdb.Book, error) {
	id, _ := strconv.ParseInt(paramID, 10, 64)
	if id < 1 {
		return nil, typrest.NewValidErr("paramID is missing")
	}
	return b.findOne(ctx, id)
}

func (b *BookSvcImpl) findOne(ctx context.Context, id int64) (*postgresdb.Book, error) {
	books, err := b.Repo.Find(ctx, dbkit.Equal(postgresdb_repo.BookTable.ID, id))
	if err != nil {
		return nil, err
	} else if len(books) < 1 {
		return nil, sql.ErrNoRows
	}
	return books[0], nil
}

// Delete book
func (b *BookSvcImpl) Delete(ctx context.Context, paramID string) error {
	id, _ := strconv.ParseInt(paramID, 10, 64)
	if id < 1 {
		return typrest.NewValidErr("paramID is missing")
	}
	_, err := b.Repo.Delete(ctx, dbkit.Equal(postgresdb_repo.BookTable.ID, id))
	return err
}

// Update book
func (b *BookSvcImpl) Update(ctx context.Context, paramID string, book *postgresdb.Book) (*postgresdb.Book, error) {
	id, _ := strconv.ParseInt(paramID, 10, 64)
	if id < 1 {
		return nil, typrest.NewValidErr("paramID is missing")
	}
	err := validator.New().Struct(book)
	if err != nil {
		return nil, typrest.NewValidErr(err.Error())
	}
	affectedRow, err := b.Repo.Update(ctx, book, dbkit.Equal(postgresdb_repo.BookTable.ID, id))
	if err != nil {
		return nil, err
	}
	if affectedRow < 1 {
		return nil, sql.ErrNoRows
	}
	return b.findOne(ctx, id)
}

// Patch book
func (b *BookSvcImpl) Patch(ctx context.Context, paramID string, book *postgresdb.Book) (*postgresdb.Book, error) {
	id, _ := strconv.ParseInt(paramID, 10, 64)
	if id < 1 {
		return nil, typrest.NewValidErr("paramID is missing")
	}
	affectedRow, err := b.Repo.Patch(ctx, book, dbkit.Equal(postgresdb_repo.BookTable.ID, id))
	if err != nil {
		return nil, err
	}
	if affectedRow < 1 {
		return nil, sql.ErrNoRows
	}
	return b.findOne(ctx, id)
}
