package bisnis

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
