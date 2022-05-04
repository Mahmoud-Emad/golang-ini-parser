package main

import (
	"log"
	"os"
	"strings"
	// "fmt"
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
		if !validateSection(section) {
			log.Fatal("Invalid section name")
		}		
		_, err := f.WriteString("[" + section + "]\n")
		if err != nil {
			log.Fatal(err)
		}
		for key, val := range items {
			if !validateKey(key) {
				log.Fatal("Invalid key name")
			}
			_, err := f.WriteString(key + " = " + val + "\n\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	
	defer f.Close()
}


func (p Parser) readFromString(str string){
	// This function to read the ini content from a string
	section := ""
	for _, line := range strings.Split(str, "\n") {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			section = strings.Trim(line, "[]")
			p[section] = make(map[string]string)
		} else {
			if strings.Contains(line, "=") {
				parts := strings.Split(line, "=")
				key := strings.Trim(parts[0], " ")
				value := strings.Trim(parts[1], " ")
				p[section][key] = value
			}
		}
	}
}

func main() {
	file := "test.txt"
	p := Parser{"test": {"key": "value"}}
	p.read(file)
	// p.write(p)
	p.readFromString("[test]\nkey = value")

	// x := Parser{"test2": {"key": "value"}}
	// x.read(file)
	// x.write(p)
}


// Validations
func validateName(name string) bool {
	wrongChars := []string{
		"#", "@", "!", "$", "%", "^", "&", "*",
		"(", ")","=", "?", ":", ";", ",", ".",
		">", "<", "/", "\\", "|", "~", "`", "\"", "'", " ",
	}
	for _, char := range wrongChars {
		if strings.Contains(name, char) {
			return false
		}
	}
	return true
}

func validateSection(section string) bool {
	//validate sections

	if len(section) < 2 {
		return false
	}

	if !validateName(section){
		return false
	}

	return true
}

func validateKey(key string) bool {
	//validate key
	//check if string is empty
	if len(key) == 0 || strings.HasPrefix(key, "[") && strings.HasSuffix(key, "]") {
		return false
	}

	if !validateName(key){
		return false
	}

	return true
}