package databaser

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"sqlcoin/services/errorchecker"

	_ "github.com/go-sql-driver/mysql"
)

/*
Input is a very simple struct
*/
type Input struct {
	InID        int    `json:"inID"`
	TxHash      string `json:"txHash"`
	PrevTxHash  string `json:"prevTxHash"`
	PrevTxIndex int    `json:"prevTxIndex"`
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
GetInputs ...
*/
func GetInputs() [][]byte {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:redacted@/sqlcoin")

	// if there is an error opening the connection, handle it
	errorchecker.CheckFileError(err)
	defer db.Close()

	var inputs [][]byte

	// Execute the query
	results, err := db.Query("SELECT * FROM inputs")
	errorchecker.CheckFileError(err)

	var input Input

	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(&input.InID, &input.TxHash, &input.PrevTxHash, &input.PrevTxIndex)
		e, err := json.Marshal(input)
		errorchecker.CheckFileError(err)
		inputs = append(inputs, e)
	}

	return inputs
}

/*
GetAllBlocks ...
*/
func GetAllBlocks() []Block {
	var blocks []Block
	var block Block

	// Open up our database connection.
	db, err := sql.Open("mysql", "root:redacted@/sqlcoin")

	// if there is an error opening the connection, handle it
	errorchecker.CheckFileError(err)
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM blocks LIMIT 10")
	errorchecker.CheckFileError(err)

	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(&block.Height, &block.Hash, &block.PrevHash, &block.Merkle, &block.Timestamp)
		blocks = append(blocks, block)
	}

	return blocks
}

/*
GetSingleBlocks ...
*/
func GetSingleBlocks(blockId string) Block {
	var block Block
	id, _ := strconv.Atoi(blockId)

	sqlStatement := `SELECT * FROM blocks WHERE height=?`

	// Open up our database connection.
	db, err := sql.Open("mysql", "root:redacted@/sqlcoin")

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
