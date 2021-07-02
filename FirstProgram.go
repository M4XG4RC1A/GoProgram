package main


import (
    "bufio"
    "fmt"
    "os"
    "strings"
  )

func say(pharase string) string{
    pharase = "You tell me to say: "+pharase
    return pharase
}

func main() {
    fmt.Printf("Hello, tell me something to say:\n")
    reader := bufio.NewReader(os.Stdin)
    readStr, _ := reader.ReadString('\n')
    readStr = strings.Replace(readStr, "\n", "", -1)
    tell := say(readStr)
    fmt.Printf(tell+"\n\nended\n\n")
}