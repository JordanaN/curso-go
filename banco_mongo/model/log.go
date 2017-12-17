package model

type RegistroLog struct {
	DataVisita string `json:"dataVisita" bson:"dataVisita"`
	Quem       string `json:"quem,omitempty" bson:"quem,omitempty"`
}
