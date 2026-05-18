package repository

import (
	"script-desktop/internal/model"

	"github.com/jmoiron/sqlx"
)

type UsuarioRepository struct {
	DB *sqlx.DB
}

func (U UsuarioRepository) CadastrarUsuario(usuario model.Usuario) error {
	query := "INSERT INTO usuarios (nome,email,senha) VALUES(:nome,:email,:senha)"
	_, err := U.DB.NamedExec(query, usuario)
	if err != nil {
		return err
	}
	return nil
}
