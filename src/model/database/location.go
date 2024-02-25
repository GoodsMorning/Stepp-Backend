package database

type LocationDb struct {
	LocationId int64  `json:"location_id"`
	Name       string `json:"name"`
	GeoPoint   string `json:"geo_point"`
	SteppCount int64  `json:"stepp_count"`
}
