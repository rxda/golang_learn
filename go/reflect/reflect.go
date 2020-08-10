package reflect

import (
	"fmt"
	"reflect"
)

func f() {
	a := []interface{}{1}
	fmt.Println(reflect.TypeOf(a).Elem())
}
