package csv_helpers

import (
	"encoding/csv"
	"golang.org/x/net/context"
	"io"
	"os"
)

func ReadGeoInfoFromCSV(filename string, context context.Context, geoInfoChan chan GeoInfoLine) (err error) {
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
		case geoInfoChan <- *csvLineToGeoInfoLine(d):
		case <-context.Done():
			return context.Err()
		}
	}
	return nil
}

func ReadCountryInfoFromCSV(filename string, context context.Context, geoInfoChan chan CountryLine) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	cr := csv.NewReader(f)
	cr.Comma = rune(',')
	//cr.Comment = rune('#')
	cr.ReuseRecord = true
	var readFirstLine bool

	for {
		d, err := cr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if !readFirstLine {
			readFirstLine = true
			continue
		}
		select {
		case geoInfoChan <- *csvLineToCountryLine(d):
		case <-context.Done():
			return context.Err()
		}
	}
	return nil
}

func csvLineToGeoInfoLine(csvLine []string) *GeoInfoLine {
	return &GeoInfoLine{
		StartIP:           csvLine[0],
		EndIP:             csvLine[1],
		CountryCode:       csvLine[2],
		Region:            csvLine[3],
		RegionCode:        csvLine[4],
		City:              csvLine[5],
		CityCode:          csvLine[6],
		ConnSpeed:         csvLine[7],
		MobileCarrier:     csvLine[8],
		MobileCarrierCode: csvLine[9],
		ProxyType:         csvLine[10],
	}
}

func csvLineToCountryLine(csvLine []string) *CountryLine {
	return &CountryLine{
		Iso3:          csvLine[0],
		Iso2:          csvLine[1],
		CountryName:   csvLine[2],
		Regions:       csvLine[3],
		ContinentCode: csvLine[4],
		ContinentName: csvLine[5],
		CountryCode:   csvLine[6],
	}
}

type GeoInfoLine struct {
	StartIP           string `csv:"start-ip"`
	EndIP             string `csv:"end-ip"`
	CountryCode       string `csv:"edge-two-letter-country"`
	Region            string `csv:"edge-region"`
	RegionCode        string `csv:"edge-region-code"`
	City              string `csv:"edge-city"`
	CityCode          string `csv:"edge-city-code"`
	ConnSpeed         string `csv:"edge-conn-speed"`
	MobileCarrier     string `csv:"mobile-carrier"`
	MobileCarrierCode string `csv:"mobile-carrier-code"`
	ProxyType         string `csv:"proxy-type"`
}

type CountryLine struct {
	Iso3          string `csv:"ISO-3"`
	Iso2          string `csv:"ISO-2"`
	CountryName   string `csv:"COUNTRY-NAME"`
	Regions       string `csv:"REGIONS"`
	ContinentCode string `csv:"CONTINENT-CODE"`
	ContinentName string `csv:"CONTINENT-NAME"`
	CountryCode   string `csv:"COUNTRY-CODE"`
}
