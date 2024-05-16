## Integrating 0x api with Golang!

### GET /zerox/sources/:chain

This route allow the user to retrieve the protocols 0x will be interacting according to the chain name.

```curl

curl http://localhost:8080/zerox/sources

```

```json
{"records":["MultiHop","Uniswap_V3"]}
```

### GET /zerox/price/:chain/:sellToken/:buyToken:sellAmount

This route allow the user to retrieve from 0x token price along with a couple of data.

```curl
http://localhost:8080/zerox/price/ethereum/0xdAC17F958D2ee523a2206206994597C13D831ec7/0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE/1000
```

```json
{
"price":"40234.034344676906158",
"value":"0",
"gas":"151000",
"estimateGas":"",
"sources":[
{
"name":"0x",
"proportion":"0"
},
{
"name":"Uniswap",
"proportion":"0"
},
{
"name":"Uniswap_V2",
"proportion":"1"
},
{
"name":"Curve",
"proportion":"0"
},
{
"name":"Balancer",
"proportion":"0"
},
{
"name":"Balancer_V2",
"proportion":"0"
},
{
"name":"BancorV3",
"proportion":"0"
},
{
"name":"SushiSwap",
"proportion":"0"
},
{
"name":"DODO",
"proportion":"0"
},
{
"name":"DODO_V2",
"proportion":"0"
},
{
"name":"CryptoCom",
"proportion":"0"
},
{
"name":"Lido",
"proportion":"0"
},
{
"name":"MakerPsm",
"proportion":"0"
},
{
"name":"KyberDMM",
"proportion":"0"
},
{
"name":"Uniswap_V3",
"proportion":"0"
},
{
"name":"Curve_V2",
"proportion":"0"
},
{
"name":"ShibaSwap",
"proportion":"0"
},
{
"name":"Synapse",
"proportion":"0"
},
{
"name":"Synthetix",
"proportion":"0"
},
{
"name":"Aave_V2",
"proportion":"0"
},
{
"name":"Compound",
"proportion":"0"
},
{
"name":"KyberElastic",
"proportion":"0"
},
{
"name":"Maverick_V1",
"proportion":"0"
},
{
"name":"PancakeSwap_V3",
"proportion":"0"
}
]
}
```

### GET /zerox/price/:chain/:sellToken/:buyToken/:sellAmount/:takerAddress/:feeRecipient/:feePercentage

This routes returns raw unsigned tx data to be signed from your preferable blockchain client.
Remember that the `takerAddress` needs to give allowance to 0x route first!