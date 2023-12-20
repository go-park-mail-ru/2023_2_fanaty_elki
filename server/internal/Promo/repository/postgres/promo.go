package repository

import (
	"database/sql"
	"server/internal/Promo/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

//PromoRepo struct
type PromoRepo struct {
	DB *sql.DB
}

//NewPromoRepo creates new promocode repository
func NewPromoRepo(db *sql.DB) repository.PromoRepositoryI {
	return &PromoRepo{
		DB: db,
	}
}

//GetPromo gets promocode from db
func (repo *PromoRepo) GetPromo(code string) (*entity.Promo, error) {
	promo := &dto.DBGetPromo{}
	row := repo.DB.QueryRow("SELECT id, code, promo_type, sale, restaurant_id, active_from, active_to FROM promo WHERE code = $1", code)
	err := row.Scan(&promo.ID, &promo.Code, &promo.PromoType, &promo.Sale, &promo.RestaurantID, &promo.ActiveFrom, &promo.ActiveTo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return dto.ToEntityGetPromo(promo), nil
}

//GetPromoByID gets promocode by id
func (repo *PromoRepo) GetPromoByID(PromoID uint) (*entity.Promo, error) {
	promo := &dto.DBGetPromo{}
	row := repo.DB.QueryRow("SELECT id, code, promo_type, sale, restaurant_id, active_from, active_to FROM promo WHERE id = $1", PromoID)
	err := row.Scan(&promo.ID, &promo.Code, &promo.PromoType, &promo.Sale, &promo.RestaurantID, &promo.ActiveFrom, &promo.ActiveTo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return dto.ToEntityGetPromo(promo), nil
}

//UsePromo uses promocode in db
func (repo *PromoRepo) UsePromo(UserID uint, PromoID uint) error {
	insertPromo := `INSERT INTO user_promo (user_id, promo_id) VALUES ($1, $2)`
	_, err := repo.DB.Exec(insertPromo, UserID, PromoID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}

//SetPromoToCart sets promocode to cart
func (repo *PromoRepo) SetPromoToCart(cartID uint, PromoID uint) error {
	insertPromo := `INSERT INTO cart_promo (cart_id, promo_id) VALUES ($1, $2)`
	_, err := repo.DB.Exec(insertPromo, cartID, PromoID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}

//CheckPromo checks promocode in db
func (repo *PromoRepo) CheckPromo(UserID uint, PromoID uint) (bool, error) {
	userPromo := &entity.UserPromo{}
	row := repo.DB.QueryRow("SELECT  user_id, promo_id FROM user_promo WHERE user_id = $1 and promo_id = $2", UserID, PromoID)
	err := row.Scan(
		&userPromo.UserID,
		&userPromo.PromoID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

//DeletePromo deletes promocode from  db
func (repo *PromoRepo) DeletePromo(UserID uint, PromoID uint) error {
	deletePromo := `DELETE FROM user_promo WHERE user_id = $1 AND promo_id = $2`
	_, err := repo.DB.Exec(deletePromo, UserID, PromoID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}

//DeletePromoFromCart deletes promocode from cart db
func (repo *PromoRepo) DeletePromoFromCart(cartID uint, PromoID uint) error {
	deletePromo := `DELETE FROM cart_promo WHERE cart_id = $1 AND promo_id = $2`
	_, err := repo.DB.Exec(deletePromo, cartID, PromoID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}
