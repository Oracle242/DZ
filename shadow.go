package main

import "fmt"

func main() {
	channels := [7]string{"NBC", "CBS", "ABC", "Fox", "CW", "Russia1", "TNT"}
	channels2 := []string{}
	tricolor := make(map[string]string)

	for i := 0; i < len(channels); i++ {
		min := channels[0]
		minIndex := 0
		for j := 0; j < len(channels); j++ {
			if channels[j] < min {
				min = channels[j]
				minIndex = j
			}
		}
		channels[minIndex] = "ZZZZ"
		channels2 = append(channels2, min)
	}

	channelSlEng := channels2[:4]
	for _, meaning := range channelSlEng {
		tricolor[meaning] = "Eng"
	}

	channeSlRus := channels2[5:7]
	for _, meaning := range channeSlRus {
		tricolor[meaning] = "rus"
	}

	// fmt.Println(channels2)
	// fmt.Println(channelSlEng)
	// fmt.Println(channeSlRus)
	fmt.Println(tricolor)
}

// разобраться с range(с разными вариациями)
// Изучить раздел misc.
// Практика слайсов.
