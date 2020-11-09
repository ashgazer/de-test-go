package utils

import (
	"sort"
	"io/ioutil"
	"fmt"
	"strconv"
)

// GetFiles helper utility to files names
func GetFiles(dir string) []int{
	files, _ := ioutil.ReadDir(dir)
	test := files
	var filenames []int = make([]int, len(files))

	for x := 0; x < len(files); x++ {
		filenames[x], _ = strconv.Atoi(test[x].Name())
	}

	sort.Ints(filenames)
	fmt.Printf("%v", filenames )
	return filenames

}
