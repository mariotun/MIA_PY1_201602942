package comandos

import(
	"fmt"
	"strings"
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
			}

		}else{
			fmt.Println("Error: NO se encontro la particion.")
		}

	}

}

func crearArchivo(path string) {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
	  var file, err = os.Create(path)
	  if existeError(err) {
		return
	  }
	  defer file.Close()
	}
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


func grafica_disco(ubicacion string,path string,extencion string){

	err,masterb:=Leer_MBR(path)	
	var cadena string=""

	if ( err != nil){
		fmt.Println(" Error: No se puede crear el reporte, no se encontro el disco. ")
	}else{

		total:=masterb.Mbr_tamano
		var usado int64

		ru:=strings.Split(path, ".") 
		crearArchivo(ru[0]+".dot")

		cadena+="digraph G{\n\n"
		cadena+="  tbl [\n    shape=box\n    label=<\n"
		cadena+="     <table border=\"0\" cellborder=\"2\" width=\"600\" height=\"200\" color=\"blue\">\n"
		cadena+="     <tr>\n"
		cadena+= "     <td height=\"200\" width=\"100\"> MBR </td>\n"

		for i:=0 ; i < 4 ; i++ {
			parcial:=masterb.Mbr_partition[i].Part_size

			if ( masterb.Mbr_partition[i].Part_start != -1){
				porcentaje_real := (parcial*100)/total
                porcentaje_aux := (porcentaje_real*500)/100
				usado += porcentaje_real
				
				if (masterb.Mbr_partition[i].Part_status != '1') {
					if (masterb.Mbr_partition[i].Part_type == 'P'){

						cadena+="  <td height=\"200\" width="+"\""+string(porcentaje_aux)+"\""+">PRIMARIA <br/>"+string(porcentaje_real)+"%"+"</td>\n"

						if (i != 3 ) {
							part1:=masterb.Mbr_partition[i].Part_start + masterb.Mbr_partition[i].Part_size
							part2:=masterb.Mbr_partition[i+1].Part_start
							if (masterb.Mbr_partition[i+1].Part_start != -1){
								if ( (part2-part1)!=0 ){//existe la fragmentacion
									fragmentacion2:=part2-part1
									porcentaje_real2:=(fragmentacion2*100)/total
									porcentaje_aux2:=(porcentaje_real*500)/100

									cadena+="<td height=\"200\" width="+"\""+string(porcentaje_aux2)+"\""+"> LIBRE <br/>"+string(porcentaje_real2)+"%"+"</td>\n"
								}
							}
						}else{
							part11:=masterb.Mbr_partition[i].Part_start + masterb.Mbr_partition[i].Part_size
							tam_mbr:=total+int64(binary.Size(masterb))

							if ( (tam_mbr-part11) != 0){
								libre:=(tam_mbr-part11)+int64(binary.Size(masterb))
								porcentaje_reall:=(libre*100)/total
								porcentaje_auxx:=(porcentaje_reall*500)/100

								cadena+="<td height=\"200\" width="+"\""+string(porcentaje_auxx)+"\""+"> LIBRE <br/>"+string(porcentaje_reall)+"%"+"</td>\n"

							}
						}

					}else{//es una extendida

						cadena+="  <td  height=\"200\" width=\""+string(porcentaje_real)+"\">\n     <table border=\"0\"  height=\"200\" WIDTH=\""+string(porcentaje_real)+"\" cellborder=\"1\">\n"
						cadena+="  <tr>  <td height=\"60\" colspan=\"15\"> EXTENDIDA </td>  </tr>\n     <tr>\n"

						err2,ebr:=Leer_EBR(path,masterb.Mbr_partition[i].Part_start)

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
											cadena+=" <td height=\"140\"> EBR </td>\n"
											cadena+=" <td height=\"140\">LOGICA<br/>"+string(porcentaje_real)+"%"+"</td>\n"
										
										}else{//espacion en disco no asignado
											cadena+=" <td height=\"150\">LIBRE 1 <br/>"+string(porcentaje_real)+"%"+"</td>\n"
										}

										if ( aux.Part_next == -1){
											parcial=(masterb.Mbr_partition[i].Part_start + masterb.Mbr_partition[i].Part_size) - ( aux.Part_start + aux.Part_size)
											porcentaje_real=(parcial*100)/total
											
											if ( porcentaje_real != 0) {
												cadena+="<td height=\"150\">LIBRE 2<br/>"+string(porcentaje_real)+"%"+"</td>\n"
											}
											break
										}/*else{
											err2,aux=Leer_EBR(path,aux.Part_next)
										}*/

									}

									err2,aux=Leer_EBR(path,aux.Part_next)

								}//final del for para recorre las logicas

								parcial=aux.Part_size
								porcentaje_real=(parcial*100)/total

									if (porcentaje_real != 0){
										if ( aux.Part_status != '1'){
											cadena+=" <td height=\"140\"> EBR </td>\n"
											cadena+=" <td height=\"140\">LOGICA<br/>"+string(porcentaje_real)+"%"+"</td>\n"
										
										}else{//espacion en disco no asignado
											cadena+=" <td height=\"150\">LIBRE 1 <br/>"+string(porcentaje_real)+"%"+"</td>\n"
										}

										if ( aux.Part_next == -1){
											parcial=(masterb.Mbr_partition[i].Part_start + masterb.Mbr_partition[i].Part_size) - ( aux.Part_start + aux.Part_size)
											porcentaje_real=(parcial*100)/total
											
											if ( porcentaje_real != 0) {
												cadena+="<td height=\"150\">LIBRE 2<br/>"+string(porcentaje_real)+"%"+"</td>\n"
											}
											break
										}

									}
						//---------------------------------------------------------------------			


							}else{
								cadena+=" <td height=\"140\"> Ocupado"+string(porcentaje_real)+"%"+"</td>"
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

										cadena+="   <td height=\"200\" width=\""+string(porcentaje_aux)+"\">LIBRE<br/>"+string(porcentaje_real)+"%"+"</td>\n"
									}
								}
							}else{
								p1:=masterb.Mbr_partition[i].Part_start + masterb.Mbr_partition[i].Part_size
								mbr_tam:=total + int64(binary.Size(masterb))

								if ( (mbr_tam - p1) != 0){//lo que esta libre
									libre:=(mbr_tam - p1) + int64(binary.Size(masterb))
									porcentaje_real:=(libre*100)/total
									porcentaje_aux:=(porcentaje_real*500)/100

									cadena+="   <td height=\"200\" width=\""+string(porcentaje_aux)+"\">LIBRE<br/>"+string(porcentaje_real)+"%"+"</td>\n"
								}

							}

						}





					}//final del tercer if 

				}else{//es el espacio que no estea asignado
					cadena+="     <td height=\"200\" width="+"\""+string(porcentaje_aux)+"\""+">LIBRE <br/>"+string(porcentaje_real)+"%"+"</td>\n"
				
				}//final del segundo if

			}//final del primer if

		}//final del for que recorre las particiones

		cadena+="</tr> \n  </table>  \n>];\n\n}"

















	}


	
}


func grafica_mbr(ubicacion string,path string,extencion string){


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

}

