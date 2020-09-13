package comandos

import(
	"fmt"
	"strings"
	"strconv"
	//"bytes"
	//"io/ioutil"
	//"log"
	"os"
	"os/exec"
	"encoding/binary"
	"../estructuras"
)

func REP(path string,ngrafica string,nidgrap string)  {
	fmt.Println("Dentro de la funcion de reportes")
	fmt.Println(" path:"+path+" \n Ngrafica:"+ngrafica+" \n ID:"+nidgrap)

	if ( path=="" || ngrafica=="" || nidgrap=="") {
		fmt.Println(" Revise el comando, no se puede ejecutar. ")
	}else{

		ubicacionP:=Lista.GetDireccion(nidgrap)
		extencion:=extencion_archivo(path)

		if ( ubicacionP != "nil"){

			var res string
			carpeta:=strings.Split(path, "/") 
			for i:=0 ;i<len(carpeta);i++{
				if (i<len(carpeta)-1){
				res+=carpeta[i]+"/"
				}
			}
			Crear_Carpeta(res)

			if ( ngrafica == "disk"){
				grafica_disco(ubicacionP,path,extencion)

			}else if ( ngrafica == "mbr"){
				grafica_mbr(ubicacionP,path,extencion)
			}else{
				fmt.Println(" Error: Nombre del reporte incorrecto. ")
			}

		}else{
			fmt.Println("Error: NO se encontro la particion.")
		}

	}

}

func crearArchivo(path string) {
	//Verifica que el archivo existe
	//var _, err = os.Stat(path)
	//Crea el archivo si no existe
	//if os.IsNotExist(err) {
	  var file, err = os.Create(path)
	  if existeError(err) {
		return
	  }
	  defer file.Close()
	//}
	fmt.Println("File Created Successfully", path)
  }

  func escribeArchivo(path string,contenido string) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0777)
	if existeError(err) {
	  return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	_, err = file.WriteString(contenido)
	if existeError(err) {
	  return
	}

	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
	  return
	}
	fmt.Println("Archivo actualizado existosamente.")
  }

  func existeError(err error) bool {
	if err != nil {
	  fmt.Println(err.Error())
	}
	return (err != nil)
  }


func extencion_archivo(path string) string{
	ext:=strings.Split(path, ".") 
	return ext[1]
}

func direccion_carpeta(path string) string{

	return "s"
}


func GoString(c []byte) string {
    n := -1
    for i, b := range c {
        if b == 0 {
            break
        }
        n = i
    }
    return string(c[:n+1])
}




