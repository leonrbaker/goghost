package goghost

/* To Do:
Need to do multiple run_string cmds
Need copyright 
*/


/*
//#include <stdlib.h>
#include <windows.h>
#include "ierrors.h"
#include "iapi.h"
#cgo LDFLAGS: gsdll64.dll

static char**makeCharArray(int size) {
        return calloc(sizeof(char*), size);
}

static void setArrayString(char **a, char *s, int n) {
        a[n] = s;
}

static void freeCharArray(char **a, int size) {
        int i;
        for (i = 0; i < size; i++)
                free(a[i]);
        free(a);
}

*/
import "C"
import "unsafe"

//import "os"
import "errors"
//import "log"
import "strconv"
//import "io/ioutil"


type RevisionStruct struct {
	product string
	copyright string
	revision int
	revisiondate int
}

type GS unsafe.Pointer

func Revision() (RevisionStruct, error) {
	var goRevision RevisionStruct
	var code C.int
	var revision C.struct_gsapi_revision_s
	code = C.gsapi_revision(&revision, C.int(unsafe.Sizeof(revision)))
	if code != 0 {
		return goRevision, errors.New("Revision structure error")
	}
	goRevision.product = C.GoString(revision.product)
	goRevision.copyright = C.GoString(revision.copyright)
	goRevision.revision = int(revision.revision)
	goRevision.revisiondate = int(revision.revisiondate)
	//C.free(unsafe.Pointer(revision.product))
	return goRevision, nil
}

func New_instance() (GS, error) {
	var minst unsafe.Pointer
	var code C.int
	code = C.gsapi_new_instance(&minst, nil)
	if code != 0 {
		return GS(minst), errors.New("gsapi_new_instance failed")
	}
	return GS(minst), nil
}

func Init_with_args(instance GS, sargs []string) error {
	//fmt.Println(sargs)
	cargs := C.makeCharArray(C.int(len(sargs)))
	defer C.freeCharArray(cargs, C.int(len(sargs)))
	for i, s := range sargs {
	        C.setArrayString(cargs, C.CString(s), C.int(i))
	}

	code := C.gsapi_init_with_args(unsafe.Pointer(instance), C.int(len(sargs)), cargs)
	if code != 0 {
		if code <= -100 {
			Exit(instance)
			return errors.New("gsapi_init_with_args failed with code " + strconv.Itoa(int(code)) + ". gsapi_exit called.")
		}
		return errors.New("gsapi_init_with_args failed with code " + strconv.Itoa(int(code)))
	}
	return nil
}

func Run_string(instance GS, str string) error {
	var exit_code C.int
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	code := C.gsapi_run_string_with_length(unsafe.Pointer(instance), cstr, C.uint(len(str)), C.int(0), &exit_code)
	if code != 0 {
		return errors.New("gsapi_run_string_with_length failed with exit code " + strconv.Itoa(int(exit_code)))
	}
	return nil
}

// I don't think we need this
// func set_arg_encoding(minst GS) {
// 	code := C.gsapi_set_arg_encoding(unsafe.Pointer(minst), C.GS_ARG_ENCODING_UTF8)
// 	if code == 0 {
// 		//code = C.gsapi_init_with_args(minst, gsargc, gsargv);
// 		fmt.Println("gsapi_set_arg_encoding returned 0")
// 	} else {
// 		fmt.Println(code)
// 	}
// }

func Exit(instance GS) error {
	code := C.gsapi_exit(unsafe.Pointer(instance))
	if code != 0 {
		return errors.New("gsapi_exit failed")
	}
	return nil
}

func Delete_instance(instance GS) {
	C.gsapi_delete_instance(unsafe.Pointer(instance))
}