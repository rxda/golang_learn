package main

import "fmt"
import "encoding/binary"

func main() {
	var mySlice = []byte{0x56, 0x36}
	data := binary.BigEndian.(mySlice)
	fmt.Println(data)
	x, n := binary.Varint(mySlice)
	if n != len(mySlice) {
		fmt.Println("Varint did not consume all of in")
	}
	fmt.Println(x)
}



