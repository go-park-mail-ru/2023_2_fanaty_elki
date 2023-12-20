package dto

import (
	"database/sql"
	"log"
	"server/internal/domain/entity"
	"strconv"
	"time"
)

//DBGetPromo dto
type DBGetPromo struct {
	ID           uint
	Code         string
	PromoType    int
	Sale         sql.NullString
	RestaurantID sql.NullString
	ActiveFrom   time.Time
	ActiveTo     time.Time
}

//RespPromo dto
type RespPromo struct {
	Type     int    `json:"Type"`
	Discount uint   `json:"Discount"`
	Promo    string `json:"Promo"`
}

//ToRespPromo transforms Promo to RespPromo
func ToRespPromo(promo *entity.Promo) *RespPromo {
	if promo == nil {
		return nil
	}
	return &RespPromo{
		Type:     promo.PromoType,
		Discount: promo.Sale,
		Promo:    promo.Code,
	}
}

//ToEntityGetPromo transforms DBGetPromo to Promo
func ToEntityGetPromo(dbPromo *DBGetPromo) *entity.Promo {
	if dbPromo == nil {
		return nil
	}
	return &entity.Promo{
		ID:           dbPromo.ID,
		Code:         dbPromo.Code,
		PromoType:    dbPromo.PromoType,
		Sale:         transformSQLStringToUint(dbPromo.Sale),
		RestaurantID: transformSQLStringToUint(dbPromo.RestaurantID),
		ActiveFrom:   dbPromo.ActiveFrom,
		ActiveTo:     dbPromo.ActiveTo,
	}
}

func transformSQLStringToUint(str sql.NullString) uint {
	if str.Valid {
		num, err := strconv.ParseUint(str.String, 10, 64)
		if err != nil {
			log.Fatal("Converting error")
		}
		return uint(num)
	}
	return 0
}
