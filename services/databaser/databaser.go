package databaser

import (
	"database/sql"
	"fmt"
	"sqlcoin/services/custommodels"
	"sqlcoin/services/errorchecker"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	_ "github.com/go-sql-driver/mysql"
)

func createLogin(dbCreds string, dbName string) string {
	var login strings.Builder

	login.WriteString(dbCreds)
	login.WriteString("@/")
	login.WriteString(dbName)

	return login.String()
}

/*
MakeInsert ...
*/
func MakeInsert(block *wire.MsgBlock) {
	for _, tx := range block.Transactions {
		insertOutputs(tx.TxOut)
	}
	// login := createLogin(dbCreds, dbName)
	// db, err := sql.Open("mysql", login)
	// errorchecker.CheckFileError(err)
	// defer db.Close()

	// dbBlock := prepareBlock(block)
	// lastBlockID := insertBlock(dbBlock, db)

	// insertTxs(block.Transactions, lastBlockID, db)
}

func prepareBlock(block *wire.MsgBlock) custommodels.DbBlock {
	return custommodels.DbBlock{
		Hash:      block.Header.BlockHash().String(),
		PrevHash:  block.Header.PrevBlock.String(),
		Merkle:    block.Header.MerkleRoot.String(),
		Timestamp: block.Header.Timestamp.Unix(),
	}
}

func insertBlock(dbBlock custommodels.DbBlock, db *sql.DB) int64 {
	insertStatement, err := db.Prepare("INSERT INTO blocks (height, hash, prevHash, merkle, timestamp) VALUES (?,?,?,?,?)")
	errorchecker.CheckFileError(err)

	result, err := insertStatement.Exec(nil, dbBlock.Hash, dbBlock.PrevHash, dbBlock.Merkle, dbBlock.Timestamp)
	errorchecker.CheckFileError(err)

	lastBlockID, _ := result.LastInsertId()
	return lastBlockID
}

func insertTxs(txs []*wire.MsgTx, lastBlockID int64, db *sql.DB) {
	for _, tx := range txs {
		dbTx := custommodels.DbTx{
			Hash:        tx.TxHash().String(),
			BlockHeight: lastBlockID,
		}

		insertStatement, err := db.Prepare("INSERT INTO txs (txID, hash, blockHeight) VALUES (?,?,?)")
		errorchecker.CheckFileError(err)

		_, err = insertStatement.Exec(nil, dbTx.Hash, dbTx.BlockHeight)
		errorchecker.CheckFileError(err)

		// insertOutputs(tx.TxOut, dbTx.Hash, db)
		// insertInputs(tx.TxIn, dbTx.Hash, db)
	}
}

func insertOutputs(outputs []*wire.TxOut) {
	for i, output := range outputs {
		_, addresses, _, err := txscript.ExtractPkScriptAddrs(output.PkScript, &chaincfg.MainNetParams)
		fmt.Println(addresses)
		errorchecker.CheckFileError(err)
		var address string
		if len(addresses) > 0 {
			address = addresses[0].EncodeAddress()

		}

		dbOutput := custommodels.DbOutput{
			TxIndex: i,
			Amount:  output.Value,
			Address: address,
			Used:    0,
		}

		fmt.Println(dbOutput)
	}
}

func insertInputs(inputs []*wire.TxIn, txHash string, db *sql.DB) {
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
