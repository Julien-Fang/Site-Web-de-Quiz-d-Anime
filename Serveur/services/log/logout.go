package log

import (
	"fmt"
	"net/http"
	"time"

	db "pc3r_projet/mongoDB"
	session "pc3r_projet/session"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	claim, _ := session.VerifyToken(w, r)
	if claim == nil {
		return
	}
	userid := claim["userid"].(string)

	_, err := db.ToggleConnectionStatus(userid, false)
	if err != nil {
		http.Error(w, "Erreur lors de la déconnexion", http.StatusInternalServerError)
		return
	}

	ClearAllCookies(w, r)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Utilisateur %s s'est déconnecté", userid)
}

func ClearAllCookies(w http.ResponseWriter, r *http.Request) {
	// Parcourir tous les cookies de la requête et définir leur expiration à un moment antérieur
	// pour les supprimer
	for _, cookie := range r.Cookies() {
		fmt.Println(cookie.Name)
		deletedCookie := http.Cookie{
			Name:     cookie.Name,
			Value:    "",
			Expires:  time.Unix(0, 0),         // Expiration à l'époque Unix (01-01-1970 00:00:00 UTC)
			MaxAge:   -1,                      // Indique au navigateur de supprimer le cookie
			Path:     "/",                     // Assure que le cookie est supprimé dans tous les chemins
			HttpOnly: true,                    // Empêcher l'accès JavaScript au cookie
			Secure:   true,                    // Cookie envoyé uniquement via HTTPS
			SameSite: http.SameSiteStrictMode, // Politique SameSite
		}
		// Ajouter le cookie supprimé à la réponse HTTP
		http.SetCookie(w, &deletedCookie)
	}
}
