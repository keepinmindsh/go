package main

import(
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		return 0, err
	}

	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

func main() {
	h1, err := getHash("text1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	h2, err := getHash("text2.txt")
	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)
}
