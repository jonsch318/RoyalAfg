package mw

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const CSRFCookieName = "xcsrf"

//ValidateCSRF compares the two csrf tokens for a double submit cookie csrf protection.
func ValidateCSRF(res *http.Request) error {
	token := res.Header.Get("X-CSRF-Token")

	if token == "" {
		return errors.New("header token is empty")
	}

	cookie, err := res.Cookie(CSRFCookieName)
	if err != nil {
		return err
	}
	if cookie.HttpOnly {
		return errors.New("csrf cookie had no http only cookie")
	}

	cookieSplit := strings.Split(cookie.Value, ":")
	if len(cookieSplit) != 2 {
		return fmt.Errorf("csrf cookie had wrong format %v", cookieSplit)
	}

	if cookieSplit[1] != token {
		return fmt.Errorf("csrf cookie and passed token where not identical")
	}
	if ok := verify(cookieSplit[0], cookieSplit[1]); !ok {
		return errors.New("token could not be reconstructed")
	}
	return nil
}

//createToken follows the implementation of the clientside csrf token library (https://github.com/pillarjs/csrf/)
func verify(secret, token string) bool {
	var index = strings.Index(token, "-")
	if index < 0 {
		return false
	}
	salt := token[:index]
	expected := createTokenFromSalt(secret, salt)
	if expected != token {
		return false
	}
	return true
}

func createTokenFromSalt(secret, salt string) string {
	h := sha1.New()
	h.Write([]byte(salt + "-" + secret))
	enc := base64.StdEncoding.EncodeToString(h.Sum(nil))
	enc = strings.ReplaceAll(enc, "/", "_")
	enc = strings.ReplaceAll(enc, "=", "")
	enc = strings.ReplaceAll(enc, "+", "-")

	return salt + "-" + enc
}
