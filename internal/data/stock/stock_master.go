package bisnis

// // JSON BASED
// func (d Data) addToRedis(ctx context.Context, data interface{}, key string) (err error) {
// 	jsoned, err := json.Marshal(data)
// 	if err != nil {
// 		return errors.Wrap(err, "[addToRedis]")
// 	}

// 	return d.rdb.Set(ctx, key, jsoned, 3600*time.Second).Err()
// }

// // JSON BASED
// func (d Data) getFromRedis(ctx context.Context, key string, dest interface{}) (err error) {
// 	result, err := d.rdb.Get(ctx, key).Bytes()
// 	if err != nil {
// 		return err
// 	}

// 	return json.Unmarshal(result, &dest)
// }

// func (d Data) GetOneStockProduct(ctx context.Context, stockcode string, stockname string, stockid string) (goldStockEntity.GetOneStock, error) {
// 	var (
// 		user goldStockEntity.GetOneStock
// 		// users []goldEntity.GetGoldUser
// 		err error
// 	)
// 	log.Println("data GetGoldUser object", stockid, stockid, stockcode, stockcode)
// 	log.Println("data GetGoldUser object12345", stockcode)
// 	rows, err := (*d.stmt)[getOneStockProduct].QueryxContext(ctx, stockcode, "%"+stockname+"%", "%"+stockname+"%", stockid, stockid)
// 	if err != nil {
// 		return user, errors.Wrap(err, "[DATA] [GetGoldUser]")
// 	}
// 	log.Println("datagolduser", user)

// 	defer rows.Close()

// 	for rows.Next() {
// 		if err = rows.StructScan(&user); err != nil {
// 			return user, errors.Wrap(err, "[DATA] [GetGoldUser]")
// 		}
// 		// users = append(users, user)
// 	}
// 	return user, err
// }

// func (d Data) GetAllStockHeaderToRedis(ctx context.Context) (users []goldStockEntity.GetOneStock, err error) {
// 	// var (
// 	// 	user  goldStockEntity.GetOneStock
// 	// users []goldStockEntity.GetOneStock
// 	// 	err   error
// 	// )

// 	var (
// 		rdbKey = "bisnis-be:getallstockheader"
// 	)

// 	err = d.getFromRedis(ctx, rdbKey, &users)
// 	if err == redis.Nil {
// 		users, err := d.GetAllStockHeader(ctx)
// 		if err != nil {
// 			return users, errors.Wrap(err, "[DATA][GetAllStockHeaderToRedis]")
// 		}

// 		if ok := d.addToRedis(ctx, users, rdbKey); ok != nil {
// 			return []goldStockEntity.GetOneStock{}, errors.Wrap(ok, "[DATA][GetAllStockHeaderToRedis]")
// 		}

// 		return users, nil

// 	} else if err != nil {
// 		return users, err
// 	}

// 	// log.Println("data GetGoldUser object")
// 	// rows, err := (*d.stmt)[getAllStockHeader].QueryxContext(ctx)
// 	// if err != nil {
// 	// 	return users, errors.Wrap(err, "[DATA] [GetAllStockHeader]")
// 	// }
// 	// log.Println("datagolduser", user)

// 	// defer rows.Close()

// 	// for rows.Next() {
// 	// 	if err = rows.StructScan(&user); err != nil {
// 	// 		return users, errors.Wrap(err, "[DATA] [GetAllStockHeader]")
// 	// 	}
// 	// 	users = append(users, user)
// 	// }
// 	return users, err
// }

// func (d Data) InsertStockSalesToRedis(ctx context.Context, stock goldStockEntity.InsertStock) (string, error) {
// 	var result string
// 	var err error

// 	_, err = (*d.stmt)[insertStockSales].ExecContext(ctx,
// 		stock.StockID,
// 		stock.StockCode,
// 		stock.StockName,
// 		stock.StockPack,
// 		stock.StockQTY,
// 		stock.StockPrice,
// 		stock.StockUpdateBy,
// 	)

// 	log.Println("data stock object", stock)

// 	if err != nil {
// 		result = "Gagal"
// 		return result, errors.Wrap(err, "[DATA][InsertStockSales]")
// 	}
// 	result = "Sukses"

// 	return result, err

// }
