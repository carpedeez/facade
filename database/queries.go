package database

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type Querier interface {
	CreateDisplay(userID uuid.UUID, title, description string) (Display, error)
	GetDisplay(displayID uuid.UUID) (Display, error)
	UpdateDisplay(displayID uuid.UUID, title, description *string) (Display, error)
	DeleteDisplay(displayID uuid.UUID) (bool, error)

	CreateItem(userID, displayID uuid.UUID, externalLink string) (Item, error)
	GetItem(itemID, displayID uuid.UUID) (Item, error)
	GetItems(displayID uuid.UUID) ([]Item, error)
	UpdateItem(itemID, displayID uuid.UUID, externalLink, socialPostLink, photoURL *string) (Item, error)
	DeleteItem(itemID, displayID uuid.UUID) (bool, error)

	IsItemOwner(userID, itemID uuid.UUID) (bool, error)
	IsDisplayOwner(userID, displayID uuid.UUID) (bool, error)

	Close() error
}

type postgresQuerier struct {
	DB *sql.DB
}

func (q postgresQuerier) CreateDisplay(userID uuid.UUID, title, description string) (Display, error) {
	d := Display{}
	//need to check sql injections
	sql, params, err := goqu.Insert("displays").Cols("user_id", "title", "descr", "photo_url").Vals(goqu.Vals{userID, title, description, ""}).Returning("*").ToSQL()
	if err != nil {
		return d, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}

	err = q.DB.QueryRow(sql, params...).Scan(&d.ID, &d.UserID, &d.Title, &d.Description, &d.PhotoURL)
	if err != nil {
		return d, fmt.Errorf("failed to execute database query: %w", err)
	}

	return d, nil
}

func (q postgresQuerier) GetDisplay(displayID uuid.UUID) (Display, error) {
	d := Display{}

	sql, params, err := goqu.Select("*").From("displays").Where(goqu.Ex{"id": displayID}).ToSQL()
	if err != nil {
		return d, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}

	err = q.DB.QueryRow(sql, params...).Scan(&d.ID, &d.UserID, &d.Title, &d.Description, &d.PhotoURL)
	if err != nil {
		return d, fmt.Errorf("failed to execute database query: %w", err)
	}

	return d, nil
}

func (q postgresQuerier) UpdateDisplay(displayID uuid.UUID, title, description *string) (Display, error) {
	d := Display{}

	u := goqu.Update("displays")
	r := goqu.Record{}
	if title != nil {
		r["title"] = *title
	}
	if description != nil {
		r["descr"] = *description
	}
	sql, params, err := u.Set(r).Where(goqu.Ex{"id": displayID}).Returning("*").ToSQL()
	if err != nil {
		return d, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}

	err = q.DB.QueryRow(sql, params...).Scan(&d.ID, &d.UserID, &d.Title, &d.Description, &d.PhotoURL)
	if err != nil {
		return d, fmt.Errorf("failed to execute database query: %w", err)
	}

	return d, nil
}

