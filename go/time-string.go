package main

import (
    "fmt"
    "time"
)

// GetCurrentTimeString returns the current time as a string in the format "yyyy_mm_dd_hh_mm_ss".
func GetCurrentTimeString() string {
    return time.Now().Format("2006_01_02_15_04_05")
}

func main() {
    currentTimeString := GetCurrentTimeString()
    fmt.Println(currentTimeString)
}
