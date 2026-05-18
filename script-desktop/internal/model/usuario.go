package model

type UsuarioCadastrado struct {
	Nome  string `json:"nome" `
	Email string `json:"email" `
	Senha string `json:"senha"`
}

type Usuario struct {
	ID    string `json:"id" db:"id"`
	Nome  string `json:"nome" db:"nome"`
	Email string `json:"email" db:"email"`
	Senha string `json:"senha" db:"senha"`
}

type UsuarioLogin struct {
	Nome  string `json:"nom"`
	Senha string `json:"senha"`
}

type UsuarioResetarSenha struct {
	Nome     string `json:"nome"`
	Resposta string `json:"resposta"`
}
