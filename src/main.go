package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"telegrambotapi"
)

func main() {
	str := TokenRead()
	var bot telegrambotapi.Bot
	bot.APIadress = str[0]
	bot.Token = str[1]
	bot.Proxy = "http://127.0.0.1:1080"
	user, err := bot.GetMe()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

func TokenRead() []string {
	var str []string
	filename := "./toko.password"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("os.Open err:", err)
	}
	defer file.Close()

	br := bufio.NewReader(file)
	i := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		str = append(str, string(a))
		i++
	}
	return str
}
