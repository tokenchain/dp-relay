package BlockSync

import (
	"dprelay/common/conf"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func getFileName(conf *conf.Config) string {
	filename := fmt.Sprintf("%ssycned_latest.json", conf.RootDir)
	return filename
}

type TransRecord struct {
	Token string `json:"token" yaml:"token"`
	From  string `json:"from" yaml:"from"`
	To    string `json:"to" yaml:"to"`
	Time  string `json:"time" yaml:"time"`
	Money string `json:"money" yaml:"money"`
	Id    string `json:"id" yaml:"id"`
}

const layoutISO = "2020-09-19 20:07:46"

func getTime(td string) time.Time {
	t, _ := time.Parse(layoutISO, td)
	return t
}

type GetTransactions struct {
	Code       int           `json:"code" yaml:"code"`
	Message    string        `json:"msg" yaml:"msg"`
	SyncedTime string        `json:"time" yaml:"time"`
	Records    []TransRecord `json:"data" yaml:"data"`
}

func makeFile(conf *conf.Config, data []byte) {
	ioutil.WriteFile(getFileName(conf), data, os.ModePerm)
}

//check file if okay then make a write request
func checkFile(conf *conf.Config, data []byte) bool {
	var req GetTransactions
	var reqlocal GetTransactions
	if err := json.Unmarshal(data, &req); err != nil {
		return true
	}
	// Open the file
	file, err := os.Open(getFileName(conf))
	if err != nil {
		conf.DebugLn(err.Error())
		return true
	}
	defer file.Close()
	datajson, _ := ioutil.ReadFile(getFileName(conf))
	if err := json.Unmarshal(datajson, &reqlocal); err != nil {
		return false
	}
	t1, _ := strconv.Atoi(req.SyncedTime)
	t2, _ := strconv.Atoi(reqlocal.SyncedTime)
	if t1 > t2 {
		return true
	}

	return true
}

func SyncLoop(config *conf.Config, db *DBOperation) {
	defer SyncLoop(config, db)
	data := makerequest(config)
	processData(config, data, db)
}

func processData(config *conf.Config, data []byte, sdb *DBOperation) bool {
	var req GetTransactions
	if err := json.Unmarshal(data, &req); err != nil {
		config.DebugLn(err.Error())
		return true
	}
	config.DebugLn("start here ... ")
	for i := 0; i < len(req.Records); i++ {
		fmt.Println(req.Records[i])
		AppendTransaction(sdb.static_db_handling_object, req.Records[i])
	}
	config.DebugLn("end here ... ")
	/*for i, record := range req.Records {
		fmt.Println(record, i)
		AppendTransaction(sdb.static_db_handling_object, record)
	}*/
	return false
}
func makerequest(f *conf.Config) []byte {
	resp, err := http.Get(f.SyncRequest)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		f.DebugLn(err.Error())
	}
	defer resp.Body.Close()
	//f.DebugLn(string(body))
	if checkFile(f, body) {
		makeFile(f, body)
	}
	return body
}

/*
func Syncdata() {
	defer Syncdata()
	BlocksModel.Db = tuuz.Db()
	lastblock := BlocksModel.Api_find_last()
	height := 1
	if len(lastblock) > 0 {
		height = Calc.Any2Int(lastblock["height"]) + 1
	}
	ret, err := CosCore.Blocks(height)
	if err != nil {
		panic(err)
	} else {
		rtt, err := Jsong.JObject(ret)
		timer := ""
		chain_id := ""
		block_hash := ""
		from_address := ""
		to_address := ""
		memo := ""
		amount := ""
		fee := ""
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
			time.Sleep(time.Second)
			return
		} else {
			block_meta, err := Jsong.ParseObject(rtt["block_meta"])
			if err != nil {
				Log.Errs(err, tuuz.FUNCTION_ALL())
			}
			block_id, err := Jsong.ParseObject(block_meta["block_id"])
			if err != nil {
				Log.Errs(err, tuuz.FUNCTION_ALL())
			} else {
				block_hash = Calc.Any2String(block_id["hash"])
			}

			block, err := Jsong.ParseObject(rtt["block"])
			if err != nil {
				Log.Errs(err, tuuz.FUNCTION_ALL())
			} else {
				header, err := Jsong.ParseObject(block["header"])
				if err != nil {
					Log.Errs(err, tuuz.FUNCTION_ALL())
				} else {
					timer = Calc.Any2String(header["time"])
					height = Calc.Any2Int(header["height"])
					if height < Calc.Any2Int(lastblock["height"]) {
						time.Sleep(time.Second)
						return
					}
					chain_id = Calc.Any2String(header["chain_id"])
				}

				data, err := Jsong.ParseObject(block["data"])
				if err != nil {
					Log.Errs(err, tuuz.FUNCTION_ALL())
				} else {
					txss, err := Jsong.ParseSlice(data["txs"])

					if err != nil || len(txss) < 1 {
						//fmt.Println(err, tuuz.FUNCTION_ALL())
					} else {
						txs, err := txs_format(Calc.Any2String(txss[0]))
						if err != nil {
							fmt.Println(err, tuuz.FUNCTION_ALL())
							time.Sleep(time.Second)
							return
						} else {
							from_address = Calc.Any2String(txs["from_address"])
							to_address = Calc.Any2String(txs["to_address"])
							memo = Calc.Any2String(txs["memo"])
							amount = Calc.Any2String(txs["amount"])
							fee = Calc.Any2String(txs["fee"])
						}
					}
				}
			}
		}
		if height != 0 {
			BlocksModel.Db = tuuz.Db()
			BlocksModel.Api_insert(height, timer, chain_id, block_hash, from_address, to_address, memo, amount, fee, ret)
		} else {
			time.Sleep(time.Second)
		}
		log.Info("sync data in one sec.")
	}
}

func txs_format(txs string) (map[string]interface{}, error) {
	arr := make(map[string]interface{})
	str, err := CosCore.Txs_Decode(txs)
	if err != nil {
		fmt.Println("txs.err")
		return nil, err
	} else {
		json, err := Jsong.JObject(str)
		if err == nil {
			result, err := Jsong.ParseObject(json["result"])
			if err != nil {
				fmt.Println("result,为空")
				Log.Errs(err, tuuz.FUNCTION_ALL())
				return nil, err
			}
			msgs, err := Jsong.ParseSlice(result["msg"])
			if err != nil {
				fmt.Println("msgs,为空")
				Log.Errs(err, tuuz.FUNCTION_ALL())
				return nil, err
			}
			msg, err := Jsong.ParseObject(msgs[0])
			if err != nil {
				fmt.Println("msg,为空")
				Log.Errs(err, tuuz.FUNCTION_ALL())
				return nil, err
			}
			value, err := Jsong.ParseObject(msg["value"])
			if err != nil {
				fmt.Println("value,,为空")
				Log.Errs(err, tuuz.FUNCTION_ALL())
				return nil, err
			}
			arr["from_address"] = value["from_address"]
			arr["to_address"] = value["to_address"]
			arr["amount"], err = Jsong.Encode(value["amount"])
			arr["fee"], err = Jsong.Encode(result["fee"])
			arr["memo"] = result["memo"]
		}
	}
	return arr, err
}
*/
