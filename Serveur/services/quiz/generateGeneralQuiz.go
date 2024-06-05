package quiz

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	conf "pc3r_projet/config"
	ds "pc3r_projet/dataStructure"
	auth "pc3r_projet/session"
)

func generateGeneralQuiz(idJoueur string) ds.QuizCollection {
	rand.Seed(time.Now().UnixNano())
	var quizCollection []ds.Quiz
	var added []string // pour ne pas ajouter deux fois le meme anime

	nbQuestions := 0
	for nbQuestions < 10 {
		var quiz ds.Quiz
		switch rand.Intn(4) {
		case 0: // deviner titre animé à partir du synposis
			url := "https://kitsu.io/api/edge/anime"
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

		case 1: // deviner la saison de sortie (winter, spring, summer, fall)
			switch rand.Intn(2) {

			case 0: // deviner la saison de sortie à partir du titre
				url := "https://kitsu.io/api/edge/anime"
				slug, season := getAnime_4_V2(url)

				if season == "" || slug == "" || contains(added, slug) {
					continue
				}

				quiz = ds.Quiz{
					Question:        "Devinez la saison de la toute première sortie de l'anime à l'aide du titre : \n" + slug,
					Anime:           slug,
					ReponsePossible: SeasonAnswers(),
					BonneReponse:    season,
				}

			case 1: // deviner la saison de sortie à partir du synopsis
				url := "https://kitsu.io/api/edge/anime"
				slug, season, synopsis := getAnime_6_V2(url)
				var b bool = findTitleInSynopsis(synopsis, slug)

				if synopsis == "" || season == "" || slug == "" || len(synopsis) < 50 || b || contains(added, slug) {
					continue
				}

				quiz = ds.Quiz{
					Question:        "Devinez la saison de la toute première sortie de l'anime à l'aide du syopsis : \n" + synopsis,
					Anime:           slug,
					ReponsePossible: SeasonAnswers(),
					BonneReponse:    season,
				}

			}

		case 2: // deviner genre d'animé
			TopGenres := AllGenres()
			genres := generateRdmGenre(TopGenres)
			var tmpgenre string
			for i := 0; i < len(genres); i++ {
				tmpgenre += genres[i]
				if i < len(genres)-1 {
					tmpgenre += ", "
				}
			}
			url := "https://kitsu.io/api/edge/anime?filter[genres]=" + tmpgenre
			switch rand.Intn(2) {
			case 0: // deviner genre d'animé a partir du titre
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

			case 1: // deviner genre d'animé par synopsis
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
			}

		//case 3: // deviner son rating

		case 3: // deviner l'annee de sortie
			switch rand.Intn(2) {
			case 0: // deviner l'année de sortie a partir du titre
				url := "https://kitsu.io/api/edge/anime"
				slug, year := getAnime_7_V2(url)

				if year == 0 || slug == "" || contains(added, slug) {
					continue
				}

				quiz = ds.Quiz{
					Question:        "Devinez l'année de sortie de l'anime à l'aide du titre : \n" + slug,
					Anime:           slug,
					ReponsePossible: CreateQuizOptions(strconv.Itoa(year), IncorrectAnswers([]string{strconv.Itoa(year)}, YearAnswers)),
					BonneReponse:    strconv.Itoa(year),
				}

			case 1: // deviner l'année de sortie a partir du synopsis
				url := "https://kitsu.io/api/edge/anime"
				slug, synopsis, year := getAnime_8_V2(url)
				var b bool = findTitleInSynopsis(synopsis, slug)
				if year == 0 || synopsis == "" || slug == "" || len(synopsis) < 50 || b || contains(added, slug) {
					continue
				}

				quiz = ds.Quiz{
					Question:        "Devinez l'année de sortie de l'anime à l'aide du synopsis : \n" + synopsis,
					Anime:           slug,
					ReponsePossible: CreateQuizOptions(strconv.Itoa(year), IncorrectAnswers([]string{strconv.Itoa(year)}, YearAnswers)),
					BonneReponse:    strconv.Itoa(year),
				}
			}
		}
		quizCollection = append(quizCollection, quiz)
		added = append(added, quiz.Anime)
		nbQuestions++
	}
	quizG := ds.QuizCollection{
		IdJoueur:        idJoueur,
		QuizCollection:  quizCollection,
		Mark:            0,
		Finished:        false,
		Number_question: 0,
	}
	return quizG

}

func PostGeneralQuizHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	claim, _ := auth.VerifyToken(w, r)
	if claim == nil {
		return
	}
	idJoueur := claim["userid"].(string)

	collection := conf.GetClient().Database("Projet_PC3R").Collection("generalQuiz")

	existQuiz, quiz := quizExists(collection, idJoueur)
	if !existQuiz {
		quizC := generateGeneralQuiz(idJoueur)
		_, err := collection.InsertOne(context.Background(), quizC)
		if err != nil {
			panic(err)
		}
		quiz = quizC
	}
	jsonData, _ := json.Marshal(quiz)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
