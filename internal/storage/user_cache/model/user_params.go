package model

// UserParamsCache represents a cache for user params.
type UserParamsCache struct {
	SexID              byte `redis:"sex"`
	PhysicalActivityID byte `redis:"physical_activity"`
	ClimateID          byte `redis:"climate"`
	TimezoneID         byte `redis:"timezone"`
	Weight             byte `redis:"weight"`
	WaterGoal          int  `redis:"water_goal"`
}