func (q postgresQuerier) DeleteDisplay(displayID uuid.UUID) (bool, error) {
	sql, params, err := goqu.Delete("displays").Where(goqu.Ex{"id": displayID}).ToSQL()
	if err != nil {
		return false, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}

	res, err := q.DB.Exec(sql, params...)
	if err != nil {
		return false, fmt.Errorf("failed to execute database query: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("no clue ngl")
	}

	if rows != 1 {
		return false, nil
	}

	return true, nil
}

func (q postgresQuerier) CreateItem(userID, displayID uuid.UUID, externalLink string) (Item, error) {
	i := Item{}

	sql, params, err := goqu.Insert("items").Cols("user_id", "display_id", "external_link", "social_post_link", "photo_url").Vals(goqu.Vals{userID, displayID, externalLink, "", ""}).Returning("*").ToSQL()
	if err != nil {
		return i, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}

	err = q.DB.QueryRow(sql, params...).Scan(&i.ID, &i.UserID, &i.DisplayID, &i.ExternalLink, &i.SocialPostLink, &i.PhotoURL)
	if err != nil {
		return i, fmt.Errorf("failed to execute database query: %w", err)
	}

	return i, nil
}

func (q postgresQuerier) GetItem(itemID, displayID uuid.UUID) (Item, error) {
	i := Item{}

	sql, params, err := goqu.Select("*").From("items").Where(goqu.Ex{"id": itemID, "display_id": displayID}).ToSQL()
	if err != nil {
		return i, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}

	err = q.DB.QueryRow(sql, params...).Scan(&i.ID, &i.UserID, &i.DisplayID, &i.ExternalLink, &i.SocialPostLink, &i.PhotoURL)
	if err != nil {
		return i, fmt.Errorf("failed to execute database query: %w", err)
	}

	return i, nil
}

func (q postgresQuerier) GetItems(displayID uuid.UUID) ([]Item, error) {
	is := []Item{}

	sql, params, err := goqu.Select("*").From("items").Where(goqu.Ex{"display_id": displayID}).ToSQL()
	if err != nil {
		return is, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}

	r, err := q.DB.Query(sql, params...)
	if err != nil {
		return is, fmt.Errorf("failed to execute database query: %w", err)
	}

	for r.Next() {
		var i Item
		_ = r.Scan(&i.ID, &i.UserID, &i.DisplayID, &i.ExternalLink, &i.SocialPostLink, &i.PhotoURL)
		is = append(is, i)
	}
	err = r.Err()
	if err != nil {
		return is, fmt.Errorf("failed to scan results: %w", err)
	}

	err = r.Close()
	if err != nil {
		return is, fmt.Errorf("failed to close rows: %w", err)
	}

	return is, nil
}

func (q postgresQuerier) UpdateItem(itemID, displayID uuid.UUID, externalLink, socialPostLink, photoURL *string) (Item, error) {
	i := Item{}

	u := goqu.Update("items")
	r := goqu.Record{}
	if externalLink != nil {
		r["external_link"] = *externalLink
	}
	if socialPostLink != nil {
		r["social_post_link"] = *socialPostLink
	}
	if photoURL != nil {
		r["photo_url"] = *photoURL
	}
	sql, params, err := u.Set(r).Where(goqu.Ex{"id": itemID, "display_id": displayID}).Returning("*").ToSQL()
	if err != nil {
		return i, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}

	err = q.DB.QueryRow(sql, params...).Scan(&i.ID, &i.UserID, &i.DisplayID, &i.ExternalLink, &i.SocialPostLink, &i.PhotoURL)
	if err != nil {
		return i, fmt.Errorf("failed to execute database query: %w", err)
	}

	return i, nil
}

func (q postgresQuerier) DeleteItem(itemID, displayID uuid.UUID) (bool, error) {
	sql, params, err := goqu.Delete("items").Where(goqu.Ex{"id": itemID, "display_id": displayID}).ToSQL()
	if err != nil {
		return false, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}

	res, err := q.DB.Exec(sql, params...)
	if err != nil {
		return false, fmt.Errorf("failed to execute database query: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("no clue ngl")
	}

	if rows != 1 {
		return false, nil
	}

	return true, nil
}

func (q postgresQuerier) IsItemOwner(userID, itemID uuid.UUID) (bool, error) {
	sql, params, err := goqu.Select("user_id").From("items").Where(goqu.Ex{"id": itemID}).ToSQL()
	if err != nil {
		return false, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}

	var uid uuid.UUID
	err = q.DB.QueryRow(sql, params...).Scan(&uid)
	if err != nil {
		return false, fmt.Errorf("failed to execute database query: %w", err)
	}

	if uid != userID {
		return false, nil
	}

	return true, nil
}

func (q postgresQuerier) IsDisplayOwner(userID, displayID uuid.UUID) (bool, error) {
	sql, params, err := goqu.Select("user_id").From("displays").Where(goqu.Ex{"id": displayID}).ToSQL()
	if err != nil {
		return false, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}

	var uid uuid.UUID
	err = q.DB.QueryRow(sql, params...).Scan(&uid)
	if err != nil {
		return false, fmt.Errorf("failed to execute database query: %w", err)
	}

	if uid != userID {
		return false, nil
	}

	return true, nil
}
