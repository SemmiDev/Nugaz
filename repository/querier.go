package repository

import (
	"context"
	"github.com/SemmiDev/Go-Scheduled/entity"
)

type Saver interface {
	Save(ctx context.Context, arg *entity.Task) (*entity.Task, error)
}

type Finder interface {
	FindById(ctx context.Context, ID string) (*entity.Task, error)
	FindAll(ctx context.Context) ([]*entity.Task, error)
}

type Updater interface {
	Update(ctx context.Context, arg *entity.Task) (*entity.Task, error)
	UpdateIsDone(ctx context.Context, ID string) (*entity.Task, error)
}

type Deleter interface {
	Delete(ctx context.Context, ID string) error
}

type Querier interface {
	Saver
	Finder
	Updater
	Deleter
}

var _ Querier = (*Queries)(nil)
