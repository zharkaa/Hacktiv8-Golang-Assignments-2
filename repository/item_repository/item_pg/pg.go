package item_pg

import (
	"database/sql"
	"fmt"
	"h8-assignment-2/entity"
	"h8-assignment-2/pkg/errs"
	"h8-assignment-2/repository/item_repository"
)

const (
	getItemsByCodes = `
		SELECT * from "items" WHERE "item_code"
	`
)

type itemPG struct {
	db *sql.DB
}

func NewItemPG(db *sql.DB) item_repository.Repository {
	return &itemPG{
		db: db,
	}
}

func (itemPG *itemPG) generateGetItemsByCodesQuery(itemAmount int) string {
	baseQuery := `SELECT "item_id", "item_code", "quantity", "description", "order_id", "created_at" FROM "items"
	WHERE "item_code" IN`

	statement := "("

	for i := 1; i <= itemAmount; i++ {

		if i == itemAmount {
			statement += fmt.Sprintf("$%d)", i)
			break
		}

		statement += fmt.Sprintf("$%d,", i)

	}

	return fmt.Sprintf("%s %s", baseQuery, statement)
}

func (itemPG *itemPG) GetItemsByCodes(itemCodes []any) ([]entity.Item, errs.Error) {
	query := itemPG.generateGetItemsByCodesQuery(len(itemCodes))

	rows, err := itemPG.db.Query(query, itemCodes...)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	items := []entity.Item{}

	for rows.Next() {

		item := entity.Item{}

		err = rows.Scan(&item.ItemId, &item.ItemCode, &item.Quantity, &item.Description, &item.OrderId, &item.CreatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		items = append(items, item)
	}

	return items, nil
}
