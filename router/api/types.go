package api

type anime struct {
	Id               string
	Title            string
	AlternativeTitle string
	Aired            int64
	Duration         int64
	CurrentStatus    int
	Season           string
	Type             string
}
