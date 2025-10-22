package bisnis

import (
	bisnisEntity "bisnis-be/internal/entity/bisnis"
	"bisnis-be/pkg/errors"
	"context"
)

// func (d Data) GetGoldUser(ctx context.Context) ([]goldEntity.GetGoldUser, error) {
// 	var (
// 		user  goldEntity.GetGoldUser
// 		users []goldEntity.GetGoldUser
// 		err   error
// 	)
// 	log.Println("data GetGoldUser object")
// 	rows, err := (*d.stmt)[getGoldUser].QueryxContext(ctx)
// 	if err != nil {
// 		return users, errors.Wrap(err, "[DATA] [GetGoldUser]")
// 	}
// 	log.Println("datagolduser", users)

// 	defer rows.Close()

// 	for rows.Next() {
// 		if err = rows.StructScan(&user); err != nil {
// 			return users, errors.Wrap(err, "[DATA] [GetGoldUser]")
// 		}
// 		users = append(users, user)
// 	}
// 	return users, err
// }

func (d Data) AddTransaction(ctx context.Context, addTransaction bisnisEntity.AddTransaction) (int, error) {
	var (
		transaction bisnisEntity.Transaction
		err         error
	)
	_, err = (*d.stmt)[insertTransaction].ExecContext(ctx,
		addTransaction.AgentID,
		addTransaction.ProductID,
		addTransaction.Nama,
		addTransaction.Usia,
		addTransaction.Premium,
	)
	if err != nil {
		return transaction.TransID, errors.Wrap(err, "[DATA][AddTransaction][insertTransaction]")
	}

	rows, err := (*d.stmt)[getTransaction].QueryxContext(ctx, addTransaction.AgentID,
		addTransaction.ProductID,
		addTransaction.Nama,
		addTransaction.Usia,
		addTransaction.Premium)
	if err != nil {
		return transaction.TransID, errors.Wrap(err, "[DATA] [AddTransaction][getTransaction]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&transaction); err != nil {
			return transaction.TransID, errors.Wrap(err, "[DATA][AddTransaction][StructScan]")
		}
	}
	return transaction.TransID, err
}

func (d Data) DeleteTransaction(ctx context.Context, deleteTransactionz bisnisEntity.DeleteTransaction) (int, string, error) {
	var (
		transaction bisnisEntity.Transaction
		result      string
		err         error
	)
	res, err := (*d.stmt)[deleteTransaction].ExecContext(ctx,
		deleteTransactionz.AgentID,
		deleteTransactionz.TransID,
	)
	if err != nil {
		return transaction.TransID, "Error ExecContext", errors.Wrap(err, "[DATA][AddTransaction][insertTransaction]")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return transaction.TransID, "Error RowsAffected", errors.Wrap(err, "[DATA][DeleteTransaction][RowsAffected]")
	}

	if rowsAffected == 0 {
		result = "Data not Found"
		return transaction.TransID, result, errors.New("no transaction deleted")
	}
	return transaction.TransID, "Success", err
}

func (d Data) UpdateTransaction(ctx context.Context, addTransaction bisnisEntity.UpdateTransaction) (int, string, error) {
	var (
		transaction bisnisEntity.Transaction
		result      string
		err         error
	)
	res, err := (*d.stmt)[updateTransaction].ExecContext(ctx,
		addTransaction.Nama,
		addTransaction.Usia,
		addTransaction.Premium,
		addTransaction.TransID,
		addTransaction.AgentID,
		addTransaction.ProductID,
	)
	if err != nil {
		return transaction.TransID, "Error ExecContext", errors.Wrap(err, "[DATA][AddTransaction][insertTransaction]")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return transaction.TransID, "Error RowsAffected", errors.Wrap(err, "[DATA][DeleteTransaction][RowsAffected]")
	}

	if rowsAffected == 0 {
		result = "Data not Updated"
		return transaction.TransID, result, errors.New("no transaction deleted")
	}

	result = "Success"

	return transaction.TransID, result, err
}
