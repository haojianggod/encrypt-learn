package main

import (
	"fmt"
	"errors"
	"bytes"
	"strconv"
)

func add(a, b int) int {
	return a + b
}

func CaesarBruteForce(encrypt string)  {
   
	for i:=0;i<26;i++ {
		err, decrypt := doDecrypt(encrypt, i)     
		if (err != nil) {
			fmt.Println("input: " + encrypt + ", not valid caesar encrypt content")
			return
		}
		fmt.Println("input: " + encrypt + ", try key: " + strconv.Itoa(i) + ", decrypt content: " + decrypt)
	}
}


func doDecrypt(input string, offset int)(error, string) {
	ret := bytes.Buffer{}
	for _, c := range []rune(input) {
		err := checkCharValid(c)
		if (err != nil) {
			return err, "not valid input"
		}
	    ret.WriteString(convert(c, offset))
	}
	return nil, ret.String()
}

func convert(c rune, offset int) string {
    if int(c) - offset < 65 {
		return string(int(c) + 26 - offset + 32)
	} else {
		return string(int(c) - offset + 32)
	}
}

func checkCharValid(c rune) error {
	if c > 90 || c < 65 {
        return errors.New("not valid char :" + string(c))
	}
	return nil
}

func main() {
	CaesarBruteForce("PELCGBTENCUL")
}
