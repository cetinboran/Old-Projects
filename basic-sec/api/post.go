package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/cetinboran/basicsec/basicsec"
	"github.com/cetinboran/basicsec/config"
	"github.com/cetinboran/basicsec/database"
	"github.com/cetinboran/basicsec/models"
	"github.com/cetinboran/basicsec/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LoginForm(c *fiber.Ctx) error {
	db := database.DBConn

	// Verileri struct'a yüklüyorum
	LoginRequest := models.LoginRequest{}
	if err := c.BodyParser(&LoginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	query := "SELECT * FROM users WHERE username = ? AND password = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer stmt.Close()

	hashedPassword := utility.ConvertToMd5(LoginRequest.Password)

	// User objesine gerekli bilgileri atıyorum
	user := models.User{}
	err = stmt.QueryRow(LoginRequest.Username, hashedPassword).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(403).Redirect("login?error=1")
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	// Token için bilgileri jwt mapi içine attık.
	day := time.Hour * 24
	claims := jwt.MapClaims{
		"ID":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(day * 1).Unix(),
	}

	// Tokeni oluşturuyoruz
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tokeni secret ile şifreliyoruz.
	cookie, err := token.SignedString([]byte(config.GetSecret()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Cookie'yi setliyorum.
	// Bu proje web app olduğu için burada yapıyoruz.
	// Normalde front'a yollanır token ve karşı taraf set eder.
	c.Cookie(&fiber.Cookie{
		Name:     "Auth",
		Value:    cookie,
		HTTPOnly: true,
	})

	return c.Status(202).Redirect("/")
}

func RegisterForm(c *fiber.Ctx) error {
	db := database.DBConn

	RegisterRequest := models.RegisterRequest{}
	if err := c.BodyParser(&RegisterRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Kontrol ediyorum böyle username var mı yok mu
	stmt, err := db.Prepare("SELECT * FROM users WHERE username = ?")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer stmt.Close()

	rows, err := stmt.Query(RegisterRequest.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer rows.Close()

	// Username var ise hata kodu yolluyorum.
	if rows.Next() {
		return c.Status(403).Redirect("register?error=1")
	}

	// Gelen veriyi kontrol ediyorum.
	if len(RegisterRequest.Username) < 3 {
		return c.Status(403).Redirect("register?error=2")
	}

	if RegisterRequest.Password != RegisterRequest.ConfirmPassword {
		return c.Status(403).Redirect("register?error=3")
	}

	RegisterRequest.Password = strings.TrimSpace(RegisterRequest.Password)
	hashedPassword := utility.ConvertToMd5(RegisterRequest.Password)

	// Güvenlilik için preapre sorgusu yazıyoruz
	// mysql de yer tutucu işareti ?
	stmt, err = db.Prepare("INSERT INTO users (username, password, email) VALUES (?, ?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(RegisterRequest.Username, hashedPassword, RegisterRequest.Email)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(202).Redirect("login")
}

func ScanForm(c *fiber.Ctx) error {
	db := database.DBConn

	SqlRequest := models.Request{}
	err := c.BodyParser(&SqlRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	file, err := c.FormFile("wordlist")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	SqlRequest.Wordlist = file

	// Önce gelen urlId den url yi çekicem
	query := fmt.Sprintf("SELECT * FROM urls WHERE id = %v", SqlRequest.UrlId)
	urlData := db.QueryRow(query)

	// Url objesini oluşturup scan ile değerleri içine atıyorum.
	newUrl := models.Url{}
	urlData.Scan(&newUrl.Id, &newUrl.UserId, &newUrl.Url)

	// Eğer gelen dosya türü txt değil ise error yolluyorum
	if file.Header.Get("Content-Type") != "text/plain" {
		url := fmt.Sprintf("/Scan?error=4&urlId=%v", newUrl.Id)
		return c.Status(fiber.StatusBadRequest).Redirect(url)
	}

	parsedUrl, err := url.Parse(newUrl.Url)
	if err != nil {
		// Eğer parse işleminde bir sıkıntı var ise index'e error yolluyorum
		// silsin diye url yi yanlış yazmış
		return c.Status(fiber.StatusBadRequest).Redirect("/?error=1")
	}

	// Url'ye gelen path'i ekliyorum.
	finalURL := *parsedUrl
	finalURL.Path += SqlRequest.Path

	// Inputlari basicsec paketimdeki request struct'ına alıyorum
	// Request işlemlerini oradan yapıcam.
	SQL := basicsec.RequestInit()
	errorId := SQL.TakeInputs(finalURL.String(), SqlRequest.Type, SqlRequest.Params, SqlRequest.Cookie, SqlRequest.Wordlist)
	if errorId != -1 {
		// Burada urlId yi yollamazsam otomatik olarak index'e atıyor.
		// Query'leri filtrelemem gerekiyor.
		url := fmt.Sprintf("/Scan?error=%v&urlId=%v", errorId, newUrl.Id)
		return c.Status(fiber.StatusBadRequest).Redirect(url)
	}

	response := SQL.Start()

	// url hatalı girerse boş geliyor
	// Alttaka boşu filtreleme yapmasın diye kontrol ediyorum boş ise direkt bitsin.
	if len(response) == 0 {
		return c.Status(fiber.StatusOK).Redirect("/")
	}

	// Bütün veriler buraya geliyor.
	// Burada içeriden tarama yapabilrsin.
	// Hepsini db ye de atabilirsin
	// düşüncene kalmış
	// İçindeki verileri kendine göre ayarla ve db ye kaydet.
	filtiredResponses := utility.FilterResponse(finalURL.String(), response)
	// fmt.Println(filtiredResponses) // Şimdi bunu dB kaydet

	userId := c.Locals("id")
	for _, r := range filtiredResponses {
		query := "INSERT INTO scanes (user_id,url_id,path,description,payload,content_length,status) VALUES (?,?,?,?,?,?,?)"

		stmt, err := db.Prepare(query)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		stmt.Exec(userId, SqlRequest.UrlId, SqlRequest.Path, SqlRequest.Description, r.Line, r.ContentLength, r.Status)
	}

	return c.Status(fiber.StatusOK).Redirect("/")
}

func UrlForm(c *fiber.Ctx) error {
	db := database.DBConn
	userId := c.Locals("id")

	urlRequest := models.UrlRequest{}
	err := c.BodyParser(&urlRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// şimdilik http içermiyorsa url hata yolluyorum
	// ileride değiştir
	if !strings.Contains(urlRequest.Url, "http") {
		return c.Status(fiber.StatusBadRequest).Redirect("/Url/Add?error=1")
	}

	stmt, err := db.Prepare("INSERT INTO urls (user_id, url) VALUES (?,?)")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Ekliyorum db ye url yi
	_, err = stmt.Exec(userId, urlRequest.Url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).Redirect("/")
}

func EditProfileForm(c *fiber.Ctx) error {
	db := database.DBConn

	userId := c.Locals("id")
	// User'ı çekiyorum. Karşılaştırmak için
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", userId)

	user := models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	newProfile := models.EditProfileRequest{}
	err = c.BodyParser(&newProfile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Redirect("/Profile")
	}

	hashedOldPassword := utility.ConvertToMd5(newProfile.OldPassword)
	if user.Password != hashedOldPassword {
		return c.Status(fiber.StatusBadRequest).Redirect("Edit?error=3")
	}

	var name, email, pass bool
	if newProfile.Username != user.Username {
		name = true
		if len(newProfile.Username) < 3 {
			return c.Status(fiber.StatusBadRequest).Redirect("Edit?error=2")
		}

		rows, err := db.Query("SELECT * FROM users WHERE username = ?", newProfile.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer rows.Close()

		if rows.Next() {
			return c.Status(fiber.StatusBadRequest).Redirect("Edit?error=1")
		}
	}

	if newProfile.Email != user.Email {
		email = true

		// basit kontrol şimdilik
		if !strings.Contains(newProfile.Email, "@") {
			return c.Status(fiber.StatusBadRequest).Redirect("Edit?error=5")
		}
	}

	if newProfile.NewPassword != "" {
		pass = true
		if newProfile.NewPassword != newProfile.ConfirmPassword {
			return c.Status(fiber.StatusBadRequest).Redirect("Edit?error=4")
		}
	}
	query := "UPDATE users SET "
	var args []interface{}
	if name {
		query += "username = ?,"
		args = append(args, newProfile.Username)
	}

	if email {
		query += "email = ?,"
		args = append(args, newProfile.Email)
	}

	if pass {
		hashedPassword := utility.ConvertToMd5(newProfile.NewPassword)

		query += "password = ?,"
		args = append(args, hashedPassword)
	}

	query = strings.TrimRight(query, ",")
	query += " WHERE id = ?"
	args = append(args, userId)

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Hepsini interface arrayine topladım ve exec içine args... ile verdim.
	_, err = stmt.Exec(args...)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).Redirect("/logout")
}
