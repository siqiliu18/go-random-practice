package flatarr

import (
	"fmt"
	"reflect"
)

// related: https://stackoverflow.com/questions/62744911/flatten-an-array-without-using-flat
func FlaterArray(input []interface{}) []int {
	res := []int{}
	for _, element := range input {
		fmt.Println(element)
		helper(element, &res)
	}
	fmt.Println(res)
	return res
}

func helper(element interface{}, res *[]int) {
	if reflect.TypeOf(element).Kind() == reflect.Array {
		*res = append(*res, element.(int))
	}
	for _, innerElem := range element.([]interface{}) {
		helper(innerElem, res)
	}
}
