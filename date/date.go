package main

import (
    "fmt"
    "time"
)
var week time.Duration
func main()  {
    t := time.Now()
    fmt.Println(t)
    fmt.Println(t.Format("02 Jan 2006 15:04"))
    fmt.Printf("%02d.%02d.%4d\n",t.Day(),t.Month(),t.Year())
    fmt.Println(t.UTC())
    week = 60 * 60 * 24 * 7 *1e9
    weekFromNow := t.Add(week)
    fmt.Println(weekFromNow)
    fmt.Println(t.Format(time.RFC822))
    fmt.Println(t.Format(time.ANSIC))
    fmt.Println(t.Format("20060102"))
}
