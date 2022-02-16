package main

import "fmt"

// function that you are looking for
func anagram(arr []string) map[string][]string {
	// result structure
	groups := make(map[string][]string, 0)

	// iterating on our array of words
	for _, w := range arr {
		tmp := []rune(w)
		// put all chars of every word into map of rune
		var chars = make(map[rune]int, 0)
		for _, c := range tmp {
			chars[c]++
		}
		// by using Sprintf i get sorted string of chars, it will be used to name a group of anagram words
		tmp_group := fmt.Sprintf("%v", chars)
		groups[tmp_group] = append(groups[tmp_group], w)

	}
	fmt.Println("result")
	//printing the result
	return groups

}

func main() {
	arr := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	groups := anagram(arr)
	for key, value := range groups {
		if len(value) == 1 {
			continue
		}
		fmt.Println(key, value)
	}
}
