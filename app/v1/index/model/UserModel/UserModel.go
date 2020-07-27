package UserModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "cos_user"

var Db = tuuz.Db()

func Api_insert(username, password, paypass, address string) bool {
	db := Db.Table(table)
	data := make(map[string]interface{})
	data["username"] = username
	data["password"] = password
	data["paypass"] = paypass
	data["address"] = address
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find_byUsername(username string) gorose.Data {
	db := Db.Table(table)
	where := make(map[string]interface{})
	where["username"] = username
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find(username, password string) gorose.Data {
	db := Db.Table(table)
	where := make(map[string]interface{})
	where["username"] = username
	where["password"] = password
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_byPaypass(username, paypass string) gorose.Data {
	db := Db.Table(table)
	where := make(map[string]interface{})
	where["username"] = username
	where["paypass"] = paypass
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
