package main

import (
	"bufio"
	"filehanding/utils"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const outputfile string = "sorted_data.csv"
const filename = "data.csv"

func createSortedFiles(output string, header string, folderPath string ){
	of, _ := os.OpenFile(output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0700)
	defer of.Close()
	of.WriteString(header + "\n")
	for _, file := range utils.GetFiles(folderPath) {
		fmt.Println(file)
		dat, _ := ioutil.ReadFile(folderPath + strconv.Itoa(file))
		of.Write(dat)
	}	
}

func main() {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	s.Scan()
	header := s.Text()

	for s.Scan() {
		t := s.Text()
		row := strings.Split(t, ",")
		shitty := fmt.Sprintf("data/%s", row[2])
		if f, err := os.Open(shitty); err != nil {
			os.Create(shitty)
		} else {
			f.Close()
		}
		f, err := os.OpenFile(shitty, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)

		if err != nil {
			fmt.Println(err)
		}

		defer f.Close()
		f.WriteString(t + "\n")
	}

		createSortedFiles(outputfile, header, "./data/")
}










