package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type (
	Key int
)

const (
	services = "Internal Service"
	LogKey   = Key(48)
)

// Data is data standard output
type Data struct {
	RequestID     string    `json:"RequestID"`
	UserId        string    `json:"user_id"`
	TimeStart     time.Time `json:"TimeStart"`
	UserCode      string    `json:"UserCode"`
	Device        string    `json:"Device"`
	Host          string    `json:"Host"`
	Endpoint      string    `json:"Endpoint"`
	RequestMethod string    `json:"RequestMethod"`
	RequestHeader string    `json:"RequestHeader"`
	StatusCode    int       `json:"StatusCode"`
	Response      string    `json:"Response"`
	ExecTime      float64   `json:"ExecutionTime"`
	Messages      []string  `json:"Messages"`
}

type TrackIP struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

type Captha struct {
	Result string `json:"result"`
}

type ReqBulkByID struct {
	ID []int `json:"id"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()

	return scope.SetColumn("ID", uuid)
}

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" schema:"ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
