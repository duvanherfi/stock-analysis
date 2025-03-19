package database

import (
	"context"
	"fmt"
	"github.com/duvanherfi/stock-analysis/internal/models"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(dbURL string) error {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	err = db.AutoMigrate(&models.StockItem{})
	if err != nil {
		return err
	}

	return nil
}

func InsertStockItem(item models.StockItem) error {

	err := crdbgorm.ExecuteTx(context.Background(), DB, nil,
		func(tx *gorm.DB) error {
			err := tx.Where(item).FirstOrCreate(&item).Error

			return err
		},
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
