package elasticsearch

import (
	"github.com/gosexy/to"
	"github.com/titusjaka/geoip-fts-testing/csv-helpers"
)

type Info struct {
	IPAddress     IpRange `json:"ip_address"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"country_code"`
	Region        string  `json:"region"`
	RegionCode    uint32  `json:"region_code"`
	City          string  `json:"city"`
	CityCode      uint32  `json:"city_code"`
	ConnSpeed     string  `json:"conn_speed"`
	ISP           string  `json:"isp"`
	MobileISP     string  `json:"mobile_isp"`
	MobileISPCode uint32  `json:"mobile_isp_code"`
	ProxyType     string  `json:"proxy_type"`
}

type IpRange struct {
	StartIP string `json:"gte"`
	EndIP   string `json:"lte"`
}

func (eo *Info) ToDataLine() *csv_helpers.DataLine {
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
		MobileCarrier:     eo.MobileISP,
		MobileCarrierCode: to.String(eo.MobileISPCode),
		ProxyType:         eo.ProxyType,
	}
}

func DataLineToElasticObject(dl *csv_helpers.DataLine) *Info {
	return &Info{
		IPAddress: IpRange{
			StartIP: dl.StartIP,
			EndIP:   dl.EndIP,
		},
		Country:       dl.Country,
		Region:        dl.Region,
		RegionCode:    uint32(to.Uint64(dl.RegionCode)),
		City:          dl.City,
		CityCode:      uint32(to.Uint64(dl.CityCode)),
		ConnSpeed:     dl.ConnSpeed,
		ISP:           dl.ISP,
		MobileISP:     dl.MobileCarrier,
		MobileISPCode: uint32(to.Uint64(dl.MobileCarrierCode)),
		ProxyType:     dl.ProxyType,
	}
}
