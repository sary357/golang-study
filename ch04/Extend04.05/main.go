package main

import (
	"fmt"
	"os"
	"strings"
)

type locale struct {
	language  string
	territory string
}

func getLocale() map[locale]bool {
	supportLocale := map[locale]bool{
		{"en", "US"}: true,
		{"zh", "TW"}: true,
	}
	return supportLocale
}

func main() {
	found := false
	localeParts := strings.Split(os.Args[1], "_") // 以底線為界分割字串
	if len(localeParts) != 2 {                    // 若分割後的切片長度不是 2, 代表語系格式錯誤
		fmt.Printf("Incorrect input: %v\n", os.Args[1])
		os.Exit(1)
	}
	passedLocale := locale{ // 建立要查詢用的語系結構
		territory: localeParts[1],
		language:  localeParts[0],
	}
	for k, _ := range getLocale() {
		if passedLocale == k {
			found = true
		}
	}
	if found {
		fmt.Printf("Locale supported: %v\n", os.Args[1])
	} else {
		fmt.Printf("Locale not supported: %v\n", os.Args[1])
	}

}
