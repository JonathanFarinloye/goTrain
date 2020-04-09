package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

func main() {
	jobs := []int {100, 250, 50, 20}
	// fixedPartition(jobs)
	fmt.Println(dynamicPartition(jobs, 300))
}


func dynamicPartition(jobs []int, memory int) []int {
	blockView := []int {}
	for _, value := range jobs {
		if value <= memory {
			memory -= value
			blockView = append(blockView, value)
		}
	}
	return blockView
}


func fixedPartition(jobs []int) {
	parts := map[string][]int {
		"firstPart": {100, 0, 0},
		"secondPart": {70, 0, 0},
		"thirdPart": {200, 0, 0},
	}
	for i := 0; i < len(jobs); i++ {
		for _, value := range parts {
			if value[0] > jobs[i] && value[2] != 1 {
				value[2] = 1
				value[1] = jobs[i]
				fmt.Println(value[0], jobs[i])
			}
		}
	}
	fmt.Println(parts)
}

// func evaporator(content float64, evapPerDay int, threshold int) int {
// 	var count int 
// 	thresholdMl := float64(threshold) * 0.01 * content
// 	for {
// 		if content <= thresholdMl {
// 			break
// 		}
// 		evapML := float64(evapPerDay) * .01 * content
// 		content = content - evapML
// 		count ++
// 	}
// 	return count
// }


// func divisors(n int) int {
// 	sum := 0
// 	for i := 1 ;i <= n; i++ {
// 		if n % i == 0 {
// 			sum ++
// 		}
// 	}
// 	return sum
// }


// func duplicateCount(s string) int {
	// Solve soon
// }

// func goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
// 	return laLigaGoals + copaDelReyGoals + championsLeagueGoals
//   }

// func nbYear(p0 int, percent float64, aug int, p int) int {
// 	years := 0
// 	for p0 < p {
// 		years ++
// 		p0 += int((float64(p0) * (percent/100))) + aug
// 	}
// 	return years
// }

// func duplicateEncode(word string) string {
// 	var newName string
// 	for _, elem := range word {
// 		letter := strings.ToLower(string(elem))
// 		if strings.Count(strings.ToLower(word), letter) != 1 {
// 			newName += ")"
// 		} else {
// 			newName += "("
// 		}
// 	}
// 	return newName
// }