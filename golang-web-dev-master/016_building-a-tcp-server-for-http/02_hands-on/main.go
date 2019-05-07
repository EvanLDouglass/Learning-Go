package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	lsn, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("problem with listen", err)
	}

	for {
		conn, err := lsn.Accept()
		if err != nil {
			log.Println("problem accepting connection", err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	html := `<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8">
		<title>Test</title>
	</head>
	<body>
		<h1>Host: %s</h1>
	</body>
	</html>`

	scan := bufio.NewScanner(conn)
	fmt.Println("Printing Request:")
	for scan.Scan() {
		ln := scan.Text()
		fmt.Println(ln)
		xln := strings.Fields(ln)
		if len(xln) > 0 && xln[0] == "Host:" {
			html = fmt.Sprintf(html, xln[1])
		}

		if ln == "" {
			break
		}
	}

	// Respond
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, html)
}
