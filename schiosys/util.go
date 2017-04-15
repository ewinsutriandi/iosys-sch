package schiosys

import (
	"log"
	"reflect"
	"runtime"
)

// ref from :http://stackoverflow.com/questions/24809287
// log error with function name and line number reference
func HandleError(err error) (b bool) {
	if err != nil {
		// notice that we're using 1, so it will actually log the where
		// the error happened, 0 = this function, we don't want that.
		pc, fn, line, _ := runtime.Caller(1)

		log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		b = true
	}
	return
}

//IsEmpty check for empty object
//reference : http://stackoverflow.com/questions/25349004/
func IsEmpty(object interface{}) bool {
	//First check normal definitions of empty
	log.Printf("dodollll masuk")
	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}
	log.Printf("dodollll keluar")
	//Then see if it's a struct
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		// and create an empty copy of the struct object to compare against
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true
		}
	}
	return false
}
