package main

import (
  "bufio"
  "crypto/sha1"
  "flag"
  "fmt"
  "os"
)

type info struct {
  name, pseudo, mail string
}

var namePtr = flag.String("name", "", "Your name")
var pseudoPtr = flag.String("pseudo", "", "Your pseudo")
var mailPtr = flag.String("mail", "", "Your mail")

func main()  {
  flag.Parse()

  infos := getInfos(namePtr, pseudoPtr, mailPtr)

  concatenatedInputs := concatenate(infos, ":")

  hash := getHash(concatenatedInputs)
  fmt.Printf("%x\n", hash)
}

func concatenate(infos info, separator string) string {
  return infos.name + separator + infos.pseudo + separator + infos.mail;
}

func getHash(input string) []byte {
  hash := sha1.New()
  hash.Write([]byte(input))
  return hash.Sum(nil)
}

func getInfos(namePtr, pseudoPtr, mailPtr *string) info {
  if *namePtr == "" {
    *namePtr = askFor("Name: ")
  }

  if *pseudoPtr == "" {
    *pseudoPtr = askFor("Pseudo: ")
  }

  if *mailPtr == "" {
    *mailPtr = askFor("Mail: ")
  }

  return info{
    *namePtr,
    *pseudoPtr,
    *mailPtr}
}

func askFor(message string) string {
  input := ""
  for input == "" {
    fmt.Print(message)
    stdinReader := bufio.NewReader(os.Stdin)
    line, err := stdinReader.ReadString('\n')
    if nil != err {
      panic(err)
    }
    if len(line) > 1 {
      input = line[:len(line) - 1]
    } else {
      input = ""
    }
  }
  return input
}
