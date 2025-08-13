package api

import (
	"fmt"
	"log"

	"github.com/cetinboran/basicsec/database"
	"github.com/cetinboran/basicsec/models"
	"github.com/cetinboran/basicsec/utility"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	db := database.DBConn
	userId := c.Locals("id")

	// Ekrana yazacağım url'leri çekiyorum
	rows, err := db.Query("SELECT * FROM urls WHERE user_id = ?", userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Bütün urlleri sql.rows dan çıkarıp bir struct arrayine atıyorum.
	var Urls []models.Url
	for rows.Next() {
		var url models.Url
		if err := rows.Scan(&url.Id, &url.UserId, &url.Url); err != nil {
			log.Fatal(err)
		}

		// Scanes'den gelen url'id ye göre scan countu hesaplayıp struct içine atıyorum
		// İşim kolaylaşsın diye.

		var count int
		query := "SELECT COUNT(*) FROM scanes WHERE url_id = ? AND user_id = ?"
		row := db.QueryRow(query, url.Id, userId)

		row.Scan(&count)
		url.ScanCount = count

		Urls = append(Urls, url)
	}

	// Auth'tan gelen bilgileri yolluyorum.
	data := fiber.Map{
		"Id":       c.Locals("id"),
		"Username": c.Locals("username"),
		"Urls":     Urls,
	}

	// fmt.Println(data)
	return c.Render("index", data)
}

func Forbidden(c *fiber.Ctx) error {
	return c.Render("forbidden", fiber.Map{})
}

func Login(c *fiber.Ctx) error {
	m := c.Queries()

	data := fiber.Map{}

	value, ok := m["error"]
	if ok {
		switch value {
		case "1":
			data["Error"] = "Invalid Credentials"
		}
	} else {
		// Burada default value'lar atıyorum.
		data["Error"] = ""
	}

	return c.Render("login", data)
}

func Register(c *fiber.Ctx) error {
	m := c.Queries()
	data := fiber.Map{}

	value, ok := m["error"]
	if ok {
		data = RegisterErrors(value)
	}

	return c.Render("register", data)
}

func Scan(c *fiber.Ctx) error {
	db := database.DBConn
	m := c.Queries()
	data := fiber.Map{}

	// Eğer error var ise buraya giriyor gereli verileri ekleyip çıkıyor.
	value, ok := m["error"]
	if ok {
		data = ScanErrors(value)
	} else {
		data = fiber.Map{}
	}

	value, ok = m["urlId"]
	if ok {
		data["urlId"] = value

		// url Id ye göre Url yi çekiyorum ekranda görünsün diye.
		stmt, err := db.Prepare("SELECT * FROM urls WHERE id = ?")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		row := stmt.QueryRow(value)

		url := models.Url{}
		row.Scan(&url.Id, &url.UserId, &url.Url)

		data["Url"] = url.Url
		data["Username"] = c.Locals("username")

	} else {
		// Eğer id gelmez ise karşıya geri git hatalı geldi.
		// Burada bir fonksiyon ile gelen id valid mi diye bakmak lazım olabilir.
		return c.Status(fiber.StatusBadRequest).Redirect("/")
	}

	return c.Render("scan", data)
}

func AddUrl(c *fiber.Ctx) error {
	m := c.Queries()
	data := fiber.Map{}

	value, ok := m["error"]
	if ok {
		data = AddUrlErrors(value)
		data["Username"] = c.Locals("username")
	}

	return c.Render("addUrl", data)
}

func DeleteUrl(c *fiber.Ctx) error {
	db := database.DBConn
	m := c.Queries()

	value, ok := m["urlId"]
	if ok {
		query := "DELETE FROM urls WHERE id = ?"
		stmt, err := db.Prepare(query)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).Redirect("/")
		}

		stmt.Exec(value)
	}

	return c.Status(fiber.StatusOK).Redirect("/")
}

