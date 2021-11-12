package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	botToken = "2108403709:AAHUN0sg0POV1Vpsc3x7ZFCG1lSsbheB0Ow"
	//channelname        = "@test_channel_for_rest_server"
	groupId   int64 = -597182447
	channelId int64 = -1001580495914
	reader          = bufio.NewReader(os.Stdin)
)

type MessageBody struct {
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type Message struct {
	msg      MessageBody
	priority string
}

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func getUpdatesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("we got to: getUpdatesHandler")
	resp, _ := http.Get("https://api.telegram.org/bot" + botToken + "/getUpdates")
	fmt.Println(resp)
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("we got to: sendMessage")

	message := new(Message)
	message.msg.ChatId = channelId

	fmt.Println("Give me a message text:")
	message.msg.Text, _ = reader.ReadString('\n')
	message.msg.Text = "Bot: " + message.msg.Text

	fmt.Println("Give me a message priority:")
	message.priority, _ = reader.ReadString('\n')

	fmt.Println(message)

	reqBytes, _ := json.Marshal(message.msg)
	url := "https://api.telegram.org/bot" + botToken + "/sendMessage"

	res, _ := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))

	fmt.Println("post result:" + res.Status)
}

func handleRequests() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/articles", getAllArticlesHandler)
	// TODO: Handle get requests for single articles searched by http variables
	// http.HandleFunc("/articlesbyid", getArticleByIdHandler)
	http.HandleFunc("/sendMessage", sendMessage)
	http.HandleFunc("/getUpdates", getUpdatesHandler)
	http.HandleFunc("/spamtheshit", spammerHandler)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

func spammerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("we got to: spammerHandler")

	message := new(Message)
	message.msg.ChatId = groupId
	message.msg.Text = "ZAFAR NIMAGAAP JALA QALISAAN"

	url := "https://api.telegram.org/bot" + botToken + "/sendMessage"

	reqBytes, _ := json.Marshal(message.msg)
	for i := 0; i < 13; i++ {
		result, _ := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
		fmt.Println("spam result: " + result.Status)
	}
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 style=\"font-family: consolas\">Welcome to the page!</h1>")
	fmt.Println("we got to: homePage")
}

func getAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("we got to: getAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

//func getArticleByIdHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("we got to: getAllArticles")
//	keys, ok := r.URL.Query()["id"]
//
//	fmt.Println("keys:", keys, "\nok: ", ok)
//
//	key := keys[0]
//	for i := 0; i < len(Articles); i++ {
//		if string(key) == strconv.Itoa(i) {
//			json.NewEncoder(w).Encode(Articles[i])
//		}
//	}
//	json.NewEncoder(w).Encode(Articles[int(key)])
//}

func main() {
	Articles = []Article{
		{
			Title:   "ArticleOne",
			Desc:    "DescriptionOne",
			Content: "ContentOne",
		},
		{
			Title:   "ArticleTwo",
			Desc:    "DescriptionTwo",
			Content: "ContentTwo",
		},
	}
	handleRequests()
}
