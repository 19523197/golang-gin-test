package service

import "belajar-gin/repository"

type FormulirService interface {
}

type formulirServiceimp struct {
	repository.FormulirRepository
}

func newFormulirService(repository repository.FormulirRepository) *formulirServiceimp {
	return &formulirServiceimp{repository}
}
