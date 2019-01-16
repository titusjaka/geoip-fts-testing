package xapiandsearch

import (
	"github.com/titusjaka/geoip-fts-testing"
	"github.com/titusjaka/geoip-fts-testing/csv-helpers"
	"math/big"
)

type XapiandObject struct {
	ID                string   `json:"id"`
	StartIP           *big.Int `json:"start_ip"`
	EndIP             *big.Int `json:"end_ip"`
	Country           string   `json:"country"`
	Region            string   `json:"region"`
	RegionCode        string   `json:"region_code"`
	City              string   `json:"city"`
	CityCode          string   `json:"city_code"`
	ConnSpeed         string   `json:"conn_speed"`
	ISP               string   `json:"isp"`
	MobileCarrier     string   `json:"mobile_carrier"`
	MobileCarrierCode string   `json:"mobile_carrier_code"`
}

func (xo *XapiandObject) ToDataLine() *csv_helpers.DataLine {
	return &csv_helpers.DataLine{
		StartIP:           geoip_fts_testing.IntToIp(xo.StartIP),
		EndIP:             geoip_fts_testing.IntToIp(xo.EndIP),
		Country:           xo.Country,
		Region:            xo.Region,
		RegionCode:        xo.RegionCode,
		City:              xo.City,
		CityCode:          xo.CityCode,
		ConnSpeed:         xo.ConnSpeed,
		ISP:               xo.ISP,
		MobileCarrier:     xo.MobileCarrier,
		MobileCarrierCode: xo.MobileCarrierCode,
	}
}

func DataLineToXapiandObject(dl *csv_helpers.DataLine) *XapiandObject {
	return &XapiandObject{
		StartIP:           geoip_fts_testing.IpToInt(dl.StartIP),
		EndIP:             geoip_fts_testing.IpToInt(dl.EndIP),
		Country:           dl.Country,
		Region:            dl.Region,
		RegionCode:        dl.RegionCode,
		City:              dl.City,
		CityCode:          dl.CityCode,
		ConnSpeed:         dl.ConnSpeed,
		ISP:               dl.ISP,
		MobileCarrier:     dl.MobileCarrier,
		MobileCarrierCode: dl.MobileCarrierCode,
	}
}
