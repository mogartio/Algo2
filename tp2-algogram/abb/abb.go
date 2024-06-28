package diccionario

import TDAPila "algogram/pila"

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz *nodoAbb[K, V]
	cant int
	cmp  func(K, K) int
}

type iterRango[K comparable, V any] struct {
	diccionario abb[K, V]
	pila        TDAPila.Pila[*nodoAbb[K, V]]
	desde       *K
	hasta       *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	dicc := new(abb[K, V])
	dicc.cmp = funcion_cmp
	return dicc
}

func crearNodo[K comparable, V any](dato K, valor V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = dato
	nodo.dato = valor
	return nodo
}
func (diccionario abb[K, V]) Pertenece(clave K) bool {
	return diccionario.pertenece(clave, diccionario.raiz)
}

func (diccionario abb[K, V]) pertenece(buscado K, nodo *nodoAbb[K, V]) bool {
	if nodo == nil {
		return false
	}
	if diccionario.cmp(nodo.clave, buscado) == 0 {
		return true
	}
	if diccionario.cmp(nodo.clave, buscado) < 0 {
		return diccionario.pertenece(buscado, nodo.derecho)
	}
	if diccionario.cmp(nodo.clave, buscado) > 0 {
		return diccionario.pertenece(buscado, nodo.izquierdo)
	}
	return false
}

func (diccionario *abb[K, V]) Guardar(clave K, valor V) {
	if diccionario.raiz == nil {
		diccionario.raiz = crearNodo[K, V](clave, valor)
		diccionario.cant++
		return
	}
	diccionario.insertar(clave, valor, diccionario.raiz)
}

func (diccionario *abb[K, V]) insertar(clave K, valor V, nodo *nodoAbb[K, V]) {
	if diccionario.cmp(clave, nodo.clave) == 0 {
		nodo.dato = valor
		return
	}
	if diccionario.cmp(clave, nodo.clave) < 0 {
		if nodo.izquierdo == nil {
			nodo.izquierdo = crearNodo[K, V](clave, valor)
			diccionario.cant++
			return
		}
		diccionario.insertar(clave, valor, nodo.izquierdo)
		return
	}
	if diccionario.cmp(clave, nodo.clave) > 0 {
		if nodo.derecho == nil {
			nodo.derecho = crearNodo[K, V](clave, valor)
			diccionario.cant++
			return
		}
		diccionario.insertar(clave, valor, nodo.derecho)
		return
	}
	nodo.dato = valor
	diccionario.cant++
	return
}

func (diccionario abb[K, V]) Obtener(clave K) V {
	return diccionario.obtenerValor(clave, diccionario.raiz)
}

func (diccionario abb[K, V]) obtenerValor(clave K, nodo *nodoAbb[K, V]) V {
	var dato V
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	if diccionario.cmp(clave, nodo.clave) == 0 {
		dato = nodo.dato
	}
	if diccionario.cmp(clave, nodo.clave) < 0 {
		dato = diccionario.obtenerValor(clave, nodo.izquierdo)
	}
	if diccionario.cmp(clave, nodo.clave) > 0 {
		dato = diccionario.obtenerValor(clave, nodo.derecho)
	}
	return dato

}

func (diccionario abb[K, V]) Cantidad() int {
	return diccionario.cant
}

func (diccionario *abb[K, V]) Borrar(clave K) V {
	return diccionario.borrar(clave, diccionario.raiz, nil)
}

func (diccionario *abb[K, V]) borrar(clave K, hijo *nodoAbb[K, V], padre *nodoAbb[K, V]) V {
	var dato V
	if hijo == nil {
		panic("La clave no pertenece al diccionario")
	}
	if diccionario.cmp(clave, hijo.clave) == 0 {
		dato = diccionario.borrarEnlace(hijo, padre)
	}
	if diccionario.cmp(clave, hijo.clave) < 0 {
		dato = diccionario.borrar(clave, hijo.izquierdo, hijo)
	}
	if diccionario.cmp(clave, hijo.clave) > 0 {
		dato = diccionario.borrar(clave, hijo.derecho, hijo)
	}
	return dato
}

