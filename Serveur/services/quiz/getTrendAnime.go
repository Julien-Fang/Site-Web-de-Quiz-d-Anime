package quiz

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
)

// getTrendAnime, obtenir les animes tendances
func getTrendAnime() []string {
	url := "https://kitsu.io/api/edge/trending/anime"
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

	var trendAnime []string
	data := jsonDataMap["data"].([]interface{})
	for i := 0; i < len(data); i++ {
		anime := data[i].(map[string]interface{})
		attributes := anime["attributes"].(map[string]interface{})
		tmp_nom_anime := attributes["slug"].(string)

		regex := regexp.MustCompile("-")
		nom_anime := regex.ReplaceAllString(tmp_nom_anime, " ")
		regex2 := regexp.MustCompile(`^\s+|\s+$`)
		nom_anime = regex2.ReplaceAllString(nom_anime, "")

		trendAnime = append(trendAnime, nom_anime)
	}
	return trendAnime
}

func GetTrendAnimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	trendAnime := getTrendAnime()
	jsonData, _ := json.Marshal(trendAnime)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
