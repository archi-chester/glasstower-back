package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

// слушаем порт
func listener(port string) {
	// начинаем слушать порт
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	// обработка соединения
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(conn net.Conn) {
			// вход в функцию коннекта
			fmt.Println("Point")

			// read the length prefix
			prefix := make([]byte, 4)
			_, err = io.ReadFull(conn, prefix)

			length := binary.BigEndian.Uint32(prefix)
			// verify length if there are restrictions

			message := make([]byte, int(length))
			_, err = io.ReadFull(conn, message)
			// вывели в консоль
			println(string(decrypt(message)))
			// // Echo all incoming data.
			// io.Copy(conn, conn)
			// // Читаем данные из порта
			// status, err := bufio.NewReader(conn).ReadString('\n')
			// if err != nil {
			// 	// handle error
			// 	fmt.Println(err)
			// }
			// fmt.Println(status)
			// Shut down the connection.
			conn.Close()
		}(conn)
	}

}

//  посылаем сообщение
func sender(port string, message []byte) {
	// создаем сендера
	conn, err := net.Dial("tcp", "127.0.0.1:"+port)
	if err != nil {
		// handle error
	}
	// create the length prefix
	prefix := make([]byte, 4)
	binary.BigEndian.PutUint32(prefix, uint32(len(message)))

	// write the prefix and the data to the stream (checking errors)
	_, err = conn.Write(prefix)
	_, err = conn.Write(message)
}
