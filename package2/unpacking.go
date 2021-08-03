package main

import (
	"errors"
	"fmt"
)

func printer(char rune, count int) (out string) {
	if char == 0 {
		return
	}
	if count == 0 {
		count++
	}
	for i := 0; i < count; i++ {
		out += string(char)
	}
	return out
}

func UnpackingString(str string) (string, error) {
	var out = ""
	var errorCode = "invalid string"
	var startNumeric = 48
	var endNumeric = 57
	var escaping = 92
	var counter int
	var escapingFlag bool
	var tmp rune

	for _, char := range str {
		if escapingFlag ||
			(int(char) <= startNumeric || int(char) >= endNumeric) && int(char) != escaping {
			out += printer(tmp, counter)
			tmp = char
			escapingFlag = false
			counter = 0
		} else if int(char) >= startNumeric && int(char) <= endNumeric {
			if tmp != 0 {
				counter = counter*10 + int(char) - startNumeric
			} else {
				return "", errors.New(errorCode)
			}
		} else if int(char) == escaping {
			escapingFlag = true
		}
	}
	out += printer(tmp, counter)
	return out, nil
}

func main() {
	testString := "ab\\\\21"
	fmt.Println(UnpackingString(testString))
}
