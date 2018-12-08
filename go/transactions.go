package daily

import (
	"bytes"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type TransactionRequest struct {
	Signature  string `json:"signature" valid:"hexadecimal,required"`
	Sender     string `json:"sender" valid:"hexadecimal,required"`
	Receiver   string `json:"receiver" valid:"hexadecimal,required"`
	FeeAmount  string `json:"feeAmount" valid:"required"`
	SendAmount string `json:"sendAmount" valid:"required"`
	Token      string `json:"token" valid:"hexadecimal"`
	Nonce      string `json:"nonce" valid:"int,required"`
}

// https://goethereumbook.org/signature-generate/
func GenerateHash(req TransactionRequest) (string, error) {
	var err error
	var TXHash common.Hash
	var methodID []byte
	var token []byte
	var receiver []byte
	var paddedValue []byte
	var paddedFee []byte
	var paddedNonce []byte

	methodID, err = hexutil.Decode("0x15420b71")
	if err != nil {
		log.Printf("Method decode error: %s\n", err.Error())
		return "", err
	}

	token, err = hexutil.Decode(req.Token)
	token = common.LeftPadBytes(token, 32)
	if err != nil {
		log.Printf("Token decode error: %s\n", err.Error())
		return "", err
	}

	receiver, err = hexutil.Decode(req.Receiver)
	receiver = common.LeftPadBytes(receiver, 32)
	if err != nil {
		log.Printf("Receiver decode error: %s\n", err.Error())
		return "", err
	}

	amount := new(big.Int)
	amount.SetString(req.SendAmount, 10)
	paddedValue = common.LeftPadBytes(amount.Bytes(), 32)

	fee := new(big.Int)
	fee.SetString(req.FeeAmount, 10)
	paddedFee = common.LeftPadBytes(fee.Bytes(), 32)

	nonce := new(big.Int)
	nonce.SetString(req.Nonce, 10)
	paddedNonce = common.LeftPadBytes(nonce.Bytes(), 32)

	if err != nil {
		log.Fatalln(err)
	}

	var data []byte

	// "15420b71": transferPreSignedHashing(address,address,uint256,uint256,uint256)
	// keccak256(abi.encodeWithSelector(bytes4(0x15420b71), _token, _to, _value, _fee, _nonce));

	data = append(data, methodID...)
	data = append(data, token...)
	data = append(data, receiver...)
	data = append(data, paddedValue...)
	data = append(data, paddedFee...)
	data = append(data, paddedNonce...)

	TXHash = crypto.Keccak256Hash(data)
	return TXHash.Hex(), err
}

// Returns
func VerifySigner(sender string, hash string, sig string) bool {
	expectedAddressBytes, _ := hexutil.Decode(sender)

	var hashB []byte
	var sigB []byte
	hashB, err := hexutil.Decode(hash)
	if err != nil {
		log.Printf("Hash decode error: %s\n", err.Error())
		return false
	}
	hashB = crypto.Keccak256([]byte("\x19Ethereum Signed Message:\n32"), []byte(hashB))

	sigB, err = hexutil.Decode(sig)
	if err != nil {
		log.Printf("Sig decode error: %s\n", err.Error())
		return false
	}
	pkey, err := crypto.SigToPub(hashB, sigB)
	if err != nil {
		log.Printf("SigToPub error: %s\n", err.Error())
		return false
	}
	actualAddress := crypto.PubkeyToAddress(*pkey)
	matches := bytes.Equal(expectedAddressBytes, actualAddress.Bytes())

	return matches
}

const (
	StatusQueued     = "QUEUED"
	StatusProcessing = "PROCESSING"
	StatusSuccess    = "SUCESS"
	StatusFailed     = "FAILED"
)

type Transaction struct {
	TXHash        string `json:"-"`
	Signature     string `json:"signature"`
	Sender        string `json:"sender"`
	Receiver      string `json:"receiver"`
	Delegate      string `json:"delegate"`
	FeeAmount     string `json:"feeAmount"`
	SendAmount    string `json:"sendAmount"`
	Token         string `json:"token"`
	Nonce         string `json:"nonce"`
	SubmittedHash string `json:"submittedHash"`
	Status        string `json:"status"`
}

type Transactions []Transaction
