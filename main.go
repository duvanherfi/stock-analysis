package main

import (
	"encoding/json"
	"fmt"
	"github.com/duvanherfi/stock-analysis/internal/api"
	"github.com/duvanherfi/stock-analysis/internal/config"
	"github.com/duvanherfi/stock-analysis/internal/database"
	"github.com/duvanherfi/stock-analysis/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	cfg := config.LoadConfig()
	err := database.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	api := api.NewAPI(database.DB)

	fetchAndStoreData(cfg)
	// ... Iniciar el servidor API ...
	e := echo.New()
	e.GET("/api/stocks", api.GetStocks)
	e.GET("/api/recommends-stocks", api.RecommendStocks)
	e.Static("/", "ui/dist")
	e.File("/", "ui/dist/index.html")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Logger.Fatal(e.Start(":3000"))

}

func fetchAndStoreData(cfg config.Config) {
	var cant_items int64
	database.DB.Model(&models.StockItem{}).Count(&cant_items)
	if cant_items >= 1043 {
		return
	}
	url := cfg.APIURL
	client := &http.Client{}

	for url != "" {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Printf("Error creating request: %v", err)
			return
		}
		req.Header.Add("Authorization", "Bearer "+cfg.APIKey)

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error making request: %v", err)
			fmt.Println("req: ", req)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			return
		}

		var response models.StockAPIResponse

		if err := json.Unmarshal(body, &response); err != nil {
			log.Printf("Error unmarshalling response: %v", err)
			return
		}

		if len(response.Items) == 0 {
			break
		}

		for _, item := range response.Items {

			from, _ := strconv.ParseFloat(strings.Replace(item.TargetFrom, "$", "", 1), 64)
			to, _ := strconv.ParseFloat(strings.Replace(item.TargetTo, "$", "", 1), 64)

			dbItem := models.StockItem{
				Ticker:     item.Ticker,
				TargetFrom: from,
				TargetTo:   to,
				Company:    item.Company,
				Action:     item.Action,
				Brokerage:  item.Brokerage,
				RatingFrom: item.RatingFrom,
				RatingTo:   item.RatingTo,
				//Time:       item.Time,
			}
			err := database.InsertStockItem(dbItem)
			if err != nil {
				fmt.Println("Error inserting item: ", err)
				return
			}
		}

		url = url + "?next_page=" + response.NextPage
	}
}
