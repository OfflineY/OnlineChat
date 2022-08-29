package modules

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

// json 结构体
type Product struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

// 客户端 获取返回数据模块
func GetChat(url string) {

	// url := "ws://localhost:8080/socket"
	// url 获取 ini 里面的内容
	// 连接 url 里的服务器
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("连接服务器失败:", err)
		GetOut()
	}

	log.Printf("连接服务器成功")

	// 设置自己的用户名
	log.Print("请输入你的用户名: ")
	inputReader := bufio.NewReader(os.Stdin)
	clientName, _ := inputReader.ReadString('\n')

	// 新用户加入全群聊通知
	err = c.WriteMessage(websocket.TextMessage, []byte("【系统消息】 新的用户加入："+clientName))
	if err != nil {
		log.Printf("发生错误：%s", err)
		GetOut()
	}

	// 将新的用户加入切片
	// addUser(strings.Join(strings.Fields(clientName), ""))

	// clientName 后面会跟一个 \t，使用 join 和 fields
	go sendMsg(c, strings.Join(strings.Fields(clientName), ""))

	// 重复监听服务端发送过来的数据
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		// log形式打印出来
		log.Printf("%s", message)
	}
}

// 客户端 发送消息模块
func sendMsg(c *websocket.Conn, clientName string) {
	for {
		// fmt.Print("Send> ")
		inputReader := bufio.NewReader(os.Stdin)
		sendMsg, _ := inputReader.ReadString('\r')

		sendMsg = strings.Join(strings.Fields(sendMsg), "")

		// 简化服务端，在客户端进行数据处理
		// 拼贴用户名和发送的消息
		p := &Product{
			Name: clientName,
			Msg:  sendMsg,
		}
		jsonP, _ := json.Marshal(p)
		// msg := gin.H{
		// 	"name":clientName,
		// 	"msg":sendMsg,
		// }
		// msg := "【" + clientName + "】" + sendMsg
		// msg := sendMsg
		// fmt.Print(len(clientName))
		// fmt.Print(msg)

		// 处理完成发送至服务端
		err := c.WriteMessage(websocket.TextMessage, jsonP)
		if err != nil {
			fmt.Println(err)
		}
	}
}
