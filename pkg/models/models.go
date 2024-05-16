package models

type Sources struct {
	Records []string `json:"records"`
}

type Configs struct {
	ZeroXAPIKey string `json:"zeroXAPIKey"`
}

type PriceRequest struct {
	SellToken string `json:"sellToken"`
	BuyToken string `json:"buyToken"`
	SellAmount string `json:"sellAmount"`
}

type PriceResponse struct {
	Price string `json:"price"`
	Value string `json:"value"`
	Gas string `json:"gas"`
	EstimateGas string `json:"estimateGas"`
	Sources []PriceResponseSources `json:"sources"`
}

type PriceResponseSources struct {
	Name string `json:"name"`
	Proportion string `json:"proportion"`
}

type QuoteRequest struct {
	SellToken string `json:"sellToken"`
	BuyToken string `json:"buyToken"`
	SellAmount string `json:"sellAmount"`
	BuyAmount string `json:"buyAmount"`
	TakerAddress string `json:"takerAddress"`
	FeeRecipient string `json:"feeRecipient"`
	BuyTokenPercentageFee string `json:"buyTokenPercentageFee"`
}

type QuoteResponse struct {
	Price string `json:"price"`
	To string `json:"to"`
	Data string `json:"data"`
	Value string `json:"value"`
	EstimateGas string `json:"estimateGas"`
	AllowanceTarget string `json:"allowanceTarget"`
}