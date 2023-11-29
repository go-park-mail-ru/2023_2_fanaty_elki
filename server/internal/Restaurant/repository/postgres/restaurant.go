package repository

import (
	"database/sql"
	"server/internal/domain/entity"
	"server/internal/domain/dto"
)

type restaurantRepo struct {
	DB *sql.DB
}

func NewRestaurantRepo(db *sql.DB) *restaurantRepo {
	return &restaurantRepo{
		DB: db,
	}
}

func (repo *restaurantRepo) GetRestaurants() ([]*entity.Restaurant, error) {
	rows, err := repo.DB.Query(`SELECT id, name, rating, comments_count, icon 
								FROM restaurant
								ORDER BY rating DESC`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
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

func (repo *restaurantRepo) GetRestaurantById(id uint) (*entity.Restaurant, error) {
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

func (repo *restaurantRepo) GetMenuTypesByRestaurantId(id uint) ([]*entity.MenuType, error) {
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

func (repo *restaurantRepo) GetCategoriesByRestaurantId(id uint) ([]*entity.Category, error) {
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

func (repo *restaurantRepo) GetRestaurantsByCategory(name string) ([]*entity.Restaurant, error) {
	rows, err := repo.DB.Query(`SELECT restaurant.id, restaurant.name, rating, comments_count, icon 
								FROM restaurant_category rc 
								INNER JOIN restaurant ON rc.restaurant_id=restaurant.id
								INNER JOIN category ON rc.category_id=category.id 
								WHERE category.name = $1
								ORDER BY rating DESC`, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
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
			} else {
				return nil, err
			}
		}
		Restaurants = append(Restaurants, restaurant)
	}
	if count == 0 {
		return nil, entity.ErrNotFound
	}
	return Restaurants, nil
}

func (repo *restaurantRepo) GetCategories() ([]*entity.Category, error) {
	rows, err := repo.DB.Query(`SELECT id, name FROM category`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
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

func (repo *restaurantRepo) SearchRestaurants(word string) ([]*entity.Restaurant, error) {
	rows, err := repo.DB.Query(`SELECT id, name, rating, comments_count, icon
							    FROM restaurant 
								WHERE LOWER(name) 
								LIKE LOWER('%' || $1 || '%')
								ORDER BY rating DESC`, word)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
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
			} else {
				return nil, err
			}
		}
		Restaurants = append(Restaurants, restaurant)
	}
	return Restaurants, nil
}

func (repo *restaurantRepo) SearchCategories(word string) ([]*entity.Restaurant, error) {
	rows, err := repo.DB.Query(`SELECT restaurant.id, restaurant.name, rating, comments_count, icon
								FROM restaurant_category rc 
								INNER JOIN restaurant on rc.restaurant_id=restaurant.id
								INNER JOIN category on rc.category_id=category.id 
								WHERE LOWER(category.name) LIKE LOWER('%' || $1 || '%')
								ORDER BY rating`, word)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
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
			} else {
				return nil, err
			}
		}
		Restaurants = append(Restaurants, restaurant)
	}
	return Restaurants, nil
}


func (repo *restaurantRepo) UpdateComments(comment *dto.ReqCreateComment) error {
	getCountCommments := `SELECT rating, comments_count
						  FROM restaurant
						  where id = $1`
	var count uint
	var rating float32
	err := repo.DB.QueryRow(getCountCommments, comment.RestaurantId).Scan(&rating, &count)

	if err != nil {
		return entity.ErrInternalServerError
	}
	
	count += 1
	rating = (rating * float32(count - 1) + float32(comment.Rating)) / float32(count)
	
	updateComment := `UPDATE restaurant
					  SET rating = $1, comments_count = $2
					  WHERE id = $3`

	_, err = repo.DB.Exec(updateComment, rating, count, comment.RestaurantId)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil

}