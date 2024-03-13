package tx

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Gwei = big.NewInt(1e9)

func (s *HttpService) DoTx(ctx context.Context, input *Input) (*Output, error) {
	client, err := ethclient.Dial(input.Host)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(input.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create private key: %v", err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fromAddressHex := fromAddress.Hex()
	if input.FromAddress != fromAddressHex {
		return nil, fmt.Errorf("from address does not match private key")
	}

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get the nonce: %v", err)
	}

	amount := input.Amount
	gasLimit := input.GasLimit
	gasPrice := input.GasPrice
	toAddress := common.HexToAddress(input.ToAddress)
	var data []byte

	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get the chain ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign the transaction: %v", err)
	}

	balance, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get the balance: %v", err)
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send the transaction: %v", err)
	}

	return &Output{
		SignedTx:    signedTx.Hash().Hex(),
		FromAddress: fromAddress.Hex(),
		Balance:     balance,
	}, nil
}
