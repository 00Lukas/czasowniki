package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	dbName = "db.txt"
	db     []string
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func readDb() {
	file, err := os.Open(dbName)
	checkError(err)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		db = append(db, scan.Text())
	}
}

func randWord() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(db))
}

func main() {
	var (
		s, x string
	)
	readDb()
	fmt.Println("Press Q for quit")
	for {
		s = db[randWord()]
		s2 := strings.Split(s, "-")
		fmt.Print(s2[0])
		fmt.Scanf("%s", &x)
		if strings.ToLower(x) == "q" {
			break
		} else {
			fmt.Println(s2[1], s2[2], s2[3])
			fmt.Println("-----------")
		}
	}
}
