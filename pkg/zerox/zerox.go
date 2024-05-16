package zerox

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	models "aggregator/pkg/models"
	config "aggregator/pkg/config"
) 

func FetchSources(chain string) (models.Sources, error) {

	var url string
	var sources models.Sources
	cfg := config.LoadEnv()

	switch chain {
		case "ethereum":
			url = "https://api.0x.org/swap/v1/sources"
		default:
			url = fmt.Sprintf("https://%s.api.0x.org/swap/v1/sources", chain)
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("0x-api-key", cfg.ZeroXAPIKey)
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)


	client := fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		fmt.Println(err, "error from client.Do")
		return models.Sources{}, err
	}

	if err := json.Unmarshal(resp.Body(), &sources); err != nil {
		fmt.Println(err, "error from json.Unmarshal")
		return models.Sources{}, err
	}

	return sources, err
}

func FetchPrice(chain string, sellToken string, buyToken string, sellAmount string) (models.PriceResponse, error) {

	var url string
	var responsePrice models.PriceResponse
	cfg := config.LoadEnv()

	switch chain {
		case "ethereum":
			url = fmt.Sprintf("https://api.0x.org/swap/v1/price?sellToken=%s&buyToken=%s&sellAmount=%s",
			sellToken, buyToken, sellAmount)
		default:
			url = fmt.Sprintf("https://%s.api.0x.org/swap/v1/price?sellToken=%s&buyToken=%s&sellAmount=%s", 
			chain, sellToken, buyToken, sellAmount)
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("0x-api-key", cfg.ZeroXAPIKey)
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		fmt.Println(err, "error from client.Do")
		return models.PriceResponse{}, err
	}

	fmt.Println(string(resp.Body()), "body")

	if err := json.Unmarshal(resp.Body(), &responsePrice); err != nil {
		fmt.Println(err, "error from json.Unmarshal")
		return models.PriceResponse{}, err
	}

	return responsePrice, err
}

func FetchQuote(chain string, sellToken string, buyToken string, sellAmount string, takerAddress string, feeRecipient string, buyTokenPercentageFee string) (models.QuoteResponse, error) {
	// Quote, err := zeroxservice.FetchQuote(chain, sellToken, buyToken, sellAmount, takerAddress, feeRecipient, buyTokenPercentageFee)
	
	var url string
	var quoteResponse models.QuoteResponse
	cfg := config.LoadEnv()

	switch chain {
		case "ethereum":
			url = fmt.Sprintf("https://api.0x.org/swap/v1/quote?buyToken=%s&sellToken=%s&sellAmount=%s&takerAddress=%s&feeRecipient=%s&buyTokenPercentageFee=%s",
			sellToken, buyToken, sellAmount, takerAddress, feeRecipient, buyTokenPercentageFee)
		default:
			url = fmt.Sprintf("https://%s.api.0x.org/swap/v1/quote?buyToken=%s&sellToken=%s&sellAmount=%s&takerAddress=%s&feeRecipient=%s&buyTokenPercentageFee=%s",
			chain, buyToken, sellToken, sellAmount, takerAddress, feeRecipient, buyTokenPercentageFee)
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("0x-api-key", cfg.ZeroXAPIKey)
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		fmt.Println(err, "error from client.Do")
		return models.QuoteResponse{}, err
	}

	if err := json.Unmarshal(resp.Body(), &quoteResponse); err != nil {
		fmt.Println(err, "error from json.Unmarshal")
		return models.QuoteResponse{}, err
	}

	fmt.Println(string(resp.Body()), "body")
	return quoteResponse, err
}