func (diccionario *abb[K, V]) borrarEnlace(hijo *nodoAbb[K, V], padre *nodoAbb[K, V]) V {
	var remplazo *nodoAbb[K, V]
	if hijo.izquierdo == nil && hijo.derecho == nil {
		remplazo = nil
	}
	if hijo.izquierdo == nil && hijo.derecho != nil {
		remplazo = hijo.derecho
	}
	if hijo.derecho == nil && hijo.izquierdo != nil {
		remplazo = hijo.izquierdo
	}
	if hijo.izquierdo != nil && hijo.derecho != nil {
		remplazo = diccionario.buscarMayorIzquierdo(hijo.izquierdo)
		diccionario.Borrar(remplazo.clave)
		diccionario.cant++ // ya que al llamar a Borrar dos veces, la cantidad decrece dos veces.
		if padre == nil {
			remplazo.derecho = diccionario.raiz.derecho
			remplazo.izquierdo = diccionario.raiz.izquierdo

		} else {
			remplazo.derecho = hijo.derecho
			remplazo.izquierdo = hijo.izquierdo
		}
	}
	if padre == nil {
		diccionario.raiz = remplazo
		diccionario.cant--
		return hijo.dato
	} else if diccionario.cmp(padre.clave, hijo.clave) > 0 {
		padre.izquierdo = remplazo
	} else if diccionario.cmp(padre.clave, hijo.clave) < 0 {
		padre.derecho = remplazo
	}
	diccionario.cant--
	return hijo.dato
}

func (diccionario abb[K, V]) buscarMayorIzquierdo(actual *nodoAbb[K, V]) *nodoAbb[K, V] {
	for actual.derecho != nil {
		actual = actual.derecho
	}
	return actual
}

func (diccionario abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	diccionario.iterarRango(nil, nil, visitar, diccionario.raiz, true)
}

func (diccionario abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	_ = diccionario.iterarRango(desde, hasta, visitar, diccionario.raiz, true)
}

func (diccionario abb[K, V]) iterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool, nodo *nodoAbb[K, V], corte bool) bool {
	if corte == false {
		return false
	}
	if nodo == nil {
		return corte
	}
	if desde == nil || diccionario.cmp(*desde, nodo.clave) < 0 {
		corte = diccionario.iterarRango(desde, hasta, visitar, nodo.izquierdo, corte)
		if corte == false {
			return false
		}
	}
	if (desde == nil || diccionario.cmp(*desde, nodo.clave) <= 0) && (hasta == nil || diccionario.cmp(*hasta, nodo.clave) >= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	if hasta == nil || diccionario.cmp(*hasta, nodo.clave) > 0 {
		corte = diccionario.iterarRango(desde, hasta, visitar, nodo.derecho, corte)
		if corte == false {
			return false
		}
	}
	return corte
}

func (diccionario abb[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := new(iterRango[K, V])
	iterador.pila = diccionario.crearPila(nil, nil)
	iterador.diccionario = diccionario
	return diccionario.IteradorRango(nil, nil)
}

func (diccionario abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := diccionario.crearPila(desde, hasta)
	iterador := new(iterRango[K, V])
	if desde != nil {
		iterador.desde = desde
	}
	if hasta != nil {
		iterador.hasta = hasta
	}
	iterador.pila = pila
	iterador.diccionario = diccionario
	return iterador
}

func (diccionario abb[K, V]) crearPila(desde *K, hasta *K) TDAPila.Pila[*nodoAbb[K, V]] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	actual := diccionario.raiz
	if desde != nil { // si el nodo actual es menor al desde, vamos hacia la derecha
		for actual != nil {
			if diccionario.cmp(actual.clave, *desde) < 0 {
				actual = actual.derecho
				continue
			}
			break
		}
	}
	for actual != nil {
		if desde == nil || diccionario.cmp(actual.clave, *desde) >= 0 {
			pila.Apilar(actual)
			actual = actual.izquierdo
			continue
		}
		actual = actual.derecho
	}
	return pila
}

func (iterador iterRango[K, V]) VerActual() (K, V) {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iterador.pila.VerTope()
	return actual.clave, actual.dato
}

func (iterador *iterRango[K, V]) HaySiguiente() bool {
	if iterador.pila.EstaVacia() {
		return false
	}
	if iterador.hasta == nil {
		return true
	}
	if iterador.diccionario.cmp((iterador.pila.VerTope()).clave, *iterador.hasta) > 0 {
		return false
	}
	return true
}

func (iterador *iterRango[K, V]) Siguiente() K {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iterador.pila.Desapilar()
	clave := actual.clave
	if actual.derecho == nil {
		return clave
	}
	actual = actual.derecho
	iterador.pila.Apilar(actual)
	actual = actual.izquierdo
	for actual != nil {
		if iterador.desde == nil || iterador.diccionario.cmp(actual.clave, *iterador.desde) >= 0 {
			iterador.pila.Apilar(actual)
			actual = actual.izquierdo
		}
	}
	return clave
}
