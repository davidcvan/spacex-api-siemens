package models

type Dragon struct {
	HeatShield struct {
		Material    string  `json:"material"`
		SizeMeters  float64 `json:"size_meters"`
		TempDegrees int     `json:"temp_degrees"`
		DevPartner  string  `json:"dev_partner"`
	} `json:"heat_shield"`
	LaunchPayloadMass struct {
		Kg int `json:"kg"`
		Lb int `json:"lb"`
	} `json:"launch_payload_mass"`
	LaunchPayloadVol struct {
		CubicMeters int `json:"cubic_meters"`
		CubicFeet   int `json:"cubic_feet"`
	} `json:"launch_payload_vol"`
	ReturnPayloadMass struct {
		Kg int `json:"kg"`
		Lb int `json:"lb"`
	} `json:"return_payload_mass"`
	ReturnPayloadVol struct {
		CubicMeters int `json:"cubic_meters"`
		CubicFeet   int `json:"cubic_feet"`
	} `json:"return_payload_vol"`
	PressurizedCapsule struct {
		PayloadVolume struct {
			CubicMeters int `json:"cubic_meters"`
			CubicFeet   int `json:"cubic_feet"`
		} `json:"payload_volume"`
	} `json:"pressurized_capsule"`
	Trunk struct {
		TrunkVolume struct {
			CubicMeters int `json:"cubic_meters"`
			CubicFeet   int `json:"cubic_feet"`
		} `json:"trunk_volume"`
		Cargo struct {
			SolarArray         int  `json:"solar_array"`
			UnpressurizedCargo bool `json:"unpressurized_cargo"`
		} `json:"cargo"`
	} `json:"trunk"`
	HeightWTrunk struct {
		Meters float64 `json:"meters"`
		Feet   float64 `json:"feet"`
	} `json:"height_w_trunk"`
	Diameter struct {
		Meters float64 `json:"meters"`
		Feet   int     `json:"feet"`
	} `json:"diameter"`
	FirstFlight      string   `json:"first_flight"`
	FlickrImages     []string `json:"flickr_images"`
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Active           bool     `json:"active"`
	CrewCapacity     int      `json:"crew_capacity"`
	SidewallAngleDeg int      `json:"sidewall_angle_deg"`
	OrbitDurationYr  int      `json:"orbit_duration_yr"`
	DryMassKg        int      `json:"dry_mass_kg"`
	DryMassLb        int      `json:"dry_mass_lb"`
	Thrusters        []struct {
		Type   string `json:"type"`
		Amount int    `json:"amount"`
		Pods   int    `json:"pods"`
		Fuel1  string `json:"fuel_1"`
		Fuel2  string `json:"fuel_2"`
		Isp    int    `json:"isp"`
		Thrust struct {
			KN  float64 `json:"kN"`
			Lbf int     `json:"lbf"`
		} `json:"thrust"`
	} `json:"thrusters"`
	Wikipedia   string `json:"wikipedia"`
	Description string `json:"description"`
	ID          string `json:"id"`
}
