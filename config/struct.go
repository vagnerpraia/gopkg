package gpconfig

type Config struct {
}

func NewConfig() *Config {

	return &Config{}
}

func (that *Config) Write(path string, output any) error {

	return Write(path, output)
}

func (that *Config) Unmarshal(path string, output any) error {

	return Unmarshal(path, output)
}
