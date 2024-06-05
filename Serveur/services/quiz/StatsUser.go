package quiz

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	conf "pc3r_projet/config"
	ds "pc3r_projet/dataStructure"
	db "pc3r_projet/mongoDB"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// incrementer le nombre de quiz joué par un utilisateur
// func IncreaseNbQuiz(w http.ResponseWriter, r *http.Request, client *mongo.Client, idJoueur string, quiz string) ds.StatsUser {
// 	collection := client.Database("Projet_PC3R").Collection("StatsUser")

// 	filter := bson.M{"idjoueur": idJoueur}

// 	var stats ds.StatsUser
// 	// err := collection.FindOne(context.Background(), filter).Decode(&stats)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	collection.FindOne(context.Background(), filter).Decode(&stats)

// 	// JE DOIS INCREMENTER ICI ET UPDATEONE
// 	switch quiz {
// 	case "QuizGeneral":
// 		stats.NbQuizGeneral++
// 	case "QuizSynopsis":
// 		stats.NbQuizSynopsis++
// 	case "QuizGenre":
// 		stats.NbQuizGenre++
// 	case "QuizImage":
// 		stats.NbQuizPicture++
// 	}

// 	_, err := collection.UpdateOne(context.Background(), filter, stats)
// 	if err != nil {
// 		panic(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	jsonData, _ := json.Marshal(stats)
// 	w.Write(jsonData)

// 	return stats
// }

// recuperer les statistiques d'un utilisateur
func GetStats(w http.ResponseWriter, r *http.Request, client *mongo.Client, idJoueur string) ds.StatsUser {
	collection := client.Database("Projet_PC3R").Collection("statsUser")

	filter := bson.M{"idjoueur": idJoueur}

	var stats ds.StatsUser
	collection.FindOne(context.Background(), filter).Decode(&stats)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(stats)
	w.Write(jsonData)

	return stats
}

func GetTop5QuizGeneral(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	collection := conf.GetClient().Database("Projet_PC3R").Collection("statsUser")

	// Récupérer tous les documents StatsUser
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())

	var allStats []ds.StatsUser
	if err := cursor.All(context.Background(), &allStats); err != nil {
		panic(err)
	}

	// Trier les données manuellement
	sort.Slice(allStats, func(i, j int) bool {
		// Comparaison selon la moyenne générale (NoteTotalGeneral / NbQuizGeneral)
		I := float64(allStats[i].NoteTotalGeneral)
		J := float64(allStats[j].NoteTotalGeneral)
		return I > J // Tri décroissant
	})

	// Limiter les résultats à 5
	var top5General []ds.StatsUser
	if len(allStats) >= 5 {
		top5General = allStats[:5]
	} else {
		top5General = allStats
	}

	// Convertir en format compatible JSON et remplacer l'ID par le pseudo
	var top5QuizGeneral []map[string]interface{}

	for _, user := range top5General {
		var score = user.NoteTotalGeneral
		var pseudoUser ds.User
		pseudoUser, _ = db.FindUserByID(user.IdJoueur)
		pUser := pseudoUser.Pseudo

		top5QuizGeneral = append(top5QuizGeneral, map[string]interface{}{
			"idjoueur":      pUser, //user.IdJoueur,
			"nbquizgeneral": user.NbQuizGeneral,
			"score":         score,
		})
	}

	// Répondre avec les données JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(top5QuizGeneral)
	w.Write(jsonData)

}

