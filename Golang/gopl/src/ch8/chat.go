package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// broadcaster
// 记录当前连接的客户端的集合
type client chan<- string

type message struct {
	text string // 消息
	user client
}

// 事件
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan message)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go chatHandleConn(conn)
	}
}

func broadcaster() {
	// 其实就是当成集合来用
	clients := make(map[client]struct{})
	for {
		select {
		case msg := <-messages:
			// 向每个 client 发送消息
			for cli := range clients {
				if msg.user == cli {
					continue
				}
				cli <- msg.text
			}
		case cli := <-entering:
			clients[cli] = struct{}{}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// 输入和发送
func chatHandleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- message{who + " has arrived", ch}
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- message{who + ": " + input.Text(), ch}
	}

	// 退出
	leaving <- ch
	messages <- message{who + " has left", ch}
	conn.Close()
}

// 信息打印
func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		//goland:noinspection GoUnhandledErrorResult
		fmt.Fprintln(conn, msg)
	}
}
