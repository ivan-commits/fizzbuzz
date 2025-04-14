package contract

import "context"

type StatsRepository interface {
	IncrementKey(c context.Context, key string) error
	GetMostUsedKey(c context.Context) (key string, count int, err error)
}
