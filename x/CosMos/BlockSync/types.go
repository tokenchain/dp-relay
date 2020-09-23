package BlockSync

import (
	"dprelay/common/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

type T_block_transaction struct {
	Id          int64     `gorm:"column:id"`
	Fid         int64     `gorm:"column:fid"`
	Coin        string    `gorm:"column:coin"`
	From        string    `gorm:"column:from"`
	To          string    `gorm:"column:to"`
	Amount      string    `gorm:"column:amount"`
	Hash        string    `gorm:"column:hash"`
	BlockHeight int64       `gorm:"column:block_height"`
	CreateTime  time.Time `gorm:"column:time"`
}
type T_dxp_address struct {
	Id         int64  `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	Addressdid string `gorm:"column:addressdid"`
	Addressdx  string `gorm:"column:addressdx"`
	Pub        string `gorm:"column:pubkey"`
	Mnemonic   string `gorm:"column:mnemonic"`
	Raw        string `gorm:"column:raw"`
}

type (
	dbsupport struct {
		f *conf.Config
		//build and develop dedicated connection to history transactions
		slotHistory *gorm.DB
		//build and develop dedicated connection to bet transactions
		slotAccountKeys *gorm.DB
		//build and develop dedicated connection to login transactions
		slotHashUpdate *gorm.DB
		//specific check for the jack pot connection
		lineJackPot *gorm.DB
		//full command for connection the db
		host_cmd string
	}

	EXPO interface {
		//get the history tranaction DB
		GetHistoryLine() *gorm.DB
		//get the history tranaction DB
		GetHashLine() *gorm.DB
		//get bet determination transaction DB
		GetAccountLine() *gorm.DB
		//get jack pot and coin specific in transaction DB
		GetJackPotLine() *gorm.DB
	}
)

var _ EXPO = &dbsupport{}

func EstablishConn(v string, conf *conf.Config) *dbsupport {
	return &dbsupport{host_cmd: v, f: conf}
}
func (v *dbsupport) connec(wg *sync.WaitGroup) (*gorm.DB, bool) {
	wg.Add(1)
	defer wg.Done()

	LN, er := gorm.Open("mysql", v.host_cmd)
	if er != nil {
		fmt.Printf("Fail to connect db server on %+v \n", v.host_cmd)
		panic(fmt.Sprintf("No error should happen when connecting to test database, but got err=%+v", er))
		return nil, false
	}
	// Disable table name's pluralization globally
	LN.SingularTable(true)
	LN.DB().SetMaxOpenConns(v.f.DB.ConnMax)
	LN.DB().SetMaxIdleConns(v.f.DB.ConnMax)
	//other operations
	return LN, true
}
func (v *dbsupport) GetHistoryLine() *gorm.DB {
	return v.slotHistory.Table("block_transaction")
}
func (v *dbsupport) GetHashLine() *gorm.DB {
	return v.slotHashUpdate
}
func (v *dbsupport) GetAccountLine() *gorm.DB {
	return v.slotAccountKeys.Table("dxp_address")
}
func (v *dbsupport) GetJackPotLine() *gorm.DB {
	return v.lineJackPot
}
