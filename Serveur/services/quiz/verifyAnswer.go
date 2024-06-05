package quiz

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	conf "pc3r_projet/config"
	ds "pc3r_projet/dataStructure"
	auth "pc3r_projet/session"

	"go.mongodb.org/mongo-driver/bson"
)

// verifyAnwser, vérifier la réponse du joueur
// quiz : le type de quiz
// pour qu'il soit generique a toutes les types de quiz
// renvoie un bool pour indiquer si la reponse est correcte ou non
func VerifyAnswerHandler(w http.ResponseWriter, r *http.Request) {
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

	var user ds.UserAnswer
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Erreur de décodage")
		return
	}

	collection := conf.GetClient().Database("Projet_PC3R").Collection(user.Quiz)

	filter := bson.M{"idjoueur": idJoueur, "finished": false}
	var quizCollection ds.QuizCollection
	collection.FindOne(context.Background(), filter).Decode(&quizCollection)

	quizCollection.QuizCollection[quizCollection.Number_question].RepJoueur = user.Answer
	ok := checkAnswer(user.Answer, quizCollection.QuizCollection[quizCollection.Number_question].BonneReponse)

	if ok {
		quizCollection.Mark++
		quizCollection.QuizCollection[quizCollection.Number_question].Valide = true
	} else {
		quizCollection.QuizCollection[quizCollection.Number_question].Valide = false
	}

	// on incrémente pour passer à la question suivante
	quizCollection.Number_question++
	fmt.Println("Nombre de question :", quizCollection.Number_question)
	if quizCollection.Number_question == 10 {
		addStats(user.Quiz, idJoueur, quizCollection.Mark)
		quizCollection.Finished = true
		fmt.Println("Fin du quiz")
	}

	// Mise à jour du document dans la base de données
	collection.ReplaceOne(context.Background(), filter, quizCollection)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(quizCollection)
	w.Write(jsonData)
}

func checkAnswer(userAnswer string, correctAnswers string) bool {
	return strings.EqualFold(userAnswer, correctAnswers)
}

func addStats(QuizType string, userid string, mark int) {
	collection := conf.GetClient().Database("Projet_PC3R").Collection("statsUser")
	filter := bson.M{"idjoueur": userid}
	var stats ds.StatsUser
	collection.FindOne(context.Background(), filter).Decode(&stats)
	fmt.Println("idjoueur :", userid)
	fmt.Println("type de quiz :", QuizType)
	switch QuizType {
	case "generalQuiz":
		stats.NbQuizGeneral++
		stats.NoteTotalGeneral += mark
	case "genreQuiz":
		stats.NbQuizGenre++
		stats.NoteTotalGenre += mark
	case "pictureQuiz":
		stats.NbQuizPicture++
		stats.NoteTotalPicture += mark
	case "synopsisQuiz":
		stats.NbQuizSynopsis++
		stats.NoteTotalSynopsis += mark
	}

	update := bson.M{"$set": stats}
	collection.UpdateOne(context.Background(), filter, update)
}
