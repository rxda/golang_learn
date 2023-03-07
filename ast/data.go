package ast

type Unit struct {
	Label string
	Type  string
	Value string `json:"value,omitempty"`
}

type Table struct {
	Label string
	Units []Unit

}

type SubTable struct {
	Label string
	Header []string

}