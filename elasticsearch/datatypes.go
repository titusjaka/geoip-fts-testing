package elasticsearch

import (
	"github.com/gosexy/to"
	"github.com/titusjaka/geoip-fts-testing/csv-helpers"
)

type Info struct {
	IPAddress     IpRange `json:"ip_address"`
	CountryCode   string  `json:"country_code"`
	Region        string  `json:"region"`
	RegionCode    uint32  `json:"region_code"`
	City          string  `json:"city"`
	CityCode      uint32  `json:"city_code"`
	ConnSpeed     string  `json:"conn_speed"`
	MobileISP     string  `json:"mobile_isp"`
	MobileISPCode uint32  `json:"mobile_isp_code"`
	ProxyType     string  `json:"proxy_type"`
}

type Country struct {
	Iso   string `json:"iso"`
	Title string `json:"title"`
}

type IpRange struct {
	StartIP string `json:"gte"`
	EndIP   string `json:"lte"`
}

func (i *Info) ToDataLine() *csv_helpers.GeoInfoLine {
	return &csv_helpers.GeoInfoLine{
		StartIP:           i.IPAddress.StartIP,
		EndIP:             i.IPAddress.EndIP,
		CountryCode:       i.CountryCode,
		Region:            i.Region,
		RegionCode:        to.String(i.RegionCode),
		City:              i.City,
		CityCode:          to.String(i.CityCode),
		ConnSpeed:         i.ConnSpeed,
		MobileCarrier:     i.MobileISP,
		MobileCarrierCode: to.String(i.MobileISPCode),
		ProxyType:         i.ProxyType,
	}
}

func GeoInfiLineToElasticObject(dl *csv_helpers.GeoInfoLine) *Info {
	return &Info{
		IPAddress: IpRange{
			StartIP: dl.StartIP,
			EndIP:   dl.EndIP,
		},
		CountryCode:   dl.CountryCode,
		Region:        dl.Region,
		RegionCode:    uint32(to.Uint64(dl.RegionCode)),
		City:          dl.City,
		CityCode:      uint32(to.Uint64(dl.CityCode)),
		ConnSpeed:     dl.ConnSpeed,
		MobileISP:     dl.MobileCarrier,
		MobileISPCode: uint32(to.Uint64(dl.MobileCarrierCode)),
		ProxyType:     dl.ProxyType,
	}
}

func CountryInfoToElasticObject(line *csv_helpers.CountryLine) *Country {
	return &Country{
		Iso:   line.Iso3,
		Title: line.CountryName,
	}
}
