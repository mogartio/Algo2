from cola import *

class GrafoIter:
	def __init__(self, grafo):
		self.grafo = grafo
		self.contador = 0
	
	def __iter__(self):
		return self
	def __next__(self):
		if self.contador >= len(self.grafo.claves) - 1:
			raise StopIteration
		self.contador += 1
		return self.grafo.claves[self.contador - 1]

class Vertice:
	def __init__(self, dato, posicion):
		self.dato  = dato
		self.id = posicion
	def __str__(self):
		return str(self.dato)
class Grafo:
	def __init__(self, dirigido):
		self.vertices = {}
		self.claves = []
		self.dirigido = dirigido

	def ver_vertice(self, posicion):
		return self.claves[posicion]

	def agregar_vertice(self, dato):
		nuevo_vertice = Vertice(dato, len(self.claves))
		self.vertices[nuevo_vertice] = {}
		self.claves.append(nuevo_vertice)

	def unir_vertices(self, origen, destino, peso):
		self.vertices[origen] = self.vertices.get(origen, {})
		self.vertices[origen][destino] = peso
		if not self.dirigido:
			self.vertices[destino] = self.vertices.get(destino, {})
			self.vertices[destino][origen] = peso

	def ver_peso(self, v, w):
		return self.vertices[v][w]
	
	def sacar_vertice(self, vertice_indeseado):
		self.claves.pop(vertice_indeseado.id)
		self.vertices[vertice_indeseado] = None
		for vertice in self.claves:
			if vertice.id > vertice_indeseado.id:
				vertice.id -= 1
			if vertice_indeseado in self.vertices[vertice]:
				del self.vertices[vertice][vertice_indeseado]


	def estan_unidos(self, origen, destino):
		unidos = (destino in self.vertices[origen] or (origen in self.vertices[destino] and not self.dirigido) )
		if unidos:
			return (unidos, self.vertices[origen][destino])
		return (False, None)

	def adyacentes(self, vertice):
		adyacentes = []
		for clave in self.vertices[vertice]:
			adyacentes.append(clave)
		return adyacentes
	def __iter__(self):
		return GrafoIter(self)

	def __str__(self):
		lista = ""
		for clave in self.claves:
			lista += str(clave)
			for arista in self.vertices[clave]:
				lista +=  ("-" + str(self.vertices[clave][arista]) +"->" + str(arista))
			lista += "\n" 
		return lista


def Trasponer(grafo):
	nuevo_grafo = Grafo(False)
	visitados = {}
	simul = []
	for v in grafo:
		if v not in visitados:
			simul.append(ciclo_bfs(v, visitados, grafo, nuevo_grafo))

	for xd in simul:
		for clave in xd:
			for valor in xd[clave]:
				nuevo_grafo.unir_vertices(clave, valor, 0)
				w = nuevo_grafo.ver_vertice(1)			
				u = nuevo_grafo.ver_vertice(0)
				print(nuevo_grafo.estan_unidos(w, u))
	print(simul)
	print(nuevo_grafo)


def dfs(grafo):
	visitados = {}
	contador_componentes_conexas = 0
	for v in grafo:
		if v not in visitados:
			contador_componentes_conexas += 1
			ciclo_dfs(grafo, v, visitados)
	print(contador_componentes_conexas)

def ciclo_dfs(grafo, vertice, visitados):
	visitados[vertice] = True
	print(vertice)
	for w in grafo.adyacentes(vertice):
		if w not in visitados:
			ciclo_dfs(grafo, w, visitados)


def se_conocen(grafo):
	orden = {}
	padres = {}
	visitados = {}
	componentes_conexas = 0
	for v in grafo:
		if v not in visitados:
			componentes_conexas += 1
			visitados[v] = True
			padres[v] = None
			orden[v] = 0
			_dfs(grafo, v, padres, visitados, orden)
	if componentes_conexas > 1:

		return False
	for claves in orden:
		for clave in orden:
			print(claves, clave)
			print(orden[clave])
			if orden[claves] - orden[clave] > 6 or orden[claves] - orden[clave] < -6 :
				return False
	return True

def _dfs(grafo, v, padres, visitados, orden):
	for w in grafo.adyacentes(v):
		if w not in visitados:
			visitados[w] = True
			orden[w] = orden[v] + 1
			_dfs(grafo, w, padres, visitados, orden)

def ciclo_bfs(vertice, visitados, grafo, nuevo_grafo):
	q = Cola()
	q.encolar(vertice)
	visitados[vertice] = True
	simul = {}
	nuevo_grafo.agregar_vertice(vertice.dato)
	while not q.esta_vacia():
		v = q.desencolar()
		visitados[v] = True
		for w in grafo.adyacentes(v):
			if w not in visitados:
				nuevo_grafo.agregar_vertice(w.dato)
				q.encolar(w)
				simul[w] = simul.get(w, [])
				simul[w].append(v)

	
	return simul




def main():
	grafo = Grafo(False)
	for i in range(0, 10):
		grafo.agregar_vertice(i)
		if i != 0:
			grafo.unir_vertices(grafo.claves[i], grafo.claves[i - 1], 10)
	vertice_aux = grafo.ver_vertice(2)
	verticee = grafo.ver_vertice(6)
	grafo.sacar_vertice(verticee)
	print(grafo)
	dfs(grafo)
	




