package errores

type ErrorUsuarioLoggeado struct{}

func (e ErrorUsuarioLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioInexistente struct{}

func (e ErrorUsuarioInexistente) Error() string {
	return "Error: usuario no existente"
}

type ErrorNadieLoggeado struct{}

func (e ErrorNadieLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorVerSiguiente struct{}

func (e ErrorVerSiguiente) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorLike struct{}

func (e ErrorLike) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorVerLikes struct{}

func (e ErrorVerLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}
