package service

import "script-desktop/internal/model"

type UsuarioRepository interface {
	CadastrarUsuario(usuario model.Usuario) error
}

type UsuarioService struct {
	Repo UsuarioRepository
}

func (U UsuarioService) CadastrarUsuario(usuarioInput model.UsuarioCadastrado) error {
	usuario := model.Usuario{
		Nome:  usuarioInput.Nome,
		Senha: usuarioInput.Senha,
		Email: usuarioInput.Email,
	}
	return U.Repo.CadastrarUsuario(usuario)
}

func (U UsuarioService) LogarUsuario(usuario model.UsuarioLogin) error {
	return nil
}

func (U UsuarioService) ResetarSenha(usuario model.UsuarioResetarSenha) error {
	return nil
}
