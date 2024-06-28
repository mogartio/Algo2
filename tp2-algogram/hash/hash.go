package diccionario

import (
	TDALista "algogram/lista"
	"fmt"
)

// CONSTANTES ----------------------------------------------
const (
	VACIO estadoCelda = iota
	OCUPADO
	BORRADO
)

const (
	FACTOR_CARGA_MINIMO = 0.3
	FACTOR_CARGA_MAXIMO = 0.7
	LARGO_TABLA_INICIAL = 10
)

// TYPES --------------------------------------------------
type estadoCelda int

type celda[V any, K comparable] struct {
	estado estadoCelda
	valor  V
	clave  K
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celda[V, K]
	cantidad int
	borrados int
	iter     TDALista.IteradorLista[celda[V, K]]
}

type iteradorDiccionarioExterno[K comparable, V any] struct {
	iterador *TDALista.IteradorLista[celda[V, K]]
}

// FUNCIONES CREADORAS: ----------------------

func crearcelda[V any, K comparable](clave K, valor V) *celda[V, K] {
	celda := new(celda[V, K])
	celda.valor = valor
	celda.clave = clave
	celda.estado = OCUPADO
	return celda
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.tabla = crearTabla[celda[V, K]](LARGO_TABLA_INICIAL)
	return hash
}

func crearTabla[V any](capacidad int) []V {
	return make([]V, capacidad)
}

// ---------------- PRIMITIVAS DEL DICCIONARIO ------------------------------------------
func (dic *hashCerrado[K, V]) Guardar(clave K, dato V) {
	// Si la clave existe, cambia el valor al nuevo dato
	posicion_clave := dic.obtenerPosClave(clave)
	if posicion_clave != -1 {
		dic.tabla[posicion_clave].valor = dato
		return

	}

	dic.cantidad++
	if dic.factorCarga() > FACTOR_CARGA_MAXIMO {
		dic.redimensionar(dic.largo() * 2)
	}
	celda := crearcelda(clave, dato)
	num_hash := f_hash(clave, dic.largo())

	if dic.tabla[num_hash].estado == VACIO { // Si está vacía, guarda sin problemas
		dic.tabla[num_hash] = *celda
		return
	} else if dic.tabla[num_hash].estado == BORRADO { // Si hay un borrado lo reemplaza
		dic.tabla[num_hash] = *celda
		dic.borrados--
		return
	}

	// Encuentra la primer posición vacía que encuentre
	posicion_vacia := dic.obtenerPosVacia(num_hash)
	dic.tabla[posicion_vacia] = *celda
}

func (dic *hashCerrado[K, V]) Pertenece(clave K) bool { //check
	posicion_clave := dic.obtenerPosClave(clave)
	return posicion_clave != -1 && dic.tabla[posicion_clave].estado == OCUPADO
}

func (dic *hashCerrado[K, V]) Obtener(clave K) V { //check
	posicion_clave := dic.obtenerPosClave(clave)
	if !(posicion_clave != -1 && dic.tabla[posicion_clave].estado == OCUPADO) {
		panic("La clave no pertenece al diccionario")
	}
	return dic.tabla[posicion_clave].valor
}

func (dic *hashCerrado[K, V]) Borrar(clave K) V {
	posicion_clave := dic.obtenerPosClave(clave)
	if !(posicion_clave != -1 && dic.tabla[posicion_clave].estado == OCUPADO) {
		panic("La clave no pertenece al diccionario")
	}
	dic.cantidad--
	dic.borrados++
	dic.tabla[posicion_clave].estado = BORRADO
	dato := dic.tabla[posicion_clave].valor
	if dic.factorCarga() <= FACTOR_CARGA_MINIMO {
		dic.redimensionar(dic.largo() / 2)
	}
	return dato
}

func (dic *hashCerrado[K, V]) Cantidad() int {
	return dic.cantidad
}

func (dic *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iteradorDiccionarioExterno[K, V])
	lista := TDALista.CrearListaEnlazada[celda[V, K]]()
	for i, elem := range dic.tabla {
		if dic.tabla[i].estado == OCUPADO {
			lista.InsertarUltimo(elem)
		}
	}
	dic.iter = lista.Iterador()
	iter.iterador = &dic.iter
	return iter
}

func (dic *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, celda := range dic.tabla {
		if celda.estado == OCUPADO {
			if !visitar(celda.clave, celda.valor) {
				return
			}
		}
	}

}

// -----------------PRIMITIVAS DEL ITERADOR ---------------------------

func (iter *iteradorDiccionarioExterno[K, V]) HaySiguiente() bool {
	return (*iter.iterador).HaySiguiente()

}

func (iter *iteradorDiccionarioExterno[K, V]) Siguiente() K {
	actual := (*iter.iterador).Siguiente()
	return actual.clave
}

func (iter *iteradorDiccionarioExterno[K, V]) VerActual() (K, V) {
	actual := (*iter.iterador).VerActual()
	return actual.clave, actual.valor
}

// --------------------------------------------------------------------
// Devuelve la posición de la clave, si la clave no existe devuelve -1
func (dic *hashCerrado[K, V]) obtenerPosClave(clave K) int {
	num_hash := f_hash(clave, dic.largo())
	if clave == dic.tabla[num_hash].clave && dic.tabla[num_hash].estado == OCUPADO {
		return num_hash
	} else if dic.tabla[num_hash].estado == VACIO {
		return -1
	}
	posicion_actual := num_hash

	for {
		if posicion_actual == dic.largo() {
			posicion_actual = 0
		}
		celda_actual := dic.tabla[posicion_actual]
		if celda_actual.estado == VACIO {
			break
		}
		if clave == celda_actual.clave {
			return posicion_actual
		}
		posicion_actual++

	}
	return -1

}

// Obtiene la primer posición vacía para una clave
func (dic *hashCerrado[K, V]) obtenerPosVacia(n_hash int) int {
	posicion_actual := n_hash + 1
	for {
		if posicion_actual == dic.largo() {
			posicion_actual = 0
		}
		if dic.tabla[posicion_actual].estado == VACIO {
			break
		}
		posicion_actual++
	}
	return posicion_actual
}

func (dic *hashCerrado[K, V]) redimensionar(capacidad int) {
	nuevo_dic := crearTabla[celda[V, K]](capacidad)
	copia_tabla := crearTabla[celda[V, K]](dic.largo())
	copy(copia_tabla, dic.tabla)
	dic.tabla = nuevo_dic
	for _, celda := range copia_tabla {
		if celda.estado != OCUPADO {
			continue
		}
		nuevo_id_hashing := f_hash(celda.clave, dic.largo())
		if dic.tabla[nuevo_id_hashing].estado == VACIO {
			dic.tabla[nuevo_id_hashing] = celda
			continue
		}
		posicion_vacia := dic.obtenerPosVacia(nuevo_id_hashing)
		dic.tabla[posicion_vacia] = celda
	}
	dic.borrados = 0
}

func (dic *hashCerrado[K, V]) largo() int {
	return len(dic.tabla)
}

func (dic *hashCerrado[K, V]) factorCarga() float32 {
	return float32(dic.cantidad+dic.borrados) / float32(dic.largo())
}

// ----------FUNCIONES PARA HASHING ---------------------------------

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func fnvHashing[K comparable](clave K, largo int) uint64 {
	var h uint64 = 14695981039346656037
	cla := convertirABytes(clave)
	for _, c := range cla {
		h *= uint64(1099511628211)
		h ^= uint64(c)
	}
	return h % uint64(largo)
}

func f_hash[K comparable](clave K, largo int) int {
	id_hashing := fnvHashing(clave, largo)
	return int(id_hashing)
}
