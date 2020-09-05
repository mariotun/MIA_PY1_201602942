
package estructuras


type ExecStr struct { 
	path string
    hpath int64
}

type MkdiskStr struct{
	size int64
	//fit [10]byte
	name [1000]byte
    unit [10]byte
    path [1000]byte
    //hsize, hfit, hunit,hpath int64
}

type RmdiskStr struct{
	path [1000]byte
    hpath int64
}


type FdiskStr struct{
	size int64
    unit [10]byte
    path [1000]byte
    typee [10]byte
    fit [10]byte
    Delete [10]byte
    name [50]byte
    add int64
    hsize, hfit, hunit,hpath,htype,hdelete,hname,hadd int64
}

type MountStr struct{
    path [1000]byte
    name [45]byte
    hname, hpath int64
}

type UnmountStr struct{
    id [1000]byte
    hid int64
}

type MbrStr struct {
    Mbr_tamano int64 //Tamano total del disco en bytes
    Mbr_fecha_creacion [20]byte  //Fecha y hora de creacion del disco
    Mbr_disk_signature int //Numero random, que identifica de forma unica cada disco
    Mbr_disk_fit [2]byte //Tipo de ajuste
    Mbr_partition [4]Partition //4 particiones
}

type Partition struct{
    Part_status byte //Indica si la particion esta activa o no
    Part_type byte //Indica el tipo de particion
    Part_fit byte //Tipo de ajuste de la particion
    Part_start int64 //Indica en que byte del disco inicia la particion
    Part_size int64 //Contiene el tamano de la particion en bytes
    Part_name [16]byte //Nombre de la particion
}

type EbrStr struct{
    Part_status byte //Indica si la particion esta activa o no
    Part_fit byte //Tipo de ajuste
    Part_start int64 //Indica en que byte del disco inicia la particion
    Part_size int64 //Contiene el tamano total de la particion en bytes
    Part_next int64 //Byte en el que esta el proxima EBR. -1 si no hay siguiente
    Part_name [16]byte //Nombre de la particion 
}
 
type RepStr struct{
    name [45]byte
    path [1000]byte
    id [8]byte
    hname, hpath, hid int64
}
