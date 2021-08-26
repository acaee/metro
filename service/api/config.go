package api


type Config struct {
	Host	 string
	Post	 int
}

const (
	DefaultHost = "localhost"
	DefaultPost = 2000
)

func NewConfig() *Config {
	return &Config{
		Host: DefaultHost,
		Post: DefaultPost,
	}
}