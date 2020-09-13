package comandos

import(
	"fmt"
	//"../funciones"
	"../estructuras"
)

	var Lista ListaSimple

//lista:=ListaSimple{}

func MOUNT(path string, name string)  {
	fmt.Println("Dentro de la funcion mount")

	if ( path=="" && name=="" ){
		fmt.Println("----------Lista de particiones montadas----------")
		Lista.Mostrar()
	}else{

	err,masterb:=Leer_MBR(path)
	
	 indicep:=buscar_particion_e(path,name)

	 if ( indicep != -1) {

		//err,masterb:=Leer_MBR(path)

		if err != nil{
			fmt.Println(" Error: No se pudo leer el disco, en el comando mount.")
			
		}else{
			masterb.Mbr_partition[indicep].Part_status='2'//el estado 2 es porque se va a montar la particion
		
			Escribir_MBR(path,masterb)

			letra:=Lista.BuscarLetra(path,name)

			if ( letra == -1 ){
				fmt.Println(" Error: La particion ya esta montada. ")

			}else{
				num:=Lista.BuscarNumero(path,name)

				ltr:=rune(letra)

				Lista.Insertar(& NodoParticiones{path:path,nombre:name,letra:byte(ltr),num:num})
				
				fmt.Println(" Mensaje: La particion se ha montano con exito. ")
			}

			
	
		}//else cuando se lee el mbr


	}else{//por si es una logica

		indicep2:=buscar_particion_l(path,name)

		if ( indicep2 != -1){

			err2,ebr:=Leer_EBR(path,indicep2)

			if ( err2 != nil ){
				fmt.Println(" Error: No se encontro el disco. ")
			}else{

				ebr.Part_status='2'
				Escribir_EBR(path,ebr)

				letra2:=Lista.BuscarLetra(path,name)

				if ( letra2 == -1){
					fmt.Println(" Mensaje: La particion ya se encuentra montada. ")
				}else{

					ltr2:=rune(letra2)

					num2:=Lista.BuscarNumero(path,name)
					Lista.Insertar(& NodoParticiones{path:path,nombre:name,letra:byte(ltr2),num:num2})

					fmt.Println(" Mensaje: La particion se ha montano con exito. ")
				}


			}

		}else {
			fmt.Println(" Error: La particion no se encuentra creada en el disco. ") 
		}
	 }



	}


}


func  buscar_particion_e(path string,name string) int {

	err,masterb:=Leer_MBR(path)

	var namee [16]byte
    copy(namee[:],name)

	if err != nil{
		fmt.Println(" Error: No se pudo leer el disco, en el comando mount.")
		return -1
	}else{

		for  i:=0 ; i < 4 ; i++ {
			
			if ( masterb.Mbr_partition[i].Part_status !='1') {//debemos encontrar una particion ya con datos
				
				if ( masterb.Mbr_partition[i].Part_name == namee){//verificamos si ya existe una particion con ese nombre
					return i
				}
			}
		}

	}

	return -1
	
}


func buscar_particion_l(path string, name string) int64{

	err,masterb:=Leer_MBR(path)

	if err != nil{
		fmt.Println(" Error: No se pudo leer el disco, en el comando mount.")
		return -1
	}else{

		var extendida int=-1

		for i:=0 ; i < 4 ; i++ {

			if ( masterb.Mbr_partition[i].Part_type =='e') {
				extendida=i
				break

			}
		}

		if (extendida != -1){

			err2,ebr:=Leer_EBR(path,masterb.Mbr_partition[extendida].Part_start)
		
			if err2 != nil {
				fmt.Println(" Error: No se pudo leer el EBR.")
				//archivo2.Close()
				panic(err)
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

				if( aux.Part_name == namee2 ){
					return (aux.Part_start)
				}
				
			}



		}


		
	}
	return -1
}