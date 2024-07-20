package model

type EnvironmentScore struct {
	CO2Value        int    `json:"co2_value"`
	CO2Stars        int    `json:"co2_stars"`
	WaterValue      int    `json:"water_value"`
	WaterStars      int    `json:"water_stars"`
	RainforestValue string `json:"rainforest_value"`
	RainforestStars int    `json:"rainforest_stars"`
}
