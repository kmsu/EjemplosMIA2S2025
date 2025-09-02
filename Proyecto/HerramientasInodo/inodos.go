package HerramientasInodos

import (
	"Proyecto/Herramientas"
	"Proyecto/Structs"
	"encoding/binary"
	"os"
	"strings"
)

// buscarinodo(0,/users.txt, superbloque, disco1.mia)
func BuscarInodo(idInodo int32, path string, superBloque Structs.Superblock, file *os.File) int32 {
	//Dividir la ruta por cada /
	stepsPath := strings.Split(path, "/")
	//el arreglo vendra [ ,val1, val2] por lo que me corro una posicion
	tmpPath := stepsPath[1:]
	//fmt.Println("Ruta actual ", tmpPath)

	//cargo el inodo a partir del cual voy a buscar
	var Inode0 Structs.Inode
	Herramientas.ReadObject(file, &Inode0, int64(superBloque.S_inode_start+(idInodo*int32(binary.Size(Structs.Inode{})))))
	//Recorrer los bloques directos (carpetas/archivos) en la raiz
	var folderBlock Structs.Folderblock
	for i := 0; i < 12; i++ {
		idBloque := Inode0.I_block[i]
		if idBloque != -1 {
			Herramientas.ReadObject(file, &folderBlock, int64(superBloque.S_block_start+(idBloque*int32(binary.Size(Structs.Folderblock{})))))
			//Recorrer el bloque actual buscando la carpeta/archivo en la raiz
			for j := 2; j < 4; j++ {
				//apuntador es el apuntador del bloque al inodo (carpeta/archivo), si existe es distinto a -1
				apuntador := folderBlock.B_content[j].B_inodo
				if apuntador != -1 {
					pathActual := Structs.GetB_name(string(folderBlock.B_content[j].B_name[:]))
					if tmpPath[0] == pathActual {
						//buscarInodo(apuntador, ruta[1:], path, superBloque, iSuperBloque, file, r)
						if len(tmpPath) > 1 {
							//return buscarIrecursivo(apuntador, tmpPath[1:], superBloque.S_inode_start, superBloque.S_block_start, file)
						} else {
							return apuntador
						}
					}
				}
			}
		}
	}
	//agregar busqueda en los apuntadores indirectos
	//i=12 -> simple; i=13 -> doble; i=14 -> triple
	//Si no encontro nada retornar 0 (la raiz)
	return idInodo
}

// Buscar inodo de forma recursiva. En el ejemplo evitar llegar a su usp
/*func buscarIrecursivo(idInodo int32, path []string, iStart int32, bStart int32, file *os.File) int32 {
	return 0
}*/
