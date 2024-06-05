package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	ds "pc3r_projet/dataStructure"
	db "pc3r_projet/mongoDB"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user ds.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	if user.Pseudo == "" || user.Connexion.Login == "" || user.Connexion.Password == "" {
		fmt.Println("Création d'un utilisateur")
		http.Error(w, "Champs manquants", http.StatusBadRequest)
		return
	}

	if err := checkUserExistence(user); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = db.InsertUser(user)
	if err != nil {
		http.Error(w, "Erreur lors de l'insertion de l'utilisateur dans la base de données", http.StatusInternalServerError)
		return
	}

	var usertmp ds.User
	usertmp, _ = db.FindUserByLogin(user.Connexion.Login)

	err = db.InsertStats(usertmp.ID)
	if err != nil {
		http.Error(w, "Erreur lors de l'insertion des statistiques de l'utilisateur dans la base de données", http.StatusInternalServerError)
		return
	}

	// Répondre au client avec un code de statut HTTP 200 OK
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Utilisateur enregistré avec succès")
}

func checkUserExistence(user ds.User) error {
	if db.IsLoginTaken(user.Connexion.Login) {
		return errors.New("le login est déjà pris")
	}

	if db.IsPseudoTaken(user.Pseudo) {
		return errors.New("le pseudo est déjà pris")
	}

	return nil
}
