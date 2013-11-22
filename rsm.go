package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {

    fmt.Println("###############################################")
    fmt.Println("########## Red Shirt Manager v0.0.1 ###########")
    fmt.Println("###############################################")

    listDirContent()

    var inputName string
    var countFalseLogin int
    var selectedUser []string

    fmt.Println("Enter your name, if you have a red shirt.")
    fmt.Scanln(&inputName)

    for stringInSlice(inputName, getDirContent()) == false {
        countFalseLogin++
        if countFalseLogin >= 3 {
            fmt.Println("Nope.")
            return
        }
        fmt.Println("You either mistyped or don't own a red shirt. Try again.")
        fmt.Scanln(&inputName)
    }

    selectedUser = getUserData(inputName)

    fmt.Println("Name: " +selectedUser[0])
    fmt.Println("Age: " +selectedUser[1])
    fmt.Println("Gender: " +selectedUser[2])
}

func listDirContent() {
    files, _ := ioutil.ReadDir("/tmp/astronauts/")
    for _, f := range files {
        fmt.Println(f.Name())
    }
}

func getDirContent() []string {
    var rsNames []string
    files, _ := ioutil.ReadDir("/tmp/astronauts/")
    for _, f := range files {
        rsNames = append(rsNames, f.Name())
    }
    return rsNames
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func getUserData(user string) []string {
    content, err := ioutil.ReadFile("/tmp/astronauts/"+user)
    if err != nil {
        fmt.Println("Error: ")
        fmt.Println(err)
    }
    csvIndex := strings.Split(string(content), ",")
    return csvIndex
}
