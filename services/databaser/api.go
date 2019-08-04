package databaser

import (
	"database/sql"
	// "encoding/json"
	"fmt"
	"strconv"
	"strings"

	"sqlcoin/services/errorchecker"

	_ "github.com/go-sql-driver/mysql"
)

/*
Input is a very simple struct
*/
type Input struct {
	InID         int    `json:"id"`
	BlockHash    string `json:"blockHash"`
	TxHash       string `json:"txHash"`
	Index        int    `json:"index"`
	Timestamp    int    `json:"timestamp"`
	PrevOutHash  string `json:"prevOutHash"`
	PrevOutIndex int    `json:"prevOutIndex"`
}

/*
Output is a very simple struct
*/
type Output struct {
	OutID     int    `json:"id"`
	BlockHash string `json:"blockHash"`
	TxHash    string `json:"txHash"`
	Index     int    `json:"index"`
	Address   string `json:"address"`
	Amount    int64  `json:"amount"`
	Timestamp int64  `json:"timestamp"`
}

/*
Block is a very simple struct
*/
type Block struct {
	Height    int    `json:"heigth"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevHash"`
	Merkle    string `json:"merkle"`
	Timestamp int64  `json:"timestamp"`
}

/*
Tx is a very simple struct
*/
type Tx struct {
	TxHash  string   `json:"hash"`
	Inputs  []Input  `json:"inputs"`
	Outputs []Output `json:"outputs"`
}

/*
GetAllInputs ...
*/
func GetAllInputs() []Input {
	var input Input
	var inputs []Input

	// Open up our database connection.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, username, database)
	db, err := sql.Open("postgres", psqlInfo)

	// if there is an error opening the connection, handle it
	errorchecker.CheckFileError(err)
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM inputs LIMIT 10")
	errorchecker.CheckFileError(err)

	for rows.Next() {
		// for each row, scan the result into our tag composite object
		err = rows.Scan(&input.InID, &input.BlockHash, &input.TxHash, &input.Index, &input.Timestamp, &input.PrevOutHash, &input.PrevOutIndex)
		inputs = append(inputs, input)
	}

	return inputs
}

/*
GetAllOutputs ...
*/
func GetAllOutputs() []Output {
	var output Output
	var outputs []Output

	// Open up our database connection.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, username, database)
	db, err := sql.Open("postgres", psqlInfo)

	// if there is an error opening the connection, handle it
	errorchecker.CheckFileError(err)
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT tx_hash, index FROM outputs ORDER BY id DESC LIMIT 1000")
	errorchecker.CheckFileError(err)

	for rows.Next() {
		err = rows.Scan(&output.TxHash, &output.Index)
		outputs = append(outputs, output)
	}

	return outputs
}

/*
GetSingleOutput ...
*/
func GetSingleOutput(outputParam string) Output {
	var output Output
	outputParams := strings.Split(outputParam, ":")
	sqlStatement := `SELECT * FROM outputs WHERE tx_hash = $1 AND index = $2`

	// Open up our database connection.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, username, database)
	db, err := sql.Open("postgres", psqlInfo)

	// if there is an error opening the connection, handle it
	errorchecker.CheckFileError(err)
	defer db.Close()

	// Execute the query
	row := db.QueryRow(sqlStatement, outputParams[0], outputParams[1])
	errorchecker.CheckFileError(err)

	err = row.Scan(&output.OutID, &output.BlockHash, &output.TxHash, &output.Index, &output.Address, &output.Amount, &output.Timestamp)
	errorchecker.CheckFileError(err)

	return output
}

/*
GetSingleTx ...
*/
func GetSingleTx(txParam string) Tx {
	var tx Tx
	var output Output
	var input Input

	sqlStatementInputs := `SELECT prev_out_hash, prev_out_index FROM inputs WHERE tx_hash = $1`
	sqlStatementOutputs := `SELECT address, amount FROM outputs WHERE tx_hash = $1`

	// Open up our database connection.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, username, database)
	db, err := sql.Open("postgres", psqlInfo)

	errorchecker.CheckFileError(err)
	defer db.Close()

	// First query (inputs)
	inputRows, err := db.Query(sqlStatementInputs, txParam)
	errorchecker.CheckFileError(err)

	for inputRows.Next() {
		err = inputRows.Scan(&input.PrevOutHash, &input.PrevOutIndex)
		tx.Inputs = append(tx.Inputs, input)
	}

	// Second query (outputs)
	outputRows, err := db.Query(sqlStatementOutputs, txParam)
	errorchecker.CheckFileError(err)

	for outputRows.Next() {
		err = outputRows.Scan(&output.Address, &output.Amount)
		tx.Outputs = append(tx.Outputs, output)
	}

	tx.TxHash = txParam

	return tx
}

/*
CountInputs ...
*/
func CountInputs() int {
	var count int

	// Open up our database connection.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, username, database)
	db, err := sql.Open("postgres", psqlInfo)

	// if there is an error opening the connection, handle it
	errorchecker.CheckFileError(err)
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT count(*) FROM inputs")
	errorchecker.CheckFileError(err)

	for results.Next() {
		err = results.Scan(&count)
		errorchecker.CheckFileError(err)
	}

	return count
}

/*
GetSingleBlocks ...
*/
func GetSingleBlocks(blockId string) Block {
	var block Block
	id, _ := strconv.Atoi(blockId)

	sqlStatement := `SELECT * FROM blocks WHERE height = ?`

	// Open up our database connection.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, username, database)
	db, err := sql.Open("postgres", psqlInfo)

	// if there is an error opening the connection, handle it
	errorchecker.CheckFileError(err)
	defer db.Close()

	// Execute the query
	row := db.QueryRow(sqlStatement, id+1)
	errorchecker.CheckFileError(err)

	err = row.Scan(&block.Height, &block.Hash, &block.PrevHash, &block.Merkle, &block.Timestamp)
	errorchecker.CheckFileError(err)
	return block
}
