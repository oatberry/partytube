//                    _         _         _
//   _ __   __ _ _ __| |_ _   _| |_ _   _| |__   ___
//  | '_ \ / _` | '__| __| | | | __| | | | '_ \ / _ \
//  | |_) | (_| | |  | |_| |_| | |_| |_| | |_) |  __/
//  | .__/ \__,_|_|   \__|\__, |\__|\__,_|_.__/ \___|
//  |_|                   |___/
//
//      an internet video jukebox using mpv + youtube-dl
//      author: thomas berryhill (c) 2018
//      license: GPLv3

package main

import (
	"bufio"
	"log"
	"net"
	"os/exec"
	"strings"
)

func main() {
	queue := make(chan string)
	go mpvListen(queue)
	tcpListen(queue)
}

func mpvListen(rx <-chan string) {
	for message := range rx {
		cmd := exec.Command("mpv", "--fullscreen", message)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("mpv returned an error: %v\n", err)
			log.Printf("output: \n%s", string(output))
		}
	}
}

func tcpListen(tx chan<- string) {
	listener, err := net.Listen("tcp", ":2338")
	if err != nil {
		log.Println(err)
		return
	}

	defer listener.Close()
	log.Println("listening on", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		go handleConnection(conn, tx)
	}
}

func handleConnection(conn net.Conn, tx chan<- string) {
	defer conn.Close()

	bufReader := bufio.NewReader(conn)
	remoteIP := conn.RemoteAddr().String()

	log.Println("info:", remoteIP, "connected")
	defer log.Println("info:", remoteIP, "disconnected")

loop:
	for {
		conn.Write([]byte("partytube > "))
		line, err := bufReader.ReadString('\n')
		if err != nil {
			log.Println("info:", remoteIP, err)
			return
		}

		line = line[:len(line)-1]

		if strings.HasPrefix(line, "http://") || strings.HasPrefix(line, "https://") {
			conn.Write([]byte("queued\n"))
			log.Println("queued:", line)
			tx <- line
			continue
		} else if len(line) == 0 {
			continue
		}

		switch line {
		case "ping":
			conn.Write([]byte("pong\n"))
		case "quit":
			break loop
		default:
			conn.Write([]byte("unrecognized input\n"))
		}
	}

	conn.Write([]byte("goodbye.\n"))
}
