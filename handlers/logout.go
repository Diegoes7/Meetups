package handlers

import (
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Remove the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // delete
		HttpOnly: true,
		Secure:   true, // match how you set it
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Logged out"}`))
}
