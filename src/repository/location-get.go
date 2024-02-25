package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	table "stepp-backend/src/constant/database"
	"stepp-backend/src/model/database"
)

func (r *repository) GetLocation(ctx context.Context, locationId int) (*database.LocationDb, error) {

	response := database.LocationDb{}
	err := r.gormDb.Table(table.LOCATION).Where("location_id = ?", locationId).First(&response).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Err: Post Not Found !")
			return nil, err
		}
		return nil, err
	}

	return &response, nil
}
