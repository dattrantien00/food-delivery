package restaurantlikemodel

type Filter struct {
	RestaurantId int `json:"-" gorm:"column:restaurant_id"`
	UserId       int `json:"-" gorm:"column:user_id"`
}
