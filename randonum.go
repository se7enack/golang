package main

import (
    "fmt"
    "math/rand"
    "os/exec"
    "strconv"
)

func main() {
    x := rand.Intn(10000)
    cmd := exec.Command("say", strconv.Itoa(x))
    cmd.Output()
    fmt.Println(strconv.Itoa(x))
}
