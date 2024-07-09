package restaurantmodel

import "testing"

func TestRestaurantCreate_Validate(t *testing.T) {
	dataTest := RestaurantCreate{
		Name: "",
	}
	err := dataTest.Validate()

	if err != ErrNameIsEmpty{
		t.Error("validate restaurant input name:",dataTest.Name,".Expect: ErrEmpty", "Output: ",err )
		return
	}
}
