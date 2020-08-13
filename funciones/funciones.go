
package funciones

import(
	"fmt"
	"strings"
	"os"
	"bufio"
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
	
	
	for scanner.Scan(){
		i++
		linea:=scanner.Text()
		fmt.Println(i,linea)
	}

}