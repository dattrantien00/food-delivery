package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"-" gorm:"column:id"`
	Url       string `json:"url" gorm:"column:url"`
	Width     int    `json:"width" gorm:"column:width"`
	Height    int    `json:"height" gorm:"column:height"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string {
	return "images"
}

func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Fail to unmashal JSONB value", value))
	}

	var img Image

	err := json.Unmarshal(bytes, &img)
	if err != nil {
		return err
	}
	*i = img

	return nil
}

func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}
	return json.Marshal(i)
}

type Images []Image

func (i *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Fail to unmashal JSONB value", value))
	}

	var img Images

	err := json.Unmarshal(bytes, &img)
	if err != nil {
		return err
	}
	*i = img

	return nil
}

func (i *Images) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}
	return json.Marshal(i)
}
