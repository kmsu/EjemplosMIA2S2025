package main

import (
	DM "Proyecto/Comandos/AdministradorDiscos" //DM -> DiskManagement (Administrador de discos)
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//MENSAJES DE INICIO
	Ms_inicio := "Bienvenido escriba un comando"
	Ms_info := "Si desesa salir escriba el comando exit"
	fmt.Println(Ms_inicio)
	fmt.Println(Ms_info)
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\n$: ")
		reader.Scan() //Lee la entrada y la prepara para ser usada mas adelante.

		entrada := strings.TrimRight(reader.Text(), " ") //Quitar espacios vacios a la derecha

		//estructura de comentarios
		// mkdisk -parametro=valor #comentarioentrada -> [comando][comentario]
		//#comentario mkdisk -parametro=valor [][comentario]
		linea := strings.Split(entrada, "#") //para ignorar comentarios desde la consola manual
		//Llamada a metodo analizar
		if strings.ToLower(linea[0]) != "exit" {
			analizar(linea[0])
		} else {
			fmt.Println("Salida exitosa")
			break
		}
	}
}

func analizar(entrada string) {
	//Separar los parametros -size=3000 -path=ruta (obtenemos la lista: size=3000, path=ruta)

	//mkdisk -size=3000 -path=ruta
	//Quitar espacios en blanco del final
	tmp := strings.TrimRight(entrada, " ")
	parametros := strings.Split(tmp, " -")

	// ----------------------------------------------------------------    Eliminar el if de execute en la explicacion de la clase  -----------------------------

	//analizamos los parametros
	//--------------------------------- ADMINISTRADOR DE DISCOS ------------------------------------------------
	if strings.ToLower(parametros[0]) == "mkdisk" {
		//MKDISK
		//crea un archivo binario que simula un disco con su respectivo MBR
		if len(parametros) > 1 {
			DM.Mkdisk(parametros)
		} else {
			fmt.Println("MKDISK ERROR: parametros no encontrados")
		}

	} else if strings.ToLower(parametros[0]) == "exit" {
		fmt.Println("Salida exitosa")
		os.Exit(0)

	} else if strings.ToLower(parametros[0]) == "" {
		//para agregar lineas con cada enter sin tomarlo como error
	} else {
		fmt.Println("Comando no reconocible")
	}
}
