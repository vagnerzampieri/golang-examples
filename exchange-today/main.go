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
        VarBid int64 `json:"varBid"`
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

        emoji := func() string {
                if response.Brl.VarBid > 0 {
                        return ":arrow_up_small:"
                } else if response.Brl.VarBid < 0 {
                        return ":arrow_down_small:"
                }
                return ":black_square_for_stop:"
        }()

        fmt.Printf("Dólar hoje: %s %s\n", response.Brl.Bid, emoji)
}
