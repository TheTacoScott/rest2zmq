//
//  Hello World server.
//  Binds REP socket to tcp://*:5555
//  Expects "Hello" from client, replies with "World"
//

package main

import (
  zmq "github.com/pebbe/zmq4"
  "net/http"
  "fmt"
  "runtime"
  "sync"
  "time"
)
var (
  publisher *zmq.Socket
  PubMutex sync.Mutex
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
    if 1 == 1 {
      PubMutex.Lock()
      publisher.Send("taco",0) //zmq socket not native thread safe
      PubMutex.Unlock()
    }
    time.Sleep(30 * time.Millisecond) //wait for results more zmq magic
    fmt.Fprintf(w, "<h1>%s</h1>", r.URL.Path)
}

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU()-1)

  publisher,_ = zmq.NewSocket(zmq.PUB)
  defer publisher.Close()
  publisher.Bind("tcp://*:5555")

  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":7777", nil)
}
