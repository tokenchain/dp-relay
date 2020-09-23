package BlockSync

import (
	"dprelay/common/conf"
	"fmt"
	"sync"
)

const (
	GORM_MYSQL_CONNECTOR = "%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local"
)
type DBOperation struct {
	f                         *conf.Config
	static_db_handling_object *dbsupport
	block_trans               []T_block_transaction
	accounts                  []T_dxp_address
}

func NewClientQ(ff *conf.Config) *DBOperation {
	return &DBOperation{
		f:           ff,
		block_trans: []T_block_transaction{},
	}
}
func (q *DBOperation) Start() *DBOperation {
	q.connectWait()
	return q
}
func (q *DBOperation) connectWait() *dbsupport {
	q.static_db_handling_object = EstablishConn(fmt.Sprintf(GORM_MYSQL_CONNECTOR, q.f.DB, q.f.DB.Password, q.f.DB.Server, q.f.DB.DbName), q.f)
	var dwg sync.WaitGroup
	var lines int
	q.f.DebugLn("🌐 db waiting.. ")
	if a, ok := q.static_db_handling_object.connec(&dwg); ok {
		q.static_db_handling_object.slotHistory = a
		lines++
	} else {
		q.f.DebugLn("cannot continue")
	}
	if a, ok := q.static_db_handling_object.connec(&dwg); ok {
		q.static_db_handling_object.slotHashUpdate = a
		lines++
	} else {
		q.f.DebugLn("cannot continue")
	}
	if a, ok := q.static_db_handling_object.connec(&dwg); ok {
		q.static_db_handling_object.lineJackPot = a
		lines++
	} else {
		q.f.DebugLn("cannot continue")
	}
	if a, ok := q.static_db_handling_object.connec(&dwg); ok {
		q.static_db_handling_object.slotAccountKeys = a
		lines++
	} else {
		q.f.DebugLn("cannot continue")
	}
	dwg.Wait()
	q.f.DebugLn("🚦 There are %d communication channels established. ", lines)
	return q.static_db_handling_object
}
