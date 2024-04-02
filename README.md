# Ethereum Blockchain Parser

The Ethereum Blockchain Parser is a Go application designed to parse Ethereum blockchain transactions and allow users to subscribe to specific addresses for transaction notifications. It interacts with the Ethereum JSONRPC interface to fetch transaction data and provides a simple interface for users to query transactions for subscribed addresses.

## Features

- **Subscription Management**: Users can subscribe to specific Ethereum addresses to receive notifications for inbound and outbound transactions.
- **Transaction Querying**: The parser allows users to retrieve a list of transactions for subscribed addresses.
- **JSONRPC Interaction**: Utilizes Ethereum JSONRPC to fetch block details and transaction data from the Ethereum blockchain.
- **Memory Storage**: Stores subscriber data and transaction history in memory for quick retrieval.

## Usage

1. Clone the repository:

```bash
git clone https://github.com/LuSicong22/tx-parser.git
```
2. Navigate to the project directory:

```bash
cd tx-parser
```
3. Build the application:
```bash
go build
```
4. Run the application:
```bash
./tx-parser 
```