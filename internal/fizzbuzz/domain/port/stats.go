package port

type StatsRepository interface {
	IncrementKey(key string) error
	GetMostUsedKey() (key string, count int, err error)
}
