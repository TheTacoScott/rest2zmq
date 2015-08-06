//
//  Request-reply client.
//  Connects REQ socket to tcp://localhost:5559
//  Sends "Hello" to server, expects "World" back
//

package main

import (
  zmq "github.com/pebbe/zmq4"
  "fmt"
)

func main() {
  requester, _ := zmq.NewSocket(zmq.REQ)
  defer requester.Close()
  requester.Connect("tcp://localhost:5555")
  major, minor, patch := zmq.Version()
  fmt.Printf("Current 0MQ version is %d.%d.%d\n", major, minor, patch)
  for request := 0; request < 10; request++ {
    requester.Send("Hello", 0)
    reply, _ := requester.Recv(0)
    fmt.Printf("Received reply %d [%s]\n", request, reply)
  }
}
