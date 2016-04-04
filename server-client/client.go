package main
import "net"
import "fmt"
import "bufio"
import "os"
import "time"

func main() {
  file := os.Args[1]
  port := ":" + os.Args[2]
  m := make(map[string]bool)
  f, err := os.Open(file)
  if err != nil {
    fmt.Println("error opening file ", err)
    os.Exit(1)
  }
  defer f.Close()
  r := bufio.NewReader(f)
  for {
    line, err := r.ReadString('\n')
    // fmt.Print(line)
    m[line] = true
    if err != nil {
      break
  }
}
  conn, err := net.Dial("tcp", "catserver"+port)
  if err != nil {
    fmt.Println("Cannot connect to server:", err.Error())
    os.Exit(1)
  }
  for i := 0; i < 10; i++ {
    _, err = conn.Write([]byte("LINE\n"))
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }
    message, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        println("Server is down:", err.Error())
        os.Exit(1)
    }
    _, ok := m[message]
    if ok {
      fmt.Println("OK")
    } else {
      fmt.Println("MISSING")
    }
    time.Sleep(3000 * time.Millisecond)
  }

}
