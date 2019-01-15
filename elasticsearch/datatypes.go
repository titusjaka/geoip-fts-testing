package elasticsearch

import (
	"github.com/titusjaka/geoip-fts-testing/csv-helpers"
)

type ElasticObject struct {
	ID                string         `json:"-"`
	IPAddress         ElasticIpRange `json:"ip_address"`
	Country           string         `json:"country"`
	Region            string         `json:"region"`
	RegionCode        string         `json:"region_code"`
	City              string         `json:"city"`
	CityCode          string         `json:"city_code"`
	ConnSpeed         string         `json:"conn_speed"`
	ISP               string         `json:"isp"`
	MobileCarrier     string         `json:"mobile_carrier"`
	MobileCarrierCode string         `json:"mobile_carrier_code"`
}

type ElasticIpRange struct {
	StartIP string `json:"gte"`
	EndIP   string `json:"lte"`
}

func (eo *ElasticObject) ToDataLine() *csv_helpers.DataLine {
	return &csv_helpers.DataLine{
		StartIP:           eo.IPAddress.StartIP,
		EndIP:             eo.IPAddress.EndIP,
		Country:           eo.Country,
		Region:            eo.Region,
		RegionCode:        eo.RegionCode,
		City:              eo.City,
		CityCode:          eo.CityCode,
		ConnSpeed:         eo.ConnSpeed,
		ISP:               eo.ISP,
		MobileCarrier:     eo.MobileCarrier,
		MobileCarrierCode: eo.MobileCarrierCode,
	}
}

func DataLineToElasticObject(dl *csv_helpers.DataLine) *ElasticObject {
	return &ElasticObject{
		IPAddress: ElasticIpRange{
			StartIP: dl.StartIP,
			EndIP:   dl.EndIP,
		},
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
