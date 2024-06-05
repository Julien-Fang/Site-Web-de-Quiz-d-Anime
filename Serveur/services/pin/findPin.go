package pin

import (
	"encoding/json"
	"fmt"
	"net/http"

	conf "pc3r_projet/config"
	ds "pc3r_projet/dataStructure"
	auth "pc3r_projet/session"

	"go.mongodb.org/mongo-driver/bson"
)

func FindPin(w http.ResponseWriter, r *http.Request) {
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

	filter := bson.M{"idjoueur": idJoueur}
	cur, err := collection.Find(r.Context(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var pinsSlug []string
	for cur.Next(r.Context()) {
		var pin ds.Pin
		cur.Decode(&pin)
		pinsSlug = append(pinsSlug, pin.Slugs...)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(pinsSlug)
	w.Write(jsonData)
	fmt.Fprint(w, "Slugs trouvés")
}
