package repository

import (
	"belajar-gin/model"
	"database/sql"
)

type FormulirRepository interface {
	FindAll() ([]model.Formulir, error)
	FindOne(int) (model.Formulir, error)
}

type formulirRepositoryImp struct {
	sql *sql.DB
}

func NewFormulirRepositoryImp(sqlDB *sql.DB) *formulirRepositoryImp {
	return &formulirRepositoryImp{sqlDB}
}

func (f *formulirRepositoryImp) FindAll() (response []model.Formulir, err error) {
	return response, nil
}

func (f *formulirRepositoryImp) FindOne(int) (response model.Formulir, err error) {
	return response, nil
}
