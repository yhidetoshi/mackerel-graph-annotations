package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mackerelio/mackerel-client-go"
)

var (
	service = "NatureRemo"
	title = "たいとる"
	description = "Description"
	mkrKey  = os.Getenv("MKRKEY")
	client  = mackerel.NewClient(mkrKey)
)
/*
const (
	timezone = "Asia/Tokyo"
	offset   = 9 * 60 * 60
)
*/

func main() {
	//jst := time.FixedZone(timezone, offset)
	//nowTime := time.Now().In(jst)
	nowTime := time.Now()
	toTime := nowTime.Add(1 * time.Minute)

	_, err := client.CreateGraphAnnotation(&mackerel.GraphAnnotation{
		Title:       title,
		Description: description,
		From:        nowTime.Unix(),
		To:          toTime.Unix(),
		Service:     service,
	})
	if err != nil {
		fmt.Println(err)
	}
}
