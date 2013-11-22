package main

import (
    "fmt"
    "io/ioutil"
)

func main() {

    fmt.Println("###############################################")
    fmt.Println("########## Red Shirt Manager v0.0.1 ###########")
    fmt.Println("###############################################")

    listDirContent()

    var inputName string

    fmt.Println("Enter your name, if you have a red shirt.")
    fmt.Scanln(&inputName)

    for stringInSlice(inputName, getDirContent()) == false {
        fmt.Println("You either mistyped or don't own a red shirt. Try again.")
        fmt.Scanln(&inputName)
    }


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
