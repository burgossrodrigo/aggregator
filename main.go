package main

import (
	zeroxservice "aggregator/pkg/zerox"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router := gin.Default()
	router.GET("/zerox/sources/:chain", getZeroxSource)
	router.GET("/zerox/price/:chain/:sellToken/:buyToken/:sellAmount", getZeroxPrice)
	router.GET("/zerox/quote/:chain/:sellToken/:buyToken/:sellAmount/:takerAddress/:feeRecipient/:buyTokenPercentageFee", getZeroxQuote)
	router.Run(":8080")
}

func getZeroxSource(c *gin.Context) {
	chain := c.Param("chain")
    sources, err := zeroxservice.FetchSources(chain)
    if err != nil {
        fmt.Println(err, "error from getZeroxSource")
        c.JSON(500, gin.H{"error": err})
    } else {
        c.JSON(200, sources)
    }
}

func getZeroxPrice(c *gin.Context) {
	chain := c.Param("chain")
	sellToken := c.Param("sellToken")
	buyToken := c.Param("buyToken")
	sellAmount := c.Param("sellAmount")
	Price, err := zeroxservice.FetchPrice(chain, sellToken, buyToken, sellAmount)
	if err != nil {
		fmt.Println(err, "error from getZeroxPrice")
		c.JSON(500, gin.H{"error": err})
	} else {
		c.JSON(200, Price)
	}
}

func getZeroxQuote(c *gin.Context) {
	chain := c.Param("chain")
	sellToken := c.Param("sellToken")
	buyToken := c.Param("buyToken")
	sellAmount := c.Param("sellAmount")
	takerAddress := c.Param("takerAddress")
	feeRecipient := c.Param("feeRecipient")
	buyTokenPercentageFee := c.Param("buyTokenPercentageFee")
	Quote, err := zeroxservice.FetchQuote(chain, sellToken, buyToken, sellAmount, takerAddress, feeRecipient, buyTokenPercentageFee)
	if err != nil {
		fmt.Println(err, "error from getZeroxQuote")
		c.JSON(500, gin.H{"error": err})
	} else {
		c.JSON(200, Quote)
	}
}