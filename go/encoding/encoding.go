package encoding

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Dog struct {
	Name string
	Age uint8
	Color string
}

func toJson(d *Dog) (string, string){
	res, _ := json.Marshal(*d)
	res2, _:= json.Marshal(d)
	return string(res), string(res2)
}

func toXml(d *Dog) string {
	res, _ :=xml.Marshal(d)
	return string(res)
}

func base64Ex(str string) string {
	b64 := base64.StdEncoding.EncodeToString([]byte(str))
	return b64
}

func md5Example(str string) string{
	md5Str := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", md5Str)
}

func md5Example2(str string) string{
	h := md5.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func md5file() string{
	f, err := os.Open(".gitignore")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}