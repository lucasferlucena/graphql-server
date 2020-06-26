package graph

import "github.com/jinzhu/gorm"

//go:generate go run github.com/99designs/gqlgen

//Resolver ...Resolver
type Resolver struct {
	DB *gorm.DB
}
