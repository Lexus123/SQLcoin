package databaser

import (
	"database/sql"
	"fmt"
	"strconv"

	"sqlcoin/services/converter"
	"sqlcoin/services/errorchecker"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = ""
	database = "sqlcoin"
)

/*
MakeConnection ...
*/
func MakeConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, username, database)
	db, err := sql.Open("postgres", psqlInfo)
	errorchecker.CheckFileError(err)
	err = db.Ping()
	errorchecker.CheckFileError(err)
	return db
}

/*
DigestBlock ...
*/
func DigestBlock(allBlocks []*wire.MsgBlock, db *sql.DB) {
	// var allUpdates string
	var lastBlock, lastTx, lastOutput bool
	var txsLen, outputsLen int

	var lastInput bool
	var inputsLen int

	outputInserts := `INSERT INTO outputs (block_hash, tx_hash, index, address, amount, timestamp) VALUES `
	inputInserts := `INSERT INTO inputs (block_hash, tx_hash, index, timestamp, prev_out_hash, prev_out_index) VALUES `

	blocksLen := len(allBlocks)

	for bIndex, block := range allBlocks {
		blockHash := block.Header.BlockHash()
		timestamp := block.Header.Timestamp.Unix()

		if bIndex+1 == blocksLen {
			lastBlock = true
			txsLen = len(block.Transactions)
		}

		for tIndex, tx := range block.Transactions {

			txHash := tx.TxHash()

			// Check whether this is the last block
			if lastBlock && tIndex+1 == txsLen {
				lastTx = true
				outputsLen = len(tx.TxOut)
				inputsLen = len(tx.TxIn)
			}

			for oIndex, output := range tx.TxOut {
				if lastTx && oIndex+1 == outputsLen {
					lastOutput = true
				}

				outputInserts += createOutputInsert(output, oIndex, blockHash, txHash, timestamp, lastOutput)
			}

			for iIndex, input := range tx.TxIn {
				if lastTx && iIndex+1 == inputsLen {
					lastInput = true
				}
				inputInserts += createInputInsert(input, blockHash, txHash, iIndex, timestamp, lastInput)
			}
		}

	}

	fmt.Println("=============== OUTPUT INSERTS ===================")
	fmt.Println(outputInserts)
	_, err := db.Exec(outputInserts)
	errorchecker.CheckFileError(err)

	fmt.Println("=============== INPUT INSERTS ===================")
	fmt.Println(inputInserts)
	_, err = db.Exec(inputInserts)
	errorchecker.CheckFileError(err)
}

func createOutputInsert(output *wire.TxOut, index int, blockHash, txHash chainhash.Hash, timestamp int64, lastOutput bool) string {
	insertStatement := `('` + blockHash.String() + `',
	'` + txHash.String() + `',
	` + strconv.Itoa(index) + `,
	'` + converter.ConvertAddress(output) + `', 
	` + strconv.FormatInt(output.Value, 10) + `,
	` + strconv.FormatInt(timestamp, 10) + `)`

	if lastOutput {
		insertStatement += "; "
	} else {
		insertStatement += ", "
	}

	return insertStatement
}

func createInputInsert(input *wire.TxIn, blockHash, txHash chainhash.Hash, index int, timestamp int64, lastInput bool) string {
	insertStatement := `('` + blockHash.String() + `',
	'` + txHash.String() + `',
	` + strconv.Itoa(index) + `,
	` + strconv.FormatInt(timestamp, 10) + `, 
	'` + input.PreviousOutPoint.Hash.String() + `',
	` + strconv.Itoa(int(input.PreviousOutPoint.Index)) + `)`

	if lastInput {
		insertStatement += "; "
	} else {
		insertStatement += ", "
	}

	return insertStatement
}
