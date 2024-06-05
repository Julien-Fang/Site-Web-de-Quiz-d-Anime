package pin

import (
	"fmt"
	"io/ioutil"
	"net/http"

	conf "pc3r_projet/config"
	ds "pc3r_projet/dataStructure"
	auth "pc3r_projet/session"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// creer ou inserer un slug à la pin
func CreatePin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	claim, _ := auth.VerifyToken(w, r)
	if claim == nil {
		fmt.Println("Session terminée")
		return
	}
	idJoueur := claim["userid"].(string)

	// Lire le slug à partir du corps de la requête
	slugBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Erreur de lecture du corps de la requête")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	slug := string(slugBytes)
	fmt.Println("Slug reçu:", slug)

	// Assurez-vous de refermer le corps de la requête
	defer r.Body.Close()

	collection := conf.GetClient().Database("Projet_PC3R").Collection("Pin")

	// Recherche de la pin existante pour cet utilisateur
	filter := bson.M{"idjoueur": idJoueur}
	var existingPin ds.Pin
	err = collection.FindOne(r.Context(), filter).Decode(&existingPin)
	if err != nil && err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err == mongo.ErrNoDocuments {
		// Si la pin n'existe pas, on la crée avec le slug
		_, err = collection.UpdateOne(
			r.Context(),
			bson.M{"idjoueur": idJoueur},
			bson.M{"$addToSet": bson.M{"slugs": slug}},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		// Si la pin existe, on vérifie d'abord si le slug est déjà présent dans la liste
		// Avant d'ajouter le slug à la liste
		existingSlug := string(slug)
		exists := false
		for _, s := range existingPin.Slugs {
			if s == existingSlug {
				exists = true
				break
			}
		}
		if !exists {
			_, err = collection.UpdateOne(
				r.Context(),
				bson.M{"idjoueur": idJoueur},
				bson.M{"$addToSet": bson.M{"slugs": slug}},
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Pin créée ou mise à jour avec succès")
}
