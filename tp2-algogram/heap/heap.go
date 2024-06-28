package cola_prioridad

const (
	CRITERIO_DISMINUICION = 4
	CAPACIDAD_INICIAL     = 10
)

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	cola := new(heap[T])
	cola.cmp = funcion_cmp
	cola.datos = make([]T, 0, CAPACIDAD_INICIAL)
	return cola
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	cola := new(heap[T])
	cola.cmp = funcion_cmp
	copia_arreglo := make([]T, len(arreglo))
	copy(copia_arreglo, arreglo)
	heapify(&copia_arreglo, cola.cmp)
	cola.datos = copia_arreglo
	cola.cantidad = len(arreglo)
	return cola
}

func swap[T any](x *T, y *T) {
	*x, *y = *y, *x
}

func heapify[T any](arreglo *[]T, cmp func(T, T) int) {
	for i := (len(*arreglo)) / 2; i >= 0; i-- {
		downHeap(i, arreglo, cmp, len(*arreglo))
	}
}

func downHeap[T any](indice int, lista *[]T, cmp func(T, T) int, tope int) {
	hijo_izq, hijo_der := calcularHijos(indice)
	if hijo_izq >= tope {
		return
	}
	if hijo_der < tope && cmp((*lista)[hijo_der], (*lista)[indice]) <= 0 && cmp((*lista)[hijo_izq], (*lista)[indice]) <= 0 {
		return
	}
	if cmp((*lista)[hijo_izq], (*lista)[indice]) > 0 {
		if hijo_der >= tope || cmp((*lista)[hijo_der], (*lista)[indice]) <= 0 {
			swap(&(*lista)[hijo_izq], &(*lista)[indice])
			downHeap(hijo_izq, lista, cmp, tope)
			return
		}
		if cmp((*lista)[hijo_der], (*lista)[indice]) > 0 {
			var remplazo int
			if cmp((*lista)[hijo_der], (*lista)[hijo_izq]) < 0 {
				remplazo = hijo_izq
			} else {
				remplazo = hijo_der
			}
			swap(&(*lista)[remplazo], &(*lista)[indice])
			downHeap(remplazo, lista, cmp, tope)
			return
		}
	}
	if hijo_der >= tope {
		return
	}
	if cmp((*lista)[hijo_der], (*lista)[indice]) > 0 {
		swap(&(*lista)[hijo_der], &(*lista)[indice])
		downHeap(hijo_der, lista, cmp, tope)
		return
	}
}

func calcularHijos(padre int) (int, int) {
	return 2*padre + 1, 2*padre + 2
}

func (cola heap[T]) EstaVacia() bool {
	return cola.Cantidad() == 0
}

func (cola heap[T]) Cantidad() int {
	return cola.cantidad
}

func (cola heap[T]) VerMax() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.datos[0]
}

func (cola *heap[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := cola.datos[0]
	swap(&cola.datos[0], &cola.datos[cola.cantidad-1])
	cola.cantidad--
	cola.datos = cola.datos[:cola.cantidad]
	downHeap(0, &cola.datos, cola.cmp, len(cola.datos))
	if cola.cantidad*CRITERIO_DISMINUICION <= len(cola.datos) {
		cola.achicar()
	}
	return dato
}

func (cola *heap[T]) achicar() {
	datos_nuevo := make([]T, cola.cantidad)
	copy(datos_nuevo, cola.datos)
	cola.datos = datos_nuevo
}

func (cola *heap[T]) Encolar(dato T) {
	cola.datos = append(cola.datos, dato)
	cola.cantidad++
	cola.upHeap(cola.datos, cola.cantidad-1)
}

func (cola heap[T]) upHeap(arreglo []T, indice int) {
	if indice == 0 {
		return
	}
	padre := (indice - 1) / 2
	if cola.cmp(arreglo[indice], arreglo[padre]) > 0 {
		swap(&arreglo[indice], &arreglo[padre])
		cola.upHeap(arreglo, padre)
	}
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	heapify(&elementos, funcion_cmp)
	cantidad := len(elementos) - 1
	for i := 0; i < len(elementos); i++ {
		swap[T](&elementos[0], &elementos[cantidad])
		downHeap[T](0, &(elementos), funcion_cmp, cantidad)
		cantidad--
	}

}
