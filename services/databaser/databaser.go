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
func DigestBlock(block *wire.MsgBlock, db *sql.DB) {
	var allInserts, allUpdates string

	blockHash := block.Header.BlockHash()
	timestamp := block.Header.Timestamp.Unix()

	for _, tx := range block.Transactions {

		txHash := tx.TxHash()

		for index, output := range tx.TxOut {
			allInserts += createInsert(output, index, blockHash, txHash, timestamp)
		}

		for index, input := range tx.TxIn {
			allUpdates += createUpdate(input, index, blockHash, txHash, timestamp)
		}
	}

	fmt.Println("===============INSERTS===================")
	fmt.Println(allInserts)
	// HIER IS DE STATEMENT VOOR ALLE INSERTS READY

	fmt.Println("===============UPDATES===================")
	fmt.Println(allUpdates)
	// HIER IS DE STATEMENT VOOR ALLE UPDATES READY

	_, err := db.Exec(allInserts)
	errorchecker.CheckFileError(err)

	_, err = db.Exec(allUpdates)
	errorchecker.CheckFileError(err)
}

func createInsert(output *wire.TxOut, index int, blockHash, txHash chainhash.Hash, timestamp int64) string {
	insertStatement := `INSERT INTO outputs (block_hash, tx_hash, index, address, amount, timestamp) VALUES (
	` + blockHash.String() + `,
	` + txHash.String() + `,
	` + strconv.Itoa(index) + `,
	` + converter.ConvertAddress(output) + `, 
	` + strconv.FormatInt(output.Value, 10) + `,
	` + strconv.FormatInt(timestamp, 10) + `);`

	return insertStatement
}

func createUpdate(input *wire.TxIn, index int, blockHash, txHash chainhash.Hash, timestamp int64) string {
	updateStatement := `UPDATE outputs SET (spending_block_hash, spending_tx_hash, spending_index, spending_timestamp) = (
	` + blockHash.String() + `,
	` + txHash.String() + `,
	` + strconv.Itoa(index) + `,
	` + strconv.FormatInt(timestamp, 10) + `) WHERE tx_hash =
	` + input.PreviousOutPoint.Hash.String() + ` AND index =
	` + strconv.FormatUint(uint64(input.PreviousOutPoint.Index), 10) + `;`

	return updateStatement
}
