package paging

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/restful-api-boilerplate/api/utils/checkdebug"
)

// GetPaginator paginator getter
func GetPaginator(db *gorm.DB, page int, limit int, data interface{}) pagination.Paginator {
	return *pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id asc"},
		ShowSQL: checkdebug.CheckDebug(),
	}, data)
}
