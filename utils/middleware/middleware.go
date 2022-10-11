package middleware 

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"go_rest_monolith/types"
)

func ParseQuery(c *fiber.Ctx) types.Filter {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	sortby := c.Query("sortby")
	if sortby == "" {
		sortby = "id"
	}
	sortdir := c.Query("sortdir")
	if sortdir == "" {
		sortdir = "asc"
	}

	res := types.Filter{
    	Page: page,
        Limit: limit,
        SortBy: sortby,
        SortDir: sortdir,
    }

	return res
}