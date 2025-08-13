package api

import "github.com/gofiber/fiber/v2"

func ScanErrors(errorId string) fiber.Map {
	data := fiber.Map{}
	switch errorId {
	case "1":
		data["Error"] = "Invalid URL"
	case "2":
		data["Error"] = "Invalid Request Type"
	case "3":
		data["Error"] = "Please Enter Params"
	case "4":
		data["Error"] = "This extention not allowed"
	case "5":
		data["Error"] = "You need to add SQL Keyword\n to the URL or Params"
	case "6":
		data["Error"] = "Please enter the parameters in the desired format"
	}

	return data
}

func RegisterErrors(errorId string) fiber.Map {
	data := fiber.Map{}

	switch errorId {
	case "1":
		data["Error"] = "This username already being used"
	case "2":
		data["Error"] = "Username length must be at least 3 character"
	case "3":
		data["Error"] = "Passwords do not match"
	}

	return data
}

func AddUrlErrors(errorId string) fiber.Map {
	data := fiber.Map{}

	switch errorId {
	case "1":
		data["Error"] = "Please Enter Valid Url"
	}

	return data
}

func EditProfileErrors(errorId string) fiber.Map {
	data := fiber.Map{}

	switch errorId {
	case "1":
		data["Error"] = "This username already being used"
	case "2":
		data["Error"] = "Username length must be at least 3 character"
	case "3":
		data["Error"] = "Old password is Invalid"
	case "4":
		data["Error"] = "Passwords do not match"
	case "5":
		data["Error"] = "E-mail Invalid Format"
	}

	return data
}
