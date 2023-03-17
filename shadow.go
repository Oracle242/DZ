package main

import "fmt"

func main() {
	channels := [7]string{"NBC", "CBS", "ABC", "Fox", "CW", "Russia1", "TNT"}

	tricolor := make(map[string]string)

	for i := 0; i < len(channels); i++ {
		for j := 0; j < len(channels)-1; j++ {
			if channels[j] > channels[j+1] {
				cloudChannels := channels[j+1]
				channels[j+1] = channels[j]
				channels[j] = cloudChannels
			}
		}
	}

	channelSlEng := channels[:4]
	for _, meaning := range channelSlEng {
		tricolor[meaning] = "Eng"
	}

	channeSlRus := channels[5:7]
	for _, meaning := range channeSlRus {
		tricolor[meaning] = "rus"
	}

	fmt.Println(channels)
	fmt.Println(channelSlEng)
	fmt.Println(channeSlRus)
	fmt.Println(tricolor)
}

// разобраться с range(с разными вариациями)
// Изучить раздел misc.
// Практика слайсов.
