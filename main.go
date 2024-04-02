package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Transaction represents a single Ethereum transaction
type Transaction struct {
	Hash        string `json:"hash"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	BlockNumber int    `json:"blockNumber"`
}

// EthereumResponse represents the response from Ethereum JSON-RPC API
type EthereumResponse struct {
	ID      int             `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
}

// Parser represents the Ethereum blockchain parser interface
type Parser struct {
	currentBlock int
	subscribed   map[string]bool
}

// NewParser creates a new instance of Parser
func NewParser() *Parser {
	return &Parser{
		currentBlock: 0,
		subscribed:   make(map[string]bool),
	}
}

// GetCurrentBlock returns the last parsed block number
func (p *Parser) GetCurrentBlock() int {
	return p.currentBlock
}

// Subscribe adds address to observer list
func (p *Parser) Subscribe(address string) bool {
	if _, ok := p.subscribed[address]; ok {
		return false // already subscribed
	}
	p.subscribed[address] = true
	return true
}

// GetTransactions returns list of inbound or outbound transactions for an address
func (p *Parser) GetTransactions(address string) ([]Transaction, error) {
	// Call Ethereum JSON-RPC API to get transactions for the given address
	body := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionsByAddress",
		"params":  []interface{}{address},
		"id":      "1",
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://cloudflare-eth.com", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	// Print out response body for debugging
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Response body:", string(bodyBytes))

	var ethResp EthereumResponse
	err = json.NewDecoder(resp.Body).Decode(&ethResp)
	if err != nil {
		return nil, err
	}

	// Parse transactions from result
	var transactions []Transaction
	err = json.Unmarshal(ethResp.Result, &transactions)
	if err != nil {
		return nil, err
	}

	// Update current block number
	var blockNumber int
	err = json.Unmarshal(ethResp.Result, &blockNumber)
	if err != nil {
		return nil, err
	}
	p.currentBlock = blockNumber

	return transactions, nil
}

func main() {
	// Example usage
	parser := NewParser()

	// Subscribe addresses
	parser.Subscribe("0x1234567890abcdef")

	// Get current block
	fmt.Println("Current block:", parser.GetCurrentBlock())

	// Get transactions for an address
	transactions, err := parser.GetTransactions("0x1234567890abcdef")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Transactions for address 0x1234567890abcdef:")
		for _, tx := range transactions {
			fmt.Println(tx)
		}
	}
}
