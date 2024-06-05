package log

import (
	"net/http"

	db "pc3r_projet/mongoDB"
)

func SessionExpiredHandler(w http.ResponseWriter, r *http.Request) {
	sessionExpired, _ := r.Cookie("sessionExpired")
	userid := sessionExpired.Value

	_, err := db.ToggleConnectionStatus(userid, false)
	if err != nil {
		http.Error(w, "Erreur lors de la mise à jour du statut de connexion", http.StatusInternalServerError)
		return
	}
	ClearAllCookies(w, r)

	// Envoyer une réponse JSON avec statut Unauthorized
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`{"error": "Session expired"}`))
}
