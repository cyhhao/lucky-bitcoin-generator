package main

import (
	"bip44"
	"fmt"
	"io/ioutil"
	"main/tools"
	"os"
	"runtime"
	"strings"
)

func writeFile(name, data string) {
	fd, _ := os.OpenFile("./ans/"+name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.WriteString(data)
	fd.Close()

}

func readFile() string {
	input, err := ioutil.ReadFile("./input")
	if err != nil {
		fmt.Println("请在当前目录下创建名为 input 的文件，每行一个目标前缀，搜索结果会放入 ans 目录下")
	}

	return string(input)
}

var (
	Mnemonic string
	Address  string
	findList []string
	tt       = tools.TimeTest{}
	index    = 0
)

func run() {
	bitSize := 256
	mnemonic, _ := bip44.NewMnemonic(bitSize)
	seedBytes, _ := mnemonic.NewSeed("")
	xKey, _ := bip44.NewKeyFromSeedBytes(seedBytes, bip44.MAINNET)
	accountKey, _ := xKey.BIP44AccountKey(bip44.BitcoinCoinType, 0, false)
	externalAddress, _ := accountKey.DeriveP2PKAddress(bip44.ExternalChangeType, 0, bip44.MAINNET)
	value := externalAddress.Value

	for _, comp := range findList {
		if strings.HasPrefix(value[1:], comp) {
			Mnemonic = mnemonic.Value
			Address = value
			ans := Address + "  " + Mnemonic + "\n"
			fmt.Println(ans)
			writeFile(comp, ans)
			//println(err)
			//ansChan <- ans
			tt.End()
			tt.Start("")

		}
	}
	index += 1
	//fmt.Println(index)
	if index%1000 == 0 {
		tt.Ding()
	}
}

func main() {

	os.Mkdir("./ans", os.ModePerm)
	inputTxt := readFile()
	inputTxt = strings.Replace(inputTxt, " ", "", -1)
	inputTxt = strings.Trim(inputTxt, "\n")

	if inputTxt == "" {
		return
	}
	findList = strings.Split(inputTxt, "\n")

	fmt.Println("Target list is", findList, "total:", len(findList))

	runtime.GOMAXPROCS(3)

	//ansChan := make(chan string)

	tt.Start("")
	for {
		run()

	}
	tt.End()

}
