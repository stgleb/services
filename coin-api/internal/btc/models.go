package btc

import "time"

type balanceRequest struct {
	Address  string `json:"address"`
	Currency string `json:"currency"`
}

type explorerTxStatus struct {
	Total         float64 `json:"total"`
	Fees          float64 `json:"fees"`
	Confirmations int64   `json:"confirmations"`

	BlockHash  string `json:"block_hash"`
	BlockIndex int64  `json:"block_index"`

	Hash      string    `json:"hash"`
	Confirmed time.Time `json:"confirmed"`
	Received  time.Time `json:"received"`
}

type txRef struct {
	TxHash        string    `json:"tx_hash"`
	BlockHeight   int       `json:"block_height"`
	TxInputN      int       `json:"tx_input_n"`
	TxOutputN     int       `json:"tx_output_n"`
	Value         int       `json:"value"`
	RefBalance    int       `json:"ref_balance"`
	Confirmations int       `json:"confirmations"`
	Confirmed     time.Time `json:"confirmed"`
	DoubleSpend   bool      `json:"double_spend"`
	Spent         bool      `json:"spent,omitempty"`
	SpentBy       string    `json:"spent_by,omitempty"`
}

// NOTE(stgleb): See https://blockcypher.github.io/documentation/#address_details
type explorerAddressResponse struct {
	Address                 string  `json:"address"`
	TotalReceived           int64   `json:"total_received"`
	TotalSent               int64   `json:"total_sent"`
	Balance                 int64   `json:"balance"`
	UnconfirmedBalance      int64   `json:"unconfirmed_balance"`
	FinalBalance            int64   `json:"final_balance"`
	NTx                     int64   `json:"n_tx"`
	UnconfirmedNTx          int64   `json:"unconfirmed_n_tx"`
	FinalNTx                int64   `json:"final_n_tx"`
	Transactions            []txRef `json:"txrefs"`
	UnconfirmedTransactions []txRef `json:"unconfirmed_txrefs"`
	HasMore                 bool    `json:"hasMore"`
	TxURL                   string  `json:"tx_url"`
}

type TxInfo struct {
	BlockHash string `json:"block_hash"`
}

type Deposit struct {
	Amount        int    `json:"amount"`
	TxHash        string `json:"tx_hash"`
	BlockHash     string `json:"block_hash"`
	Confirmations int    `json:"confirmations"`
	Spent         bool   `json:"spent"`
	SpentBy       string `json:"spent_by,omitempty"`
	Height        int    `json:"height"`
}

type BalanceResponse struct {
	Address     string    `json:"address"`
	Balance     int64     `json:"balance"`
	Utxo        []Deposit `json:"utxo"`
	PendingUtxo []Deposit `json:"pending_utxo"`
}

type TxStatus struct {
	Amount        float64 `json:"amount"`
	Confirmations int64   `json:"confirmations"`
	Fee           float64 `json:"fee"`

	BlockHash  string `json:"blockhash"`
	BlockIndex int64  `json:"block_index"`

	Hash      string `json:"hash"`
	Confirmed int64  `json:"confirmed"`
	Received  int64  `json:"received"`
}
