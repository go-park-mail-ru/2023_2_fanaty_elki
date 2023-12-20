package dto

import (
	"database/sql"
	"log"
	"server/internal/domain/entity"
	"strconv"
	"time"
)

type DBGetPromo struct {
	ID           uint
	Code         string
	PromoType    int
	Sale         sql.NullString
	RestaurantId sql.NullString
	ActiveFrom   time.Time
	ActiveTo     time.Time
}

type RespPromo struct {
	Type     int  `json:"Type"`
	Discount uint `json:"Discount"`
}

func ToRespPromo(promo *entity.Promo) *RespPromo {
	if promo == nil {
		return nil
	}
	return &RespPromo{
		Type:     promo.PromoType,
		Discount: promo.Sale,
	}
}

func ToEntityGetPromo(dbPromo *DBGetPromo) *entity.Promo {
	if dbPromo == nil {
		return nil
	}
	return &entity.Promo{
		ID:           dbPromo.ID,
		Code:         dbPromo.Code,
		PromoType:    dbPromo.PromoType,
		Sale:         transformSqlStringToUint(dbPromo.Sale),
		RestaurantID: transformSqlStringToUint(dbPromo.RestaurantId),
		ActiveFrom:   dbPromo.ActiveFrom,
		ActiveTo:     dbPromo.ActiveTo,
	}
}

func transformSqlStringToUint(str sql.NullString) uint {
	if str.Valid {
		num, err := strconv.ParseUint(str.String, 10, 64)
		if err != nil {
			log.Fatal("Converting error")
		}
		return uint(num)
	}
	return 0
}
