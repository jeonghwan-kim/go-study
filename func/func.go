package main

import (
    "fmt"
    "strconv"
)

func displayInt(s string) {
    if v, err := strconv.Atoi(s); err != nil {
        fmt.Printf("%s", err)
    } else {
        fmt.Printf("정수 값은 %d입니다.\n", v)
    }

}

func main() {
    displayInt("two")
    displayInt("2")
}
