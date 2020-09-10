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

	
	
	 indicep:=buscar_particion_e(path,name)

	 if ( indicep != -1) {

		err,masterb:=Leer_MBR(path)

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

				Lista.Insertar(& NodoParticiones{path:path,nombre:name,letra:'d',num:num})
				
				fmt.Println(" Mensaje: La particion se ha montano con exito. ")
			}

			
	
		}//else cuando se lee el mbr


	}else{//por si es una logica

		indicep2:=buscar_particion_l(path,name)

		if (){

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
						return 1
					}
					err2,aux=Leer_EBR(path,aux.Part_next)
				}

				if(aux.Part_name==namee2){
					return 1
				}
				
			}



		}


		
	}
	return -1
}