package src

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LogParser() {
	in := bufio.NewScanner(os.Stdin)

	logParser := make(map[string]int)
	for in.Scan() {
		info := strings.Fields(in.Text())
		domain := info[0]
		noOfTimes, _ := strconv.Atoi(info[1])
		logParser[domain] += noOfTimes
	}
	fmt.Println(logParser)
}
