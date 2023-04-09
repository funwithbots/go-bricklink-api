package entity

type Entity interface {
	PrimaryKey() int
	Label() Label
}
