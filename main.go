package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
	//"os/exec"

	//"./comandos"
	"./funciones"
)

func main()  {
	//fmt.Println("proyecto de archivos")
	salida:=1
	
	encabezado:=`
	-----------------------------------------		
	 UNIVERSIDAD DE SAN CARLOS DE GUATEMALA
	 FACULTAD DE INGENIERA
	 MANEJO E IMPLEMENTACION DE ARCHIVOS
	 MARIO TUN - 201602942
	-----------------------------------------
	`
	for salida ==1{
		fmt.Print(encabezado+">> ")
		reader:=bufio.NewReader(os.Stdin)

		entrada,_:=reader.ReadString('\n')//leer hasta el separador de saldo de linea
		eleccion:= strings.TrimRight(entrada,"\r\n")//remover el salto de linea de ...

		//funciones.LineaComando(eleccion)
		funciones.LeerArchivoEntrada(eleccion)

		/*switch eleccion {
			
			case "1":
				fmt.Println("es la opcion 1")
				comandos.MKDISK()
			case "2":
				fmt.Println("es la opcion 2")
				comandos.RMDISK()
			case "exit":
				salida=0
			case "clear":
				c := exec.Command("clear") 
				c.Stdout = os.Stdout 
				c.Run()
			default:
				fmt.Println("Comando incorrecta!!! ")
		
			
		}*/
salida=0
	}


	



}