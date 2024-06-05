package quiz

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	conf "pc3r_projet/config"
	ds "pc3r_projet/dataStructure"
	auth "pc3r_projet/session"
)

func generateQuizByPicture(idJoueur string) ds.QuizCollection {
	rand.Seed(time.Now().UnixNano())
	var quizCollection []ds.Quiz
	var added []string // pour ne pas ajouter deux fois le meme anime

	url := "https://kitsu.io/api/edge/anime"
	nbQuestions := 0

	for nbQuestions < 10 {
		var quiz ds.Quiz
		switch rand.Intn(3) {
		case 0: // deviner titre animé grâce à une image
			slug, titles, image := getAnime_3_V2(url)

			if titles == nil || image == "" || slug == "" || contains(added, slug) {
				continue
			}

			quiz = ds.Quiz{
				Question:        "Devinez l'anime à l'aide de l'image :\n",
				Anime:           slug,
				ReponsePossible: CreateQuizOptions(slug, IncorrectAnswers([]string{slug}, TitleAnswers)),
				BonneReponse:    slug,
				Image:           image,
			}

		case 1: // deviner le genre de l'anime grâce à une image
			// getAnime_9_V2 ne trouve pas le genre, à rectifier
			slug, id, image := getAnime_9_V2(url)
			genresAnime := getGenresAnime(id)

			if len(genresAnime) == 0 || image == "" || slug == "" || contains(added, slug) {
				continue
			}

			fmt.Println(genresAnime)
			var chosen string = genresAnime[rand.Intn(len(genresAnime))]

			quiz = ds.Quiz{
				Question:        "Devinez un des genres de cet animé à l'aide de l'image :\n",
				Anime:           slug,
				ReponsePossible: CreateQuizOptions(chosen, IncorrectAnswers(genresAnime, AllGenres())),
				BonneReponse:    chosen,
				Image:           image,
			}

		case 2: // deviner la saison de sortie de l'anime grâce à une image
			slug, saison, image := getAnime_10_V2(url)

			if saison == "" || image == "" || contains(added, slug) {
				continue
			}

			quiz = ds.Quiz{
				Question:        "Devinez la saison de sortie de cet animé à l'aide de l'image :\n",
				Anime:           slug,
				ReponsePossible: SeasonAnswers(),
				BonneReponse:    saison,
				Image:           image,
			}
		}
		quizCollection = append(quizCollection, quiz)
		added = append(added, quiz.Anime)
		nbQuestions++
	}

	quizP := ds.QuizCollection{
		IdJoueur:        idJoueur,
		QuizCollection:  quizCollection,
		Mark:            0,
		Finished:        false,
		Number_question: 0,
	}

	return quizP

}

func PictureQuizHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	claim, _ := auth.VerifyToken(w, r)
	if claim == nil {
		return
	}
	idJoueur := claim["userid"].(string)

	collection := conf.GetClient().Database("Projet_PC3R").Collection("pictureQuiz")

	existQuiz, quiz := quizExists(collection, idJoueur)
	if !existQuiz {
		quizP := generateQuizByPicture(idJoueur)
		_, err := collection.InsertOne(context.Background(), quizP)
		if err != nil {
			panic(err)
		}
		quiz = quizP
	}
	jsonData, _ := json.Marshal(quiz)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
