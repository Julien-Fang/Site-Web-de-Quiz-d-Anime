package log

import (
	"encoding/json"
	"net/http"

	db "pc3r_projet/mongoDB"
	session "pc3r_projet/session"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	claim, _ := session.VerifyToken(w, r)
	userid := claim["userid"].(string)
	user, err := db.FindUserByID(userid)

	if err != nil {
		http.Error(w, "Erreur lors de la recherche de l'utilisateur", http.StatusInternalServerError)
		return

	}

	jsonData, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
