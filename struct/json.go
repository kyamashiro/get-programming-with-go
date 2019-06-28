package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// jsonパッケージから見えるようにフィールド名を大文字にする.
	type location struct {
		// 構造体タグ key:"value"という形に変換する
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	curiosity := location{-4.5, 137.4}

	bytes, err := json.Marshal(curiosity)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))

	bytes, err = json.MarshalIndent(curiosity, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))

}
