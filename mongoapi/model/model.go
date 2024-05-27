package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

type Color struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Part struct {
	Name     string `json:"name"`
	Material string `json:"material"`
}

type Product struct {
	ID                  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name                string             `json:"name"`
	Description         string             `json:"description"`
	Price               string             `json:"price"`
	Variants            []string           `json:"variants"`
	Colors              []Color            `json:"colors"`
	Image               string             `json:"image"`
	DetailedDescription string             `json:"detailedDescription"`
	Parts               []Part             `json:"parts"`
}
