package main

import (
	"context"

	"script-desktop/internal/model"
	"script-desktop/internal/repository"
	"script-desktop/internal/service"

	"github.com/jmoiron/sqlx"
)

// App struct
type App struct {
	ctx context.Context
	DB  *sqlx.DB
}

// NewApp creates a new App application struct
func NewApp(db *sqlx.DB) *App {
	return &App{DB: db}
}

// startup is called when the app starts. The context is saved
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) CadastrarUsuario(usuario model.UsuarioCadastrado) (string, error) {
	repository := repository.UsuarioRepository{DB: a.DB}
	service := service.UsuarioService{Repo: repository}
	err := service.CadastrarUsuario(usuario)
	if err != nil {
		return err.Error(), err
	}

	return "Sucesso", nil
}
