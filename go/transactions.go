package daily

type TransactionRequest struct {
	Sender     string `json:"sender" valid:"hexadecimal,required"`
	Receiver   string `json:"receiver" valid:"hexadecimal,required"`
	Delegate   string `json:"delegate" valid:"hexadecimal"`
	FeeAmount  string `json:"feeAmount" valid:"required"`
	SendAmount string `json:"sendAmount" valid:"required"`
	Token      string `json:"token" valid:"uppercase,required"`
	Nonce      string `json:"nonce" valid:"int,required"`
}

func GenerateHash(req TransactionRequest) string {
	// var TXHash string

	return req.Nonce
}

const (
	StatusQueued     = "QUEUED"
	StatusProcessing = "PROCESSING"
	StatusSuccess    = "SUCESS"
	StatusFailed     = "FAILED"
)

type Transaction struct {
	TXHash     string `json:"txHash"`
	Sender     string `json:"sender"`
	Receiver   string `json:"receiver"`
	Delegate   string `json:"delegate"`
	FeeAmount  string `json:"feeAmount"`
	SendAmount string `json:"sendAmount"`
	Token      string `json:"token"`
	Nonce      string `json:"nonce"`
	Status     string `json:"status"`
}

type Transactions []Transaction
