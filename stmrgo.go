package stmrgo

/*
 #include <stdlib.h>
 #include "stmr.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func Stem(word string) string {
	cword := C.CString(word)
	defer C.free(unsafe.Pointer(cword))
	end := int64(C.stem(cword, C.int(0), C.int(len(word)-1)))
	fmt.Println(end)
	return word[:end+1]
}
