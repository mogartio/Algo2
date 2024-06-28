package usuario

type Usuario interface {
	Nombre() string

	Publicar(string, int) Post

	VerProximo() Post

	HayOtroPost() bool

	ActualizarFeed(post Post)

	Indice() int
}
