package api

import (
	"fmt"
	"github.com/duvanherfi/stock-analysis/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type API struct {
	DB *gorm.DB
}

func NewAPI(db *gorm.DB) *API {
	return &API{DB: db}
}

func (api *API) GetStocks(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		fmt.Println("error: ", err)
		page = 1
	}
	var stockItems []models.StockItem
	pagination := Pagination{
		Page:  page,
		Limit: 10,
		Sort:  "time desc",
		Rows:  stockItems,
	}

	paginate(&stockItems, &pagination, api.DB).Find(&pagination.Rows)

	err = c.JSON(http.StatusOK, pagination)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return err
}

func (api *API) RecommendStocks(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	var stockItems []models.StockItem
	pagination := Pagination{
		Page:  page,
		Limit: 10,
		Sort:  "target_to desc",
		Rows:  stockItems,
	}

	paginate(&stockItems, &pagination, api.DB.Where("target_to > target_from")).Find(&pagination.Rows)

	err = c.JSON(http.StatusOK, pagination)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return err
}

// ... Implementar otros handlers para obtener detalles de acciones específicas ...
