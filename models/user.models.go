package models

import (
	"go_rest_monolith/types"
	"go_rest_monolith/config/database"
	"gorm.io/gorm"
	"go_rest_monolith/utils/middleware/pagination"
	"time"
)

func FindUsers(params types.Filter) types.Result {
	var usersResponse []types.UserResponse
	db := database.Gorm2

	paginator := pagination.Paging(&pagination.Param{
		DB:      db.Where("is_deleted = ?", 0),
		Table:	 "users",
		Select:  "id, name",
		Page:    params.Page,
		Limit:   params.Limit,
		OrderBy: []string{params.SortBy+" "+params.SortDir},
		ShowSQL: false,
	}, &usersResponse)

	statusCode := 200
	if len(usersResponse) < 1 {
		statusCode = 404
	}

	res := types.Result{
    	Status: "success",
        StatusCode: statusCode,
        Data: paginator,
    }

	return res
}

func FindUser(id int) types.Result {
	var usersResponse []types.UserResponse
	db := database.Gorm2

	paginator := pagination.Paging(&pagination.Param{
		DB:      db.Where("id = ?", id),
		Table: 	 "users",
		Select:  "id, name",
		Page:    1,
		Limit:   1,
		OrderBy: []string{"id asc"},
		ShowSQL: false,
	}, &usersResponse)

	statusCode := 200
	if len(usersResponse) < 1 {
		statusCode = 404
	}

	res := types.Result{
    	Status: "success",
        StatusCode: statusCode,
        Data: paginator,
    }

	return res
}

func CreateUser(user *types.User) *gorm.DB {
	return database.Gorm2.Create(user)
}

func UpdateUser(id int, user interface{}) *gorm.DB {
	return database.Gorm2.Model(&user).Where("id = ?", id).Updates(user)
}

func DeleteUser(id int) *gorm.DB {
	user := new(types.User)
	datetime := time.Now()

	data := types.User{
		Is_deleted: 1,
		Deleted_at: datetime,
	}
	return database.Gorm2.Model(&user).Where("id = ?", id).Updates(data)
}

func GetUserById(id int) *gorm.DB {
	return database.Gorm2.Model(&types.User{}).Where("id = ?", id).First(&types.User{})
}