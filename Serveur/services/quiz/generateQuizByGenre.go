package quiz

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	conf "pc3r_projet/config"
	ds "pc3r_projet/dataStructure"
	auth "pc3r_projet/session"
)

// generateQuizByGenre, générer un quiz où l'utilisateur doit deviner le genre de l'anime
func generateQuizByGenre(idJoueur string, genre string) ds.QuizCollection {
	rand.Seed(time.Now().UnixNano())
	var quizCollection []ds.Quiz
	var added []string // pour ne pas ajouter deux fois le meme anime

	nbQuestions := 0
	url := "https://kitsu.io/api/edge/anime?filter[genres]=" + genre

	//Il faudrait indiquer aux joueurs les genres possibles
	for nbQuestions < 10 {
		var quiz ds.Quiz
		switch rand.Intn(3) {
		case 0:
			slug, id := getAnime_1_V2(url)
			var genresAnime []string = getGenresAnime(id)

			if len(genresAnime) == 0 || slug == "" || contains(added, slug) {
				continue
			}

			var chosen string = genresAnime[rand.Intn(len(genresAnime))]
			quiz = ds.Quiz{
				Question:        "Devinez un des genres de cet animé à l'aide du titre : \n" + slug,
				Anime:           slug,
				ReponsePossible: CreateQuizOptions(chosen, IncorrectAnswers(genresAnime, AllGenres())),
				BonneReponse:    chosen,
			}

		case 1:
			slug, synopsis, id := getAnime_2_V2(url)
			var genresAnime []string = getGenresAnime(id)
			var b bool = findTitleInSynopsis(synopsis, slug)

			if len(genresAnime) == 0 || synopsis == "" || slug == "" || b || contains(added, slug) {
				continue
			}

			var chosen string = genresAnime[rand.Intn(len(genresAnime))]
			quiz = ds.Quiz{
				Question:        "Devinez un des genres de cet animé à l'aide du synopsis : \n" + synopsis,
				Anime:           slug,
				ReponsePossible: CreateQuizOptions(chosen, IncorrectAnswers(genresAnime, AllGenres())),
				BonneReponse:    chosen,
			}

		case 2:
			slug, id := getAnime_1_V2(url)
			var genresAnime []string = getGenresAnime(id)
			image := getPictureAnime(id)

			if len(genresAnime) == 0 || image == "" || slug == "" || contains(added, slug) {
				continue
			}

			var chosen string = genresAnime[rand.Intn(len(genresAnime))]
			quiz = ds.Quiz{
				Question:        "Devinez un des genres de cet animé à l'aide de l'image :\n",
				Anime:           slug,
				ReponsePossible: CreateQuizOptions(chosen, IncorrectAnswers(genresAnime, AllGenres())),
				Image:           image,
				BonneReponse:    chosen,
			}
		}

		quizCollection = append(quizCollection, quiz)
		added = append(added, quiz.Anime)
		nbQuestions++
	}

	quizGenre := ds.QuizCollection{
		IdJoueur:        idJoueur,
		QuizCollection:  quizCollection,
		Mark:            0,
		Finished:        false,
		Number_question: 0,
	}

	return quizGenre
}

func PostGenreQuizHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	claim, _ := auth.VerifyToken(w, r)
	if claim == nil {
		return
	}
	idJoueur := claim["userid"].(string)

	TopGenres := AllGenres()
	rdmGenre := rand.Intn(len(TopGenres))
	genre := TopGenres[rdmGenre]
	collection := conf.GetClient().Database("Projet_PC3R").Collection("genreQuiz")

	existQuiz, quiz := quizExists(collection, idJoueur)
	if !existQuiz {
		quizGenre := generateQuizByGenre(idJoueur, genre)
		_, err := collection.InsertOne(context.Background(), quizGenre)
		if err != nil {
			panic(err)
		}
		quiz = quizGenre
	}
	jsonData, _ := json.Marshal(quiz)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
