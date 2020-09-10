
package lexico

import (
	"strings"
	//"strconv"
	)

type Tokenn struct {
	pause string
	comando string
	size int64
	path string
	name string
	namefd string
	unit string
	tipo string
	fit string
	delete string
	add string
	namegrafic string
}

func (t *Tokenn) Limpiar(){
	t.pause=""
	t.comando=""
	t.size=0
	t.path=""
	t.name=""
	t.namefd=""
	t.unit=""
	t.tipo=""
	t.fit=""
	t.delete=""
	t.add=""
	t.namegrafic=""
}

//---------------------------------------------------------------
func (t *Tokenn) Set_Pause(npause string) {
	t.pause = npause
}
func (t *Tokenn) Get_Pause() string{
	return t.pause
}
//--------------------------------------------------------------
func (t *Tokenn) Set_Comando(ncomando string) {
	t.comando = ncomando
}
func (t *Tokenn) Get_Comando() string{
	return t.comando
}
//--------------------------------------------------------------
func (t *Tokenn) Set_Size(nsize int64) {
	t.size = nsize
}
func (t *Tokenn) Get_Size() int64{
	return t.size
}
//--------------------------------------------------------------
func (t *Tokenn) Set_Path(npath string) {
	t.path = npath
	//fmt.Println("del metodo set:",string(t.path))
}
func (t *Tokenn) Get_Path() string{
	return t.path //string(data[:])
}
//--------------------------------------------------------------
func (t *Tokenn) Set_Name(nname string) {
	t.name = nname
}
func (t *Tokenn) Get_Name() string{
	return t.name
}
//--------------------------------------------------------------
func (t *Tokenn) Set_Unit(nunit string) {
	t.unit = nunit
}
func (t *Tokenn) Get_Unit() string{
	return t.unit
}
//--------------------------------------------------------------
func (t *Tokenn) Set_Tipo(ntipo string) {
	t.tipo =ntipo
}
func (t *Tokenn) Get_Tipo() string{
	return t.tipo
}
//--------------------------------------------------------------
func (t *Tokenn) Set_Namefd(nnamefd string) {
	t.namefd = nnamefd
}
func (t *Tokenn) Get_Namefd() string{
	return t.namefd
}
//--------------------------------------------------------------
func (t *Tokenn) Set_Fit(nfit string) {
	t.fit = nfit
}
func (t *Tokenn) Get_Fit() string{
	return t.fit
}
//--------------------------------------------------------------
func (t *Tokenn) Set_Delete(ndelete string) {
	t.delete = ndelete
}
func (t *Tokenn) Get_Delete() string{
	return t.delete
}
//--------------------------------------------------------------
func (t *Tokenn) Set_Add(nadd string) {
	
	dato:=strings.Split(nadd,"->")

	//if num,err:=strconv.ParseInt(dato[1],10,64);err==nil{  t.add = num }else{ }
	t.add = dato[1]	
}
func (t *Tokenn) Get_Add() string{
	return t.add
}
//--------------------------------------------------------------
func (t *Tokenn) Set_NameGrafic(namegra string) {
	t.namegrafic=namegra
	
}
func (t *Tokenn) Get_NameGrafic() string{
	return t.namegrafic
}
