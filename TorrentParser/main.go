package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"unicode"
)

var piece int = 0

func parseInt(data string, index int) (int, int) {
	i := index
	value := ""
	for data[i] != 'e' {
		value += string(data[i])
		i++
	}
	val, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Failed to parse")
		return -1, i
	}
	return val, i
}

func parseString(data string, index int, prev int) (string, int) {
	k := index - 1
	for k >= 0 && unicode.IsDigit(rune(data[k])) && k > prev {
		k--
	}
	l, err := strconv.Atoi(data[k+1 : index])
	if err != nil {
		return " ", index + 1
	}
	if l == 0 {
		return " ", index + 2
	}
	if piece == 1 {
		piece = 2
	}
	if data[index+1:index+l+1] == "pieces" {
		piece = 1
	}
	if piece == 2 {
		temp := ""
		s := data[index+1 : index+l+1]
		for m := 0; m < len(s); m++ {
			temp += fmt.Sprintf("%x ", s[m])
		}
		piece = 0
		return temp, index + l
	} else {
		return data[index+1 : index+l+1], index + l
	}
}

func parseList(data string, index int) ([]interface{}, int) {
	i := index
	arr := make([]interface{}, 0)
	prev := index
	for i < len(data) && data[i] != 'e' {
		if data[i] == 'i' {
			val, j := parseInt(data, i+1)
			arr = append(arr, val)
			i = j
		} else if data[i] == ':' && unicode.IsDigit(rune(data[i-1])) {
			s, j := parseString(data, i, prev)
			arr = append(arr, s)
			i = j - 1
			if s != " " {
				prev = j
				arr = append(arr, s)
			}
		} else if data[i] == 'd' {
			s, j := parseDict(data, i+1)
			arr = append(s)
			i = j
		}
		i++
	}
	return arr, i
}

func parseDict(data string, index int) ([]interface{}, int) {
	i := index
	dict := make([]interface{}, 0)
	prev := index - 1
	for data[i] != 'e' {
		if data[i] == ':' && unicode.IsDigit(rune(data[i-1])) {
			s, j := parseString(data, i, prev)
			i = j - 1
			if s != " " {
				prev = j
				dict = append(dict, s)
			}
			// if s[len(s)-1] == 'e' {
			// 	i++
			// }
		} else if data[i] == 'i' {
			s, j := parseInt(data, i+1)
			i = j
			dict = append(dict, s)
		} else if data[i] == 'l' {
			s, j := parseList(data, i)
			i = j
			dict = append(dict, s)
		} else if data[i] == 'd' {
			s, j := parseDict(data, i+1)
			i = j
			dict = append(dict, s)
		}
		i++
	}
	return dict, i
}

func main() {
	fname := os.Args[1]
	d, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("err")
		return
	}
	data := string(d)
	prev := -1
	for i := 0; i < len(data); i++ {
		if data[i] == ':' && unicode.IsDigit(rune(data[i-1])) {
			s, j := parseString(data, i, prev)
			i = j - 1
			if s != " " {
				prev = j
			}
			if piece == 2 {
				for m := 0; m < len(s); m++ {
					fmt.Printf("%x ", s[m])
				}
				piece = 0
			} else {
				fmt.Println(s)
			}
		} else if data[i] == 'i' {
			s, j := parseInt(data, i+1)
			i = j
			fmt.Println(s)

		} else if data[i] == 'l' {
			s, j := parseList(data, i)
			i = j
			fmt.Println(s)
		} else if data[i] == 'd' {
			s, j := parseDict(data, i+1)
			i = j
			fmt.Println(s)
		} else {
			continue
		}
	}
	return
}
