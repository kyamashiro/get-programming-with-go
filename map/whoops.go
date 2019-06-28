package main

import "fmt"

func main() {
	planets := map[string]string{
		"地球": "Sector ZZ9",
		"火星": "Sector ZZ9",
	}

	planetsMarksii := planets
	planets["地球"] = "whoops"

	fmt.Println(planets)
	// mapは同じ基底配列を共有するので値が書き換わる
	fmt.Println(planetsMarksii)

	delete(planets, "地球")
	fmt.Println(planetsMarksii)
}
