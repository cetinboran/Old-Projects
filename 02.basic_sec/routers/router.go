package routers

import (
	"github.com/cetinboran/basicsec/api"
	"github.com/gofiber/fiber/v2"
)

func SetRouters(app *fiber.App) {

	app.Get("/Forbidden", api.Forbidden)

	// JWTAuth ile kontrol ediyorum. İzni varsa giriyor.

	app.Get("/", api.JwtAuth, api.Index)
	app.Get("/Logout", api.JwtAuth, api.Logout)

	Profile := app.Group("/Profile")
	Profile.Get("/", api.JwtAuth, api.Profile)
	Profile.Get("/Edit", api.JwtAuth, api.EditProfile)
	Profile.Get("Delete", api.JwtAuth, api.DeleteProfile)
	Profile.Post("/EditProfileForm", api.JwtAuth, api.EditProfileForm)

	URL := app.Group("/Url")
	URL.Get("/Add", api.JwtAuth, api.AddUrl)
	URL.Post("/Form", api.JwtAuth, api.UrlForm)
	URL.Get("/Delete", api.JwtAuth, api.DeleteUrl)

	Scan := app.Group("/Scan")
	Scan.Get("/", api.JwtAuth, api.Scan)
	Scan.Post("/ScanForm", api.JwtAuth, api.ScanForm) // This is for search form.
	Scan.Get("/View", api.JwtAuth, api.View)
	Scan.Get("/Delete", api.JwtAuth, api.DeleteScan)
	Scan.Get("/DeleteAll", api.JwtAuth, api.DeleteAllScan)

	Auth := app.Group("/Auth")

	Auth.Get("/Login", api.Login)
	Auth.Post("/LoginForm", api.LoginForm)

	Auth.Get("/Register", api.Register)
	Auth.Post("/RegisterForm", api.RegisterForm)
}
