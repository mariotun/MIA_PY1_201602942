
package comandos

import (
	"os"
	"bytes"
	"fmt"
	
	"encoding/binary"
	"log"
	"../estructuras"
)

func Escribir_MBR(path string,mbr estructuras.MbrStr ){

	archivo, err := os.OpenFile(path,os.O_RDWR,0777)
	defer archivo.Close()
	if err != nil { log.Fatal(err) }
	
	var binario3 bytes.Buffer
	binary.Write(&binario3, binary.BigEndian, mbr)

	_, err = archivo.Write(binario3.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}

func Leer_MBR(path string) (error,estructuras.MbrStr){

	file, err := os.Open(path)
	defer file.Close()

	if err!= nil {
		//log.Fatal(err)
		return err,estructuras.MbrStr{ }
	}

	m :=  estructuras.MbrStr{ }
	var size int = int(binary.Size(m))
	
	/*data := make([]byte, size)
	
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}*/
	data:=Leer_Bytes(file,size)
	buffer:=bytes.NewBuffer(data)
	
	err = binary.Read(buffer, binary.BigEndian, &m)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}
	
	return err,m
	
}

func Escribir_EBR(path string, ebr estructuras.EbrStr){

	archivo, err := os.OpenFile(path,os.O_RDWR,0777)
	defer archivo.Close()
	if err != nil { log.Fatal(err) }

	archivo.Seek(ebr.Part_start,0)

	extendida:=&ebr

	var binario_ebr bytes.Buffer
	binary.Write(&binario_ebr,binary.BigEndian,extendida)

	_,err=archivo.Write(binario_ebr.Bytes())
	if err!=nil{ 
		log.Fatal(err) 
	}
	//Escribir_Bytes(archivo,binario_ebr.Bytes())
}

func Leer_EBR(path string,start int64) (error,estructuras.EbrStr){

	file,err:=os.Open(path)

	defer file.Close()
	if err!= nil { 
		return err,estructuras.EbrStr{}
		//log.Fatal(err) 
	}

	file.Seek(start,0)

	ebr:=estructuras.EbrStr{ }
	var size int=int(binary.Size(ebr))

	data := Leer_Bytes(file, size)
	buffer := bytes.NewBuffer(data)

	err = binary.Read(buffer, binary.BigEndian, &ebr)
	if err != nil {
		fmt.Println(" Error al leer binario ")
		//log.Fatal("binary.Read failed", err)
	}

	return err,ebr
}








/*
func Escribir_Bytes(archivo *os.File,bytes []byte){
	_,err:=archivo.Write(bytes)
	if err!=nil{ log.Fatal(err) }
}*/

func Leer_Bytes(file *os.File, number int) []byte {
	bytes := make([]byte, number) //array de bytes

	_, err := file.Read(bytes) // Leido -> bytes
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}




	/*archivo.Seek(0,0)
	disco:=&mbr
	var binario3 bytes.Buffer
	binary.Write(&binario3,binary.BigEndian,disco)
	Escribir_Bytes(archivo,binario3.Bytes())*/

	/*file,err:=os.Open(path)

	defer file.Close()
	if err!= nil { log.Fatal(err) }

	m:=estructuras.MbrStr{ }
	var size int=int(binary.Size(m))

	data := LeerBytes(file, size)
	buffer := bytes.NewBuffer(data)

	err = binary.Read(buffer, binary.BigEndian, &m)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	return m*/