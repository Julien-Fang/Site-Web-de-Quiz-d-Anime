package pin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	conf "pc3r_projet/config"
	ds "pc3r_projet/dataStructure"
	auth "pc3r_projet/session"

	"go.mongodb.org/mongo-driver/bson"
)

// GET
func GetAllPin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	claim, _ := auth.VerifyToken(w, r)
	if claim == nil {
		fmt.Println("Session terminée")
		return
	}
	idJoueur := claim["userid"].(string)

	collection := conf.GetClient().Database("Projet_PC3R").Collection("Pin")

	// Recherche de la pin existante pour cet utilisateur
	filter := bson.M{"idjoueur": idJoueur}
	var existingPin ds.Pin
	err := collection.FindOne(r.Context(), filter).Decode(&existingPin)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "[]")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Si la pin n'existe pas, renvoyer un tableau vide
	if len(existingPin.Slugs) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "[]")
		return
	}

	// Sinon, renvoyer la liste des slugs
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	animeNames := make([]string, len(existingPin.Slugs))
	regex := regexp.MustCompile(`"anime"\s*:\s*"([^"]+)"`)

	for i, anime := range existingPin.Slugs {
		matches := regex.FindStringSubmatch(anime)
		if len(matches) > 1 {
			animeNames[i] = matches[1]
		}
	}
	json.NewEncoder(w).Encode(animeNames)
}
