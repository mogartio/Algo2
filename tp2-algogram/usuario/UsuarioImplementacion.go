package usuario

import (
	heap "algogram/heap"
	"math"
)

type usuarioImplementacion struct {
	nombre   string
	indice   int
	loggeado bool
	feed     heap.ColaPrioridad[*post]
}

func CrearUsuario(nombre string, posicion int) Usuario { //Será llamada una vez por cada id que nos den en main
	nuevo_usuario := new(usuarioImplementacion)
	nuevo_usuario.nombre = nombre
	nuevo_usuario.indice = posicion
	nuevo_usuario.feed = heap.CrearHeap[*post](func(A *post, B *post) int {
		diferenciaA := int(math.Abs(float64((A.IndicePublicador())) - float64(nuevo_usuario.indice)))
		diferenciaB := int(math.Abs(float64((B.IndicePublicador())) - float64(nuevo_usuario.indice)))
		if diferenciaA > diferenciaB {
			return -1
		}
		if diferenciaA < diferenciaB {
			return 1
		}
		if A.VerID() > B.VerID() {
			return -1
		}
		return 1
	})
	return nuevo_usuario
}

func (usuario usuarioImplementacion) Nombre() string {
	return usuario.nombre
}

func (usuario usuarioImplementacion) Publicar(contenido string, id int) Post {
	return CrearPost(id, usuario.nombre, contenido, usuario.indice)
}

func (usuario *usuarioImplementacion) ActualizarFeed(post_nuevo Post) {
	post_aux := (post_nuevo).(*post)
	if post_nuevo.Publicador() == usuario.nombre {
		return
	}
	usuario.feed.Encolar(post_aux)
}

func (usuario *usuarioImplementacion) VerProximo() Post {
	nuevo_post := (usuario.feed.Desencolar())
	return nuevo_post
}

func (usuario usuarioImplementacion) Indice() int {
	return usuario.indice
}

func (usuario usuarioImplementacion) HayOtroPost() bool {
	return usuario.feed.Cantidad() >= 1
}
