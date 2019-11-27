package slice

import "fmt"

func A(input *[]int ){
	fmt.Printf("%p\n", input)
	for i:=0;i<100;i++{
		*input = append(*input, i)
	}

	fmt.Println(input)
	fmt.Printf("%p\n", input)
}

func B(){
	sl := make([]int, 0, 8)
	fmt.Println(sl)
	fmt.Printf("B:%p\n", sl)
	A(&sl)
	fmt.Println(sl)
	fmt.Printf("B:%p\n", sl)
}
