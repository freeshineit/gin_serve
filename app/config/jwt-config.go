package config

type JWTConfig struct {
	Secret string
	Issuer string
}

func GetJWTConfig() *JWTConfig {
	secret := Conf.GetString("jwt.secret")
	issuer := Conf.GetString("jwt.issuer")

	if secret == "" {
		secret = "xiaoshaoqq@gmail.com,.<>?"
	}

	if issuer == "" {
		issuer = "xiaoshaoqq@gmail.com"
	}

	return &JWTConfig{
		Secret: secret,
		Issuer: issuer,
	}
}
