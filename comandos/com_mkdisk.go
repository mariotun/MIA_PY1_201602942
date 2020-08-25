
package comandos

import(

	"fmt"
	"bytes"
	"os"
	//"os/exec"

	"log"
	"encoding/binary"
	//"unsafe"

	"../estructuras"
	//"strings"
	//"strconv"
	"time"
	//"math/rand"
	

)

 var nrandom int=10
 //var nsize int64=0

func MKDISK(size int64,path string,name string,unit string){
	/*fmt.Println("Dentro de la funcion mkdisk :\n"+" size:"+size+"\n path:"+path+
	"\n name:"+name+"\n unit:"+unit+"\n")*/
	Escribir_Archivo(size,path,name,unit)

}

func Escribir_Archivo(size int64,path string,name string,unit string){	

	if size>0 && (unit=="k" || unit=="m" || unit==""){

	DiscoDuro:=estructuras.MbrStr{ }//inicializamos el struct para el mbr
	//fmt.Println("-->",path)
	Crear_Carpeta(path)

	dir:=path+name
	archivo,err:=os.Create(dir) //se crea el archivo binario con el nombre del paramtro
	defer archivo.Close()
	if err!=nil{ 
		log.Fatal(err)
	}

	
	/*tam, err := strconv.ParseInt(size, 10, 64)
	if err == nil { //fmt.Printf("%d of type %T", n, n) 
	}else{  fmt.Println("NO es un numero") }*/


	var ceros int8=0
	valor:=&ceros
	//fmt.Println(" *Tamaño de la variable ceros:",unsafe.Sizeof(ceros))

	//se va a escribir un 0 en el inicio del archivo
	 var binario bytes.Buffer
	 binary.Write(&binario,binary.BigEndian,valor)
	 Escribir_Bytes(archivo,binario.Bytes())

	if unit=="k"{
		nsize:=(size*1024)
		archivo.Seek(nsize-1,0)
		DiscoDuro.Mbr_tamano=int64(nsize)//ingresamos el tamaño del disco al mbr
	}else if unit=="m" || unit==""{
		nsize:=(size*1024*1024)
		archivo.Seek(nsize-1,0)
		DiscoDuro.Mbr_tamano=int64(nsize)//ingresamos el tamaño del disco al mbr
	}
	//se va a escribir un 0 al final del archivo
	 var binario2 bytes.Buffer
	 binary.Write(&binario2,binary.BigEndian,valor)
	 Escribir_Bytes(archivo,binario2.Bytes())

	//ahora toca escribir el struct(mbr) en el archivo 
	t := time.Now()
	fecha := fmt.Sprintf("%d/%02d/%02d - %02d:%02d:%02d",
		t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute(), t.Second())
	copy(DiscoDuro.Mbr_fecha_creacion[:],fecha)//se ingresa la fecha y hora actual


	nrandom++
	DiscoDuro.Mbr_disk_signature=nrandom//se ingresa un numero random para identificar a cada disco
	fmt.Println("random:",DiscoDuro.Mbr_disk_signature)
	fmt.Println("Tamaño:",DiscoDuro.Mbr_tamano)
	fmt.Println("hora-fecha:",string(DiscoDuro.Mbr_fecha_creacion[:]))

	//nos posicionamos en el incio del archivo para escribir el mbr
	archivo.Seek(0,0)

	disco:=&DiscoDuro
	var binario3 bytes.Buffer
	binary.Write(&binario3,binary.BigEndian,disco)
	Escribir_Bytes(archivo,binario3.Bytes())

	fmt.Println("--Mensaje: Se creo el disco correctamente.")

	}else{
		fmt.Print(" Error:\n (1)El valor de Size no es el correcto para la creacion del disco."+
			     "\n (2)La letra para las unidades no es la correcta.")
	}


}

func Escribir_Bytes(archivo *os.File,bytes []byte){
	_,err:=archivo.Write(bytes)
	if err!=nil{ log.Fatal(err) }
}