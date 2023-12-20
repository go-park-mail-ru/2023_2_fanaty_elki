package repository

import (
	"database/sql"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

//RestaurantRepo struct
type RestaurantRepo struct {
	DB *sql.DB
}

//NewRestaurantRepo creates new object of Restaruant repo
func NewRestaurantRepo(db *sql.DB) *RestaurantRepo {
	return &RestaurantRepo{
		DB: db,
	}
}

//GetRestaurants gets restaurants
func (repo *RestaurantRepo) GetRestaurants() ([]*entity.Restaurant, error) {
	rows, err := repo.DB.Query(`SELECT id, name, rating, comments_count, icon 
								FROM restaurant
								ORDER BY rating DESC`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} 
		return nil, err
	}
	defer rows.Close()
	var Restaurants = []*entity.Restaurant{}
	for rows.Next() {
		restaurant := &entity.Restaurant{}
		err = rows.Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Rating,
			&restaurant.CommentsCount,
			&restaurant.Icon,
		)
		if err != nil {
			return nil, err
		}
		Restaurants = append(Restaurants, restaurant)
	}
	return Restaurants, nil
}

//GetRestaurantByID gets restaurant by id
func (repo *RestaurantRepo) GetRestaurantByID(id uint) (*entity.Restaurant, error) {
	restaurant := &entity.Restaurant{}
	row := repo.DB.QueryRow(`SELECT id, name, rating, comments_count, icon 
							 FROM restaurant 
							 WHERE id = $1`, id)
	err := row.Scan(
		&restaurant.ID,
		&restaurant.Name,
		&restaurant.Rating,
		&restaurant.CommentsCount,
		&restaurant.Icon,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrNotFound
		}
		return nil, entity.ErrInternalServerError
	}
	return restaurant, nil
}

//GetRestaurantByName gets restaurant by name
func (repo *RestaurantRepo) GetRestaurantByName(name string) (*entity.Restaurant, error) {
	restaurant := &entity.Restaurant{}
	row := repo.DB.QueryRow(`SELECT id, name, rating, comments_count, icon 
							 FROM restaurant 
							 WHERE name = $1`, name)
	err := row.Scan(
		&restaurant.ID,
		&restaurant.Name,
		&restaurant.Rating,
		&restaurant.CommentsCount,
		&restaurant.Icon,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrNotFound
		}
		return nil, entity.ErrInternalServerError
	}
	return restaurant, nil
}

//GetMenuTypesByRestaurantID gets menu from restaurant
func (repo *RestaurantRepo) GetMenuTypesByRestaurantID(id uint) ([]*entity.MenuType, error) {
	rows, err := repo.DB.Query(`SELECT id, name, restaurant_id FROM menu_type WHERE restaurant_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var MenuTypes = []*entity.MenuType{}
	var count = 0
	for rows.Next() {
		count++
		menuType := &entity.MenuType{}
		err = rows.Scan(
			&menuType.ID,
			&menuType.Name,
			&menuType.RestaurantID,
		)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		MenuTypes = append(MenuTypes, menuType)
	}
	if count == 0 {
		return nil, entity.ErrNotFound
	}
	return MenuTypes, nil
}

//GetCategoriesByRestaurantID gets categories from restaurant
func (repo *RestaurantRepo) GetCategoriesByRestaurantID(id uint) ([]*entity.Category, error) {
	rows, err := repo.DB.Query(`SELECT category.id, category.name 
								FROM restaurant_category rc 
								INNER JOIN category ON rc.category_id=category.id 
								WHERE restaurant_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Categories = []*entity.Category{}
	var count = 0
	for rows.Next() {
		count++
		category := &entity.Category{}
		err = rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		Categories = append(Categories, category)
	}
	if count == 0 {
		return nil, entity.ErrNotFound
	}
	return Categories, nil
}

//GetRestaurantsByCategory gets restaurtants by categories
func (repo *RestaurantRepo) GetRestaurantsByCategory(name string) ([]*entity.Restaurant, error) {
	rows, err := repo.DB.Query(`SELECT restaurant.id, restaurant.name, rating, comments_count, icon 
								FROM restaurant_category rc 
								INNER JOIN restaurant ON rc.restaurant_id=restaurant.id
								INNER JOIN category ON rc.category_id=category.id 
								WHERE category.name = $1
								ORDER BY rating DESC`, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} 
		return nil, err
	}
	defer rows.Close()
	var Restaurants = []*entity.Restaurant{}
	var count = 0
	for rows.Next() {
		count++
		restaurant := &entity.Restaurant{}
		err = rows.Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Rating,
			&restaurant.CommentsCount,
			&restaurant.Icon,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			} 
			return nil, err
		}
		Restaurants = append(Restaurants, restaurant)
	}
	if count == 0 {
		return nil, entity.ErrNotFound
	}
	return Restaurants, nil
}

