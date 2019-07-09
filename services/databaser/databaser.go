package databaser

import (
	"database/sql"
	"fmt"
	"sqlcoin/services/custommodels"
	"sqlcoin/services/errorchecker"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	db, err := sql.Open("postgres", psqlInfo)
	errorchecker.CheckFileError(err)
	err = db.Ping()
	errorchecker.CheckFileError(err)
	return db
}

/*
MakeInsert ...
*/
func MakeInsert(block *wire.MsgBlock, db *sql.DB) {
	dbBlock := prepareBlock(block)
	insertBlock(dbBlock, db)

	// insertTxs(block.Transactions, lastBlockID, db)
}

func prepareBlock(block *wire.MsgBlock) custommodels.DbBlock {
	return custommodels.DbBlock{
		Hash:      block.Header.BlockHash(),
		PrevHash:  block.Header.PrevBlock,
		Merkle:    block.Header.MerkleRoot,
		Timestamp: block.Header.Timestamp.Unix(),
	}
}

func insertBlock(dbBlock custommodels.DbBlock, db *sql.DB) {
	sqlStatement := `
		INSERT INTO blocks (hash)
		VALUES (` + fmt.Sprint(dbBlock.Hash) + `)`
	_, err := db.Exec(sqlStatement)
	errorchecker.CheckFileError(err)
}

func insertTxs(txs []*wire.MsgTx, lastBlockID int64, db *sql.DB) {
	// txn, _ := db.Begin()

	// txStatement, _ := txn.Prepare(pq.CopyIn("txs", "id", "hash", "block_id"))
	// outputStatement, _ := txn.Prepare(pq.CopyIn("outputs", "id", "block_id", "tx_hash", "index", "address", "amount", "spending_block_id", "spending_tx_hash", "spending_index"))

	// for _, tx := range txs {
	// 	dbTx := custommodels.DbTx{
	// 		Hash:    tx.TxHash(),
	// 		BlockId: lastBlockID,
	// 	}

	// 	_, err := txStatement.Exec(nil, dbTx.Hash, dbTx.BlockId)
	// 	errorchecker.CheckFileError(err)

	// 	insertOutputs(tx.TxOut, lastBlockID, dbTx.Hash, outputStatement)
	// 	// updateOutputs(tx.TxIn, dbTx.Hash, db)
	// }

	// _, err := txStatement.Exec()
	// errorchecker.CheckFileError(err)

	// err = txStatement.Close()
	// errorchecker.CheckFileError(err)

	// err = txn.Commit()
	// errorchecker.CheckFileError(err)
}

func insertOutputs(outputs []*wire.TxOut, lastBlockID int64, txHash [32]byte, outputStatement *sql.Stmt) {
	for i, output := range outputs {
		_, addresses, _, err := txscript.ExtractPkScriptAddrs(output.PkScript, &chaincfg.MainNetParams)
		errorchecker.CheckFileError(err)
		var address string
		if len(addresses) > 0 {
			address = addresses[0].String()
		}

		dbOutput := custommodels.DbOutput{
			BlockId:         lastBlockID,
			TxHash:          txHash,
			Index:           i,
			Address:         address,
			Amount:          output.Value,
			SpendingBlockId: 0,
			SpendingTxHash:  txHash,
			SpendingIndex:   0,
		}

		_, err = outputStatement.Exec(nil, dbOutput.BlockId, dbOutput.TxHash, dbOutput.Index, dbOutput.Address, dbOutput.Amount, dbOutput.SpendingBlockId, dbOutput.SpendingTxHash, dbOutput.SpendingIndex)
		errorchecker.CheckFileError(err)
	}

	_, err := outputStatement.Exec()
	errorchecker.CheckFileError(err)

	err = outputStatement.Close()
	errorchecker.CheckFileError(err)
}

func updateOutputs(inputs []*wire.TxIn, txHash string, db *sql.DB) {
	for _, input := range inputs {
		dbInput := custommodels.DbInput{
			TxHash:      txHash,
			PrevTxHash:  input.PreviousOutPoint.Hash.String(),
			PrevTxIndex: input.PreviousOutPoint.Index,
		}

		insertStatement, err := db.Prepare("INSERT INTO inputs (inID, txHash, prevTxHash, prevTxIndex) VALUES (?,?,?,?)")
		errorchecker.CheckFileError(err)

		_, err = insertStatement.Exec(nil, dbInput.TxHash, dbInput.PrevTxHash, dbInput.PrevTxIndex)
		errorchecker.CheckFileError(err)
	}
}
