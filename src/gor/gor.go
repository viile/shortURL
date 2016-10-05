package main

import (
	"fmt"
	"os"
	"bufio"
	"net/http"
	"io/ioutil"
)
var inputReader * bufio.Reader
var input string
var err error
func formatinput(str string) string {
	slen := len(str)
	if slen < 2 {
		return ""
	}
	if str[slen-2:] == "\r\n" {
		return str[:slen-2]
	}
	if str[slen-1:] == "\n" {
		return str[:slen-1]
	}
	return str
}
func main() {
	fmt.Println("Please input uid or phone or email to find user info:")
	inputReader = bufio.NewReader(os.Stdin)
	for {
		input, err = inputReader.ReadString('\n')
		input = formatinput(input)
		fmt.Println(input)
		if len(input) < 1 {
			fmt.Println("input not null")
			continue
		}
		url := "https://children-account.putao.com/u?keyword=" + input
		fmt.Println(url)
		response,_ := http.Get(url)
		defer response.Body.Close()
		body,_ := ioutil.ReadAll(response.Body)
		fmt.Println("result : ",string(body))
	}
}
