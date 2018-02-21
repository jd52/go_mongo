package database

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
)

//ParseCSV is used to imort a CSV file.
func ParseCSV(mf multipart.File) []MongoDevice {
	fmt.Println("in parse CSV")
	csvFile := mf
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var devices []MongoDevice
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		devices = append(devices, MongoDevice{
			Hostname:   line[0],
			IPAddress:  line[1],
			DeviceType: line[2],
		})
	}
	devicesJSON, _ := json.Marshal(devices)
	fmt.Println(string(devicesJSON))

	return devices
}