func grafica_disco(ubicacion string,path string,extencion string){

	err,masterb:=Leer_MBR(ubicacion)	

	var cadena string=""

	if ( err != nil){
		fmt.Println(" Error: No se puede crear el reporte, no se encontro el disco. ")
	}else{

		total:=masterb.Mbr_tamano
		var usado int64

		//ru:=strings.Split(path, ".") 
		//crearArchivo(ru[0]+".dot")
	//	crearArchivo("disco.dot")

		cadena+="digraph G{\n\n"
		cadena+=" subgraph cluster_1{\n"
		cadena+="  label=\" DISCO \" \n fontsize=25 \n style=filled \n fillcolor=\"olivedrab3\"  " 
		cadena+="  tbl [\n    shape=box\n    label=<\n"
		cadena+="     <table border=\"0\" cellborder=\"2\" width=\"600\" height=\"200\" color=\"blue\">\n"
		cadena+="     <tr>\n"
		cadena+= "     <td height=\"200\" width=\"100\" bgcolor=\"orange\"> MBR </td>\n"

		for i:=0 ; i < 4 ; i++ {
			parcial:=masterb.Mbr_partition[i].Part_size

			if ( masterb.Mbr_partition[i].Part_start != -1){
				porcentaje_real := (parcial*100)/total
                porcentaje_aux := (porcentaje_real*500)/100
				usado += porcentaje_real
				
				if (masterb.Mbr_partition[i].Part_status != '1') {
					if (masterb.Mbr_partition[i].Part_type == 'p'){
						
						nam:=GoString(masterb.Mbr_partition[i].Part_name[:])
						cadena+="  <td height=\"200\" bgcolor=\"orange\" width="+"\""+strconv.FormatInt(porcentaje_aux,10)+"\""+">PRIMARIA <br/>"+nam+"<br/>"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>\n"

						if (i != 3 ) {
							fmt.Println("UNO UNO UNO")
							part1:=masterb.Mbr_partition[i].Part_start + masterb.Mbr_partition[i].Part_size
							part2:=masterb.Mbr_partition[i+1].Part_start
							if (masterb.Mbr_partition[i+1].Part_start != -1){
								fmt.Println("LIBRE1:",(part2 - part1))
								if ( (part2-part1)!=0 ){//existe la fragmentacion
									fragmentacion2:=part2-part1
									porcentaje_real:=(fragmentacion2*100)/total
									porcentaje_aux:=(porcentaje_real*500)/100

									cadena+="<td height=\"200\" bgcolor=\"orange\" width="+"\""+strconv.FormatInt(porcentaje_aux,10)+"\""+"> LIBRE <br/>"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>\n"
								}
							}
						}else{
							part11:=masterb.Mbr_partition[i].Part_start + masterb.Mbr_partition[i].Part_size
							tam_mbr:=total+int64(binary.Size(masterb))
								fmt.Println("LIBRE2:",(tam_mbr - part11))
							if ( (tam_mbr - part11) != 0){
								libre:=(tam_mbr-part11)+int64(binary.Size(masterb))
								porcentaje_reall:=(libre*100)/total
								porcentaje_auxx:=(porcentaje_reall*500)/100

								cadena+="<td height=\"200\" bgcolor=\"orange\" width="+"\""+strconv.FormatInt(porcentaje_auxx,10)+"\""+"> LIBRE <br/>"+strconv.FormatInt(porcentaje_reall,10)+"%"+"</td>\n"

							}
						}

					}else{//es una extendida
						nam:=GoString(masterb.Mbr_partition[i].Part_name[:])
						cadena+=" \n\n\n <td  height=\"200\" width=\""+strconv.FormatInt(porcentaje_real,10)+"\">\n     <table border=\"0\"  height=\"200\" WIDTH=\""+strconv.FormatInt(porcentaje_real,10)+"\" cellborder=\"1\">\n"
						cadena+="  <tr>  <td height=\"60\" bgcolor=\"orange\" colspan=\"15\"> EXTENDIDA <br/>"+nam+" </td>  </tr>\n     <tr>\n"

						err2,ebr:=Leer_EBR(ubicacion,masterb.Mbr_partition[i].Part_start)

						if ( err2 != nil ){
							fmt.Println(" Error: No se pudo leer el EBR al intentar graficar. ")
						}else{

							if ( ebr.Part_size != 0){//por si hay logicas

								var aux estructuras.EbrStr
								aux=ebr
						//---------------------------------------------------------------------	
								for ( aux.Part_next != -1 ) {

									parcial=aux.Part_size
									porcentaje_real=(parcial*100)/total

									if (porcentaje_real != 0){
										if ( aux.Part_status != '1'){
											nam:=GoString(aux.Part_name[:])
											cadena+=" <td height=\"140\" bgcolor=\"orange\"> EBR </td>\n"
											cadena+=" <td height=\"140\" bgcolor=\"orange\">LOGICA <br/>"+nam+" <br/>"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>\n"
										
										}else{//espacion en disco no asignado
											cadena+=" <td height=\"150\" bgcolor=\"orange\">LIBRE 1 <br/>"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>\n"
										}

										/*if ( aux.Part_next == -1){
											parcial=(masterb.Mbr_partition[i].Part_start + masterb.Mbr_partition[i].Part_size) - ( aux.Part_start + aux.Part_size)
											porcentaje_real=(parcial*100)/total
											
											if ( porcentaje_real != 0) {
												cadena+="<td height=\"150\">LIBRE 2<br/>"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>\n"
											}
											break
										}*/
										/*else{
											err2,aux=Leer_EBR(path,aux.Part_next)
										}*/

									}

									err2,aux=Leer_EBR(ubicacion,aux.Part_next)

								}//final del for para recorre las logicas

								/*if ( aux.Part_next == -1){
									parcial=(masterb.Mbr_partition[i].Part_start + masterb.Mbr_partition[i].Part_size) - ( aux.Part_start + aux.Part_size)
									porcentaje_real=(parcial*100)/total
									
									if ( porcentaje_real != 0) {
										cadena+="<td height=\"150\">LIBRE 2<br/>"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>\n"
										//cadena+="\n </tr> \n </table> \n </td> \n\n"
									}
									//break
								}*/
								

								parcial=aux.Part_size
								porcentaje_real=(parcial*100)/total

									if (porcentaje_real != 0){
										if ( aux.Part_status != '1'){
											nam:=GoString(aux.Part_name[:])
											cadena+=" <td height=\"140\" bgcolor=\"orange\"> EBR </td>\n"
											cadena+=" <td height=\"140\" bgcolor=\"orange\">LOGICA <br/>"+nam+" <br/>"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>\n"
										
										}else{//espacion en disco no asignado
											cadena+=" <td height=\"150\" bgcolor=\"orange\">LIBRE 1 <br/>"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>\n"
										}

										if ( aux.Part_next == -1){
											parcial=(masterb.Mbr_partition[i].Part_start + masterb.Mbr_partition[i].Part_size) - ( aux.Part_start + aux.Part_size)
											porcentaje_real=(parcial*100)/total
											
											if ( porcentaje_real != 0) {
												cadena+="<td height=\"150\" bgcolor=\"orange\">LIBRE 2<br/>"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>\n"
											}
											//break
										}

									}
						//---------------------------------------------------------------------			


							}else{
								cadena+=" <td height=\"140\" bgcolor=\"orange\"> Ocupado"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>"
							}

							cadena+="</tr>\n </table>\n </td>\n"

							//ver que no haya fragmentacion
							if ( i != 3){
								p1:=masterb.Mbr_partition[i].Part_start + masterb.Mbr_partition[i].Part_size
								p2:=masterb.Mbr_partition[i+1].Part_start

								if (masterb.Mbr_partition[i+1].Part_start != -1){
									if ( (p2-p1)!=0){//existe fragmentacion
										fragmentacion:=p2-p1
										porcentaje_real:=(fragmentacion*100)/total
										porcentaje_aux:=(porcentaje_real*500)/100

										cadena+="   <td height=\"200\" bgcolor=\"orange\" width=\""+strconv.FormatInt(porcentaje_aux,10)+"\">LIBRE<br/>"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>\n"
									}
								}
							}else{
								p1:=masterb.Mbr_partition[i].Part_start + masterb.Mbr_partition[i].Part_size
								mbr_tam:=total + int64(binary.Size(masterb))

								if ( (mbr_tam - p1) != 0){//lo que esta libre
									libre:=(mbr_tam - p1) + int64(binary.Size(masterb))
									porcentaje_real:=(libre*100)/total
									porcentaje_aux:=(porcentaje_real*500)/100

									cadena+="   <td height=\"200\" bgcolor=\"orange\" width=\""+strconv.FormatInt(porcentaje_aux,10)+"\">LIBRE<br/>"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>\n"
								}

							}

						}


									


					}//final del tercer if 

				}else{//es el espacio que no estea asignado
					cadena+="     <td height=\"200\" bgcolor=\"orange\" width="+"\""+strconv.FormatInt(porcentaje_aux,10)+"\""+">LIBRE <br/>"+strconv.FormatInt(porcentaje_real,10)+"%"+"</td>\n"
				
				}//final del segundo if

			}//final del primer if

		}//final del for que recorre las particiones

		cadena+="</tr> \n  </table>  \n>]; }\n\n}"


		/*crearArchivo("/home/graficas/disco.dot")
		escribeArchivo("/home/graficas/disco.dot",cadena)

		_,err:= exec.Command("dot","/home/graficas/disco.dot","-o","/home/graficas/disco.png","-Tpng").Output()
		if (err!=nil){ fmt.Println(" error con el comando1")}

		_,err2:= exec.Command("xdg-open","/home/graficas/disco.png").Output()
		if (err2!=nil){ fmt.Println(" error con el comando2")}*/
		ru:=strings.Split(path, ".") 
		crearArchivo(ru[0]+".dot")
		escribeArchivo(ru[0]+".dot",cadena)

		_,err:= exec.Command("dot",ru[0]+".dot","-o",ru[0],"-T"+extencion).Output()
		if (err!=nil){ fmt.Println(" error con el comando1")}

		_,err2:= exec.Command("xdg-open",ru[0]).Output()
		if (err2!=nil){ fmt.Println(" error con el comando2")}

		fmt.Println("Mensaje: Se genero el repoorte del disco correctamente. ")

	}


	
}


