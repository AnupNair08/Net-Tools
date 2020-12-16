package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"unicode"
)

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
	for unicode.IsDigit(rune(data[k])) && k > prev {
		k--
	}
	l, err := strconv.Atoi(data[k+1 : index])
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to parse")
		return "", index
	}
	return data[index+1 : index+l+1], index + l
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
			prev = j
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
			prev = j
			dict = append(dict, s)
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
	d, err := ioutil.ReadFile("file.torrent")
	if err != nil {
		fmt.Println("err")
		return
	}
	data := string(d)
	// data := "d8:announce40:udp://tracker.leechers-paradise.org:696913:announce-liste"
	// data := "d4:infod5:filesld6:lengthi140e4:pathl21:Big Buck Bunny.en.srted6:lengthi276134947e4:pathl18:Big Buck Bunny.mp4ed6:lengthi310380e4:pathl10:poster.jpgeeeeeee"
	prev := 0
	for i := 0; i < len(data); i++ {
		if data[i] == ':' && unicode.IsDigit(rune(data[i-1])) {
			s, j := parseString(data, i, prev)
			i = j - 1
			prev = j
			fmt.Println(s)
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
		} else if !unicode.IsPrint(rune(data[i])) {
			fmt.Printf("%x", data[i])
		}
	}
	return
}
