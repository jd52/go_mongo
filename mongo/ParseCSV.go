package mongo

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
)

func ParseCSV(mf multipart.File) {
	fmt.Println("in parse CSV")
	csvFile := mf
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var devices []Device
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		devices = append(devices, Device{
			Hostname:   line[0],
			IPAddress:  line[1],
			DeviceType: line[2],
		})
	}
	devicesJSON, _ := json.Marshal(devices)
	fmt.Println(string(devicesJSON))
}
