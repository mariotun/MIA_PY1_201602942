package comandos

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	//"../estructuras"
)

func FDISK(size int64,unit string,path string,tipo string,fit string,delete string,name string,add string)  {
	fmt.Println(" Dentro de la funcion fdisk")

	if size > 0 {
		if (delete!="" || add!=""){
			fmt.Println(" Mensaje: La creacion de una particon no acepta los parametros Delete y Add.")
		}else{
			Crear_Particiones(size,unit,path,tipo,fit,name)
		}

	}else if add!="" {
		if (size>0 || delete!=""){
			fmt.Println(" Mensaje: La Modificacion del tamaño de una particon no acepta los parametros Delete y Size.")
		}else{
			Agregar_Quitar_Particiones(path,name,add,unit)
		}

	}else if delete!=""{
		if ( size>0 || add!="" || fit!="" || tipo!=""){
			fmt.Println(" Mensaje: La Eliminacion de una particon no acepta los parametros Size,Add,Fit y Tipo.")
		}else{
			fmt.Println(" ¿Esta seguro que desea eliminar la particion? S(si)/N(no).")
			fmt.Print(">> ")
			reader:=bufio.NewReader(os.Stdin)
			entrada,_:=reader.ReadString('\n')//leer hasta el separador de saldo de linea
			eleccion:= strings.TrimRight(entrada,"\r\n")

			if strings.ToLower(eleccion)=="s"{
				Eliminar_Particion(path,name,delete)
			}else if strings.ToLower(eleccion)=="n"{
				fmt.Println(" Mensaje: No se elimino ninguna particion.")
			
			}else{
				fmt.Println(" Mensaje: Opcion Incorrecta")
				}
				
		
		}

	}else{
		fmt.Println(" Mensaje: Hay parametros que no son permitidos para el comando a ejecutar.")
	}

	
}

func Crear_Particiones(size int64,unit string,path string,tipo string,fit string,name string){

	if tipo=="p"{
		if ExisteArchivo(path)==true{
			Realizar_Particion_Primaria(path,name,size,fit,unit)
		}else{
			fmt.Println(" Error: El disco donde se desea crear la particion no existe.")
		}

	}else if tipo=="e"{
		if ExisteArchivo(path)==true{
			Realizar_Particion_Extendida(path,name,size,fit,unit)
		}else{
			fmt.Println(" Error: El disco donde se desea crear la particion no existe.")
		}

	}else if tipo=="l"{
		if ExisteArchivo(path)==true{
			Realizar_Particion_Logica(path,name,size,fit,unit)
		}else{
			fmt.Println(" Error: El disco donde se desea crear la particion no existe.")
		}


	}
	
}

func ExisteArchivo(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func Realizar_Particion_Primaria(path string,name string,size int64,fit string,unit string){

//	MBR:=estructuras.MbrStr{ }



}


func Realizar_Particion_Extendida(path string,name string,size int64,fit string,unit string){

}


func Realizar_Particion_Logica(path string,name string,size int64,fit string,unit string){


}

func Agregar_Quitar_Particiones(path string,name string,add string,unit string){


}

func Eliminar_Particion(path string,name string,delete string){


}