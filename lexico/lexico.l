
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
	"fmt"
	

	"modernc.org/golex/lex"
)

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
	IDENT
	dir
	

    VARI = iota + 0xE002
    DIAGONAL = iota + 0xE002
	
	concatenar

//---------------------------------------------	
	exec
	pathh
	direccion
	flecha
	pause
//----------------------------------------------
	mkdisk
	archivo


	rmdisk
	fdisk
	mount
	unmount

	sizee
	
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
	case IDENT:
		return "IDENT"
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
	case flecha:
		return "flecha"
	case archivo:
		return "archivo"

	case exec:
		return "Reservada"
	case pause:
		return "Reservada"

	case mkdisk:
		return "Reservada"
	case rmdisk:
		return "Reservada"
    case fdisk:
		return "Reservada"
    case mount:
		return "Reservada"
    case unmount:
		return "Reservada"

	case sizee:
		return "size"
	case pathh:
		return "path"



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

/* φ

func X42() int { return 314 }
func saludo () {
    fmt.Println("hola guatemala")
}

mkdisk
vari=-56*/



const src = `
mkdisk
/Escritorio/entrada.txt
/home/Desktop/calificacion.mia
pause
tonto
-size
-path
exec-path->/home/Desktop/calificacion.mia
entrada.sdk

`

func Alexico() { // main

	


	fset := token.NewFileSet()
	file := fset.AddFile("example.go", -1, len(src))
	src := bytes.NewBufferString(src)
	lx, err := lex.New(file, src, lex.RuneClass(rune2Class))
	if err != nil {
		panic(err)
	}

	l := &lexer{lx}

	for {
		c := l.scan()
		fmt.Printf("%v --> %s\n", str(c.Rune), l.TokenBytes(nil))
		
		if str(c.Rune)=="Error"{

			return 
			
			}
		
		if c.Rune == lex.RuneEOF {
			return
		}
		
	}
	// Output:
	// example.go:3:1: FUNC "func"
	// example.go:3:6: IDENT "Xφ42"
	// example.go:3:11: '(' "("
	// example.go:3:12: ')' ")"
	// example.go:3:14: IDENT "int"
	// example.go:3:18: '{' "{"
	// example.go:3:20: IDENT "return"
	// example.go:3:27: INT "314"
	// example.go:3:31: '}' "}"
	// example.go:4:2: EOF "\xff"
}


func (l *lexer) scan() lex.Char {
	c := l.Enter()
%}

digit		[0-9]|{unicodeDigit}
identifier	{letter}({letter}|{digit})* 
int		[0-9]+
negativo "-"+{int}
dir "/"+{identifier}
archivo {identifier}+"."{identifier}
nom "/"+{identifier}+"."{identifier}
direccion  {dir}({dir}|{nom})*


letter		    [_a-zA-Z]
unicodeDigit	\x81
unicodeLetter	\x80


%%

	c = l.Rule0()

[ \t\r\n]+

func			return l.char(FUNC)
vari        	return l.char(VARI)
     
exec			return l.char(exec)
mkdisk			return l.char(mkdisk)
rmdisk			return l.char(rmdisk)	
fdisk			return l.char(fdisk)
mount			return l.char(mount)
unmount			return l.char(unmount)

pause       	return l.char(pause)
-size			return l.char(sizee)
-path			return l.char(pathh)

"/"         	return l.char(DIAGONAL)
"->"	    	return l.char(flecha)
"\*"			return l.char(concatenar)

{identifier}	return l.char(IDENT)
{int}			return l.char(INT)
{negativo}  	return l.char(NEG)
{dir} 			return l.char(dir)
{archivo}		return l.char(archivo)   
{direccion}		return l.char(direccion)


%%


	if c, ok := l.Abort(); ok {
		return l.char(c)
	}

	goto yyAction
}