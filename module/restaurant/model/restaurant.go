package restaurantmodel

import (
	"errors"
	"food-delivery/common"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel
	// Id     int    `json:"id" gorm:"column:id;"`
	Name  string         `json:"name" gorm:"column:name;"`
	Addr  string         `json:"addr" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

type RestaurantCreate struct {
	common.SQLModel
	Name string        `json:"name" gorm:"column:name;"`
	Addr string        `json:"addr" gorm:"column:addr;"`
	Logo *common.Image `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string { return "restaurants" }
func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

func (c *RestaurantCreate) Validate() error {
	c.Name = strings.TrimSpace(c.Name)
	if c.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

type RestaurantUpdate struct {
	Name *string       `json:"name" gorm:"column:name;"`
	Addr *string       `json:"addr" gorm:"column:addr;"`
	Logo *common.Image `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName(string) string { return Restaurant{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("name cannot be empty")
)
