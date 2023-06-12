package services

import (
	"HolaaPlanet/configs"
	"HolaaPlanet/entity"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ReplyMsg struct {
	Code    int    `json:"code"`
	Content string `json:"content"`
}

type Client struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}

type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}

type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client), // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	Broadcast:  make(chan *Broadcast),
	Register:   make(chan *Client),
	Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}

func Handler(c *gin.Context) {
	uid := c.Query("user_id")
	sendId := c.Query("send_id")
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // CheckOrigin解决跨域问题
			return true
		}}).Upgrade(c.Writer, c.Request, nil) // 升级成ws协议
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	//创建一个用户实例
	client := &Client{
		ID:     uid,    //1->2
		SendID: sendId, //2->1
		Socket: conn,
		Send:   make(chan []byte),
	}
	//用户注册到用户管理上
	Manager.Register <- client
	go client.Read()
	go client.Write()
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PongHandler()
		msgType, content, err := c.Socket.ReadMessage() // 使用 ReadMessage 方法接收消息
		if err != nil {
			log.Println("read error:", err)
			return
		}
		if msgType == websocket.TextMessage { // 如果消息类型是纯文本消息
			text := string(content) // 将 []byte 类型的 message 转换为 string 类型
			//发送消息
			log.Println(c.ID, "发送消息", text)
			Manager.Broadcast <- &Broadcast{
				Client:  c,
				Message: []byte(text), //发送过来的消息
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Println(c.ID, "接受消息:", string(message))
			replyMsg := ReplyMsg{
				Code:    1,
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(replyMsg)
			err := c.Socket.WriteMessage(websocket.TextMessage, msg) //发送消息到客户端
			if err != nil {
				log.Println("write error:", err)
				return
			}
		}
	}
}

func (manager *ClientManager) Start() {
	for {
		fmt.Println("<---监听管道通信--->")
		select {
		case conn := <-Manager.Register:
			log.Printf("建立新连接: %v", conn.ID)
			Manager.Clients[conn.ID] = conn //把连接放到用户管理上
			replyMsg := ReplyMsg{
				Code:    1,
				Content: "已连接至服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister: //断开连接
			log.Printf("连接失败:%v", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				replyMsg := &ReplyMsg{
					Code:    0,
					Content: "连接已断开",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)                 //关闭发送通道
				delete(Manager.Clients, conn.ID) //删除连接
			}
		case broadcast := <-Manager.Broadcast:
			message := broadcast.Message
			sendId := broadcast.Client.SendID
			flag := false //默认对方不在线
			for id, conn := range Manager.Clients {
				if id != sendId {
					continue
				}
				select {
				case conn.Send <- message:
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
			id := broadcast.Client.ID
			if flag {
				log.Println("对方在线应答")
				replyMsg := &ReplyMsg{
					Code:    1,
					Content: "对方在线应答",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				if ok := InsertMsg(id, sendId, string(message)); ok {
					log.Println("insert_success")
				} else {
					log.Println("insert_fail")
				}
			} else {
				log.Println("对方不在线")
				replyMsg := ReplyMsg{
					Code:    0,
					Content: "对方不在线应答",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
			}
		}
	}
}

func InsertMsg(id string, sendId string, context string) bool {
	result := configs.DB.Find(&entity.SendUserMessage{})
	id1, _ := strconv.Atoi(id)
	sendId1, _ := strconv.Atoi(sendId)
	var maxvalue int
	configs.DB.Table("send_user_messages").Select("MAX(message_id)").Row().Scan(&maxvalue)
	if result.Error != nil {
		log.Print(result.Error)
		return false
	} else {
		if result.RowsAffected == 0 {
			user := entity.SendUserMessage{
				SendUser:      id1,
				ReceiveUser:   sendId1,
				Message:       context,
				SendTime:      time.Now(),
				UserMessageID: 1,
			}
			createRe := configs.DB.Create(&user)
			if createRe.Error != nil {
				log.Print(createRe.Error)
				return false
			}
		} else {
			user := entity.SendUserMessage{
				SendUser:      id1,
				ReceiveUser:   sendId1,
				Message:       context,
				SendTime:      time.Now(),
				UserMessageID: maxvalue + 1,
			}
			createRe := configs.DB.Create(&user)
			if createRe.Error != nil {
				log.Print(createRe.Error)
				return false
			}
		}
	}
	return true
}
