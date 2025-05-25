package handlers

import (
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Remove the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    "", // set cookie to empty, so no user
		Path:     "/",
		MaxAge:   -1, // delete
		HttpOnly: true,
		Secure:   false, // match how you set it
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Logged out"}`))
}
