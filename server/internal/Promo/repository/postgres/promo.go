package repository

import (
	"database/sql"
	"server/internal/Promo/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

type PromoRepo struct {
	DB *sql.DB
}

func NewPromoRepo(db *sql.DB) repository.PromoRepositoryI {
	return &PromoRepo{
		DB: db,
	}
}

func (repo *PromoRepo) GetPromo(code string) (*entity.Promo, error) {
	promo := &dto.DBGetPromo{}
	row := repo.DB.QueryRow("SELECT id, code, promo_type, sale, restaurant_id, active_from, active_to FROM promo WHERE code = $1", code)
	err := row.Scan(&promo.ID, &promo.Code, &promo.PromoType, &promo.Sale, &promo.RestaurantId, &promo.ActiveFrom, &promo.ActiveTo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return dto.ToEntityGetPromo(promo), nil
}

func (repo *PromoRepo) GetPromoById(promoId uint) (*entity.Promo, error) {
	promo := &dto.DBGetPromo{}
	row := repo.DB.QueryRow("SELECT id, code, promo_type, sale, restaurant_id, active_from, active_to FROM promo WHERE id = $1", promoId)
	err := row.Scan(&promo.ID, &promo.Code, &promo.PromoType, &promo.Sale, &promo.RestaurantId, &promo.ActiveFrom, &promo.ActiveTo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return dto.ToEntityGetPromo(promo), nil
}

func (repo *PromoRepo) UsePromo(userId uint, promoId uint) error {
	insertPromo := `INSERT INTO user_promo (user_id, promo_id) VALUES ($1, $2)`
	_, err := repo.DB.Exec(insertPromo, userId, promoId)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}

func (repo *PromoRepo) SetPromoToCart(cartId uint, promoId uint) error {
	insertPromo := `INSERT INTO cart_promo (cart_id, promo_id) VALUES ($1, $2)`
	_, err := repo.DB.Exec(insertPromo, cartId, promoId)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}

func (repo *PromoRepo) CheckPromo(userId uint, promoId uint) (bool, error) {
	userPromo := &entity.UserPromo{}
	row := repo.DB.QueryRow("SELECT  user_id, promo_id FROM user_promo WHERE user_id = $1 and promo_id = $2", userId, promoId)
	err := row.Scan(
		&userPromo.UserId,
		&userPromo.PromoId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (repo *PromoRepo) DeletePromo(userId uint, promoId uint) error {
	deletePromo := `DELETE FROM user_promo WHERE user_id = $1 AND promo_id = $2`
	_, err := repo.DB.Exec(deletePromo, userId, promoId)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}

func (repo *PromoRepo) DeletePromoFromCart(cartId uint, promoId uint) error {
	deletePromo := `DELETE FROM cart_promo WHERE cart_id = $1 AND promo_id = $2`
	_, err := repo.DB.Exec(deletePromo, cartId, promoId)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}
