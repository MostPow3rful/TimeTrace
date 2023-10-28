package config

func (c *Config) Auth(username, passowrd string) (userExists bool) {
	for _, s := range c.Users {
		if (s.Name == username) && (s.Password == passowrd) {
			return true
		}
	}

	return false
}
