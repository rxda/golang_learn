package main

import (
	"fmt"
	"golang_learn/go/reflect/some_package"
	"reflect"
	"unsafe"
)

//读取私有成员变量
func ReadPrivate() {
	device := some_package.NewDevice("迎泽", "迎泽大街", [2]float64{110.123, 35.4343})

	//无法直接读取私有成员变量
	//fmt.Println(device.id)
	// 获取reflect.Value
	idValue := reflect.ValueOf(device).FieldByName("id")
	// 已知类型，根据类型转换
	fmt.Println("device id is:", idValue.Int())
	// private value不能调用 Interface()方法
	//fmt.Println("device id is: ", idValue.Interface())
	// 未知类型
	switch idValue.Kind() {
	case reflect.String:
		fmt.Printf("device id is: %s, type is string", idValue.String())
	case reflect.Int:
		fmt.Printf("device id is: %d, type is int", idValue.Int())
	}
}

//修改成员变量
func EditFieldByName() {
	device := some_package.NewDevice("迎泽", "迎泽大街", [2]float64{110.123, 35.4343})

	// device 不是指针类型，无法使用下面的方法修改成员变量
	// reflect.ValueOf(device).FieldByName("Name").SetString("迎泽(通过reflect 修改)")

	// Indirect返回pointer value指向的value
	v := reflect.Indirect(reflect.ValueOf(&device))
	// 若已知value是指针，则可直接调用Elem,和上句作用相同
	//v := reflect.ValueOf(&device).Elem()
	v.FieldByName("Name").SetString("迎泽(通过reflect 修改)")

	v.FieldByName("Name").Set(reflect.ValueOf("迎泽(通过reflect 修改)"))

	fmt.Println(device.Name)
}
// output
// 迎泽(通过reflect 修改)


func EditPrivate() {
	device := some_package.NewDevice("迎泽", "迎泽大街", [2]float64{110.123, 35.4343})

	v := reflect.ValueOf(&device).Elem().FieldByName("Location")
	//v.FieldByName("id").SetInt(1234567)
	//不能直接使用Set修改reflect: reflect.Value.SetInt using value obtained using unexported field
	ptrToV := unsafe.Pointer(v.UnsafeAddr())
	p := (*[2]float64)(ptrToV)
	*p = [2]float64{1,2}
	fmt.Println(device)
}

func CallMethod()  {
	device := some_package.NewDevice("迎泽", "迎泽大街", [2]float64{110.123, 35.4343})
	// MethodByName只能获取public 方法
	//function := reflect.ValueOf(&device).MethodByName("change") // -- Zero Value
	// 注意，方法的接收器为指针类型，reflect.ValueOf()就必须传入指针值
	//function := reflect.ValueOf(device).MethodByName("Change") // -- 找不到方法
	function := reflect.ValueOf(&device).MethodByName("Change")

	// In 可获取参数类型，Out可获取返回类型
	param1Type := function.Type().In(0)
	param2Type := function.Type().In(1)

	// 根据reflect.Type 创建 reflect.Value
	param1 := reflect.New(param1Type).Elem()  // New创建的为指针，如果要求的参数不是指针，则需调用Elem取值
	param1.Set(reflect.ValueOf([2]float64{110,110}))

	param2 := reflect.New(param2Type).Elem()
	param2.Set(reflect.ValueOf("小店"))


	resps := function.Call([]reflect.Value{param1, param2})
	if err, ok := resps[1].Interface().(error); ok{
		if err != nil{
			panic(err)
		}
	}
	fmt.Println(resps[0].Interface().(string))
	fmt.Println(device)
}
// output
// device location is [110 110]
// 小店迎泽大街
// {5577006791947779410 小店 [110 110] 迎泽大街}