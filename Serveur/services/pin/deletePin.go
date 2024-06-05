package pin

import (
	"fmt"
	"io"
	"net/http"

	conf "pc3r_projet/config"
	auth "pc3r_projet/session"

	"go.mongodb.org/mongo-driver/bson"
)

func DeletePin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	claim, _ := auth.VerifyToken(w, r)
	if claim == nil {
		fmt.Println("Session terminée")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	idJoueur := claim["userid"].(string)

	// Lire le slug à partir du corps de la requête
	slugBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Erreur de lecture du corps de la requête")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Assurez-vous de refermer le corps de la requête
	defer r.Body.Close()

	slug := string(slugBytes)
	fmt.Println("Slug suppr:", slug)

	collection := conf.GetClient().Database("Projet_PC3R").Collection("Pin")

	filter := bson.M{"idjoueur": idJoueur, "slugs": string(slug)} // Filtrer le document par idJoueur et slug
	update := bson.M{"$pull": bson.M{"slugs": string(slug)}}      // Supprimer le slug de la liste des slugs

	// Mettre à jour le document en retirant le slug de la liste des slugs
	_, err = collection.UpdateOne(r.Context(), filter, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Erreur lors de la suppression du slug")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Slug supprimé avec succès")
}
