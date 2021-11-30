package models

type Index struct {
	ID    string `bson:"_id"`
	Key   string
	Value int
}
