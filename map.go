package main

import "fmt"

func main() {
	capitals := make(map[string] string)

	capitals["日本"] = "東京"
	capitals["アメリカ"] = "ワシントン"
	capitals["中国"] = "北京"

	fmt.Println(capitals["日本"])
	fmt.Println(capitals["アメリカ"])
	fmt.Println(capitals["中国"])
	
	fmt.Println()

	capitals["イギリス"] = "ロンドン"

	for country ,capital := range capitals {
		fmt.Println(country, capital)
	}



	capital, ok:= capitals["イギリス"]
	if ok {
		fmt.Println("登録済み", capital)
	} else {
		fmt.Println("未登録")
	}

	fmt.Println()
	delete(capitals, "イギリス")
	capital, ok = capitals["イギリス"]
	if ok {
		fmt.Println("登録済み", capital)
	} else {
		fmt.Println("未登録")
	}

	fmt.Println()
	for country ,capital := range capitals {
		fmt.Println(country, capital)
	}



}
