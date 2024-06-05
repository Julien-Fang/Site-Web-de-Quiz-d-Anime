package quiz

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	conf "pc3r_projet/config"
	ds "pc3r_projet/dataStructure"
	auth "pc3r_projet/session"
)

// generateQuizBySynopsis, générer un quiz où le joueur n'a que le synopsis
func generateQuizBySynopsis(idJoueur string) ds.QuizCollection {
	rand.Seed(time.Now().UnixNano())
	var quizCollection []ds.Quiz
	var added []string // pour ne pas ajouter deux fois le meme anime

	nbQuestions := 0
	url := "https://kitsu.io/api/edge/anime"

	for nbQuestions < 10 {
		var quiz ds.Quiz
		switch rand.Intn(3) {
		case 0: // deviner le titre a partir du synopsis
			slug, titles, synopsis := getAnime_5_V2(url)
			var b1 bool = findTitleInSynopsis(synopsis, slug)
			var b2 bool = findTitlesInSynopsis(synopsis, titles)

			if synopsis == "" || titles == nil || slug == "" || len(synopsis) < 50 || b1 || b2 || contains(added, slug) {
				continue
			}

			quiz = ds.Quiz{
				Question:        "Devinez l'anime à l'aide du synopsis : \n" + synopsis,
				Anime:           slug,
				ReponsePossible: CreateQuizOptions(slug, IncorrectAnswers([]string{slug}, TitleAnswers)),
				BonneReponse:    slug,
			}

		case 1: // deviner un des genres
			slug, synopsis, id := getAnime_2_V2(url)
			var genresAnime []string = getGenresAnime(id)
			var b bool = findTitleInSynopsis(synopsis, slug)

			if len(genresAnime) == 0 || synopsis == "" || slug == "" || len(synopsis) < 50 || b || contains(added, slug) {
				continue
			}
			var chosen string = genresAnime[rand.Intn(len(genresAnime))]

			quiz = ds.Quiz{
				Question:        "Devinez un des genres de cet animé à l'aide du synopsis : \n" + synopsis,
				Anime:           slug,
				ReponsePossible: CreateQuizOptions(chosen, IncorrectAnswers(genresAnime, AllGenres())),
				BonneReponse:    chosen,
			}

		case 2: //deviner la saison de sortie ?
			slug, season, synopsis := getAnime_6_V2(url)
			var b bool = findTitleInSynopsis(synopsis, slug)

			if synopsis == "" || season == "" || slug == "" || len(synopsis) < 50 || b || contains(added, slug) {
				continue
			}

			quiz = ds.Quiz{
				Question:        "Devinez la saison de sortie de cet animé à l'aide du synopsis : \n" + synopsis,
				Anime:           slug,
				ReponsePossible: SeasonAnswers(),
				BonneReponse:    season,
			}
		}
		quizCollection = append(quizCollection, quiz)
		added = append(added, quiz.Anime)
		nbQuestions++
	}

	quizS := ds.QuizCollection{
		IdJoueur:        idJoueur,
		QuizCollection:  quizCollection,
		Mark:            0,
		Finished:        false,
		Number_question: 0,
	}
	fmt.Println(quizS)
	return quizS
}

func SynopsisQuizHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	claim, _ := auth.VerifyToken(w, r)
	if claim == nil {
		return
	}
	idJoueur := claim["userid"].(string)

	collection := conf.GetClient().Database("Projet_PC3R").Collection("synopsisQuiz")

	existQuiz, quiz := quizExists(collection, idJoueur)
	if !existQuiz {
		quizS := generateQuizBySynopsis(idJoueur)
		_, err := collection.InsertOne(r.Context(), quizS)
		if err != nil {
			panic(err)
		}
		quiz = quizS
	}
	jsonData, _ := json.Marshal(quiz)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
