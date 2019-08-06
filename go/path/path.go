package main

import (
	"fmt"
	"path"
)

//last
func base(){
	fmt.Println(path.Base("/a/b/c/d/e/f/../e"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))
}

func clean(){
	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
		"",
	}

	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}
}

func dir(){
	fmt.Println(path.Dir("/a/b/c"))
	fmt.Println(path.Dir("a/b/c"))
	fmt.Println(path.Dir("/a/"))
	fmt.Println(path.Dir("a/"))
	fmt.Println(path.Dir("/"))
	fmt.Println(path.Dir(""))
}

func main(){
	//base()
	//clean()
	//dir()

}