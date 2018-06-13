package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"telegrambotapi"
	"time"
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

	for {
		update, err := bot.GetUpdates()
		if err != nil {
			fmt.Println(err)
		}
		for a, _ := range update {
			fmt.Println(update[a].Message.From.FirstName, update[a].Message.From.LastName, "  say:", update[a].Message.Text)

			//fmt.Println(update[a].Message.CaptionEntites)
			//fmt.Println(update[a].Message.ChannelChatCreated)
			fmt.Println(update[a].Message.Chat.Title, update[a].Message.Chat.ID, update[a].Message.Chat.Type)
			// fmt.Println(update[a].Message.ConnectedWebsite)
			// fmt.Println(update[a].Message.Contact)
			// fmt.Println(update[a].Message.Date)
			// fmt.Println(update[a].Message.DeleteChatPhoto)
			// fmt.Println(update[a].Message.Document)
			// fmt.Println(update[a].Message.EditDate)
			// fmt.Println(update[a].Message.Entites)
			// fmt.Println(update[a].Message.ForwardDate)
			// fmt.Println(update[a].Message.ForwardFrom)
			// fmt.Println(update[a].Message.ForwardFromChat)
		}
		time.Sleep(time.Second * 10)
		print("\033[H\033[2J")
		fmt.Println("------------------------")
	}
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
