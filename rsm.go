package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {

    fmt.Println("###############################################")
    fmt.Println("########## Red Shirt Manager v0.0.2 ###########")

    var inputName string
    var inputItemNumber int
    var countFalseLogin int
    var selectedUser []string
    var menuItem int
    var userExistCache []string = getDirContent()

    fmt.Println("Enter your name.")
    fmt.Scanln(&inputName)

    for stringInSlice(inputName, getDirContent()) == false {
        countFalseLogin++
        if countFalseLogin >= 3 {
            fmt.Println("No, go away.")
            return
        }
        fmt.Println("You either mistyped or aren't authorized. Try again.")
        fmt.Scanln(&inputName)
        if countFalseLogin == 1 && inputName == "remindme" {
            fmt.Println("Let me help you there...")
            listDirContent()
            fmt.Println("Now, try again.")
            fmt.Scanln(&inputName)
        }
    }

    selectedUser = getUserData(inputName)

    printUserData(selectedUser)

// get list of accounts
    for _, f := range userExistCache {
        fmt.Println(strconv.Itoa(menuItem) +" " +f)
        menuItem++
    }

    fmt.Println("Choose an account by its number.")
    fmt.Scanln(&inputItemNumber)
    for inputItemNumber >= len(userExistCache) {
        fmt.Println("Accountnumber doesn't exist.")
        fmt.Scanln(&inputItemNumber)
    }
    selectedUser = getUserData(userExistCache[inputItemNumber])

    printUserData(selectedUser)

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

func printUserData(userData []string) {
    fmt.Println("#####################")
    fmt.Println("Name: " +userData[0])
    fmt.Println("Age: " +userData[1])
    fmt.Println("Gender: " +userData[2])
    fmt.Println("#####################")
}
