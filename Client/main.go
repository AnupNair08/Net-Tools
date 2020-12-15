package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func parse(response string, filename string) {
	data := strings.Split(response, "\r\n\r\n")
	ioutil.WriteFile(filename, []byte(data[1]), 0777)
	return
}

func getFile() {
	sitename := os.Args[1]
	pathname := os.Args[2]
	filename := os.Args[3]
	addr, err := net.ResolveTCPAddr("tcp4", sitename+":80")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, addr)
	checkError(err)
	_, err = conn.Write([]byte("GET " + pathname + filename + " HTTP/1.0\r\n\r\n"))
	checkError(err)
	res, err := ioutil.ReadAll(conn)
	checkError(err)
	parse(string(res), filename)
}

func main() {
	getFile()
}
