package autovpn

import (
	"log"
	"net"
)

func handleConn(connection net.Conn) {
	//write to socket
	//check how podman deals with its socket
}

func StartSocket(sockAddr string) error {
	l, err := net.Listen("unix", sockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()
	for {
		// Accept new connections, dispatching them to echoServer
		// in a goroutine.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go handleConn(conn)
	}
	//will not arrive here
	return nil
}
