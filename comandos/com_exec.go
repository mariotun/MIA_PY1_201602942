package comandos

import(
	"fmt"
	"os"
	//"os/exec"
	"bufio"
	"strings"
	"../lexico"
	
	//"strconv"
	//"../funciones"
)

func Exec(ruta string)  {
	fmt.Println("Dentro de la funcion exec")
	i:=0
	archivo,error:=os.Open(ruta)

	//fmt.Println("iiiii"+ruta)
	if error!=nil{
		fmt.Println("Hubo un error al leer el archivo")
	}

	scanner:=bufio.NewScanner(archivo)
	
	ln:=""
	for scanner.Scan(){
		i++
		
		linea:=scanner.Text()
		if strings.HasSuffix(linea, "*")==true{
			 ln+=linea[0:len(linea)-2]
		}else{
	
			entrada:=Corregir_Entrada(ln+linea)
		LeerEntrada(strings.ToLower(entrada))
		//LeerEntrada(linea)
		fmt.Println(i,entrada)
		ln=""
		}
	}


}

func Corregir_Entrada(entrada string)string{

	salida:=""
	splitFunc := func(r rune) bool {
		return strings.ContainsRune(" ,\n,\r,\t", r)
	}

	palabras := strings.FieldsFunc(entrada, splitFunc)
	for _, palabra := range palabras {
		//fmt.Printf("Palabra %d es: %s\n", idx, palabra)
		salida+=palabra
	}
	return salida
}



func LeerEntrada(entrada string)  {

	estado:=lexico.Alexico(entrada)

	if estado==1{
		fmt.Println("	Error: Comando no reconocido totalmente(Dentro del archivo).")

	}else if estado ==0{

		//fmt.Println("	Comando escrito correctamente(exec).")
		
		ElegirComando(entrada)
	}


}

func Crear_Carpeta(directorio string){

	if _, err := os.Stat(directorio); os.IsNotExist(err) {

		err = os.MkdirAll(directorio, os.ModePerm)

		if err != nil {
		  
		  panic(err)
		}
	  }

}

func ElegirComando(entrada string){

	ncomando:=lexico.Parametros.Get_Comando()
	fmt.Println("--->"+ncomando)

	if ncomando=="pause"{
		bufio.NewReader(os.Stdin).ReadBytes('\n') 

	}else if ncomando=="mkdisk"{

		nsize:=lexico.Parametros.Get_Size()
		npath:=lexico.Parametros.Get_Path()
		nname:=lexico.Parametros.Get_Name()
		nunit:=lexico.Parametros.Get_Unit()
		
		if nsize!=0 && npath!="" && nname!="" {

		MKDISK(nsize,npath,nname,nunit)
		
		}else{
			fmt.Println(" Error: El valor de algun parametro no es el correcto para el comando. ")
		}
		lexico.Parametros.Limpiar()

	}else if ncomando=="rmdisk"{

		npath:=lexico.Parametros.Get_Path()
		if npath!=""{
		RMDISK(npath)
		}else{
			fmt.Println(" Error: No se encontro la ruta del disco a eliminar.")
		}
		lexico.Parametros.Limpiar()

	}else if ncomando=="fdisk"{
		nsize:=lexico.Parametros.Get_Size()
		npath:=lexico.Parametros.Get_Path()
		nnamefd:=lexico.Parametros.Get_Namefd()
		nunit:=lexico.Parametros.Get_Unit()
		ntipo:=lexico.Parametros.Get_Tipo()
		nfit:=lexico.Parametros.Get_Fit()
		ndelete:=lexico.Parametros.Get_Delete()
		nadd:=lexico.Parametros.Get_Add()

		fmt.Println("size:",nsize)
		fmt.Println(" path:"+npath+" name:"+nnamefd+" unit:"+nunit+" tipo:"+ntipo)
		fmt.Println(" fit:"+nfit+" delete:"+ndelete)
		fmt.Print("add:",nadd)
		
		if strings.HasPrefix(npath,"\"")==true{
			npath=npath[1:len(npath)-1]
		}

		FDISK(nsize,nunit,npath,ntipo,nfit,ndelete,nnamefd,nadd)

		lexico.Parametros.Limpiar()

	}else if ncomando=="mount"{

		npath:=lexico.Parametros.Get_Path()
		nnamefd:=lexico.Parametros.Get_Namefd()

		if strings.HasPrefix(npath,"\"")==true{
			npath=npath[1:len(npath)-1]
		}

		MOUNT(npath,nnamefd)
		
		lexico.Parametros.Limpiar()

	}else if ncomando=="unmount"{



		UNMOUNT()
		lexico.Parametros.Limpiar()

	}else if ncomando=="rep"{
		npath:=lexico.Parametros.Get_Path()
		nnamegrafic:=lexico.Parametros.Get_NameGrafic()
		
		REP(npath,nnamegrafic)
		lexico.Parametros.Limpiar()
	}
	/*else{
		fmt.Println(" (*) No hay un comando con ese combre")
	}*/
	

}


/*func ElegirComando2(entrada string)  {


	var comandoArray []string
	comandoArray = strings.Split(entrada, "-")

	/*if comandoArray[0]=="exec"{
		fmt.Println("El exec con los demas comandos")

	}else 
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


type ParamComandos struct {
    nombre string
    ruta   string
}

var parametros ParamComandos

func nada(){
parametros.nombre="mario"
parametros.ruta="/home/mario"

Datos(parametros)

}

func Datos(dat ParamComandos) ParamComandos{

	return dat
}*/