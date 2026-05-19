package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		splittedArg := strings.Split(arg, "=")
		addrs := splittedArg[1]

		conn, err := net.Dial("tcp", addrs)

		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		mustCopy(os.Stdout, conn)
	}
}

func mustCopy(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}
