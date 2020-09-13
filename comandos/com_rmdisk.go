package comandos

import(

	"fmt"
	"os"
	"strings"
	"bufio"
)

func RMDISK(path string){
	fmt.Println("Dentro de la funcion rmdisk")
	
	nombreArchivo :=path // El nombre o ruta absoluta del archivo

	if strings.HasPrefix(nombreArchivo,"\"")==true{
		nombreArchivo=nombreArchivo[1:len(nombreArchivo)-1]
	}


	fmt.Println(" ¿Esta seguro que desea eliminar el DISCO(rmdisk) en la ruta escrita? S(si)/N(no).")
	fmt.Print(">> ")
	reader:=bufio.NewReader(os.Stdin)
	entrada,_:=reader.ReadString('\n')//leer hasta el separador de saldo de linea
	eleccion:= strings.TrimRight(entrada,"\r\n")
	
	if strings.ToLower(eleccion)=="s"{

	err := os.Remove(nombreArchivo)
	if err != nil {
		fmt.Printf(" Error al intentar eliminar el archivo(DISCO): %v\n", err)
	} else {
		fmt.Println(" Se Elimino el DISCO correctamente")
	}

	}else if strings.ToLower(eleccion)=="n"{
		fmt.Println(" Mensaje: No se elimino ningun DISCO.")

	}else{
		fmt.Println(" Mensaje: Opcion Incorrecta")
	}


}