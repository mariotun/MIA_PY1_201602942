package comandos

import(
	"fmt"
	"os"
	//"os/exec"
	"bufio"
	"strings"
	"../lexico"
	//"../funciones"
)

func Exec(ruta string)  {
	fmt.Println("Dentro de la funcion exec")
	i:=0
	archivo,error:=os.Open(ruta)

	fmt.Println("iiiii"+ruta)
	if error!=nil{
		fmt.Println("Hubo un error al leer el archivo")
	}

	scanner:=bufio.NewScanner(archivo)
	
	
	for scanner.Scan(){
		i++
		linea:=scanner.Text()
		LeerEntrada(strings.ToLower(linea))
		fmt.Println(i,linea)
	}


}



func LeerEntrada(entrada string)  {

	estado:=lexico.Alexico(entrada)

	if estado==1{
		fmt.Println("	Error: Comando no reconocido totalmente(exec).")

	}else if estado ==0{

		//fmt.Println("	Comando escrito correctamente(exec).")
		ElegirComando(entrada)
	}


}

func ElegirComando(entrada string)  {


	var comandoArray []string
	comandoArray = strings.Split(entrada, "-")

	/*if comandoArray[0]=="exec"{
		fmt.Println("El exec con los demas comandos")

	}else */
	if comandoArray[0]=="#"{
		fmt.Println("es un comentario")
	}
	if comandoArray[0]=="mkdisk"{
		//fmt.Println("Dentro de opcion mkdisk")
		MKDISK()

	}else if comandoArray[0]=="pause"{
		fmt.Println("	Presione enter para continuar!!")
		bufio.NewReader(os.Stdin).ReadBytes('\n') 
		

	}else if comandoArray[0]=="rmdisk"{
		//fmt.Println("Dentro de opcion rmdisk")
		RMDISK()

	}else if comandoArray[0]=="fdisk"{
		//fmt.Println("Dentro de opcion fdisk ")
		FDISK()

	}else if comandoArray[0]=="mount"{
		//fmt.Println("Dentro de opcion mount")
		MOUNT()

	}else if comandoArray[0]=="unmount"{
		//fmt.Println("Dentro de opcion unmount")
		UNMOUNT()

	}

	
}