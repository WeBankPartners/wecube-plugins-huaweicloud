package products

import (
	"github.com/huaweicloud/golangsdk"
)

// GetResponse response
type GetResponse struct {
	Products []Product `json:"products"`
}

// Product for dcs
type Product struct {
	Price        float64 `json:"price"`
	Currency     string  `json:"currency"`
	ProductID    string  `json:"product_id"`
	SpecCode     string  `json:"spec_code"`
	SpecDetails  string  `json:"spec_details"`
	ChargingType string  `json:"charging_type"`
	SpecDetails2 string  `json:"spec_details2"`
	ProdType     string  `json:"prod_type"`

	CacheMode    string   `json:"cache_mode"`
	CpuType      string   `json:"cpu_type"`
	StorageType  string   `json:"storage_type"`
	Detail       string   `json:"detail"`
	Engine       string   `json:"engine"`
	EngineVersions string  `json:"engine_versions"`
	Flavors    []Flavor    `json:"flavors"`
}

type Flavor struct {
	Capacity string `json:"capacity"`
	Unit     string `json:"unit"`
	Azs    []string `json:"available_zones"`
}

// GetResult contains the body of getting detailed
type GetResult struct {
	golangsdk.Result
}

// Extract from GetResult
func (r GetResult) Extract() (*GetResponse, error) {
	var s GetResponse
	err := r.Result.ExtractInto(&s)
	return &s, err
}
