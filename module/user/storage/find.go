package userstorage

import (
	"context"
	"food-delivery/common"
	usermodel "food-delivery/module/user/model"

	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

func (store *sqlStore) FindUserByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	relations ...string,
) (*usermodel.User, error) {
	_, span := trace.StartSpan(ctx, "store_user.find_user")
	defer span.End()
	var user usermodel.User
	db := store.db.Table(usermodel.User{}.TableName())

	if err := db.Where(conditions).First(&user).Error; err != nil {

		if err != gorm.ErrRecordNotFound {

			return nil, common.ErrDb(err)
		}

		return nil, err
	}
	
	return &user, nil
}