func GetTop5QuizGenre(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetTop5QuizGenre")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	collection := conf.GetClient().Database("Projet_PC3R").Collection("statsUser")

	// Récupérer tous les documents StatsUser
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())

	var allStats []ds.StatsUser
	if err := cursor.All(context.Background(), &allStats); err != nil {
		panic(err)
	}

	// Trier les données manuellement
	sort.Slice(allStats, func(i, j int) bool {
		// Comparaison selon la moyenne générale (NoteTotalGeneral / NbQuizGeneral)
		I := float64(allStats[i].NoteTotalGenre)
		J := float64(allStats[j].NoteTotalGenre)
		return I > J // Tri décroissant
	})

	// Limiter les résultats à 5
	var top5Genre []ds.StatsUser
	if len(allStats) >= 5 {
		top5Genre = allStats[:5]
	} else {
		top5Genre = allStats
	}

	// Convertir en format compatible JSON et remplacer l'ID par le pseudo
	var top5QuizGenre []map[string]interface{}

	for _, user := range top5Genre {
		var score = float64(user.NoteTotalGenre)
		var pseudoUser ds.User
		pseudoUser, _ = db.FindUserByID(user.IdJoueur)
		pUser := pseudoUser.Pseudo

		top5QuizGenre = append(top5QuizGenre, map[string]interface{}{
			"idjoueur":      pUser, //user.IdJoueur,
			"nbquizgeneral": user.NbQuizGenre,
			"score":         score,
		})
	}

	// Répondre avec les données JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(top5QuizGenre)
	w.Write(jsonData)
}

func GetTop5QuizPicture(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	collection := conf.GetClient().Database("Projet_PC3R").Collection("statsUser")

	// Récupérer tous les documents StatsUser
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())

	var allStats []ds.StatsUser
	if err := cursor.All(context.Background(), &allStats); err != nil {
		panic(err)
	}

	// Trier les données manuellement
	sort.Slice(allStats, func(i, j int) bool {
		// Comparaison selon la moyenne générale (NoteTotalGeneral / NbQuizGeneral)
		I := float64(allStats[i].NoteTotalPicture)
		J := float64(allStats[j].NoteTotalPicture)
		return I > J // Tri décroissant
	})

	// Limiter les résultats à 5
	var top5Picture []ds.StatsUser
	if len(allStats) >= 5 {
		top5Picture = allStats[:5]
	} else {
		top5Picture = allStats
	}

	// Convertir en format compatible JSON et remplacer l'ID par le pseudo
	var top5QuizPicture []map[string]interface{}

	for _, user := range top5Picture {
		var score = user.NoteTotalPicture
		var pseudoUser ds.User
		pseudoUser, _ = db.FindUserByID(user.IdJoueur)
		pUser := pseudoUser.Pseudo

		top5QuizPicture = append(top5QuizPicture, map[string]interface{}{
			"idjoueur":      pUser, //user.IdJoueur,
			"nbquizgeneral": user.NbQuizPicture,
			"score":         score,
		})
	}

	// Répondre avec les données JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(top5QuizPicture)
	w.Write(jsonData)

}

func GetTop5QuizSynopsis(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	collection := conf.GetClient().Database("Projet_PC3R").Collection("statsUser")

	// Récupérer tous les documents StatsUser
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())

	var allStats []ds.StatsUser
	if err := cursor.All(context.Background(), &allStats); err != nil {
		panic(err)
	}

	// Trier les données manuellement
	sort.Slice(allStats, func(i, j int) bool {
		// Comparaison selon la moyenne générale (NoteTotalGeneral / NbQuizGeneral)
		I := float64(allStats[i].NoteTotalSynopsis)
		J := float64(allStats[j].NoteTotalSynopsis)
		return I > J // Tri décroissant
	})

	// Limiter les résultats à 5
	var top5Synopsis []ds.StatsUser
	if len(allStats) >= 5 {
		top5Synopsis = allStats[:5]
	} else {
		top5Synopsis = allStats
	}

	// Convertir en format compatible JSON et remplacer l'ID par le pseudo
	var top5QuizSynopsis []map[string]interface{}

	for _, user := range top5Synopsis {
		var score = user.NoteTotalSynopsis
		var pseudoUser ds.User
		pseudoUser, _ = db.FindUserByID(user.IdJoueur)
		pUser := pseudoUser.Pseudo

		top5QuizSynopsis = append(top5QuizSynopsis, map[string]interface{}{
			"idjoueur":      pUser, //user.IdJoueur,
			"nbquizgeneral": user.NbQuizSynopsis,
			"score":         score,
		})
	}

	// Répondre avec les données JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(top5QuizSynopsis)
	w.Write(jsonData)

}
