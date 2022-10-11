package types

import (
	"gorm.io/gorm"
	"time"
	"go_rest_monolith/utils/middleware/pagination"
)

type User struct {
	gorm.Model
	User_group_id 	int 		`gorm:"default:2"`
	Name     		string 		`gorm:"default:null"`
	Email    		string 		`gorm:"not null"`
	Password 		string 		`gorm:"not null"`
	Created_at 		time.Time 	`gorm:"default:null"`
	Updated_at 		time.Time 	`gorm:"default:null"`
	Deleted_at 		time.Time 	`gorm:"default:null"`
	Is_deleted 		int 		`gorm:"default:0"`
}

type User_groups struct {
	gorm.Model
	Name     		string		`gorm:"default:null"`
	Description    	string		`gorm:"default:null"`
	Created_at 		time.Time 	`gorm:"default:null"`
	Updated_at 		time.Time 	`gorm:"default:null"`
	Deleted_at 		time.Time 	`gorm:"default:null"`
	Is_deleted 		int  		`gorm:"default:0"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"-"`
	Password string `json:"-"`
}

type UserResp struct {
	Status      string      `json:"status"`
    StatusCode  int         `json:"statusCode"`
    Data        []*pagination.Paginator `json:"data"`
}