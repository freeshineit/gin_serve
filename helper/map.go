package helper

import "fmt"

// merge map
func MergeMap[K comparable, V any](mObj ...map[K]V) map[K]V {

	newObj := make(map[K]V)

	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}

	fmt.Println(newObj)

	return newObj
}