func grafica_mbr(ubicacion string,path string,extencion string){

	var cadena string=""
	err,masterb:=Leer_MBR(ubicacion)	

	if err != nil {
		fmt.Println(" Error: No se pudo leer el disco para graficar el MBR. ")
	}else{

		tamano:=masterb.Mbr_tamano

		cadena+="digraph G{ \n"
		cadena+="subgraph cluster{\n label=\"MBR\" \n fontsize=25"
		cadena+="\n style=filled; \n fillcolor=olivedrab3"
		cadena+="\n tbl[shape=box,label=<\n"
		cadena+="<table color=\"black\"  border=\"0\" cellborder=\"1\" cellspacing=\"0\" width=\"300\"  height=\"200\" >\n"
		cadena+="<tr>  <td width=\"150\" bgcolor=\"darkorange\"> <b>NOMBRE</b> </td> <td width=\"150\" bgcolor=\"darkorange\"> <b>VALOR</b> </td>  </tr>\n"

		cadena+="<tr>  <td><b>mbr_tamaño</b></td><td><font color='navy'>"+strconv.FormatInt(tamano,10)+"</font></td>  </tr>\n"

		fecha:=GoString(masterb.Mbr_fecha_creacion[:])
		cadena+="<tr>  <td><b>mbr_fecha_creacion</b></td> <td><font color='navy'>"+fecha+"</font></td>  </tr>\n"

		cadena+="<tr>  <td><b>mbr_disk_signature</b></td> <td><font color='navy'>"+strconv.FormatInt(masterb.Mbr_disk_signature,10)+"</font></td>  </tr>\n"

		var i_extendida int=-1

		for i:=0 ; i < 4 ; i++ {
			
			if ( (masterb.Mbr_partition[i].Part_start != -1) && ( masterb.Mbr_partition[i].Part_status != '1')){

				if ( masterb.Mbr_partition[i].Part_type == 'e' ){
					i_extendida=i
				}
				var status string
				if ( masterb.Mbr_partition[i].Part_status == '0') {
					status="0"
				
				}else if ( masterb.Mbr_partition[i].Part_status == '2') {
					status="2"
				}

				cadena+="<tr>  <td><b>part_status_"+strconv.Itoa(i+1)+"</b></td> <td><font color='navy'>"+status+"</font></td>  </tr>\n"
				cadena+="<tr>  <td><b>part_type_"+strconv.Itoa(i+1)+"</b></td> <td><font color='navy'>"+string(masterb.Mbr_partition[i].Part_type)+"</font></td>  </tr>\n"
				cadena+="<tr>  <td><b>part_fit_"+strconv.Itoa(i+1)+"</b></td> <td><font color='navy'>"+string(masterb.Mbr_partition[i].Part_fit)+"</font></td>  </tr>\n"
				cadena+="<tr>  <td><b>part_start_"+strconv.Itoa(i+1)+"</b></td> <td><font color='navy'>"+strconv.FormatInt(masterb.Mbr_partition[i].Part_start,10)+"</font></td>  </tr>\n"
				cadena+="<tr>  <td><b>part_size_"+strconv.Itoa(i+1)+"</b></td> <td><font color='navy'>"+strconv.FormatInt(masterb.Mbr_partition[i].Part_size,10)+"</font></td>  </tr>\n"
				name:=GoString(masterb.Mbr_partition[i].Part_name[:])
				cadena+="<tr>  <td><b>part_name_"+strconv.Itoa(i+1)+"</b></td> <td><font color='navy'>"+name+"</font></td>  </tr>\n"

			}

		}//final del for que recorre el disco

				cadena+="</table>\n"
				cadena+=" >]; \n } \n"

		if ( i_extendida != -1){//para las extendidas

			var indice_ebr int=1

			err2,ebr:=Leer_EBR(ubicacion,masterb.Mbr_partition[i_extendida].Part_start)

			if ( err2 != nil){
				fmt.Println(" Error: No se pudo leer el EBR para mostrar el reporte del MBR.")
			}else{


				var aux estructuras.EbrStr
				aux=ebr
						
				var status2 string

				for ( aux.Part_next != -1 ) {

					if ( aux.Part_status != '1' ) {

						cadena+="subgraph cluster_"+strconv.Itoa(indice_ebr)+"{\n label=\"EBR_"+strconv.Itoa(indice_ebr)+"\"\n"
						cadena+="fontsize=20 \n style=filled \n fillcolor=olivedrab3 "
						cadena+="\ntbl_"+strconv.Itoa(indice_ebr)+"[shape=box, label=<\n "
						cadena+="<table color=\"black\" border=\"0\" cellborder=\"1\" cellspacing=\"0\"  width=\"300\" height=\"160\" >\n "
						cadena+="<tr>  <td width=\"150\" bgcolor=\"darkorange\"><b>NOMBRE</b></td> <td width=\"150\" bgcolor=\"darkorange\"><b>VALOR</b></td>  </tr>\n"
					
						if ( aux.Part_status == '0' ){
							status2="0"
						}else if ( aux.Part_status == '2') {
							status2="2"
						}

						cadena+="<tr>  <td><b>part_status</b></td> <td><font color='navy'>"+status2+"</font></td>  </tr>\n"
						cadena+="<tr>  <td><b>part_fit</b></td> <td><font color='navy'>"+string(aux.Part_fit)+"</font></td>  </tr>\n"
						cadena+="<tr>  <td><b>part_start</b></td> <td><font color='navy'>"+strconv.FormatInt(aux.Part_start,10)+"</font></td>  </tr>\n"
						cadena+="<tr>  <td><b>part_size</b></td> <td><font color='navy'>"+strconv.FormatInt(aux.Part_size,10)+"</font></td>  </tr>\n"
						cadena+="<tr>  <td><b>part_next</b></td> <td><font color='navy'>"+strconv.FormatInt(aux.Part_next,10)+"</font></td>  </tr>\n"
						name:=GoString(aux.Part_name[:])
						cadena+="<tr>  <td><b>part_name</b></td> <td><font color='navy'>"+name+"</font></td>  </tr>\n"

						cadena+="</table>\n"
						cadena+=">];\n}\n"

						indice_ebr++
					}

					/*if ( aux.Part_next == -1){
						break
					}*/
			
					err2,aux=Leer_EBR(ubicacion,aux.Part_next)
				
				}//final del for para recorrer las logicas
				
				if ( aux.Part_next == -1){
					cadena+="subgraph cluster_"+strconv.Itoa(indice_ebr)+"{\n label=\"EBR_"+strconv.Itoa(indice_ebr)+"\"\n"
						cadena+="fontsize=20 \n style=filled \n fillcolor=olivedrab3 "
						cadena+="\ntbl_"+strconv.Itoa(indice_ebr)+"[shape=box, label=<\n "
						cadena+="<table color=\"black\" border=\"0\" cellborder=\"1\" cellspacing=\"0\"  width=\"300\" height=\"160\" >\n "
						cadena+="<tr>  <td width=\"150\" bgcolor=\"darkorange\"><b>NOMBRE</b></td> <td width=\"150\" bgcolor=\"darkorange\"><b>VALOR</b></td>  </tr>\n"
					
						if ( aux.Part_status == '0' ){
							status2="0"
						}else if ( aux.Part_status == '2') {
							status2="2"
						}

						cadena+="<tr>  <td><b>part_status</b></td> <td><font color='navy'>"+status2+"</font></td>  </tr>\n"
						cadena+="<tr>  <td><b>part_fit</b></td> <td><font color='navy'>"+string(aux.Part_fit)+"</font></td>  </tr>\n"
						cadena+="<tr>  <td><b>part_start</b></td> <td><font color='navy'>"+strconv.FormatInt(aux.Part_start,10)+"</font></td>  </tr>\n"
						cadena+="<tr>  <td><b>part_size</b></td> <td><font color='navy'>"+strconv.FormatInt(aux.Part_size,10)+"</font></td>  </tr>\n"
						cadena+="<tr>  <td><b>part_next</b></td> <td><font color='navy'>"+strconv.FormatInt(aux.Part_next,10)+"</font></td>  </tr>\n"
						name:=GoString(aux.Part_name[:])
						cadena+="<tr>  <td><b>part_name</b></td> <td><font color='navy'>"+name+"</font></td>  </tr>\n"

						cadena+="</table>\n"
						cadena+=">];\n}\n"
					
				}


			} 

		
		}//fial del if para las extendidas


					cadena+="\n } \n"

					ru:=strings.Split(path, ".") 
					crearArchivo(ru[0]+".dot")
					escribeArchivo(ru[0]+".dot",cadena)

					_,err:= exec.Command("dot",ru[0]+".dot","-o",ru[0],"-T"+extencion).Output()
					if (err!=nil){ fmt.Println(" error con el comando1")}

					_,err2:= exec.Command("xdg-open",ru[0]).Output()
					if (err2!=nil){ fmt.Println(" error con el comando2")}

					fmt.Println("Mensaje: Se genero el repoorte del disco correctamente. ")


	}//final del else principal
	 
}//final del metodo



/*
var cadena string=""

	cadena+="digraph G {\n"

	cadena+="hola->mario\n"
	cadena+="}"

	Crear_Carpeta("/home/graficas/")
	crearArchivo("/home/graficas/uno.dot")
	escribeArchivo("/home/graficas/uno.dot",cadena)

	_,err:= exec.Command("dot","/home/graficas/uno.dot","-o","/home/graficas/uno.png","-Tpng").Output()
	if (err!=nil){ fmt.Println(" error con el comando1")}

	_,err2:= exec.Command("xdg-open","/home/graficas/uno.png").Output()
	if (err2!=nil){ fmt.Println(" error con el comando2")}
	
	*/