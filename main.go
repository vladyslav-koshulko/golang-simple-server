package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	server_ip := os.Getenv("SERVER_IP")
	server_port := os.Getenv("SERVER_PORT")
	httpServer(server_ip, server_port)
}

func httpServer(domain, port string) {
	fmt.Println("Server started")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			http.Error(writer, "Method is not support", http.StatusMethodNotAllowed)
			slog.Error("Mothod is not support")
		}
		switch request.URL.Path {

		case "/":
			_, _ = fmt.Fprintf(writer, "<h1>Main page</h1></br><a href=\"http://localhost:8080/about\">About</a>")
		case "/about":
			_, _ = fmt.Fprintf(writer, "<h1>About page</h1></br><a href=\"http://localhost:8080/\">About</a>")

		}

	})

	if err := http.ListenAndServe(domain+":"+port, nil); err != nil {
		slog.Error("Server was shutdown")

	}
}
