package main

import (
  "fmt"

  "github.com/go-vgo/robotgo"
)

func main() {
  robotgo.TypeStr("Hello World")
  robotgo.TypeStr("?????")
  robotgoHello.TypeString("?????")
  

  robotgo.TypeStr("Hi galaxy. ???????.")
  robotgo.Sleep(1)

  // ustr := uint32(robotgo.CharCodeAt("Test", 0))
  // robotgo.UnicodeType(ustr)

  robotgo.KeyTap("enter")
  // robotgo.TypeString("en")
  robotgo.KeyTap("i", "alt", "command")

  arr := []string{"alt", "command"}
  robotgo.KeyTap("i", arr)

  robotgo.WriteAll("Test")
  text, err := robotgo.ReadAll()
  if err == nil {
    fmt.Println(text)
  }
}