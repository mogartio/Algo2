package usuario

import (
	abb "algogram/abb"
	errores "algogram/errores"
	"fmt"
	"os"
	"strings"
)

type post struct {
	contenido         string
	id                int
	publicador        string
	indice_publicador int //sera usado para calcular la afinidad
	cantidad_likes    int
	likeadores        *abb.DiccionarioOrdenado[string, string] //Es un puntero para que el struct post sea comparable.
}

func CrearPost(id int, publicador string, contenido string, indice_publicador int) Post {
	nuevo_post := new(post)
	nuevo_post.id = id
	nuevo_post.contenido = contenido
	likeadores := abb.CrearABB[string, string](func(A string, B string) int { return strings.Compare(A, B) })
	nuevo_post.likeadores = &likeadores
	nuevo_post.publicador = publicador
	nuevo_post.indice_publicador = indice_publicador
	return nuevo_post
}

func (post *post) LikeadoPor(usuario Usuario) {
	if !(*post.likeadores).Pertenece(usuario.Nombre()) {
		post.cantidad_likes++
		(*post.likeadores).Guardar(usuario.Nombre(), usuario.Nombre())
	}
}

func (post post) VerLikes() {
	if post.CantidadLikes() == 0 {
		error := &errores.ErrorVerLikes{}
		fmt.Fprintf(os.Stdout, "%s\n", error.Error())
		return
	}
	iterador := (*post.likeadores).Iterador()
	fmt.Fprintf(os.Stdout, "El post tiene %d likes:\n", post.CantidadLikes())
	for iterador.HaySiguiente() {
		fmt.Fprintf(os.Stdout, "	%s\n", iterador.Siguiente())
	}
}

func (post post) CantidadLikes() int {
	return post.cantidad_likes
}

func (post post) VerID() int {
	return post.id
}

func (post post) Publicador() string {
	return post.publicador
}

func (post post) IndicePublicador() int {
	return post.indice_publicador
}

func (post post) Contenido() string {
	return post.contenido
}
