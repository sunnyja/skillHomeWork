package main
 
import (
    "fmt"
    "log"
)
 
func main() {
    var n string
    fmt.Print("Введите какие-нибудь данные: ")
    _, err := fmt.Scan(&n)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Вы ввели следующие данные: %v\n", n)
}