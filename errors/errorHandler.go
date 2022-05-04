package errorHandler

import (
	"strings"
)


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

	//check presence of more than one '='
	numberOfEqualSigns := strings.Count(key, "=")
	if numberOfEqualSigns != 1 {
		return false
	}

	if !validateName(key){
		return false
	}

	return true
}