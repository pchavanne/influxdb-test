// https://www.influxdata.com/blog/getting-started-with-the-influxdb-go-client/
// https://github.com/influxdata/influxdb-client-go
package main

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	url    = "http://192.168.1.205:8086"
	token  = "aOToMeQiEtEMduD6fLFrUfIPI4__5RcEMGBGmWfL8DQMvd4VYt-iRakFZgYi5xjy6J3IpzSysdqu348MpCIb4A=="
	org    = "scm"
	bucket = "a_bucket"
)

var (
	// Create new client with default option for server url authenticate by token
	client = influxdb2.NewClient(url, token)
)

func Create() {
	// user blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking(org, bucket)
	// create point using full params constructor
	start := time.Now()
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	// write point immediately
	writeAPI.WritePoint(context.Background(), p)
	fmt.Printf("pointtook %s\n", time.Since(start))
	time.Sleep(20 * time.Millisecond)

	// create point using fluent style
	start = time.Now()
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45).
		SetTime(time.Now())
	writeAPI.WritePoint(context.Background(), p)
	fmt.Printf("fluenttook %s\n", time.Since(start))
	time.Sleep(20 * time.Millisecond)
	// Or write directly line protocol
	start = time.Now()
	line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%v", 23.5, 45)
	writeAPI.WriteRecord(context.Background(), line)
	fmt.Printf("line protocol took %s\n", time.Since(start))
}

func Read() {}

func Update() {}

func Delete() {}

func main() {
	defer client.Close()

	Create()

}
