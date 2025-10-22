package redis

import (
	"bisnis-be/pkg/errors"
	"context"
	"encoding/json"
	"time"
)

func (d Data) AddToRedis(ctx context.Context, data interface{}, key string, ttl time.Duration) (err error) {
	jsoned, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "[addToRedis]")
	}

	// return d.rdb.Set(ctx, key, jsoned, 3600*time.Second).Err()
	return d.rdb.Set(ctx, key, jsoned, ttl).Err()
}

// JSON BASED
func (d Data) GetFromRedis(ctx context.Context, key string, dest interface{}) (err error) {
	result, err := d.rdb.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(result, &dest)
}
