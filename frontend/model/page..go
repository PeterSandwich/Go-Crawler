package model

type SearchReqult struct{
	Hits int64
	Start int
	Query string
	PrevFrom int
	NextFrom int
	Items []interface{}
}
