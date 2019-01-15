package csv_helpers

import (
	"encoding/csv"
	"golang.org/x/net/context"
	"io"
	"os"
)

func ReadDataFromCSV(filename string, context context.Context, dataChan chan DataLine) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	cr := csv.NewReader(f)
	cr.Comma = rune(';')
	cr.Comment = rune('#')
	cr.ReuseRecord = true

	for {
		d, err := cr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		select {
		case dataChan <- *csvLineToDataLine(d):
		case <-context.Done():
			return context.Err()
		}
	}
	return nil
}

func csvLineToDataLine(csvLine []string) *DataLine {
	return &DataLine{
		StartIP:           csvLine[0],
		EndIP:             csvLine[1],
		Country:           csvLine[2],
		Region:            csvLine[3],
		RegionCode:        csvLine[4],
		City:              csvLine[5],
		CityCode:          csvLine[6],
		ConnSpeed:         csvLine[7],
		ISP:               csvLine[8],
		MobileCarrier:     csvLine[9],
		MobileCarrierCode: csvLine[10],
	}
}

type DataLine struct {
	StartIP           string `csv:"start-ip"`
	EndIP             string `csv:"end-ip"`
	Country           string `csv:"edge-two-letter-country"`
	Region            string `csv:"edge-region"`
	RegionCode        string `csv:"edge-region-code"`
	City              string `csv:"edge-city"`
	CityCode          string `csv:"edge-city-code"`
	ConnSpeed         string `csv:"edge-conn-speed"`
	ISP               string `csv:"isp-name"`
	MobileCarrier     string `csv:"mobile-carrier"`
	MobileCarrierCode string `csv:"mobile-carrier-code"`
}
