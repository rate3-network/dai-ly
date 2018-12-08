package ethereum

import (
	"context"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	daily "github.com/rate-engineering/dai-ly/go"
	"github.com/rate-engineering/dai-ly/go/contract"
)

type Client struct {
	*ethclient.Client
	*contract.DelegatableDai
}

func NewClient(contractAddres string, url string) (*Client, error) {
	client := Client{}
	var err error

	c, err := ethclient.Dial(url)
	if err != nil {
		return &client, err
	}

	client.Client = c

	address := common.HexToAddress(contractAddres)
	d, err := contract.NewDelegatableDai(address, client)
	if err != nil {
		return &client, err
	}

	client.DelegatableDai = d

	return &client, err
}

func (c *Client) ExecuteTransaction(tx daily.Transaction) (*types.Transaction, error) {
	fromAddress := common.HexToAddress("0x570932869143c8a6e07b4aa10e0b30814cf45ff0") //delegate
	pendingNonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	privKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY")[2:])
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privKey)
	auth.Nonce = big.NewInt(int64(pendingNonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	receiver := common.HexToAddress(tx.Receiver)
	sigB, err := hexutil.Decode(tx.Signature)
	if err != nil {
		return nil, err
	}

	amount := new(big.Int)
	amount.SetString(tx.SendAmount, 10)

	fee := new(big.Int)
	fee.SetString(tx.FeeAmount, 10)

	nonce := new(big.Int)
	nonce.SetString(tx.Nonce, 10)

	submitted, err := c.TransferPreSigned(auth, sigB, receiver,
		amount, fee, nonce)
	if err != nil {
		return nil, err
	}

	return submitted, nil
}
