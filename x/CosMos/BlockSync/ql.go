package BlockSync

import (
	"time"
)

func GetAllTransactions(connection *dbsupport) []T_block_transaction {
	ress := []T_block_transaction{}
	connection.GetHistoryLine().Select("*").Scan(&ress)
	return ress
}

func GetAllAddresses(connection *dbsupport) []T_dxp_address {
	ress := []T_dxp_address{}
	connection.GetHistoryLine().Select("*").Scan(&ress)
	return ress
}

func GetTransactionByHash(connection *dbsupport, hash string) []T_block_transaction {
	ress := []T_block_transaction{}
	connection.GetHistoryLine().Select("*").Where("hash = ", hash).Scan(&ress)
	return ress
}

func HashComplete(connection *dbsupport, r TransRecord, bh int64, hash string) int64 {
	ress := T_block_transaction{}
	connection.GetHistoryLine().Select("*").Where("fid = ?", r.Id).First(&ress)
	if ress.Fid == 0 {
		return -1
	} else {
		ress.BlockHeight = bh
		ress.Hash = hash
		connection.GetHistoryLine().Save(ress)
		return 0
	}
}
func AppendTransaction(connection *dbsupport, r TransRecord) int64 {
	hash_rec := &T_block_transaction{}
	connection.GetHistoryLine().Select("*").Where("fid = ?", r.Id).Scan(&hash_rec)
	if hash_rec.Id > 0 {
		hash_rec.CreateTime = time.Now()
		hash_rec.To = r.To
		hash_rec.From = r.From
		hash_rec.Amount = r.Money
		hash_rec.Coin = r.Token
		connection.GetHistoryLine().Save(hash_rec)
		return hash_rec.Fid
	} else {
		detailTransact := T_block_transaction{
			Coin:       r.Token,
			To:         r.To,
			From:       r.From,
			Amount:     r.Money,
			CreateTime: time.Now(),
		}
		if connection.GetHistoryLine().NewRecord(detailTransact) {
			connection.GetHistoryLine().Create(&detailTransact)
			rh := &T_block_transaction{}
			connection.GetHistoryLine().Last(&rh)
			//common.Info(fmt.Sprintf("Recorded HASH result GID %d", rh.Id))
			return rh.Id
		} else {
			//common.Error("sql failure to append prize record.")
			return -1
		}
	}
}
