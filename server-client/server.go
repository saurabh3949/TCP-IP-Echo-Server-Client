package main
import "net"
import "fmt"
import "bufio"
import "strings"
import "os"

func main() {
  file := os.Args[1]
  f, err := os.Open(file)
  if err != nil {
    fmt.Println("error opening file ", err)
    os.Exit(1)
  }
  fi, err := f.Stat()
  if err != nil {
    fmt.Println("Cant get file size. Exiting...", err)
    os.Exit(1)
  }
  if fi.Size() <= 1 {
    fmt.Println("File is empty. Exiting..")
    os.Exit(1)
  }
  r := bufio.NewReader(f)
  fmt.Println("Launching server!")
  // file := os.Args[1]
  port := ":" + os.Args[2]
  ln, _ := net.Listen("tcp", port)
  conn, _ := ln.Accept()
  for {
    message, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
      fmt.Println("Client is down. ", err)
      os.Exit(1)
    }
    if message == "LINE\n" {
      line, err := r.ReadString('\n')
      fmt.Println("Read:", line)
      if err != nil {
        f.Close()
        f, _ = os.Open(file)
        r.Reset(f)
        line, err = r.ReadString('\n')
        fmt.Println("Read:", line)
      }
      newmessage := strings.ToUpper(line)
      conn.Write([]byte(newmessage + "\n"))
    }

  }
}
