package models

import "mime/multipart"

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

/* LOGIN */
type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

/* REGISTER */

type RegisterRequest struct {
	Username        string `json:"username" form:"username"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword"`
}

/* Scanner INPUT */

type Request struct {
	UrlId       string `json:"urlId" form:"urlId"`
	Path        string `json:"path" form:"path"`
	Type        string `json:"type" form:"type"`
	Params      string `json:"params" form:"params"`
	Cookie      string `json:"cookie" form:"cookie"`
	Wordlist    *multipart.FileHeader
	ParamsType  string `json:"paramsType" form:"paramsType"`
	Description string `json:"description" form:"description"`
}

/* Edit Profile Request */

type EditProfileRequest struct {
	Username        string `json:"username" form:"username"`
	Email           string `json:"email" form:"email"`
	OldPassword     string `json:"oldPassword" form:"oldPassword"`
	NewPassword     string `json:"newPassword" form:"newPassword"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword"`
}

/* Add Url INPUT */

type UrlRequest struct {
	Url string `json:"url" form:"url"`
}

/* FOR INDEX */

type Url struct {
	Id        int
	UserId    int
	Url       string
	ScanCount int
}

// For view page.
type Scanes struct {
	ScanId        int
	UserId        int
	UrlId         int
	Path          string
	Description   string
	Payload       string
	ContentLength int
	Status        int
}

// For pages in viewUrl.html
type Page struct {
	UrlId     string
	PageCount []int
	Start     int
	End       int
}
