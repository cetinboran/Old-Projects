package api

import (
	"github.com/cetinboran/basicsec/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuth(c *fiber.Ctx) error {

	// Cookie'den auth'u çekiyorum.
	tokenString := c.Cookies("Auth")

	// Authorization başlığı yoksa veya uygun biçimde değilse hata döndürün
	if tokenString == "" {
		// Eğer token yok ise index'de ise register'a gitsin
		// Başka bir yerde ise forbidden'a
		if c.Path() == "/" {
			return c.Redirect("Auth/Register")
		}
		return c.Status(fiber.StatusUnauthorized).Redirect("/forbidden")
	}

	// Burada tokeni parse ediyoruz secret'a göre
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := []byte(config.GetSecret())
		return secret, nil
	})

	// Eğer parse yaparken bir hata çıktıysa zaten valid değildir.
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).Redirect("/forbidden")
	}

	// Bilgileri local'e atıyorum ki ileride kullanabileyim.
	// token valid değil mi diye önceden baktığım için ikinci arg _
	claims, _ := token.Claims.(jwt.MapClaims)

	c.Locals("id", claims["ID"])
	c.Locals("username", claims["username"])

	return c.Next()
}
