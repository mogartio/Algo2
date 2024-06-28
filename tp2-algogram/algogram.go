package main

import (
	comandos "algogram/comandos"
	TDAHash "algogram/hash"
	TDAUsuario "algogram/usuario"
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	scanner_archivo, _ := validar_archivos()
	lista_usuarios := CrearListaUsuarios(scanner_archivo)
	if lista_usuarios.Cantidad() == 0 {
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	var id_siguiente_post = 0
	var lista_posts []TDAUsuario.Post
	var lines []string
	var usuario_loggeado TDAUsuario.Usuario
	for scanner.Scan() {
		line := scanner.Text()
		lines = strings.Split(line, " ")
		comandos.LeerInput(lines, lista_usuarios, &usuario_loggeado, &id_siguiente_post, &lista_posts)
	}
}

func CrearListaUsuarios(scanner bufio.Scanner) TDAHash.Diccionario[string, TDAUsuario.Usuario] {
	diccionario := TDAHash.CrearHash[string, TDAUsuario.Usuario]()
	var indice = 1
	for scanner.Scan() {
		usuario_nuevo := TDAUsuario.CrearUsuario(scanner.Text(), indice)
		diccionario.Guardar(usuario_nuevo.Nombre(), usuario_nuevo)
		indice++
	}
	return diccionario
}

func validar_archivos() (bufio.Scanner, error) {

	// Devuelve una lista de archivos que pueden leerse, si alguno no pudo abrir devuelve un error
	var scanner bufio.Scanner
	args := os.Args
	if len(args) < 2 {
		return scanner, io.EOF
	}
	archivo := args[1]
	usuarios, error := os.Open(archivo)
	scanner = *bufio.NewScanner(usuarios)
	if error != nil {
		return scanner, io.EOF
	}
	return scanner, nil
}
