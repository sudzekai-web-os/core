package core

type IConfiguration interface {
	GetString(key string) *string
	GetBool(key string) *bool
	GetInt(key string) *int

	GetValue(key string) any

	GetOptions() map[string]func() any
}
