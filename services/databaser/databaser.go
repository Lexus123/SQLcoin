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

	allInserts := `INSERT INTO outputs (block_hash, tx_hash, index, address, amount, timestamp) VALUES `

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

			if lastBlock && tIndex+1 == txsLen {
				lastTx = true
				outputsLen = len(tx.TxOut)
			}

			for oIndex, output := range tx.TxOut {
				if lastTx && oIndex+1 == outputsLen {
					lastOutput = true
				}

				allInserts += createInsert(output, oIndex, blockHash, txHash, timestamp, lastOutput)
			}

			// for index, input := range tx.TxIn {
			// 	if input.PreviousOutPoint.Hash != [32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} && input.PreviousOutPoint.Index != 4294967295 {
			// 		allUpdates += createUpdate(input, index, blockHash, txHash, timestamp)
			// 	}
			// }
		}

	}

	fmt.Println("===============INSERTS===================")
	fmt.Println(allInserts)
	// HIER IS DE STATEMENT VOOR ALLE INSERTS READY

	// fmt.Println("===============UPDATES===================")
	// fmt.Println(allUpdates)
	// HIER IS DE STATEMENT VOOR ALLE UPDATES READY

	_, err := db.Exec(allInserts)
	errorchecker.CheckFileError(err)

	// _, err = db.Exec(allUpdates)
	// errorchecker.CheckFileError(err)
}

func createInsert(output *wire.TxOut, index int, blockHash, txHash chainhash.Hash, timestamp int64, lastOutput bool) string {
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

func createUpdate(input *wire.TxIn, index int, blockHash, txHash chainhash.Hash, timestamp int64) string {
	updateStatement := `UPDATE outputs SET (spending_block_hash, spending_tx_hash, spending_index, spending_timestamp) = (
	'` + blockHash.String() + `',
	'` + txHash.String() + `',
	` + strconv.Itoa(index) + `,
	` + strconv.FormatInt(timestamp, 10) + `) WHERE tx_hash =
	'` + input.PreviousOutPoint.Hash.String() + `' AND index =
	` + strconv.FormatUint(uint64(input.PreviousOutPoint.Index), 10) + `;`

	return updateStatement
}
