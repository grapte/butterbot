package types

// RecordsRequest gets multiple transaction requests after a timestamp in accordance with WhaleAlert API.
// Struct tags are used by the package "github.com/go-playground/form" to encode
type RecordsRequest struct {
	Start    int64  `form:"start"` // API: Start is a required field in unix timestamp
	End      int64  `form:"end,omitempty"`
	Cursor   string `form:"cursor,omitempty"`    // API: Pagination key from the previous response. Recommended when retrieving transactions in intervals.
	MinValue int    `form:"min_value,omitempty"` // API: Minimum transaction value ($500k for Free)
	Limit    int    `form:"limit,omitempty"`     // API: Maximum number of results returned. Default 100, max 100.
	Currency string `form:"currency,omitempty"`  // API: Returns transactions for a specific currency code. Returns all currencies by default.
}

type RecordsResponse struct {
	Result       string `json:"result"`
	Cursor       string `json:"cursor"`
	Count        int    `json:"count"`
	Transactions []struct {
		Blockchain      string `json:"blockchain"`
		Symbol          string `json:"symbol"`
		TransactionType string `json:"transaction_type"`
		Hash            string `json:"hash"`
		From            struct {
			Address   string `json:"address"`
			OwnerType string `json:"owner_type"`
		} `json:"from"`
		To struct {
			Address   string `json:"address"`
			Owner     string `json:"owner"`
			OwnerType string `json:"owner_type"`
		} `json:"to"`
		Timestamp        int     `json:"timestamp"`
		Amount           float64 `json:"amount"`
		AmountUsd        float64 `json:"amount_usd"`
		TransactionCount int     `json:"transaction_count"`
	} `json:"transactions"`
}

//type SendDiscordMessage func(rec *RecordsResponse)
