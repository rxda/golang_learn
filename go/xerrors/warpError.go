package myError

import (
	"errors"
	"fmt"
)

func warpError(){
	err1 := errors.New("1st error\n")
	//err2 := errors.New("2nd error")
	err3 := fmt.Errorf("wrap了一个错误\n%w", err1)
	err4 := fmt.Errorf("warp了err3\n%w", err3)

	//fmt.Println(err1)
	//fmt.Println(err3)
	//fmt.Println(errors.Is(err4, err1))
	//fmt.Println(err4)
	panic(err4)
	//fmt.Println(errors.Unwrap(err4))
	//fmt.Println(errors.Unwrap(errors.Unwrap(err4)))

}