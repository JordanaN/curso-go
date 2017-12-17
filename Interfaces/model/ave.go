package model

type Galinha interface {
	Cacareja() string
}

type Pata interface {
	Quack() string
}

type Ave struct {
	Nome string
}

func (a Ave) Cacareja() string {
	return "Cócóricó"
}

func (a Ave) Quack() string {
	return "Quack, quack, quack"
}
