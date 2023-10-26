package whalealert

import (
	"butterbot/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/form"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	ctx                context.Context
	apiKey             string
	lastCursor         string
	sendDiscordMessage func(rec *types.RecordsResponse)
)

func Serve(c context.Context, key string, callback func(rec *types.RecordsResponse)) {
	ctx = c
	apiKey = key
	sendDiscordMessage = callback

	// get the initial cursor at the end of 100 transactions.
	getResponse(EndpointTransactions)
	go pullAPI()
}

var (
	EndpointWA           = "https://api.whale-alert.io"
	EndpointStatus       = EndpointWA + "/v1/status"
	EndpointTransactions = EndpointWA + "/v1/transactions"
)

func pullAPI() {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			rec := getResponse(EndpointTransactions)
			//fmt.Printf("%v\n", rec)
			if rec.Count == 0 {
				fmt.Printf("No new transactions\n")
			} else {
				switch rec.Result {
				case "success":
					go sendDiscordMessage(rec)
				case "error":
					fmt.Printf("Error: RESPONSE ERROR")
				}
			}
			time.Sleep(7 * time.Second)
		}
	}
}

func NewRequest() *types.RecordsRequest {
	ret := types.RecordsRequest{
		Start:    time.Now().Add(-time.Hour).Unix() + 1,
		Cursor:   lastCursor,
		MinValue: 500000,
		Limit:    100,
	}

	return &ret
}

var encoder = form.NewEncoder()

func makeURL(endpoint string, req *types.RecordsRequest) (u *url.URL) {
	req.Cursor = lastCursor

	u, err := url.Parse(endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}

	query, err := encoder.Encode(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	u.RawQuery = query.Encode()
	return
}

func makeDefaultURL(endpoint string) (u *url.URL) {
	req := NewRequest()

	u = makeURL(endpoint, req)
	return
}

func fetchTransactions(u *url.URL) (rec *types.RecordsResponse) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-WA-API-KEY", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)

	//fmt.Printf("%v\n", string(b))
	rec = &types.RecordsResponse{}
	json.Unmarshal(b, rec)
	if rec.Result == "error" {
		fmt.Printf("Json Unmarshal Error: %s\n", b)
		return
	}
	lastCursor = rec.Cursor

	return
}

func getResponse(endpoint string) (rec *types.RecordsResponse) {
	u := makeDefaultURL(endpoint)
	rec = fetchTransactions(u)
	return
}
