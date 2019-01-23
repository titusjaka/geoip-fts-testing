package elasticsearch

import (
	"github.com/gosexy/to"
	"github.com/titusjaka/geoip-fts-testing/csv-helpers"
)

type ElasticObject struct {
	ID                string         `json:"-"`
	IPAddress         ElasticIpRange `json:"ip_address"`
	Country           string         `json:"country"`
	Region            string         `json:"region"`
	RegionCode        uint32         `json:"region_code"`
	City              string         `json:"city"`
	CityCode          uint32         `json:"city_code"`
	ConnSpeed         string         `json:"conn_speed"`
	ISP               string         `json:"isp"`
	MobileCarrier     string         `json:"mobile_carrier"`
	MobileCarrierCode uint32         `json:"mobile_carrier_code"`
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
		RegionCode:        to.String(eo.RegionCode),
		City:              eo.City,
		CityCode:          to.String(eo.CityCode),
		ConnSpeed:         eo.ConnSpeed,
		ISP:               eo.ISP,
		MobileCarrier:     eo.MobileCarrier,
		MobileCarrierCode: to.String(eo.MobileCarrierCode),
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
		RegionCode:        uint32(to.Uint64(dl.RegionCode)),
		City:              dl.City,
		CityCode:          uint32(to.Uint64(dl.CityCode)),
		ConnSpeed:         dl.ConnSpeed,
		ISP:               dl.ISP,
		MobileCarrier:     dl.MobileCarrier,
		MobileCarrierCode: uint32(to.Uint64(dl.MobileCarrierCode)),
	}
}
