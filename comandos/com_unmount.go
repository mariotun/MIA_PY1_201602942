package comandos

import(
	"fmt"
	"strings"
)

func UNMOUNT(entrada string)  {
	fmt.Println("Dentro de la funcion unmount",entrada)

	parametros := strings.Split(entrada, ",")
	actual:=Lista.first
	fmt.Println("longitud: ",len(parametros))

	if ( actual != nil) {

		for i:=1 ; i < len(parametros) ; i++{
			fmt.Println("datos: ",parametros[i])

		eliminar:=Lista.EliminarNodo(parametros[i])
		if ( eliminar == 1){
			fmt.Println(" Mensaje: La particion se desmoto con exito. ")
		}else{
			fmt.Println( " Error: La particion "+parametros[i]+" no esta montada. ")
		}

			}

	}else{
		fmt.Println( " Error: No existe ninguna unidad montada. ")
	}

	

/*	parametros := strings.Split(entrada, "")

	for i:=0 ; i < len(parametros) ; i++{
		if(parametros[i]!="unmount" || parametros[i]!="-" || parametros[i]!="->"){
		fmt.Println(parametros[i])
		}
	}

	fmt.Println("-----------------------------------------------")

	/*parametros2:=strings.Split(parametros[:], "-")
	for i:=0 ; i < len(parametros2) ; i++{
		fmt.Println(parametros2[i])
	}*/
}