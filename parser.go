package main
import (
	"log"
	"os"
	"github.com/Mahmoud-Emad/golang-ini-parser/errors/errorHandler"
	// "fmt"
	// "bufio"
	// "errors"
)


type Parser map[string]map[string]string


func (p Parser) read(filename string) {
	file, err := os.Open(filename)
    if err != nil {
		if os.IsNotExist(err) {
			return
		}
        log.Fatal(err)
    }
	
	defer file.Close()
}

func (p Parser) write(data Parser) {
	// This function to write the data to a file
	file := "example.ini"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	for section, items := range p {
		if !errorHandler.validateSection(section) {
			log.Fatal("Invalid section name")
		}
		_, err := f.WriteString("[" + section + "]\n")
		if err != nil {
			log.Fatal(err)
		}
		for key, val := range items {
			_, err := f.WriteString(key + " = " + val + "\n\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	
	defer f.Close()
}

func main() {
	file := "test.txt"
	p := Parser{"test": {"key": "value"}}
	p.read(file)
	p.write(p)
	x := Parser{"test2": {"key": "value"}}
	x.read(file)
	x.write(p)
}