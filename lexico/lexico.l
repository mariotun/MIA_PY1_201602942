
%yyc c
%yyn c = l.Next()
%yym l.Mark()

%{

package lexico

import (
	"bytes"
	"fmt"
	"go/token"
	"unicode"
	"strconv"
	
	
	"container/list"
	"modernc.org/golex/lex"
)

		
	var LValores *list.List
	//var list0 *list.List
  	
/*
type ParamComandos struct {
    nombre string
    ruta   string
}

var parametros ParamComandos

func Retorno() ParamComandos{
	return parametros
}
func Datos(dat ParamComandos){
	 parametros=dat
}
*/

var Parametros Tokenn


// Allocate Character classes anywhere in [0x80, 0xFF].
const (
	classUnicodeLeter = iota + 0x80
	classUnicodeDigit
	classOther
)

// Parser token values.
const (
	FUNC = iota + 0xE002
	INT
    NEG
	dir
	

    VARI = iota + 0xE002
    DIAGONAL = iota + 0xE002
	
	concatenar
	comentario
	identificador
	espacio
//---------------------------------------------	
	exec
	path
	direccion
	direccomi
	flecha
	pause
//----------------------------------------------
	mkdisk
	archivo
	name
	sizee
	unit

	rmdisk
	fdisk
	mount
	unmount


	
)



func valor(val int)int{
	return val
}


// For pretty printing.
func str(r rune) string {
	switch r {
	case FUNC:
		return "FUNC"
	case INT:
		return "INT"
    case VARI:
		return "VARI"
    case NEG:
		return "NEG"
    case DIAGONAL:
		return "DIAGONAL"
	case dir:
		return "dir"
	case direccion:
		return "direccion"
	case direccomi:
		return "direccomi"
	case flecha:
		return "flecha"
	case archivo:
		return "archivo"
	case comentario:
		return "comentario"
	case identificador:
		return "identificador"
	case concatenar:
		return "concatenar"

	case exec:
		return "exec"
	case pause:
		return "pause"
	case mkdisk:
		return "mkdisk"
	case rmdisk:
		return "rmdisk"
    case fdisk:
		return "fdisk"
    case mount:
		return "mount"
    case unmount:
		return "unmount"

	case sizee:
		return "size"
	case path:
		return "path"
	case name:
		return "name"
	case unit:
		return "unit"


	case lex.RuneEOF:
		return "EOF"
	
    default :
        return "Error"
	}

	//return fmt.Sprintf("%q", r)
}

type lexer struct {
	*lex.Lexer
}

func (l *lexer) char(r int) lex.Char {
	return lex.NewChar(l.First.Pos(), rune(r))
}

func rune2Class(r rune) int {
	if r >= 0 && r < 0x80 { // Keep ASCII as it is.
		return int(r)
	}

	if unicode.IsLetter(r) {
		return classUnicodeLeter
	}

	if unicode.IsDigit(r) {
		return classUnicodeDigit
	}

	return classOther
}


func Alexico(entrada string) int{ 

	

	fset := token.NewFileSet()
	file := fset.AddFile("example.go", -1, len(entrada))
	src := bytes.NewBufferString(entrada)
	lx, err := lex.New(file, src, lex.RuneClass(rune2Class))
	if err != nil {
		panic(err)
	}

	l := &lexer{lx}

	LValores = list.New()
	//LValores := New()
	
	
	

	for {
		c := l.scan()
		//fmt.Printf("%v --> %s\n", str(c.Rune), l.TokenBytes(nil))


		//LValores.PushBack(str(c.Rune)) 
		LValores.PushBack(string(l.TokenBytes(nil)))

		/*if str(c.Rune)=="path"{
			//fmt.Printf("*** dentro del metodo: %s \n",l.TokenBytes(nil))
			//san:=string(l.Token())
		Parametros.Set_Path(l.TokenBytes(nil)) 
		}*/
		
		if str(c.Rune)=="Error"{
			fmt.Printf("%v --> %s\n", str(c.Rune), l.TokenBytes(nil))
			return 1
			
			}
		
		if c.Rune == lex.RuneEOF {
	
			/*for temp := LValores.Front(); temp != nil; temp = temp.Next() {
        fmt.Println(">>>>>",temp.Value)
    			}*/
			return 0
		}
		
	}

		
}





func (l *lexer) scan() lex.Char {
	c := l.Enter()
%}

digit			[0-9]|{unicodeDigit}
identifier		{letter}({letter}|{digit})* 
int				[0-9]+
negativo 		"-"+{int}
dir 			{identifier}+"/"
archivo 		{identifier}+".dsk"
nom				{identifier}+"."{identifier}
direccion		"/"+{dir}({dir}|{nom})* 
direccomi		"\""+[^\n]+"\""
comentario		"#"[^\n]*			 
letter		    [_a-zA-Z]	
unicodeDigit	\x81
unicodeLetter	\x80


%%

	c = l.Rule0()

[ \t\r\n]+

pause       	{ Parametros.Set_Comando(string(l.TokenBytes(nil))); return l.char(pause) }
exec			{ /*LValores.PushFront("exec") ;*/  Parametros.Set_Comando(string(l.TokenBytes(nil))); return l.char(exec) }
mkdisk			{ Parametros.Set_Comando(string(l.TokenBytes(nil))); return l.char(mkdisk) }
rmdisk			{ Parametros.Set_Comando(string(l.TokenBytes(nil))); return l.char(rmdisk)	}
fdisk			{ Parametros.Set_Comando(string(l.TokenBytes(nil))); return l.char(fdisk) }
mount			{ Parametros.Set_Comando(string(l.TokenBytes(nil))); return l.char(mount) }
unmount			{ Parametros.Set_Comando(string(l.TokenBytes(nil))); return l.char(unmount) }


size			{ /*Parametros.Set_Size(string(l.TokenBytes(nil)));*/ return l.char(sizee) }
path			{ /*Parametros.Set_Path(string(l.TokenBytes(nil)));*/ return l.char(path) }
name			{ /*Parametros.Set_Name(string(l.TokenBytes(nil)));*/ return l.char(name) }
unit			{ /*Parametros.Set_Unit(string(l.TokenBytes(nil)));*/ return l.char(unit) }

"/"         	return l.char(DIAGONAL)
"->"	    	return l.char(flecha)
"\*"			return l.char(concatenar)
"-"				{ }


{int}			{ if num,err:=strconv.ParseInt(string(l.TokenBytes(nil)),10,64);err==nil{  Parametros.Set_Size(num) }else{ } ;
				  return l.char(INT) }
{negativo}  	return l.char(NEG)
{dir} 			return l.char(dir)
{archivo}		{ Parametros.Set_Name(string(l.TokenBytes(nil))); return l.char(archivo)}   
{direccion}		{ Parametros.Set_Path(string(l.TokenBytes(nil))) ; return l.char(direccion) }
{direccomi}		{ Parametros.Set_Path(string(l.TokenBytes(nil))) ; return l.char(direccomi) }
{comentario}	return l.char(comentario)
{identifier}	{ Parametros.Set_Unit(string(l.TokenBytes(nil))); return l.char(identificador) }


%%

	if c, ok := l.Abort(); ok { return l.char(c) }
	goto yyAction
}