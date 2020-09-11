package comandos

import(
	"fmt"
	"strings"
	//"io/ioutil"
	//"log"
	"os"
	"os/exec"
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

						cadena+="     <td height=\"200\" width=\"%.1f\">PRIMARIA <br/>%.1f%c</td>\n",porcentaje_aux,porcentaje_real,"%"


					}else{//es una extendida



					}//final del tercer if 

				}else{//es el espacio que no estea asignado
					cadena+="     <td height=\"200\" width=\"%.1f\">LIBRE <br/> %.1f%c</td>\n"+","+string(porcentaje_aux)+","+string(porcentaje_real)+"%"
				
				}//final del segundo if

			}//final del primer if

		}//final del for



















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