func DeleteScan(c *fiber.Ctx) error {
	db := database.DBConn
	m := c.Queries()

	var urlId string
	value, ok := m["scanId"]
	if ok {
		userId := m["userId"]

		// Başka kullanıcının scan'ını silmesin diye userId de ekledim.
		_, err := db.Exec("DELETE FROM scanes WHERE scan_id = ? AND user_id = ?", value, userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		urlId = m["urlId"]
	} else {
		return c.Status(fiber.StatusBadRequest).Redirect("/")
	}

	// Geri view'e girebilsin diye urlId yi de ekliyorum.
	redirectValue := fmt.Sprintf("/Scan/View?urlId=%v", urlId)
	return c.Status(fiber.StatusOK).Redirect(redirectValue)
}

func View(c *fiber.Ctx) error {
	db := database.DBConn
	m := c.Queries()

	data := fiber.Map{}

	value, ok := m["urlId"]
	if ok {
		// Getting the scane results

		query := "SELECT * FROM scanes WHERE url_id = ?"
		rows, err := db.Query(query, value)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var scanes []models.Scanes
		for rows.Next() {
			scan := models.Scanes{}
			err := rows.Scan(&scan.ScanId, &scan.UserId, &scan.UrlId, &scan.Path, &scan.Description, &scan.Payload, &scan.ContentLength, &scan.Status)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			scanes = append(scanes, scan)
		}

		// Gettin the url of the scan that we are viewing.
		var url string
		urlRow := db.QueryRow("SELECT url FROM urls WHERE id = ?", value)
		err = urlRow.Scan(&url)

		page, ok := m["page"]
		if ok {
			data["Page"] = page
		} else {
			data["Page"] = "1"
			page = "1"
		}

		pages := utility.Pages(scanes, page)
		pages.UrlId = value
		// For pages

		data["Pages"] = pages
		data["Scanes"] = scanes
		data["Url"] = url
		data["UrlId"] = value

		data["Id"] = c.Locals("id")
		data["Username"] = c.Locals("username")
	} else {
		// Eğer urlid gelmediye burada işi yok geri dönsün.
		return c.Status(fiber.StatusBadRequest).Redirect("/")
	}

	return c.Render("viewScanes", data)
}

func DeleteAllScan(c *fiber.Ctx) error {
	db := database.DBConn
	m := c.Queries()

	value, ok := m["urlId"]
	if ok {
		userId := m["userId"]
		query := "DELETE FROM scanes WHERE url_id = ? AND user_id = ?"
		stmt, err := db.Prepare(query)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).Redirect("/")
		}

		stmt.Exec(value, userId)
	} else {
		return c.Status(fiber.StatusBadRequest).Redirect("/")
	}

	redirectValue := fmt.Sprintf("/Scan/View?urlId=%v", value)
	return c.Status(fiber.StatusOK).Redirect(redirectValue)
}

func Profile(c *fiber.Ctx) error {
	db := database.DBConn
	userId := c.Locals("id")

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", userId)

	user := models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	data := fiber.Map{
		"UserId":   user.ID,
		"Username": user.Username,
		"Email":    user.Email,
	}

	return c.Render("profile", data)
}

func EditProfile(c *fiber.Ctx) error {
	m := c.Queries()
	db := database.DBConn
	userId := c.Locals("id")

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", userId)

	user := models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var data fiber.Map

	value, ok := m["error"]
	if ok {
		data = EditProfileErrors(value)
	} else {
		// data ' ya girmezse hata veriyor boş yolluyoruz.
		data = fiber.Map{}
	}

	data["UserId"] = user.ID
	data["Username"] = user.Username
	data["Email"] = user.Email

	return c.Render("profileEdit", data)
}

func Logout(c *fiber.Ctx) error {
	c.ClearCookie("Auth")

	return c.Status(fiber.StatusOK).Redirect("/Auth/Login")
}

func DeleteProfile(c *fiber.Ctx) error {
	db := database.DBConn
	userId := c.Locals("id")

	_, err := db.Exec("DELETE FROM urls WHERE user_id = ?", userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err = db.Exec("DELETE FROM scanes WHERE user_id = ?", userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err = db.Exec("DELETE FROM users WHERE id = ?", userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).Redirect("/Logout")
}
