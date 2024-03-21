package tx

import "math/big"

type Input struct {
	Host        string
	ToAddress   string
	FromAddress string
	PrivateKey  string
	Amount      *big.Int
	GasLimit    uint64
	GasPrice    *big.Int
}

type Output struct {
	SignedTx    string
	FromAddress string
	Balance     *big.Int
}

type Request struct {
	Host        string `json:"host"`
	ToAddress   string `json:"toAddress"`
	FromAddress string `json:"fromAddress"`
	PrivateKey  string `json:"privateKey"`
	Amount      int    `json:"amount"`
	GasLimit    int    `json:"gasLimit"`
	GasPrice    int    `json:"gasPrice"`
}

type Response struct {
	Output *Output `json:"output"`
	Error  string  `json:"error"`
}
