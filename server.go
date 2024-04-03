package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Server is running")
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go HandleServerConnection(clientConn)

	}

}

func HandleServerConnection(client net.Conn) {
	// defer client.Close()
	var size uint32
	err := binary.Read(client, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}
	bytMsg := make([]byte, size)
	_, err = client.Read(bytMsg)
	if err != nil {
		panic(err)
	}
	strMsg := string(bytMsg)
	fmt.Printf("Received: %s\n", strMsg)

	var reply string
	if strings.HasSuffix(strMsg, ".zip") {
		reply = "File has been Received"
	} else if strings.Contains(strMsg, ".") {
		reply = "only zip file can be uploaded"
	} else {
		reply = "message has been received"
	}

	binary.Write(client, binary.LittleEndian, uint32(len(reply)))
	if err != nil {
		panic(err)
	}
	client.Write([]byte(reply))
	if err != nil {
		panic(err)
	}

}
