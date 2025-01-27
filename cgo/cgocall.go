package cgo

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "greet.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func CgoCall() {
	fmt.Println("cgoCall.........")
	name := C.CString("Gopher")
	name_struct := C.CString("Gopher Struct")
	defer C.free(unsafe.Pointer(name))
	defer C.free(unsafe.Pointer(name_struct))

	year := C.int(2018)
	year_struct := C.int(2020)

	g := C.struct_Greetee{
		name: name_struct,
		year: year_struct,
	}

	ptr := C.malloc(C.sizeof_char * 1024)
	ptr_struct := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))
	defer C.free(unsafe.Pointer(ptr_struct))

	size := C.greet(name, year, (*C.char)(ptr))
	size_struct := C.greet_struct(&g, (*C.char)(ptr_struct))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))

	c := C.GoBytes(ptr_struct, size_struct)
	fmt.Println(string(c))
}
