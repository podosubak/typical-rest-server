package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-rest-server/app/repository"
	"github.com/typical-go/typical-rest-server/app/service"
)

func init() {
	typapp.AppendConstructor(
		repository.NewBookRepo,
		repository.NewMusicRepo,
		service.NewBookService,
		service.NewMusicService,
	)
}