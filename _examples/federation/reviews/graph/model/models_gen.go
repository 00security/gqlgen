// Code generated by github.com/00security/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/00security/gqlgen/_examples/federation/accounts/graph/model"
)

type Manufacturer struct {
	ID string `json:"id"`
}

func (Manufacturer) IsEntity() {}

type Product struct {
	ID           string        `json:"id"`
	Manufacturer *Manufacturer `json:"manufacturer"`
	Reviews      []*Review     `json:"reviews"`
}

func (Product) IsEntity() {}

type Review struct {
	Body    string      `json:"body"`
	Author  *model.User `json:"author"`
	Product *Product    `json:"product"`
}
