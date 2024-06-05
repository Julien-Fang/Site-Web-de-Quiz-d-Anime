package quiz

import (
	"encoding/json"
	"fmt"
	"net/http"

	conf "pc3r_projet/config"
	auth "pc3r_projet/session"
)

// POST, renvoie true si aucuns quiz n'est en cours
func UnfinishedQuiz(w http.ResponseWriter, r *http.Request) {
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

	collectionGeneral := conf.GetClient().Database("Projet_PC3R").Collection("generalQuiz")
	collectionGenre := conf.GetClient().Database("Projet_PC3R").Collection("genreQuiz")
	collectionPicture := conf.GetClient().Database("Projet_PC3R").Collection("pictureQuiz")
	collectionSynopsis := conf.GetClient().Database("Projet_PC3R").Collection("synopsisQuiz")

	existQuizGeneral, _ := quizExists(collectionGeneral, idJoueur)
	existQuizGenre, _ := quizExists(collectionGenre, idJoueur)
	existQuizPicture, _ := quizExists(collectionPicture, idJoueur)
	existQuizSynopsis, _ := quizExists(collectionSynopsis, idJoueur)

	var unfinishedQuizzes []string
	if existQuizGeneral {
		unfinishedQuizzes = append(unfinishedQuizzes, "generalQuiz")
	}
	if existQuizGenre {
		unfinishedQuizzes = append(unfinishedQuizzes, "genreQuiz")
	}
	if existQuizPicture {
		unfinishedQuizzes = append(unfinishedQuizzes, "pictureQuiz")
	}
	if existQuizSynopsis {
		unfinishedQuizzes = append(unfinishedQuizzes, "synopsisQuiz")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(unfinishedQuizzes)
	fmt.Println(unfinishedQuizzes)
}
