package main

import (
    "bufio"
    "fmt"
    "net/http"
)

func main() {
    // document.querySelectorAll(".dfx-rate[data-change-scale=\"2\"")[7].dataset.value
    resp, err := http.Get("https://www.dailyfx.com/forex-rates#indices")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 50; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
