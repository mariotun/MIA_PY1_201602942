
package lexico

//import ("strings")

type Tokenn struct {
	pause string
	comando string
	size int64
	path string
	name string
	unit string
}

func (t *Tokenn) Limpiar(){
	t.pause=""
	t.comando=""
	t.size=0
	t.path=""
	t.name=""
	t.unit=""
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