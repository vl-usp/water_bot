package model

// Timezone is a model that represents timezone
type Timezone struct {
	ID        byte
	Name      string
	Cities    string
	UTCOffset int
}
