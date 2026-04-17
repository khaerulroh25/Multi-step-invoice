package seed

import (
	"backend/config"
	"backend/models"
)

func SeedItems() {
	items := []models.Item{
		{Code: "BRG-001", Name: "Barang A", Price: 10000},
		{Code: "BRG-002", Name: "Barang B", Price: 20000},
	}

	for _, item := range items {
		config.DB.FirstOrCreate(&item, models.Item{Code: item.Code})
	}
}