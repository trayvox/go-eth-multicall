# go-eth-multicall
Minimal golang ethereum multicall implementation

The implementation only supports mainnet but can it is just a matter of pointing to another RPC and address to support other networks.

The module exports a `EthMultiCaller` struct that is bound to a client and is used to make the queries.
```go
type EthMultiCaller struct {
    Signer          *bind.TransactOpts
    Client          *ethclient.Client
    Abi             abi.ABI
    ContractAddress common.Address
}
```

Requests/reads are to be defined using the `Call` struct, and its response will be a `CallResponse`
```go
type Call struct {
    Name     string         `json:"name"`
    Target   common.Address `json:"target"`
    CallData []byte         `json:"call_data"`
}

type CallResponse struct {
    Success    bool         `json:"success"`
    ReturnData []byte       `json:"returnData"`
}
```

With a slice of `Call`s, one can perform `Execute(calls)` which will return map using the `Call.Name` as key and `CallResponse` as value.

Note that you must yourself handle the bytes that is in the `ReturnData`, e.g., load `big.Int` and similar.

# Example

```go
var erc20Abi, _ = abi.JSON(strings.NewReader(IERC20.IERC20ABI))

func GetBalanceCall(name string, tokenAddress common.Address, userAddress common.Address) Call {
    parsedData, err := erc20Abi.Pack("balanceOf", userAddress)
    if err != nil {
        panic(err)
    }
    return Call{
        Target:   tokenAddress,
        CallData: parsedData,
        Name:     name,
    }
}

func main() {
    caller := New(https://mainnet.infura.io/v3/<infurakey>)
    
    userAddress := common.HexToAddress("0xb1adceddb2941033a090dd166a462fe1c2029484")
    
    
    var userCalls = make([]Call, 0)
    userCalls = append(userCalls, GetBalanceCall("DillBalance", common.HexToAddress("0xbBCf169eE191A1Ba7371F30A1C344bFC498b29Cf"), userAddress))
    userCalls = append(userCalls, GetBalanceCall("PickleBalance", common.HexToAddress("0x429881672B9AE42b8EbA0E26cD9C73711b891Ca5"), userAddress))
    
    response := caller.Execute(userCalls)
    
    for key, value := range response {
        val := new(big.Int).SetBytes(value.ReturnData)
        println(key, val.String())
    }
}
```

