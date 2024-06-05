package quiz

import (
	"context"

	ds "pc3r_projet/dataStructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// quiz reprensente le type de quiz et correspond egalement a sa collection
// Retourne un quiz créer ou reprendre un quiz deja existant
// func createQuiz(w http.ResponseWriter, r *http.Request, client *mongo.Client, idJoueur string, quiz string) ds.QuizCollection {

// 	collection := client.Database("Projet_PC3R").Collection(quiz)

// 	existQ := quizExists(collection, idJoueur)
// 	if existQ.IdJoueur != "" {
// 		return existQ
// 	}

// 	var quizCollection ds.QuizCollection
// 	switch quiz {
// 	case "QuizGeneral":
// 		quizCollection = generateGeneralQuiz(idJoueur)
// 	case "QuizSynopsis":
// 		quizCollection = generateQuizBySynopsis(idJoueur)
// 	case "QuizGenre":
// 		TopGenres := []string{"Action", "Adventure", "Comedy", "Drama", "Fantasy", "Horror", "Mecha", "Mystery", "Psychological", "Romance", "Sci-Fi", "Slice of Life", "Sports", "Supernatural"}
// 		rdmGenre := rand.Intn(len(TopGenres))
// 		genre := TopGenres[rdmGenre]
// 		quizCollection = generateQuizByGenre(idJoueur, genre)
// 	case "QuizImage":
// 		quizCollection = generateQuizByPicture(idJoueur)

// 	}

// 	return quizCollection
// }

// -------------------
func quizExists(collection *mongo.Collection, idJoueur string) (bool, ds.QuizCollection) {
	filter := bson.M{"idjoueur": idJoueur, "finished": false}

	var quizC ds.QuizCollection
	err := collection.FindOne(context.Background(), filter).Decode(&quizC)
	if err != nil {
		return false, ds.QuizCollection{}
	}

	return true, quizC
}
