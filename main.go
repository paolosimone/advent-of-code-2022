package main

import (
	"fmt"

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
	results := advent.RunAllDays()
	report := advent.BuildReport(results)
	fmt.Println()
	fmt.Println(report.Sprint())
}
