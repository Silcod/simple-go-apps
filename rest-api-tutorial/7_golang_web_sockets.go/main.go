package main

/*
  Working with Web Sockets in go: Web sockets provide us with an alternate option to communicate between application as opposed to the standard
  RESTful APIs solution. With sockets we can do cool things like real-time communications between thousands of different clients without having
  to incure the expense of thousands of additional REST API calls hitting your servers every minute. Real life example: Chat servers.
  To implement web sockets in golang, there are a number of different options to choose from:
    1. one of the prominent libraries is socket.io
 */

 import (
   "fmt"
   "log"
   "net/http"

   socketio "github.com/googollee/go-socket.io"
 )
func main() {
  fmt.Println("Hello World")

  server, err := socketio.NewServer(nil)
  if err != nil {
    log.Fatal(err)
  }

  server.On("connection", func(so socketio.Socket) {
    log.Println("New connection")

    so.Join("chat")

    so.On("chat message", func(msg string) {
      log.Println("Message Received From Client: ", + msg)
      so.BroadcastTo("chat", "chat message", msg)
    })
  })

  fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)

  http.Handle("/socket.io/", server)
  log.Fatal(http.ListenAndServe(":5000", nil))
}
