package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type estudiante struct {
	Id        int64
	Nombre    [25]byte // Nombre fijo de 25 bytes
	Siguiente int64
}

func main() {
	menu()
}

func menu() {
	var opcion = 0

	for {
		fmt.Println("---------------------------------------------------")
		fmt.Println("-----------------Escoja una opción-----------------")
		fmt.Println("---------------------------------------------------")
		fmt.Println("1. Crear archivo binario.")
		fmt.Println("2. Eliminar archivo binario.")
		fmt.Println("3. Crear estudiante.")
		fmt.Println("4. Leer estudiantes.")
		fmt.Println("5. Salir.")
		fmt.Scanf("%d\n", &opcion)

		switch opcion {
		case 1:
			crearArchivo()
		case 2:
			eliminarArchivo()
		case 3:
			crearEstudiante()
		case 4:
			leerEstudiantes()
		case 5:
			salir()
		}

	}
}

func crearArchivo() {
	file, err := os.Create("estudiantes.bin")
	if err != nil {
		fmt.Println("Error creando el archivo:", err)
		return
	}
	defer file.Close()

	fmt.Println("Archivo creado exitosamente.")
}

func crearEstudiante() {
	var e estudiante
	fmt.Print("Ingrese ID: ")
	fmt.Scan(&e.Id)

	fmt.Print("Ingrese nombre: ")
	var nombre string
	fmt.Scan(&nombre)

	// Copiar el nombre a los 25 bytes
	copy(e.Nombre[:], []byte(nombre))

	// Por ahora ponemos Siguiente como -1 (o lo que necesites)
	e.Siguiente = -1

	//os.O_APPEND abre el archivo al final.
	file, err := os.OpenFile("estudiantes.bin", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error abriendo el archivo:", err)
		return
	}
	defer file.Close()

	//file, es el archivo donde escribire los datos
	//binary.LittleEndian especifica el orden de los bytes. aquí se usa el byte menos significativo se escribe primero
	err = binary.Write(file, binary.LittleEndian, &e)
	if err != nil {
		fmt.Println("Error escribiendo estudiante:", err)
		return
	}

	fmt.Println("Estudiante guardado.")
}

func leerEstudiantes() {
	file, err := os.Open("estudiantes.bin")
	if err != nil {
		fmt.Println("Error abriendo el archivo:", err)
		return
	}
	defer file.Close()

	for {
		var e estudiante
		err := binary.Read(file, binary.LittleEndian, &e)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error leyendo estudiante:", err)
			break
		}

		// Convertir nombre a string sin bytes vacíos
		nombre := string(bytes.Trim(e.Nombre[:], "\x00"))
		fmt.Printf("ID: %d, Nombre: %s, Siguiente: %d\n", e.Id, nombre, e.Siguiente)
	}
}

func eliminarArchivo() {
	err := os.Remove("estudiantes.bin")
	if err != nil {
		fmt.Println("Error al eliminar el archivo:", err)
		return
	}
	fmt.Println("Archivo eliminado correctamente.")
}

func salir() {
	fmt.Println("Saliendo del programa.")
	os.Exit(0)
}
