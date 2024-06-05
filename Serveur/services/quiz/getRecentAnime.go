package quiz

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// getRecentAnime, obtenir les animes récents
func getRecentAnime() []string {
	url := "https://kitsu.io/api/edge/anime?sort=-startDate"
	return getRecentAnimePage(url, 0)
}

func getRecentAnimePage(url string, cpt int) []string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	jsonData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var jsonDataMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &jsonDataMap); err != nil {
		panic(err)
	}

	today := time.Now().Truncate(24 * time.Hour)
	var recentAnime []string
	data := jsonDataMap["data"].([]interface{})
	for i := 0; i < len(data); i++ {
		anime := data[i].(map[string]interface{})
		attributes := anime["attributes"].(map[string]interface{})
		tmp_nom_anime := attributes["slug"].(string)

		regex := regexp.MustCompile("-")
		nom_anime := regex.ReplaceAllString(tmp_nom_anime, " ")
		regex2 := regexp.MustCompile(`^\s+|\s+$`)
		nom_anime = regex2.ReplaceAllString(nom_anime, "")

		date, _ := attributes["startDate"].(string)
		startDate_anime, err := time.Parse("2006-01-02", date)
		if err != nil {
			panic(err)
		}
		if startDate_anime.Equal(today) || startDate_anime.Before(today) {
			recentAnime = append(recentAnime, nom_anime)
			cpt++
		}
	}
	if cpt < 10 {
		links := jsonDataMap["links"].(map[string]interface{})
		nextLink := links["next"].(string)
		recentAnime = append(recentAnime, getRecentAnimePage(nextLink, cpt)...)
	}
	return recentAnime
}

func GetRecentAnimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	recentAnime := getRecentAnime()
	jsonData, _ := json.Marshal(recentAnime)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
