package main

/*
#include <stdio.h>

static void SayHello(const char* s) {
    puts(s);
}
*/
import "C"


func main() {
	C.puts(C.CString("Hello, World\n"))
}