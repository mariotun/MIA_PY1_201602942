
package estructuras


type ExecStr struct { 
	path [1000]byte
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
	size int64;
    unit [10]byte;
    path [1000]byte;
    typee [10]byte;
    fit [10]byte;
    Delete [10]byte;
    name [50]byte;
    add int64;
    hsize, hfit, hunit,hpath,htype,hdelete,hname,hadd int64

}

type MountStr struct{
    
}

