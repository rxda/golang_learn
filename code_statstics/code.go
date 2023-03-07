package code_statstics

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Statistics(dir string) {
	//cmd := exec.Command("git", "ls-files", "|", "xargs", "cat", "|", "wc", "-l")
	cmd := exec.Command("git", "ls-files","| xargs cat | wc -l")
	println(cmd.String())
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("combined out:\n%s\n", string(out))
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		fmt.Println(string(out))
		return nil
	})
	if err != nil{
		panic(err)
	}
}
