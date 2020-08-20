package funciones

import(
	"fmt"
	"strconv"
)


/*type NodoParticiones struct{
	path string
	nombre string
	letra [20]byte
	num int64
	next *NodoParticiones
	datos interface{}
}

type Lista struct{
	inicio *NodoParticiones
}

func (L *Lista) Insertar (datos interface{ }){
	lista:=&NodoParticiones{
		next:L.inicio,
		datos: datos,
	}

	if L.inicio != nil{
		L.inicio
	}
}*/

type NodoParticiones struct{
	path string
	nombre string
	letra string
	num int
	next *NodoParticiones

}

type ListaSimple struct{
	first *NodoParticiones
	last *NodoParticiones
}

func (L *ListaSimple) Insertar (datos *NodoParticiones){
	
	if L.first == nil{
		L.first=datos
		L.first.next=nil
		L.last=L.first
	}else{
		L.last.next=datos
		datos.next=nil
		L.last=datos
	}
}

func (L *ListaSimple) Mostrar(){

	actual:=L.first

	for actual!=nil{
		fmt.Print("[",actual.path,",",actual.num,",",actual.letra,",",actual.nombre,"]-->")
		actual=actual.next
	}

}

/*func (L *ListaSimple) BuscarLetra (path string, nombre string) int{
	aux:=L.first
	retorno:=1

	for aux !=nil{
		if ( (path==aux.path) && (nombre==aux.nombre)){
			return -1
		}else{
			if path==aux.path{
				return aux.letra
			}else if retorno <= aux.letra{
				retorno++;
			}
		}
		aux=aux.next
	}
	return retorno
}*/

func (L *ListaSimple) BuscarNumero(path string, nombre string) int{
	aux:=L.first
	retorno:=1

	for aux!=nil{
		if ( (path==aux.path) && (retorno==aux.num)){
			retorno++
			fmt.Println("retorno: ",retorno)
		}
		aux=aux.next	
	}
	return retorno
}

func (L *ListaSimple) GetDireccion(id string) string {
	aux:=L.first
	
	for aux!=nil{
		temp_id:="vd"
		temp_id+=aux.letra+strconv.Itoa(aux.num)
		
		if id==temp_id{
			return aux.path
		}
	}
	return "nil"
}

func (L *ListaSimple) BuscarNodo(path string,nombre string) bool{
	aux:=L.first

	for aux!=nil{
		if((aux.path==path) && (aux.nombre==nombre)){
			return true
		}
		aux=aux.next
	}

	return false
}

func (L *ListaSimple) EliminarNodo(id string) int{
	aux:=L.first

	temp_id:="vd"
	temp_id+=aux.letra+strconv.Itoa(aux.num)

	if id==temp_id{
		L.first=aux.next
		return 1
	}else{
		var aux2 *NodoParticiones

		for aux!=nil{
			tempId:="vd"
			tempId+=aux.letra+strconv.Itoa(aux.num)
			
			if id==tempId{
				aux2.next=aux.next
				return 1
			}
			aux2=aux
			aux=aux.next
		}
	}
	return 0
}


func Ejecutar(){

	/*uno:=NodoParticiones{}
	uno.path="/honme/mario"
	uno.nombre="mariotun"
	//uno.letra=nil
	uno.num=76*/

	ls:=ListaSimple{}
	//ls.Insertar(&uno)
	ls.Insertar( & NodoParticiones{path:"//mario",nombre:"yonathan",letra:"d",num:76} )
	ls.Insertar( & NodoParticiones{path:"//cristan",nombre:"humberto",letra:"r",num:89} )
	ls.Insertar( & NodoParticiones{path:"//lidia",nombre:"veronica",letra:"s",num:4524} )
	ls.Insertar( & NodoParticiones{path:"//leonel",nombre:"esteban",letra:"a",num:678} )
	ls.Insertar( & NodoParticiones{path:"//nicolas",nombre:"nicolas",letra:"t",num:326} )
	ls.Mostrar()
	fmt.Println(ls.BuscarNumero("//mario","yonathan"))
	fmt.Println(ls.BuscarNumero("//mario","yonathan"))
	fmt.Println(ls.BuscarNumero("//mario","yonathan"))
}