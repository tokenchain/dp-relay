package x

import (
	"encoding/json"
	"fmt"
	"github.com/tokenchain/ixo-blockchain/x/did/exported"
	"testing"
)

const sample_did_01_mem = "... add your 24 words key ..."

func Test_generator(t *testing.T) {
	setPrefix()
	fmt.Println("========Seed==========")
	total_accounts := uint32(99)
	account_index := uint32(177)
	list := []string{
		"singularity",
		"blackhole",
		"cosmos",
		"cosmic",
		"darkness",
		"nova",
		"proton",
		"rednova",
		"bitcm",
		"dollar",
		"kbs",
		"coindom",
		"coinex",
		"nasdaq",
		"tokenkingdom",
	}
	var i uint32
	var name string
	filename := "_importCerts.sh"
	writer := NewCertWriter(certPath, filename).ClearFile()
	for i = 1; i < total_accounts; i++ {

		if i < uint32(len(list)+1) {
			name = list[i-1]
		} else {
			name = "cosmos---"
			continue
		}

		did := exported.NewDidGeneratorBuilder().
			WithName(name).
			RecoverBIP44(sample_did_01_mem, "", account_index, i)

		jsonString, _ := json.Marshal(did)
		makeFile(name, jsonString)
		appendImportCert(writer, name)
		fmt.Println("========Check The DID ===========")
		fmt.Println(did.Did)
		fmt.Println("========DID Recover  =========")
		fmt.Println(did)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
	}
}