//GetCategories gets categories from db
func (repo *RestaurantRepo) GetCategories() ([]*entity.Category, error) {
	rows, err := repo.DB.Query(`SELECT id, name FROM category`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} 
		return nil, err
	}
	defer rows.Close()
	var Categories = []*entity.Category{}
	var count = 0
	for rows.Next() {
		count++
		category := &entity.Category{}
		err = rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}
		Categories = append(Categories, category)
	}
	if count == 0 {
		return nil, entity.ErrNotFound
	}
	return Categories, nil
}

//SearchRestaurants selects restaurants from db
func (repo *RestaurantRepo) SearchRestaurants(word string) ([]*entity.Restaurant, error) {
	rows, err := repo.DB.Query(`SELECT id, name, rating, comments_count, icon
							    FROM restaurant 
								WHERE LOWER(name) 
								LIKE LOWER('%' || $1 || '%')
								ORDER BY rating DESC`, word)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} 
		return nil, err
	}
	defer rows.Close()
	var Restaurants = []*entity.Restaurant{}
	for rows.Next() {
		restaurant := &entity.Restaurant{}
		err = rows.Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Rating,
			&restaurant.CommentsCount,
			&restaurant.Icon,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			} 
			return nil, err
		}
		Restaurants = append(Restaurants, restaurant)
	}
	return Restaurants, nil
}

//SearchCategories select restaurants with necessary category
func (repo *RestaurantRepo) SearchCategories(word string) ([]*entity.Restaurant, error) {
	rows, err := repo.DB.Query(`SELECT restaurant.id, restaurant.name, rating, comments_count, icon
								FROM restaurant_category rc 
								INNER JOIN restaurant on rc.restaurant_id=restaurant.id
								INNER JOIN category on rc.category_id=category.id 
								WHERE LOWER(category.name) LIKE LOWER('%' || $1 || '%')
								ORDER BY rating`, word)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} 
		return nil, err
	}
	defer rows.Close()
	var Restaurants = []*entity.Restaurant{}
	for rows.Next() {
		restaurant := &entity.Restaurant{}
		err = rows.Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Rating,
			&restaurant.CommentsCount,
			&restaurant.Icon,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			} 
			return nil, err
		}
		Restaurants = append(Restaurants, restaurant)
	}
	return Restaurants, nil
}

//UpdateComments updates comments in db
func (repo *RestaurantRepo) UpdateComments(comment *dto.ReqCreateComment) error {
	getCountCommments := `SELECT rating, comments_count
						  FROM restaurant
						  where id = $1`
	var count uint
	var rating float32
	err := repo.DB.QueryRow(getCountCommments, comment.RestaurantID).Scan(&rating, &count)

	if err != nil {
		return entity.ErrInternalServerError
	}

	count++
	rating = (rating*float32(count-1) + float32(comment.Rating)) / float32(count)

	updateComment := `UPDATE restaurant
					  SET rating = $1, comments_count = $2
					  WHERE id = $3`

	_, err = repo.DB.Exec(updateComment, rating, count, comment.RestaurantID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil

}
