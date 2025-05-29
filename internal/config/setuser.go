package config

func (conf *Config) SetUser(current_user_name string) error {
	conf.CurrentUserName = current_user_name
	write(*conf)
	return nil
}
