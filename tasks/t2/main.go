package main

import (
	"fmt"
	"log"
)

// DEPRECATED first variation, works different , multiplies next char NOT prev
func unpack_str(str string) string {
	res := []rune{}
	tmp := []rune(str)
	skip := false
	multiply := 0
	for idx, char := range tmp {
		if char == '/' && skip == false {
			skip = true
			log.Println("SKIPPING NEXT CHAR")
			continue
		}
		if skip == true {
			log.Println("CHAR WAS /..ated so it will not affect")
			skip = false
			res = append(res, char)
			continue
		}
		if skip == false {
			if char >= '0' && char <= '9' && multiply == 0 {
				log.Println("char is numeric", idx, char)
				multiply = int(char) - 48
				log.Println(multiply)
			} else if multiply != 0 && !(char >= '0' && char <= '9') {
				log.Println("char IS NOT numeric and will be MULTIPLIED ", idx, char, multiply)
				for i := 0; i < multiply; i++ {
					res = append(res, char)
				}
				multiply = 0

			} else if multiply != 0 && char >= '0' && char <= '9' {
				log.Println("TWO NUMBERS STAY TOGETHER", idx)
				return ""
			} else if multiply == 0 && !(char >= '0' && char <= '9') {
				log.Println("JUST SINGLE CHAR AND BE ADDED", idx)
				res = append(res, char)
			} else {
				log.Println("THIS IS NOT OK")
			}
		}

	}
	return string(res)
}

// this function works properly
func unpack_str2(str string) string {
	res := []rune{}
	tmp := []rune(str)
	//skip := false
	var buff rune
	buff = 0
	for _, char := range tmp {
		if !(char >= '0' && char <= '9') && buff == 0 {
			buff = char
		} else if !(char >= '0' && char <= '9') && buff != 0 {
			res = append(res, buff)
			buff = char
		} else if (char >= '0' && char <= '9') && buff != 0 {
			multiply := int(char) - 48
			for i := 1; i <= multiply; i++ {
				res = append(res, buff)
			}
			buff = 0
		}
	}
	if buff != 0 {
		res = append(res, buff)
	}
	return string(res)
}

func main() {

	str := "a4bc2d5e"
	res := unpack_str2(str)
	fmt.Println(res)
}
