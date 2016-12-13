package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/Microsoft/go-winio"
)

func setupPipe(name string) (net.Listener, error) {
	winioPipeConfig := winio.PipeConfig{
		SecurityDescriptor: "S:(ML;;NW;;;LW)D:(A;;0x12019f;;;WD)",
		MessageMode:        true,
		InputBufferSize:    4096,
		OutputBufferSize:   4096,
	}

	listener, err := winio.ListenPipe(name, &winioPipeConfig)

	if err != nil {
		fmt.Println("Error listening to the pipe")
		return nil, err
	}

	return listener, err
}

func handleConnection(conn net.Conn) {
	fmt.Println("Got connection!")

	var buffer [4096]byte
	conn.Read(buffer[0:])

	response := []byte("{\"error\": null}")
	conn.Write(response)
}

func main() {
	var name string
	flag.StringVar(&name, "name", "mock", "name for the pipe")

	address := "//./pipe/" + name

	listener, err := setupPipe(address)

	if err != nil {
		fmt.Println("Error while setting up pipe")
		return
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Got connection, but an error occured!")
			continue
		}

		go handleConnection(conn)
	}
}
