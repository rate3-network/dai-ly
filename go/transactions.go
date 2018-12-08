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

const (
	StatusQueued     = "QUEUED"
	StatusProcessing = "PROCESSING"
	StatusSuccess    = "SUCESS"
	StatusFailed     = "FAILED"
)

type Transaction struct {
	Sender     string
	Receiver   string
	Delegate   string
	FeeAmount  string
	SendAmount string
	Token      string
	Nonce      string
	Status     string
}

type Transactions []Transaction
