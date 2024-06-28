package usuario

type Post interface {
	LikeadoPor(usuario Usuario)

	VerLikes()

	CantidadLikes() int

	VerID() int

	Publicador() string

	IndicePublicador() int

	Contenido() string
}
