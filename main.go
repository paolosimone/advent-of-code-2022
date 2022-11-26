package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/paolosimone/advent-of-code-2022/advent"
)

const banner string = `
  ___   __  _  _  ____  _  _  _  _    __ _   __  ____  __   
 / __) /  \( \/ )(  __)/ )( \( \/ )  (  ( \ /  \(  __)(  )  
( (_ \(  O ))  /  ) _) ) \/ ( )  (   /    /(  O )) _) / (_/\
 \___/ \__/(__/  (____)\____/(_/\_)  \_)__) \__/(____)\____/

`

func main() {
	fmt.Print(banner)
	results := runDays()
	report := advent.BuildReport(results)
	fmt.Println()
	fmt.Println(report.Sprint())
}

func runDays() []advent.DayResult {
	if len(os.Args) == 1 {
		return advent.RunAllDays()
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Print(err)
		return []advent.DayResult{}
	}

	return []advent.DayResult{advent.RunDay(number)}
}
