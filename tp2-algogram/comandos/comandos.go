package comandos

import (
	errores "algogram/errores"
	TDAHash "algogram/hash"
	TDAUsuario "algogram/usuario"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	VALIDACION_LOGIN    = "Hola"
	VALIDACION_PUBLICAR = "Post publicado"
	VALIDACION_LOGOUT   = "Adios"
	VALIDACION_LIKE     = "Post likeado"

	LOGIN       = "login"
	LOGOUT      = "logout"
	PUBLICAR    = "publicar"
	VER_PROXIMO = "ver_siguiente_feed"
	LIKE        = "likear_post"
	VER_LIKES   = "mostrar_likes"
)

func LeerInput(input []string, lista_usuarios TDAHash.Diccionario[string, TDAUsuario.Usuario],
	usuario_loggeado *TDAUsuario.Usuario, id_siguiente_post *int, lista_posts *[]TDAUsuario.Post) {
	switch input[0] {

	case LOGIN:
		if len(input) < 2 {
			error := &errores.ErrorUsuarioInexistente{}
			fmt.Fprintf(os.Stdout, "%s\n", error.Error())
			return
		}
		if *usuario_loggeado != nil {
			error := &errores.ErrorUsuarioLoggeado{}
			fmt.Fprintf(os.Stdout, "%s\n", error.Error())
			return
		}
		var nombre = input[1]
		for i := 2; i < len(input); i++ {
			if i == len(input)-1 {
				input[i] = strings.Trim(input[i], "\n")
			}
			nombre = nombre + " " + input[i]
		}
		if !lista_usuarios.Pertenece(nombre) {
			error := &errores.ErrorUsuarioInexistente{}
			fmt.Fprintf(os.Stdout, "%s\n", error.Error())
			return
		}
		*usuario_loggeado = lista_usuarios.Obtener(nombre)
		fmt.Fprintf(os.Stdout, "%s %s\n", VALIDACION_LOGIN, nombre)

	case LOGOUT:
		if *usuario_loggeado == nil {
			error := &errores.ErrorNadieLoggeado{}
			fmt.Fprintf(os.Stdout, "%s\n", error.Error())
			return
		}
		*usuario_loggeado = nil
		fmt.Fprintf(os.Stdout, "%s\n", VALIDACION_LOGOUT)

	case PUBLICAR:

		if *usuario_loggeado == nil {
			error := &errores.ErrorNadieLoggeado{}
			fmt.Fprintf(os.Stdout, "%s\n", error.Error())
			return
		}
		post := (*usuario_loggeado).Publicar(strings.Join(input[1:], " "), *id_siguiente_post)
		iterador := lista_usuarios.Iterador()
		for iterador.HaySiguiente() {
			_, usuario := iterador.VerActual()
			usuario.ActualizarFeed(post)
			iterador.Siguiente()
		}
		*lista_posts = append(*lista_posts, post)
		*id_siguiente_post++
		fmt.Fprintf(os.Stdout, "%s\n", VALIDACION_PUBLICAR)

	case VER_PROXIMO:
		if *usuario_loggeado == nil || !(*usuario_loggeado).HayOtroPost() {
			error := &errores.ErrorVerSiguiente{}
			fmt.Fprintf(os.Stdout, "%s\n", error.Error())
			return
		}
		post := (*usuario_loggeado).VerProximo()
		fmt.Fprintf(os.Stdout, "Post ID %d\n%s dijo: %s\nLikes: %d\n",
			post.VerID(),
			post.Publicador(),
			post.Contenido(),
			post.CantidadLikes())

	case LIKE:
		id_post := strings.Trim(input[1], "\n")
		id_post_int, _ := strconv.Atoi(id_post)
		if len(input) < 2 || *usuario_loggeado == nil || (id_post != "0" && id_post_int == 0) || id_post_int >= len(*lista_posts) {
			error := &errores.ErrorLike{}
			fmt.Fprintf(os.Stdout, "%s\n", error.Error())
			return
		}
		(*lista_posts)[id_post_int].LikeadoPor(*usuario_loggeado)
		fmt.Fprintf(os.Stdout, "%s\n", VALIDACION_LIKE)

	case VER_LIKES:
		id_post := strings.Trim(input[1], "\n")
		id_post_int, _ := strconv.Atoi(id_post)
		if len(input) < 2 || (id_post != "0" && id_post_int == 0) || id_post_int >= len(*lista_posts) {
			error := &errores.ErrorVerLikes{}
			fmt.Fprintf(os.Stdout, "%s\n", error.Error())
			return
		}
		(*lista_posts)[id_post_int].VerLikes()
	}
}
