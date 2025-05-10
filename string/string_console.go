package string

import (
	"fmt"
	"unicode/utf8"
)

func MainStingLen() {
	str := "ddЯЙ漢"

	//Так как Кириллица и иероглифы имеют больше байт, len выводит > 5
	fmt.Println("Длина через len", len(str))                                   // > 5
	fmt.Println("Длина через RuneCountInStrings", utf8.RuneCountInString(str)) // 5

	str2 := "G👳‍o"
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("%c\n", str2[i]) // выведет "G", "непонятные символы", "o"
	//}

	// Исправленный вариант
	for _, val := range str2 {
		fmt.Println(val)
	}

}
