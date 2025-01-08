package constants

import "github.com/go-telegram/fsm"

// Constants for FSM States
const (
	StateDefault                    fsm.StateID = "default"
	StateOnboardingWaterNorm        fsm.StateID = "onboarding_water_norm"
	StateOnboardingWeight           fsm.StateID = "onboarding_weight"
	StateOnboardingSex              fsm.StateID = "onboarding_sex"
	StateOnboardingPhysicalActivity fsm.StateID = "onboarding_physical_activity"
	StateOnboardingClimate          fsm.StateID = "onboarding_climate"
	StateOnboardingTimezone         fsm.StateID = "onboarding_timezone"
)
