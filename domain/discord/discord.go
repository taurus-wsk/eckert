package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/mmcdole/gofeed"
	"strings"
)

func GetNews() {
	// 创建一个新的 gofeed.Parser
	parser := gofeed.NewParser()
	//https://news.google.com/topics/CAAqIQgKIhtDQkFTRGdvSUwyMHZNRGQwTWpFU0FtVnVLQUFQAQ?hl=en-US&gl=US&ceid=US%3Aen
	// 获取订阅源的内容
	feed, err := parser.ParseURL("https://nl.ign.com/feed.xml")
	if err != nil {
		panic(err)
	}

	// 打印订阅源的标题
	fmt.Println("Feed len:", len(feed.Items))

	// 打印订阅源中的文章
	for _, item := range feed.Items {
		if containsFilmOrSeries(item.Categories) {
			fmt.Println("Title:", item.Title)
			fmt.Println("Description:", item.Description)
			fmt.Println("Published Date:", item.Published)
			fmt.Println("Author:", item.Authors[0].Name)
			fmt.Println("Link:", item.Link)
			fmt.Println()
		}
	}

	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//// 读取响应体
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 打印返回值
	//fmt.Println(string(body))
}

func containsFilmOrSeries(categories []string) bool {
	for _, category := range categories {
		if strings.Contains(category, "Films") {
			return true
		}
	}
	return false
}

func News2() {
	Token := "MTE3MTY4MDA1NzcyMzQ2OTkwNQ.G5scQm.jhHggr5zO2Qq5ZLNhVgaCig6jxfC94UGlVh8yY"
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// 添加消息创建事件的处理函数
	//dg.AddHandler(messageCreate)

	// 打开 Discord session
	//err = dg.Open()
	//if err != nil {
	//	fmt.Println("Error opening Discord session: ", err)
	//	return
	//}
	ChannelID := "950974820827398235"
	// 获取频道上的消息
	messages, err := dg.ChannelMessages(ChannelID, 10, "", "", "")
	if err != nil {
		fmt.Println("Error retrieving channel messages: ", err)
		return
	}

	// 遍历消息并处理
	for _, message := range messages {
		// 处理消息
		fmt.Println("Message Content:", message.Content)
		fmt.Println("Author:", message.Author.Username)
		fmt.Println("Timestamp:", message.Timestamp)
		fmt.Println("-----")
	}
}
