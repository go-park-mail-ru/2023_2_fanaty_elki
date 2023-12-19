package repository

import (
	"database/sql"
	"fmt"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

//AddressRepo struct
type AddressRepo struct {
	DB *sql.DB
}

//NewAddressRepo creates address repo
func NewAddressRepo(db *sql.DB) *AddressRepo {
	return &AddressRepo{
		DB: db,
	}
}

//CreateAddress creates address in db
func (repo *AddressRepo) CreateAddress(address *dto.DBReqCreateUserAddress) error {

	insertAddress := `INSERT INTO address (city, street, house_number, flat_number)
				      VALUES ($1, $2, $3, $4)
					  RETURNING ID`
	var addressID uint
	err := repo.DB.QueryRow(insertAddress, address.City, address.Street, address.House, address.Flat).Scan(&addressID)
	if err != nil {
		return entity.ErrInternalServerError
	}

	updateAddress := `UPDATE users_address
					  SET current = false
					  WHERE current = true AND user_id = $1`
	_, err = repo.DB.Exec(updateAddress, address.UserID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	insertUserAddress := `INSERT INTO users_address (user_id, address_id, current)
				      	  VALUES ($1, $2, $3)`
	_, err = repo.DB.Exec(insertUserAddress, address.UserID, addressID, true)
	if err != nil {
		return entity.ErrInternalServerError
	}

	return nil
}

//DeleteAddress deletes address from db
func (repo *AddressRepo) DeleteAddress(address *dto.DBReqDeleteUserAddress) error {
	deleteAddress := `DELETE FROM users_address
					WHERE user_id = $1 AND address_id = $2`
	result, err := repo.DB.Exec(deleteAddress, address.UserID, address.AddressID)
	if err != nil {
		return entity.ErrInternalServerError
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return entity.ErrInternalServerError
	}
	if rowsAffected == 0 {
		return entity.ErrNotFound
	}
	return nil
}

//GetAddresses get address from db
func (repo *AddressRepo) GetAddresses(UserID uint) (*dto.RespGetAddresses, error) {
	getAddresses := `SELECT a.id, a.city, a.street, a.house_number, a.flat_number, ua.current
			 	 FROM address a
				 JOIN users_address ua on a.id = ua.address_id
				 JOIN users u on u.id = ua.user_id
				 WHERE ua.user_id = $1
				 ORDER BY a.created_at DESC`

	rows, err := repo.DB.Query(getAddresses, UserID)

	if err != nil {
		fmt.Println("getRepo", err)
		return nil, entity.ErrInternalServerError
	}
	defer rows.Close()
	respAddresses := &dto.RespGetAddresses{}
	var addresses = []*dto.RespGetAddress{}
	for rows.Next() {
		current := false
		address := &dto.RespGetAddress{}
		err = rows.Scan(
			&address.ID,
			&address.City,
			&address.Street,
			&address.House,
			&address.Flat,
			&current,
		)
		fmt.Println("Address", address)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		if current == true {
			respAddresses.CurrentAddressesID = address.ID
		}
		addresses = append(addresses, address)
	}
	respAddresses.Addresses = addresses
	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}

	return respAddresses, nil
}

//SetAddress sets address in db
func (repo *AddressRepo) SetAddress(address *dto.DBReqUpdateUserAddress) error {
	updateOldAddress := `UPDATE users_address
					  SET current = false
					  WHERE current = true AND user_id = $1`
	_, err := repo.DB.Exec(updateOldAddress, address.UserID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	updateAddress := `UPDATE users_address
					  SET current = true
					  WHERE user_id = $1 AND address_id = $2`
	_, err = repo.DB.Exec(updateAddress, address.UserID, address.AddressID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}
