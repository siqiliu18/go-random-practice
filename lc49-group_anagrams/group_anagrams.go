package lc49groupanagrams

import "sort"

type Anagrams struct {
	name    string
	letters map[string]int
	visited bool
	others  []string
	count   int
}

func groupAnagramsBFS(strs []string) [][]string {
	res := [][]string{}

	dataStruct := make(map[string]*Anagrams)

	for i := 0; i < len(strs); i++ {
		others := []string{}
		others = append(others, strs[:i]...)
		others = append(others, strs[i+1:]...)
		if _, foundWord := dataStruct[strs[i]]; !foundWord {
			letters := make(map[string]int)
			for _, c := range strs[i] {
				if _, foundLetter := letters[string(c)]; !foundLetter {
					letters[string(c)] = 1
				} else {
					letters[string(c)] += 1
				}
			}
			dataStruct[strs[i]] = &Anagrams{
				name:    strs[i],
				others:  others,
				letters: letters,
			}
		} else {
			dataStruct[strs[i]].count += 1
		}
	}

	for word, property := range dataStruct {
		if property.visited {
			continue
		}
		queue := []string{word}
		arr := []string{}

		for len(queue) != 0 {
			curr := queue[0]
			arr = append(arr, dataStruct[curr].name)
			for i := 1; i <= property.count; i++ {
				arr = append(arr, dataStruct[curr].name)
			}
			queue = queue[1:]
			dataStruct[curr].visited = true

			for _, other := range dataStruct[curr].others {
				if !dataStruct[other].visited {
					pass := true
					if dataStruct[curr].name == "" {
						pass = false
					}
					for k, cv := range dataStruct[curr].letters {
						if ov, foundLetter := dataStruct[other].letters[k]; len(dataStruct[curr].letters) != len(dataStruct[other].letters) || !foundLetter || cv != ov {
							pass = false
						}
					}
					if pass {
						dataStruct[other].visited = true
						queue = append(queue, other)
					}
				}
			}
		}
		res = append(res, arr)
	}

	return res
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	wordMap := make(map[string][]string)
	for _, word := range strs {
		toSort := []rune(word)
		sort.SliceStable(toSort, func(i, j int) bool {
			return toSort[i] < toSort[j]
		})
		wordMap[string(toSort)] = append(wordMap[string(toSort)], word)
	}
	for _, v := range wordMap {
		res = append(res, v)
	}
	return res
}
