package dbContext

import (
	"context"
	"database/sql"
	"time"

	"codeid.revampacademy/models/features"
)

type RecordTransactionUser struct {
	TrpaCodeNumber   string          `db:"trpa_code_number"`
	TrpaModifiedDate *time.Time      `db:"trpa_modified_date"`
	TrpaDebit        sql.NullFloat64 `db:"trpa_debit"`
	TrpaCredit       sql.NullFloat64 `db:"trpa_credit"`
	TrpaNote         string          `db:"trpa_note"`
	TrpaOrderNumber  sql.NullString  `db:"trpa_order_number"`
	TrpaFromID       string          `db:"trpa_from_id"`
	TrpaToID         string          `db:"trpa_to_id"`
	TrpaType         string          `db:"trpa_type"`
	UserName         string          `db:"user_name"`
}

const listPaymentTransaction_payment = `-- name: ListPaymentTransaction_payment :many
SELECT 
    trpa.trpa_code_number, 
    trpa.trpa_modified_date,
    trpa.trpa_debit,
    trpa.trpa_credit, 
    trpa.trpa_note,
    trpa.trpa_order_number,
    trpa.trpa_from_id, 
    trpa.trpa_to_id, 
    trpa.trpa_type,
    usr.user_name
FROM 
    payment.transaction_payment trpa
JOIN
    users.users usr
ON 
    trpa.trpa_user_entity_id = usr.user_entity_id
ORDER BY 
    trpa.trpa_code_number;
`

func (q *Queries) ListPaymentTransaction_payment(ctx context.Context) ([]RecordTransactionUser, error) {
	rows, err := q.db.QueryContext(ctx, listPaymentTransaction_payment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RecordTransactionUser
	for rows.Next() {
		var i RecordTransactionUser
		if err := rows.Scan(
			&i.TrpaCodeNumber,
			&i.TrpaModifiedDate,
			&i.TrpaDebit,
			&i.TrpaCredit,
			&i.TrpaNote,
			&i.TrpaOrderNumber,
			&i.TrpaFromID,
			&i.TrpaToID,
			&i.TrpaType,
			&i.UserName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

const getPaymentTransaction_payment = `-- name: GetPaymentTransaction_payment :many
SELECT 
    trpa.trpa_code_number, 
    trpa.trpa_modified_date,
    trpa.trpa_debit,
    trpa.trpa_credit, 
    trpa.trpa_note,
    trpa.trpa_order_number,
    trpa.trpa_from_id, 
    trpa.trpa_to_id, 
    trpa.trpa_type,
    usr.user_name
FROM 
    payment.transaction_payment trpa
JOIN
    users.users usr
ON 
    trpa.trpa_user_entity_id = usr.user_entity_id
WHERE 
	trpa.trpa_user_entity_id = $1
ORDER BY 
    trpa.trpa_code_number
	LIMIT $2 OFFSET $3;
	`

// LIMIT 5 OFFSET ($2 - 1) * $3;

// payment.transaction_payment
func (q *Queries) GetPaymentTransaction_payment(ctx context.Context, metadata *features.Metadata) ([]RecordTransactionUser, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentTransaction_payment, metadata.SearchBy, metadata.PageSize, metadata.PageNo)
	// *metadata.PageSize
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RecordTransactionUser
	for rows.Next() {
		var i RecordTransactionUser
		if err := rows.Scan(
			&i.TrpaCodeNumber,
			&i.TrpaModifiedDate,
			&i.TrpaDebit,
			&i.TrpaCredit,
			&i.TrpaNote,
			&i.TrpaOrderNumber,
			&i.TrpaFromID,
			&i.TrpaToID,
			&i.TrpaType,
			&i.UserName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const recordPaymentTransactionUser = `-- name: RecordPaymentTransactionUser :exec
INSERT INTO
    payment.transaction_payment (
        trpa_debit,
        trpa_credit,
        trpa_type,
        trpa_note,
        trpa_from_id,
        trpa_to_id,
        trpa_user_entity_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING
		trpa_code_number,
		trpa_modified_date,
		trpa_debit,
		trpa_credit,
		trpa_note,
		trpa_order_number,
		trpa_from_id,
		trpa_to_id,
		trpa_type,
		(SELECT user_name FROM users.users WHERE user_entity_id = $7) AS user_name;
`

type RecordTransactionUserParams struct {
	TrpaDebit        sql.NullFloat64 `db:"trpa_debit" json:"trpaDebit"`
	TrpaCredit       sql.NullFloat64 `db:"trpa_credit" json:"trpaCredit"`
	TrpaType         string          `db:"trpa_type" json:"trpaType"`
	TrpaNote         string          `db:"trpa_note" json:"trpaNote"`
	TrpaFromID       string          `db:"trpa_from_id" json:"trpaFromId"`
	TrpaToID         string          `db:"trpa_to_id" json:"trpaToId"`
	TrpaUserEntityID int32           `db:"trpa_user_entity_id" json:"trpaUserEntityId"`
}

func (q *Queries) RecordPaymentTransactionUser(ctx context.Context, params RecordTransactionUserParams) (*RecordTransactionUser, error) {
	row := q.db.QueryRowContext(ctx, recordPaymentTransactionUser,
		params.TrpaDebit,
		params.TrpaCredit,
		params.TrpaType,
		params.TrpaNote,
		params.TrpaFromID,
		params.TrpaToID,
		params.TrpaUserEntityID,
	)
	var i RecordTransactionUser
	err := row.Scan(
		&i.TrpaCodeNumber,
		&i.TrpaModifiedDate,
		&i.TrpaDebit,
		&i.TrpaCredit,
		&i.TrpaNote,
		&i.TrpaOrderNumber,
		&i.TrpaFromID,
		&i.TrpaToID,
		&i.TrpaType,
		&i.UserName,
	)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

const deletePaymentTransaction_payment = `-- name: DeletePaymentTransaction_payment :exec
DELETE FROM payment.transaction_payment WHERE trpa_code_number = $1
`

func (q *Queries) DeletePaymentTransaction_payment(ctx context.Context, trpaCodeNumber string) error {
	_, err := q.db.ExecContext(ctx, deletePaymentTransaction_payment, trpaCodeNumber)
	return err
}
