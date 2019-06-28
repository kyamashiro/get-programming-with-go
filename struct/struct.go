package main

import "fmt"

/*
構造体の初期化について
基本的に構造体を作成する際はフィールドを指定する。
そうすればフィールドが追加されたりフィールドの順序が変わっても問題ないので変更に強い。
*/
func main() {
	type location struct {
		lat, long float64
	}

	// 複合リテラルで初期化
	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity)

	// 構造体のフィールドを指定する場合は順序は任意
	insight := location{long: 135.9, lat: 4.5}
	fmt.Println(insight)

	// フィールドを指定しない場合、構造体定義のリストと同じ順序で初期化される
	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit)

	// 初期化しない場合、フィールドの値はそれぞれの型のゼロになる
	none := location{}
	fmt.Println(none)

	// コピーが渡される
	bradbury := location{lat: -4.5895, long: 137.4417}
	curiosity := bradbury
	curiosity.long = 0
	fmt.Println(bradbury, curiosity)
}
