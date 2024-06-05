package log

import (
	"encoding/json"
	"net/http"
	"time"

	conf "pc3r_projet/config"
	ds "pc3r_projet/dataStructure"
	db "pc3r_projet/mongoDB"
	session "pc3r_projet/session"

	"golang.org/x/crypto/bcrypt"
)

// Créer une structure pour la réponse JSON
type ResponseData struct {
	UserID  string `json:"userid"`
	Message string `json:"message"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user ds.UserLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Login == "" || user.Password == "" {
		http.Error(w, "Champs manquants", http.StatusBadRequest)
		return
	}

	result, err := db.FindUserByLogin(user.Login)
	if err != nil {
		http.Error(w, "Identifiant erroné", http.StatusUnauthorized)
		return
	}

	if !verifyPassword(result.Connexion.Password, user.Password) {
		http.Error(w, "Mot de passe erroné", http.StatusUnauthorized)
		return
	}

	if result.Connexion.Connected {
		http.Error(w, "Utilisateur déjà connecté", http.StatusForbidden)
		return
	}
	cookies := createCookies(result.ID, conf.GetAccessDuration(), conf.GetRefreshDuration(), conf.GetSessionExpiredDuration())

	for _, cookie := range cookies {
		http.SetCookie(w, cookie)
	}

	_, err = db.ToggleConnectionStatus(result.ID, true)
	if err != nil {
		http.Error(w, "Erreur lors de la mise à jour du statut de connexion", http.StatusInternalServerError)
		return
	}

	// Créer une instance de la structure ResponseData avec l'UserID
	responseData := ResponseData{
		UserID:  result.ID,
		Message: "Utilisateur " + result.Connexion.Login + " connecté avec succès"}

	jsonData, _ := json.Marshal(responseData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func verifyPassword(hashedPassword string, password string) bool {
	hashedPasswordBytes := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, []byte(password))
	return err == nil
}

func createCookies(userID string, accessDuration, refreshDuration, sessionExpiredDuration time.Duration) []*http.Cookie {
	accessToken, _ := session.GenerateToken(userID, accessDuration)
	refreshToken, _ := session.GenerateToken(userID, refreshDuration)

	accessTokenCookie := http.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		Expires:  time.Now().Add(accessDuration),
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	refreshTokenCookie := http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Expires:  time.Now().Add(refreshDuration),
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	sessionExpiredCookie := http.Cookie{
		Name:     "sessionExpired",
		Value:    userID,
		Expires:  time.Now().Add(sessionExpiredDuration),
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	return []*http.Cookie{&accessTokenCookie, &refreshTokenCookie, &sessionExpiredCookie}
}
