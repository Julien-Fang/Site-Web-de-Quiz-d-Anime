package quiz

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateQuizOptions(correctAnswer string, allAnswers []string) []string {
	rand.Seed(time.Now().UnixNano())

	answers := make([]string, 0)
	answers = append(answers, correctAnswer)

	for len(answers) < 4 {
		rdmAnswer := allAnswers[rand.Intn(len(allAnswers))]

		if rdmAnswer != correctAnswer {
			answers = append(answers, rdmAnswer)
		}
	}

	rand.Shuffle(len(answers), func(i, j int) {
		answers[i], answers[j] = answers[j], answers[i]
	})

	return answers
}

func IncorrectAnswers(userAnswer, answers []string) []string {
	// Créer une map pour stocker les genres de genresAnime
	animeGenreMap := make(map[string]bool)
	for _, genre := range userAnswer {
		animeGenreMap[genre] = true
	}

	incorrectAnswers := make([]string, 0)
	for _, genre := range answers {
		// Tous les genres qui ne sont pas présent dans animeGenreMap sont des mauvaises réponses
		if !animeGenreMap[genre] {
			incorrectAnswers = append(incorrectAnswers, genre)
		}
	}

	return incorrectAnswers
}

func AllGenres() []string {
	return []string{
		"Action",
		"Adventure",
		"Comedy",
		"Drama",
		"Fantasy",
		"Horror",
		"Mecha",
		"Mystery",
		"Psychological",
		"Romance",
		"Sci-Fi",
		"Slice of Life",
		"Sports",
		"Supernatural",
		"Thriller",
		"Crime",
		"Historical",
		"Music",
		"School",
		"Seinen",
		"Shoujo",
		"Shounen",
		"Super Power",
		"Magic",
		"Military",
		"Police",
		"Space",
		"Vampire",
		"Demons",
		"Ecchi",
		"Game",
		"Harem",
		"Josei",
		"Parody",
		"Samurai",
		"Superhero",
	}
}

var TitleAnswers = generateTitleAnswers()

func generateTitleAnswers() []string {
	url := "https://kitsu.io/api/edge/anime"
	titres := make([]string, 0)
	for i := 2; i < 7; i++ {
		fmt.Println(url)
		jsonData := getJSONData(url)
		nbAnimeInPage := getNbAnimeInPage_V2(jsonData)
		for j := 0; j < nbAnimeInPage; j++ {
			slug, _ := getAnime_1_V2_bis(jsonData, j)
			titres = append(titres, slug)
		}
		lastURL := getLastURL(url)
		url = replacePageNumberInURL(lastURL, i)
	}
	return titres
}

func SeasonAnswers() []string {
	return []string{
		"Winter",
		"Spring",
		"Summer",
		"Fall",
	}
}

var YearAnswers = generateYearAnswers()

func generateYearAnswers() []string {
	annees := make([]string, 0)
	for i := 1980; i < 2025; i++ {
		annees = append(annees, fmt.Sprintf("%d", i))
	}
	return annees
}
