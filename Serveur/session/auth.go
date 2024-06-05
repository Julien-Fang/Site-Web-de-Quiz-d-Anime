package session

import (
	"errors"
	"net/http"
	conf "pc3r_projet/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("lIOA1H46hHIlGaww88xLOftWAqbVgLxEQC7qeNkPBQo") // Clé secrète pour signer les tokens JWT
var accessDuration = conf.GetAccessDuration()

// Fonction pour générer un AccessToken et RefreshToken JWT
func GenerateToken(userid string, expiresIn time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"userid": userid,
		"exp":    time.Now().Add(expiresIn).Unix(), // Expiration en 30 jours
	}
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(w http.ResponseWriter, r *http.Request) (jwt.MapClaims, error) {
	// Extraire les cookies de la requête HTTP
	accessTokenCookie, err := r.Cookie("accessToken")
	if err != nil {
		refreshTokenCookie, err := r.Cookie("refreshToken")
		if err != nil {
			// Si refreshToken n'est pas présent, rediriger vers la page sessionExpired
			http.Redirect(w, r, "/user/sessionExpired", http.StatusSeeOther)
			return nil, errors.New("refreshToken cookie not found")
		}

		// Si refreshToken est présent, décodez-le pour obtenir les informations utilisateur
		claims, _ := DecodeToken(refreshTokenCookie.Value)

		// Si le refreshToken est valide, générez un nouvel accessToken
		newAccessToken, _ := GenerateToken(claims["userid"].(string), accessDuration)

		accessTokenCookie := http.Cookie{
			Name:     "accessToken",                                // Nom du cookie
			Value:    newAccessToken,                               // Valeur du cookie (votre accessToken)
			Expires:  time.Now().Add(time.Second * accessDuration), // Expiration du cookie (1 heure dans cet exemple)
			HttpOnly: true,                                         // Empêcher l'accès JavaScript au cookie
			Path:     "/",                                          // Chemin du cookie (toutes les routes dans cet exemple)
			Secure:   true,                                         // Cookie envoyé uniquement via HTTPS
			SameSite: http.SameSiteStrictMode,                      // Politique SameSite
		}

		http.SetCookie(w, &accessTokenCookie)
		return claims, nil
	}
	return DecodeToken(accessTokenCookie.Value)
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	// Parse le token avec les revendications
	token, _ := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil // jwtSecret est la clé secrète utilisée pour signer le token
	})

	// Vérifier si le token est valide
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Récupérer les revendications (claims) du token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
