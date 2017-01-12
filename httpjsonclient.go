package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	last = flag.Duration("last", 1*time.Minute, "last amount of time to get logs for")
)

func main() {
	fmt.Println("hello Http json API Example")
	// Here we parse our flags
	flag.Parse()

	end := time.Now()
	start := end.Add(*last * -1)
	fmt.Printf("Selected from %s to %s", start, end)

	// our api takes the time arguments as time in unix nano seconds
	url := fmt.Sprintf("http://192.168.1.181:25000/api/date/%d/%d", start.UnixNano(), end.UnixNano())
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Request Failed: ", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	// Always defer the close of the response body
	defer resp.Body.Close()

	// Create a new array that will hold our strutured data
	var data []SensorData

	// decode json into our struct
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println(err)
	}

	// print the results
	for i := range data {
		fmt.Printf("%+v\n", data[i])
	}
}

type SensorData struct {
	UnixTime   int64                `json:"unixtime"`
	IR         IRSensorData         `json:"infrared"`
	Humidity   HumiditySensorData   `json:"humidity"`
	Barometric BarometricSensorData `json:"barometric"`
	Optical    OpticalSensorData    `json:"optical"`
	Motion     MotionSensorData     `json:"motion"`
	RSSI       float64              `json:"rssi"`
	Battery    float64              `json:"battery"`
}
type IRSensorData struct {
	Ambient float64 `json:"ambient"`
	Object  float64 `json:"object"`
}
type HumiditySensorData struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
type BarometricSensorData struct {
	Temperature float64 `json:"temperature"`
	Pressure    float64 `json:"pressure"`
}

type OpticalSensorData struct {
	Light float64 `json:"light"`
}

type MotionSensorData struct {
	Accelerometer AccelerometerSensorData `json:"accelerometer"`
	Gyroscope     GyroscopeSensorData     `json:"gyroscope"`
	Magnetometer  MagnetometerSensorData  `json:"magnetometer"`
}

type AccelerometerSensorData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type GyroscopeSensorData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
type MagnetometerSensorData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
