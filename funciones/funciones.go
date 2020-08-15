
package funciones

import(
	"fmt"
	"strings"
	"os"
	"bufio"
	"../lexico"
	"../comandos"
	//"regexp"
)

func LineaComando(comando string) {
	var commandArray []string
	commandArray = strings.Split(comando, " ")
	fmt.Println(commandArray[0])
  //ejecutarComando(commandArray) 
}

func ejecutarComando(commandArray []string) {
	data := strings.ToLower(commandArray[0])
	  if data == "crear" {
		  fmt.Println("Creando un archivo")
	  }else {
		  fmt.Println("Otro Comando")
	  }
  }

func LeerArchivoEntrada(ruta string){

	i:=0
	archivo,error:=os.Open(ruta)

	if error!=nil{
		fmt.Println("Hubo un error al leer el archivo")
	}

	scanner:=bufio.NewScanner(archivo)
	//scanner.Split(bufio.ScanWords)
	
	for scanner.Scan(){
		i++

		//re := regexp.MustCompile(`[- ->]`)


		linea:=scanner.Text()
		split := strings.Split(linea, "-")
		//res:=re.Split(linea, -1)
		fmt.Println(i,split)
		//fmt.Println(i,res)
	}

}


func LeerEntrada2(entrada string)  {

	estado:=lexico.Alexico(entrada)

	if estado==1{
		fmt.Println("	Error: Comando no reconocido totalmente.")

	}else if estado ==0{

		//fmt.Println("	Comando escrito correctamente.")
		ElegirComando2(entrada)
	}


}

func ElegirComando2(comando string)  {


	//componentes := strings.Split(comando, "-","->")

	//fmt.Println(componentes)
	var comandoArray []string
	comandoArray = strings.Split(comando, "-")
	

	if comandoArray[0]=="exec"{
		
		comandos.Exec("/home/mario/Escritorio/entrada.txt")

	}else {

		comandos.ElegirComando(comandoArray[0])
	}
	
	
	
	
	


	
}