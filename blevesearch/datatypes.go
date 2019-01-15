package blevesearch

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	"github.com/titusjaka/geoip-fts-testing"
	"github.com/titusjaka/geoip-fts-testing/csv-helpers"
	"math/big"
)

// BleveInfoField is used as enum
type BleveInfoField string

func (bf BleveInfoField) String() string {
	return string(bf)
}

// ToDo: rewrite on custom type tag
// Constants used as enum
const (
	InfoID                BleveInfoField = "ID"
	InfoStartIP           BleveInfoField = "StartIP"
	InfoEndIP             BleveInfoField = "EndIP"
	InfoCountry           BleveInfoField = "Country"
	InfoRegion            BleveInfoField = "Region"
	InfoRegionCode        BleveInfoField = "RegionCode"
	InfoCity              BleveInfoField = "City"
	InfoCityCode          BleveInfoField = "CityCode"
	InfoConnSpeed         BleveInfoField = "ConnSpeed"
	InfoISP               BleveInfoField = "ISP"
	InfoMobileCarrier     BleveInfoField = "MobileCarrier"
	InfoMobileCarrierCode BleveInfoField = "MobileCarrierCode"
)

func DalaLineToBleveInfoObject(dl *csv_helpers.DataLine) *BleveInfoObject {
	return &BleveInfoObject{
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

type BleveInfoObject struct {
	ID                string
	StartIP           *big.Int
	EndIP             *big.Int
	Country           string
	Region            string
	RegionCode        string
	City              string
	CityCode          string
	ConnSpeed         string
	ISP               string
	MobileCarrier     string
	MobileCarrierCode string
}

func (bo *BleveInfoObject) ToDataLine() *csv_helpers.DataLine {
	return &csv_helpers.DataLine{
		StartIP:           geoip_fts_testing.IntToIp(bo.StartIP),
		EndIP:             geoip_fts_testing.IntToIp(bo.EndIP),
		Country:           bo.Country,
		Region:            bo.Region,
		RegionCode:        bo.RegionCode,
		City:              bo.City,
		CityCode:          bo.CityCode,
		ConnSpeed:         bo.ConnSpeed,
		ISP:               bo.ISP,
		MobileCarrier:     bo.MobileCarrier,
		MobileCarrierCode: bo.MobileCarrierCode,
	}
}

func (BleveInfoObject) Type() string {
	return "geoip-info"
}

func (BleveInfoObject) GetDocumentMapping() *mapping.DocumentMapping {
	geoInfoMapping := bleve.NewDocumentStaticMapping()

	geoInfoMapping.AddFieldMappingsAt(InfoStartIP.String(), getNumericMapping())
	geoInfoMapping.AddFieldMappingsAt(InfoEndIP.String(), getNumericMapping())
	geoInfoMapping.AddFieldMappingsAt(InfoCountry.String(), getStandardMapping())
	geoInfoMapping.AddFieldMappingsAt(InfoRegion.String(), getStandardMapping())
	geoInfoMapping.AddFieldMappingsAt(InfoRegionCode.String(), getKeywordMapping())
	geoInfoMapping.AddFieldMappingsAt(InfoCity.String(), getStandardMapping())
	geoInfoMapping.AddFieldMappingsAt(InfoCityCode.String(), getKeywordMapping())
	geoInfoMapping.AddFieldMappingsAt(InfoConnSpeed.String(), getKeywordMapping())
	geoInfoMapping.AddFieldMappingsAt(InfoISP.String(), getKeywordMapping())
	geoInfoMapping.AddFieldMappingsAt(InfoMobileCarrier.String(), getKeywordMapping())
	geoInfoMapping.AddFieldMappingsAt(InfoMobileCarrierCode.String(), getKeywordMapping())

	return geoInfoMapping
}
