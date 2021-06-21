package main

import (
        "encoding/json"
	"net/http"
        "fmt"
        io "io/ioutil"
        "log"
)

type Exchange struct {
        Ask string `json:"ask"`
        Bid string `json:"bid"`
        Code string `json:"code"`
        Codein string `json:"codein"`
        CreatedDate string `json:"create_date"`
        High string `json:"high"`
        Low string `json:"low"`
        Name string `json:"name"`
        PctChange string `json:"pctChange"`
        Timestamp string `json:"timestamp"`
        VarBid string `json:"varBid"`
}

type ExchangeUsdBrl struct {
        Exchange
}

type Response struct {
        Brl ExchangeUsdBrl `json:"USDBRL"`
}

func main() {
	resp, err := http.Get("https://economia.awesomeapi.com.br/last/USD-BRL")
        if err != nil {
                log.Fatal(err)
        }

        defer resp.Body.Close()

        data, _ := io.ReadAll(resp.Body)

        var response Response
        json.Unmarshal(data, &response)

        fmt.Printf("Dólar hoje: %s\n", response.Brl.Bid)
}
