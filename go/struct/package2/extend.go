package main

type animal struct {

}

type monkey struct {

}

type human struct {

}

type Writer interface{
	Write()
}

type Reader interface{
	Read()
}

type Teacher interface{
	Reader
	Writer
}

type Author struct {
	name string
}

func (a Author) Read(){

}

func main()  {
	a := new(Author)
	a.Read()
}