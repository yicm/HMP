package blog

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestOperation(t *testing.T) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic("failed to read directory")
	}
	for _, f := range files {
		if f.IsDir() {
			fmt.Println("dir: ", f.Name())
		} else {
			fmt.Println("file: ", f.Name(), f.Size())
			if f.Name() == "type.go" {
				file, err := os.Open("./" + f.Name())
				if err != nil {
					log.Printf("Cannot open text file: %s, err: [%v]", f.Name(), err)
					return
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					line := scanner.Text() // or
					//line := scanner.Bytes()

					//do_your_function(line)
					fmt.Printf("%s\n", line)
				}
			}
		}

	}
}
