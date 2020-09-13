package comandos

import(
	"fmt"
	"os"
	//"io"
	
	//"io/ioutil"
	"encoding/binary"
	
	//"bytes"
	"bufio"
//	"log"
	//"unsafe"
	"strings"
	"strconv"
	"../estructuras"
)

func FDISK(size int64,unit string,path string,tipo string,fit string,delete string,name string,add string)  {
	fmt.Println(" Dentro de la funcion fdisk")

	if size > 0 {
		if (delete!="" || add!=""){
			fmt.Println(" Mensaje: La creacion de una particon no acepta los parametros Delete y Add.")
		}else{
			fmt.Println(" EN PARTICIONES")
			Crear_Particiones(size,unit,path,tipo,fit,name)
		}

	}else if add!="" {
		if (size>0 || delete!=""){
			fmt.Println(" Mensaje: La Modificacion del tamaño de una particon no acepta los parametros Delete y Size.")
		}else{
			fmt.Println(" EN QUITAR PARTICIONES ")
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
				fmt.Println(" EN ELIMINAR PARTICIONES")
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

	if ( tipo == "p" || tipo == "" ) {
		if ExisteArchivo(path)==true{
			fmt.Println(" R_P_P")
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
    /*if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }else{ 
			return true
		}
    }
	return false*/
	filename := name
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("file does not exist")
		return false
	}else{
	return true }
}

func Realizar_Particion_Primaria(path string,name string,size int64,fit string,unit string){


	var size_completo int64=1024
	//var buffer byte='1'
	var fit_aux byte

	if ( fit == "bf" ){
		fit_aux='b'
	}else if ( fit == "wf" || fit == "" ){
		fit_aux='w'
	}else if ( fit == "ff"){
		fit_aux='f'
	}

	if ( unit == "b" ){
		size_completo=size
	}else if ( unit == "k" || unit == "" ){
		size_completo=(size*1024)
	}else if ( unit == "m" ){
		size_completo=(size*1024*2024)
	}


	
 //masterb:=estructuras.MbrStr{ }
	
 //archivo,err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0777)
 //fmt.Println("--->",path)
	err,masterb:=Leer_MBR(path)	

 	if err != nil {
		fmt.Println(" Error: El disco no existe.")
		//archivo.Close()
		panic(err)
	}else{
		
		//io.WriteString(f, "+ test string")
		var fparticion bool=false
		var numparticion int=0
		
		/*
		archivo.Seek(0,0)
		var sizemb int = int(binary.Size(masterb))//tamaño del MBR
		data := LeerBytes(archivo, sizemb)//Lee la cantidad de <size> bytes del archivo
		*/

		//ver si existe espacio para otra particion
		for i:=0 ; i < 4 ; i++ {
			if ( masterb.Mbr_partition[i].Part_start ==-1 || (masterb.Mbr_partition[i].Part_status =='1' && masterb.Mbr_partition[i].Part_size >= size_completo) ) {
				fparticion=true
				numparticion=i
				break
			}
		}
		
		if (fparticion == true){

			var espacio_utilizado int64=0

			for i:=0 ; i < 4 ; i++ {//revisa el espacio libre del disco
				if(masterb.Mbr_partition[i].Part_status != '1' ){
					espacio_utilizado+=masterb.Mbr_partition[i].Part_size//espacio de cada particion ocupada
				}
			}

			var tam int64
			tam=masterb.Mbr_tamano-int64(binary.Size(masterb))


			fmt.Println(" Espacio disponible: ",(tam-espacio_utilizado)," bytes")
			fmt.Println(" Espacio reqerido: ",size_completo," bytes")

			//ver si hay espacio para crear la particion
			if ( (tam - espacio_utilizado) >= size_completo) {
				
				if( !(ParticionExiste(path,name)) ){

					var namen [16]byte
					copy(namen[:],name)
					 
					masterb.Mbr_partition[numparticion].Part_type='p'
					masterb.Mbr_partition[numparticion].Part_fit=fit_aux
				
					if ( numparticion == 0){
						masterb.Mbr_partition[numparticion].Part_start=int64(binary.Size(masterb))
					}else{
						masterb.Mbr_partition[numparticion].Part_start=masterb.Mbr_partition[numparticion-1].Part_start + masterb.Mbr_partition[numparticion-1].Part_size
					}

					masterb.Mbr_partition[numparticion].Part_size = size_completo
                    masterb.Mbr_partition[numparticion].Part_status = '0';
					masterb.Mbr_partition[numparticion].Part_name=namen
					//archivo.Close()


					Escribir_MBR(path,masterb)

					fmt.Println(" Mensaje: Particion Primaria fue creada con exito.")

				}else{
					fmt.Println(" Mensaje: Ya hay una particion con ese nombre.")
				}

			}else{
				fmt.Println(" Mensaje: La particion a crear es mas grande que el espacion libre en el disco. ")
			}

		}else{
			fmt.Println(" Mensaje: No se puede crear otra particion, ya existen 4.")	
		}

		//archivo.Close()
	}

	


}


func Realizar_Particion_Extendida(path string,name string,size int64,fit string,unit string){

	var fit_aux byte
	var size_completo int64=1024

	if ( fit == "bf" ) {
		fit_aux='b'
	}else if ( fit == "ff" ){
		fit_aux='f'
	}else if ( fit == "wf" || fit == ""){
		fit_aux='w'
	}

	if ( unit == "b" ){
		size_completo=size
	}else if ( unit == "k" || unit == "" ){
		size_completo=(size*1024)
	}else if ( unit == "m"){
		size_completo=(size*1024*1024)
	}

	err,masterb:=Leer_MBR(path)	

 	if err != nil {
		fmt.Println(" Error: El disco no existe.")
		//archivo.Close()
		panic(err)
	}else{
	
		var fparticion bool=false
		var fextendida bool=false
		var numparticion int=0
		
		
		//ver si ya existe una particion extendida
		for i:=0 ; i < 4 ; i++ {
			if ( masterb.Mbr_partition[i].Part_type=='e') {
				fextendida=true
				break
			}
		}


		if ( fextendida == false ){

			for i:=0 ; i < 4 ; i++ {//ver si hay espacio en el disco para otra particion
				if ( masterb.Mbr_partition[i].Part_start ==-1 || (masterb.Mbr_partition[i].Part_status =='1' && masterb.Mbr_partition[i].Part_size >= size_completo) ) {
					fparticion=true
					numparticion=i
					break
				}
			}
			
			if (fparticion == true){
	
				var espacio_utilizado int64=0
	
				for i:=0 ; i < 4 ; i++ {//revisa el espacio libre del disco
					if(masterb.Mbr_partition[i].Part_status != '1' ){//busca las particiones que no estan activos(ya estan creadas)
						espacio_utilizado+=masterb.Mbr_partition[i].Part_size
					}
				}

				var tam int64
				tam=masterb.Mbr_tamano-int64(binary.Size(masterb))
	
				fmt.Println(" Espacio disponible: ",(tam-espacio_utilizado)," bytes")
				fmt.Println(" Espacio reqerido: ",size_completo," bytes")
	
				//ver si hay espacio para crear la particion
				if ( (tam- espacio_utilizado) >= size_completo) {
					
					if( !(ParticionExiste(path,name)) ){
	
						var namen [16]byte
						copy(namen[:],name)

						var na [16]byte
						copy(na[:],"")
						 
						masterb.Mbr_partition[numparticion].Part_type='e'
						masterb.Mbr_partition[numparticion].Part_fit=fit_aux
					
						if ( numparticion == 0){
							masterb.Mbr_partition[numparticion].Part_start=int64(binary.Size(masterb))
						}else{
							masterb.Mbr_partition[numparticion].Part_start=masterb.Mbr_partition[numparticion-1].Part_start + masterb.Mbr_partition[numparticion-1].Part_size
						}
	
						masterb.Mbr_partition[numparticion].Part_size = size_completo
						masterb.Mbr_partition[numparticion].Part_status = '0';
						masterb.Mbr_partition[numparticion].Part_name=namen
					//	archivo.Close()
	
						Escribir_MBR(path,masterb)

						Ebr:=estructuras.EbrStr{}
						Ebr.Part_fit=fit_aux
						Ebr.Part_status='1'
						Ebr.Part_start=masterb.Mbr_partition[numparticion].Part_start
						Ebr.Part_size=0
						Ebr.Part_next=-1
						Ebr.Part_name=na

						Escribir_EBR(path,Ebr)


						fmt.Println(" Mensaje: Particion Extendida fue creada con exito.")
	
					}else{
						fmt.Println(" Mensaje: Ya hay una particion con ese nombre.")
					}
	
				}else{
					fmt.Println(" Mensaje: La particion a crear es mas grande que el espacion libre en el disco. ")
				}
	
			}else{
				fmt.Println(" Mensaje: No se puede crear otra particion, ya existen 4.")	
			}



		}else {
			fmt.Println(" Error: Ya existe una particion extendida en el disco. ")
		}

		//archivo.Close()

	}

}


func Realizar_Particion_Logica(path string,name string,size int64,fit string,unit string){


	var fit_aux byte
	var size_completo int64=1024

	if ( fit == "bf" ) {
		fit_aux='b'
	}else if ( fit == "ff" ){
		fit_aux='f'
	}else if ( fit == "wf" || fit == ""){
		fit_aux='w'
	}

	if ( unit == "b" ){
		size_completo=size
	}else if ( unit == "k" || unit == "" ){
		size_completo=(size*1024)
	}else if ( unit == "m"){
		size_completo=(size*1024*1024)
	}

	err,masterb:=Leer_MBR(path)

	if err !=nil{
		fmt.Println(" Error: El disco no existe.")
		//archivo.Close()
		//panic(err)

	}else {
		var numextendida int=-1
		//ver si ya existe una particion extendida
		for i:=0 ; i < 4 ; i++ {
			if ( masterb.Mbr_partition[i].Part_type =='e') {
				numextendida=i
				break
			}
		}

		if ( !(ParticionExiste(path,name)) ){

			if( numextendida != -1 ){
				
			//	archivo.Close()

				err2,ebr:=Leer_EBR(path,masterb.Mbr_partition[numextendida].Part_start)

				if ( err2 != nil){
					fmt.Println(" Error: No se pudo leer el disco correctamente.")
				//	archivo.Close()
					//panic(err)

				}else{
					//Ebr:=estructuras.EbrStr{ }

					if ( ebr.Part_size == 0){//si es la primera logica a crear
						fmt.Println(" PRIMERA LOGICA ")

						//var tam int64
					    //tam=masterb.Mbr_partition[numextendida].Part_size-int64(binary.Size(ebr))
						
						if (masterb.Mbr_partition[numextendida].Part_size < size_completo){
							fmt.Println(" Error: La particion logica a crear excede al espacio disponible de la particion extendida. ")
						}else{

							var namen [16]byte
							copy(namen[:],name)

							//se actualiza el ebr que ya estaba creado cuando se creo la extendida
							ebr.Part_status='0'
							ebr.Part_fit=fit_aux
							ebr.Part_start=masterb.Mbr_partition[numextendida].Part_start
							ebr.Part_size=size_completo
							ebr.Part_next=-1
							ebr.Part_name=namen

							Escribir_EBR(path,ebr)

							fmt.Println(" Mensaje: Particion logica creada con exito. ")
						}

					}else{//si ya hay particiones logicas dentro de la extendida
						fmt.Println(" SIGUIENTES LOGICAS ")
						ebr2:=estructuras.EbrStr{ }//es para crear otro ebr y unirlo a los demas 

						var aux estructuras.EbrStr
						aux=ebr

						for (aux.Part_next != -1) {
							err,aux=Leer_EBR(path,aux.Part_next)
						}
						
						var espacionNecesario=aux.Part_start + aux.Part_size + size_completo//se toma la posicion en donde termina el ultimo mbr y se le suma el tamaño del nuevo mbr
						
						if ( espacionNecesario <= ( masterb.Mbr_partition[numextendida].Part_start + masterb.Mbr_partition[numextendida].Part_size ) ){

							aux.Part_next=aux.Part_start + aux.Part_size//se actualizo el atributo next del ebr  
							ebr=aux
							Escribir_EBR(path,ebr)//se escribio el ebr actualizado

							var namen2 [16]byte
							copy(namen2[:],name)

							ebr2.Part_status='0'
							ebr2.Part_fit=fit_aux
							ebr2.Part_start=aux.Part_start + aux.Part_size
							ebr2.Part_size=size_completo
							ebr2.Part_next=-1
							ebr2.Part_name=namen2

							Escribir_EBR(path,ebr2)

							fmt.Println(" Mensaje: Particion logica creada con exito. ")



						}else{
							fmt.Println("Error: La particion logica a crear excede al espacio disponible de la particion extendida. ")
						}
						


					}


				}

			}else{
				fmt.Println(" Error: Se necesita una particion extendida para poder guardar la logica. ")				
			}

		}else{
			fmt.Println(" Error: Ya existe una particion con ese mombre. ")
		}




		
	}


	//archivo.Close()

}

func Agregar_Quitar_Particiones(path string,name string,addn string,unit string){

	
	var size_completo int64=0
	var tipo string=""
	//var auxUnit byte;
	//var add int64=0*/
	
	add, _:= strconv.ParseInt(addn, 10, 64) 

	if ( add > 0 ){
		tipo="add"
	}

	if ( tipo != "add" ) {
		add=add*(-1)
	}


	if ( unit == "b" ){
		size_completo=add
	}else if ( unit == "k" || unit == "" ){
		size_completo=(add*1024)
	}else if ( unit == "m"){
		size_completo=(add*1024*1024)
	}


	err,masterb:=Leer_MBR(path)

	if ( err != nil ) {
		fmt.Println(" Error: El disco no existe (add). ")
	}else{

		mount:=Lista.BuscarNodo(path,name)

		if ( !mount ){
			var index int=-1
			var index_e int=-1
			var fextendida bool=false
			var namen [16]byte
			copy(namen[:],name)


			for i:=0 ; i < 4 ; i++ {

				if ( masterb.Mbr_partition[i].Part_name == namen){
					index=i
					if ( masterb.Mbr_partition[i].Part_type == 'e'){
						fextendida=true
					}
					break
				}else if ( masterb.Mbr_partition[i].Part_type == 'e'){
					index_e=i
				}
			}

			if ( index != -1) {
				if ( !fextendida){//particiones primarias
					if ( tipo == "add"){
						fmt.Println( " en primarias")
						if ( index != 3) {
							fmt.Println( " en != 3")
							p1:=masterb.Mbr_partition[index].Part_start + masterb.Mbr_partition[index].Part_size
							p2:=masterb.Mbr_partition[index+1].Part_start

							if ( (p2 - p1) != 0 ){
								fragmentacion:=p2-p1
								if ( fragmentacion >= size_completo ){
									masterb.Mbr_partition[index].Part_size=masterb.Mbr_partition[index].Part_size + size_completo
									Escribir_MBR(path,masterb)
									fmt.Println(" Mensaje: Se acaba de agregar espacio a la particion correctamente. ")
								}else{
									fmt.Println(" Error: No se puede agregar espacio a la particion , porque no hay suficiente espacio disponible en el disco.")
								}
							}else{
								fmt.Println( " en este else")
								if (masterb.Mbr_partition[index + 1].Part_status == '0'){
                                    if (masterb.Mbr_partition[index + 1].Part_size >= size_completo){
                                        masterb.Mbr_partition[index].Part_size = masterb.Mbr_partition[index].Part_size + size_completo
                                        masterb.Mbr_partition[index + 1].Part_size = (masterb.Mbr_partition[index + 1].Part_size - size_completo)
                                        masterb.Mbr_partition[index + 1].Part_start = masterb.Mbr_partition[index + 1].Part_start + size_completo
                                      
									  	Escribir_MBR(path,masterb)
                                        fmt.Println(" Mensaje: Se acaba de agregar espacio a la particion correctamente. ")
                                    }else{
										fmt.Println(" Error: No se puede agregar espacio a la particion , porque no hay suficiente espacio disponible en el disco.")
                                    }
                                }

							}

						}else{
							fmt.Println( " no es  != 3")
							p := masterb.Mbr_partition[index].Part_start + masterb.Mbr_partition[index].Part_size;
                            total := masterb.Mbr_tamano + int64(binary.Size(masterb))
                            if ((total-p) != 0){
                                fragmentacion := total - p;
                                if (fragmentacion >= size_completo){
                                    masterb.Mbr_partition[index].Part_size = masterb.Mbr_partition[index].Part_size + size_completo
									
									Escribir_MBR(path,masterb)
									fmt.Println(" Mensaje: Se acaba de agregar espacio a la particion correctamente. ")
									}else{
										fmt.Println(" Error: No se puede agregar espacio a la particion , porque no hay suficiente espacio disponible en el disco.")
                                }
                            }else{
                                fmt.Println(" Error: No se puede agregar espacio a la particion , porque no hay suficiente espacio disponible en el disco.")
                            }
							
						}

					}else{//quitar espacio en el disco

						if (size_completo >= masterb.Mbr_partition[index].Part_size){
                            fmt.Println(" ERROR: no es posible quitarle esta cantidad de espacio a la particion porque la borraria. ")
                        }else{
                            masterb.Mbr_partition[index].Part_size = masterb.Mbr_partition[index].Part_size - size_completo
							
							Escribir_MBR(path,masterb)
							fmt.Println(" Mensaje: Se quito espacio a la particion de manera exitosa. ")
							
                        }

					}

				}else{//extendidas

					if (tipo == "add"){//Agregar
                        //Verificar que exista espacio libre a la derecha
                        if (index!=3){
                             p1 := masterb.Mbr_partition[index].Part_start + masterb.Mbr_partition[index].Part_size
                             p2 := masterb.Mbr_partition[index+1].Part_start
                            if ( (p2 - p1) != 0 ){//Hay fragmentacion
                                 fragmentacion := p2-p1
                                if (fragmentacion >= size_completo){
									masterb.Mbr_partition[index].Part_size = masterb.Mbr_partition[index].Part_size + size_completo
									
									Escribir_MBR(path,masterb)
									fmt.Println(" Mensaje: Se acaba de agregar espacio a la particion correctamente. ")
                                   
                                }else{
									fmt.Println(" Error: No se puede agregar espacio a la particion , porque no hay suficiente espacio disponible en el disco.")
                                }
                            }else{
                                if ( masterb.Mbr_partition[index + 1].Part_status == '1' ){
                                    if ( masterb.Mbr_partition[index + 1].Part_size >= size_completo ){
                                        masterb.Mbr_partition[index].Part_size = masterb.Mbr_partition[index].Part_size + size_completo
                                        masterb.Mbr_partition[index + 1].Part_size = (masterb.Mbr_partition[index + 1].Part_size - size_completo)
                                        masterb.Mbr_partition[index + 1].Part_start= masterb.Mbr_partition[index + 1].Part_start + size_completo
									   
										Escribir_MBR(path,masterb)
										fmt.Println(" Mensaje: Se acaba de agregar espacio a la particion correctamente. ")
									
                                    }else{
                                        fmt.Println(" Error: No se puede agregar espacio a la particion , porque no hay suficiente espacio disponible en el disco.")
                                    }
                                }
                            }
                        }else{
                             p := masterb.Mbr_partition[index].Part_start + masterb.Mbr_partition[index].Part_size;
                             total := masterb.Mbr_tamano + int64(binary.Size(masterb))
                            if ((total-p) != 0){//Hay fragmentacion
                                 fragmentacion := total - p;
                                if (fragmentacion >= size_completo){
                                    masterb.Mbr_partition[index].Part_size = masterb.Mbr_partition[index].Part_size + size_completo
									
									Escribir_MBR(path,masterb)
									fmt.Println(" Mensaje: Se acaba de agregar espacio a la particion correctamente. ")

                                }else{
                                    fmt.Println(" Error: No se puede agregar espacio a la particion , porque no hay suficiente espacio disponible en el disco.")
                                }
                            }else{
                                fmt.Println(" Error: No se puede agregar espacio a la particion , porque no hay suficiente espacio disponible en el disco.")
                            }
                        }
                    }else{//Quitar espacio
                        //Que no borre la particion
                        if (size_completo >= masterb.Mbr_partition[index_e].Part_size){
							fmt.Println(" ERROR: no es posible quitarle esta cantidad de espacio a la particion porque la borraria. ")
                           
                        }else{

							err,ebr:=Leer_EBR(path,masterb.Mbr_partition[index_e].Part_start)

							if ( err != nil){
								fmt.Println(" Error: No se puedo leer el EBR (add)")
							
							}else{

								var aux estructuras.EbrStr
								aux=ebr

								for (aux.Part_next != -1) {
									err,aux=Leer_EBR(path,aux.Part_next)
								}

								ultimaLogica := aux.Part_start + aux.Part_size
								aux2 := (masterb.Mbr_partition[index_e].Part_start + masterb.Mbr_partition[index_e].Part_size) - size_completo

								if (aux2 > ultimaLogica){//No toca ninguna logica

									masterb.Mbr_partition[index_e].Part_size = masterb.Mbr_partition[index_e].Part_size - size_completo

									Escribir_MBR(path,masterb)
									fmt.Println(" Mensaje: Se quito espacio a la particion de manera exitosa. ")

								}else{
									fmt.Println(" ERROR: si quita ese espacio danaria una logica.")
								}

							}


                        }
					}
					

				}

			}else{//particiones logicas


				if ( index_e != -1){
                     logica := buscarParticionL(path, name);
                    if (logica != -1){
                        if (tipo == "add"){

							err,_:=Leer_EBR(path,logica)
							if ( err != nil ){
								fmt.Println(" Error: NO se pudo leer el EBR. ")
							}
                            

                        }else{//Quitar
                            //Verificar que no la elimine
                            err,ebr:=Leer_EBR(path,logica)
							if ( err != nil ){
								fmt.Println(" Error: NO se pudo leer el EBR. ")
							}else{

                            if (size_completo >= ebr.Part_size){
                                fmt.Println(" ERROR: si quita ese espacio eliminaria la logica.")
                            }else{
								ebr.Part_size = ebr.Part_size - size_completo
								
								Escribir_EBR(path,ebr)
								fmt.Println(" Mensaje: Se quito espacio a la particion de manera exitosa. ")
                                
                            }
							
							}
						}
					
						
                    }else{
                        fmt.Println(" ERROR: no se encuentra la particion.")
                    }
                }else{
                    fmt.Println(" ERROR: no se encuentra la particion a redimensionar. ")
                }

			}

		}else{
			fmt.Println(" Error: Primero se tiene que desmontar la particion  para poder redimensionar.")
		}


	}
	


}

func Eliminar_Particion(path string,name string,delete string){


}

func ParticionExiste(path string, name string) bool{

	var extendida int =-1

	err,masterb:=Leer_MBR(path)	

 	if err != nil {
		fmt.Println(" Error: No se pudo leer el disco.")
		panic(err)
	}else{
		var namee [16]byte
		copy(namee[:],name)
		
		for i:=0 ; i < 4 ; i++ {
		
			if (masterb.Mbr_partition[i].Part_name == namee ){
				//archivo.Close()
				return true

			}else if (masterb.Mbr_partition[i].Part_type =='e'){
				extendida=i
				//archivo.Close()
			}
		}

		if (extendida != -1){

			err2,ebr:=Leer_EBR(path,masterb.Mbr_partition[extendida].Part_start)
		
			if err2 != nil {
				fmt.Println(" Error: No se pudo leer el EBR.")
				//archivo2.Close()
				//panic(err)
			}else{

				var namee2 [16]byte
				copy(namee2[:],name)

				var aux estructuras.EbrStr
				aux=ebr

				for ( aux.Part_next != -1 ) {

					if(aux.Part_name == namee2){
						//archivo2.Close()
						return true
					}
					err2,aux=Leer_EBR(path,aux.Part_next)
				}

				if(aux.Part_name==namee2){
					return true
				}
				
			}

		}
	}
	//archivo.Close()
	return false
}


func buscarParticionL(path string , name string) int64{

	err,masterb:=Leer_MBR(path)

	if ( err != nil){
		fmt.Println(" Error: NO se pudo leer el disco .")
	}else{

		var extendida int=-1

		for i:=0 ; i < 4 ; i++ {
			
			if ( masterb.Mbr_partition[i].Part_type == 'e' ){
				extendida=i
				break
			}
		}

		if ( extendida != -1){

			err2,ebr:=Leer_EBR(path,masterb.Mbr_partition[extendida].Part_start)

			if ( err2 != nil ){
				fmt.Println(" Error: No se pudo el EBR")
			}else{

				var namee2 [16]byte
				copy(namee2[:],name)

				var aux estructuras.EbrStr
				aux=ebr

				for ( aux.Part_next != -1 ) {

					if(aux.Part_name == namee2){
						//archivo2.Close()
						return (aux.Part_start)
					}
					err2,aux=Leer_EBR(path,aux.Part_next)
				}

				if ( aux.Part_next == -1){
					return (aux.Part_start)
				}


			}


		}

	}
	
	return -1
}



/*
func Buscar_Nombre_Logica(path string,name string ,inicioebr int64) int{




}

func Get_next_ebr(path string,inicio int64) (ebr estructuras.EbrStr){

	ebrtemp:=estructuras.EbrStr{ }
	ebrtemp.Part_size=-1

	err2,archivo2,ebr:=Leer_EBR(path,inicio)

	if err2 !=nil{

		archivo2.Close()
		return ebr
	}
	archivo2.Close()
	return ebr
}*/