package models

//Capsule are struct for Capsules
type Capsule struct {
	ReuseCount    int      `json:"reuse_count"`
	WaterLandings int      `json:"water_landings"`
	LandLandings  int      `json:"land_landings"`
	LastUpdate    string   `json:"last_update"`
	Launches      []string `json:"launches"`
	Serial        string   `json:"serial"`
	Status        string   `json:"status"`
	Type          string   `json:"type"`
	ID            string   `json:"id"`
}
