package auth

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
)

type credentials struct {
	Email    string `json:"email" xml:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" validate:"required,min=8"`
}

func credentialsFromRequest(r *http.Request) (credentials, error) {
	var c credentials

	switch strings.SplitN(r.Header.Get("Content-Type"), ";", 2)[0] {
	case "application/json":
		return c, json.NewDecoder(r.Body).Decode(&c)
	case "application/xml":
		return c, xml.NewDecoder(r.Body).Decode(&c)
	case "application/x-www-form-urlencoded":
		if err := r.ParseForm(); err != nil {
			return c, err
		}
		c.Email = r.FormValue("email")
		c.Password = r.FormValue("password")
		return c, nil
	default:
		return c, nil
	}
}
