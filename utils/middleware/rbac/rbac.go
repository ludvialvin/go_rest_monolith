package rbac

import (
	"go_rest_monolith/config/database"
	"github.com/harranali/authority"
)

var Auth *authority.Authority

func InitRBAC() {
	db := database.Gorm2

	auth := authority.New(authority.Options{
	    TablesPrefix: "authority_",
	    DB:           db,
	})

	Auth = auth
}