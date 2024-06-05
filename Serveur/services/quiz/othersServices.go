package quiz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var regex = regexp.MustCompile("-")
var regex2 = regexp.MustCompile(`^\s+|\s+$`)

// getPageNumberFromURL, obtenir le nombre max de page
func getPageNumberFromURL(urlStr string) int {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	query := parsedURL.Query()
	pageStr := query.Get("page[number]")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		panic(err)
	}
	return page
}

func getNbAnimeInPage_V2(jsonDataMap map[string]interface{}) int {
	nbAnime := jsonDataMap["data"].([]interface{})
	return len(nbAnime)
}

// getLastURL, l'attribut last correspond à la dernière page
func getLastURL(url string) string {
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

	links := jsonDataMap["links"].(map[string]interface{})
	lastLink := links["last"].(string)
	return lastLink
}

// replacePageNumberInURL, remplacer le numéro de page dans l'URL
func replacePageNumberInURL(urlStr string, newPageNumber int) string {
	parsedURL, _ := url.Parse(urlStr)
	query := parsedURL.Query()
	query.Set("page[number]", strconv.Itoa(newPageNumber))
	parsedURL.RawQuery = query.Encode()
	return parsedURL.String()
}

// ---------------------------------------------------------

func contains(slice []string, anime string) bool {
	for _, a := range slice {
		if a == anime {
			return true
		}
	}
	return false
}

func getSeason(startDate time.Time) string {
	month := startDate.Month()
	switch month {
	case time.January, time.February, time.March:
		return "winter"
	case time.April, time.May, time.June:
		return "spring"
	case time.July, time.August, time.September:
		return "summer"
	case time.October, time.November, time.December:
		return "fall"
	default:
		return "unknown"
	}
}

func generateRdmGenre(genres []string) []string {
	var added []string
	rand.Seed(time.Now().UnixNano())
	rdmTour := rand.Intn(3) + 1 // de 1 à 3 genres, on ne prendra pas les 14
	var res []string
	var i int = 0
	for i < rdmTour {
		rdm := rand.Intn(len(genres))
		if !contains(added, genres[rdm]) {
			res = append(res, genres[rdm])
			added = append(added, genres[rdm])
			i++
		}
	}
	return res
}

// getAllGenres, obtenir les 10 premiers genres d'une anime a partir de son id
func getGenresAnime(id string) []string {
	url := fmt.Sprintf("https://kitsu.io/api/edge/anime/%s/genres", id)
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

	data := jsonDataMap["data"].([]interface{})
	var genres []string
	for i := 0; i < len(data); i++ {
		attributes := data[i].(map[string]interface{})
		genre := attributes["attributes"].(map[string]interface{})
		genres = append(genres, genre["name"].(string))
	}
	return genres

}

func getPictureAnime(id string) string {
	url := fmt.Sprintf("https://kitsu.io/api/edge/anime/%s", id)
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

	data := jsonDataMap["data"].(map[string]interface{})
	attributes := data["attributes"].(map[string]interface{})
	posterImage := attributes["posterImage"].(map[string]interface{})

	if posterImage["original"] != nil {
		original := posterImage["original"].(string)
		return original
	}

	if posterImage["medium"] != nil {
		medium := posterImage["medium"].(string)
		return medium
	}

	return ""
}

func getPictureAnime_v2(jsonDataMap map[string]interface{}) string {
	data := jsonDataMap["data"].(map[string]interface{})
	attributes := data["attributes"].(map[string]interface{})
	posterImage := attributes["posterImage"].(map[string]interface{})

	if posterImage["original"] != nil {
		original := posterImage["original"].(string)
		return original
	}

	if posterImage["medium"] != nil {
		medium := posterImage["medium"].(string)
		return medium
	}

	return ""
}

func getJSONData(url string) map[string]interface{} {
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

	return jsonDataMap
}

// recupere le titre
func getAnime_1_V2(url string) (string, string) {
	lastURL := getLastURL(url)
	nbPageTotal := getPageNumberFromURL(lastURL)
	rdmPage := rand.Intn(nbPageTotal)
	newURL := replacePageNumberInURL(lastURL, rdmPage)

	jsonData := getJSONData(newURL)
	nbAnimeInPage := getNbAnimeInPage_V2(jsonData)
	return getAnime_1_V2_bis(jsonData, nbAnimeInPage-1)
}

func getAnime_1_V2_bis(jsonDataMap map[string]interface{}, rdm int) (string, string) {
	data := jsonDataMap["data"].([]interface{})
	anime := data[rdm].(map[string]interface{})
	attributes := anime["attributes"].(map[string]interface{})
	tmp_slug := attributes["slug"].(string)
	id := anime["id"].(string)

	// regex := regexp.MustCompile("-")
	slug := regex.ReplaceAllString(tmp_slug, " ")
	// regex2 := regexp.MustCompile(`^\s+|\s+$`)
	slug = regex2.ReplaceAllString(slug, "")

	return slug, id
}

// recupere le titre, le synopsis
func getAnime_2_V2(url string) (string, string, string) {
	lastURL := getLastURL(url)
	nbPageTotal := getPageNumberFromURL(lastURL)
	rdmPage := rand.Intn(nbPageTotal)
	newURL := replacePageNumberInURL(lastURL, rdmPage)

	jsonData := getJSONData(newURL)
	nbAnimeInPage := getNbAnimeInPage_V2(jsonData)
	return getAnime_2_V2_bis(jsonData, nbAnimeInPage-1)

}

func getAnime_2_V2_bis(jsonDataMap map[string]interface{}, rdm int) (string, string, string) {
	data := jsonDataMap["data"].([]interface{})
	anime := data[rdm].(map[string]interface{})
	attributes := anime["attributes"].(map[string]interface{})
	synopsis := attributes["synopsis"].(string)
	tmp_slug := attributes["slug"].(string)
	id := anime["id"].(string)

	// regex := regexp.MustCompile("-")
	slug := regex.ReplaceAllString(tmp_slug, " ")
	// regex2 := regexp.MustCompile(`^\s+|\s+$`)
	slug = regex2.ReplaceAllString(slug, "")

	return slug, synopsis, id
}

// recupere les titres
func getAnime_3_V2(url string) (string, []string, string) {
	lastURL := getLastURL(url)
	nbPageTotal := getPageNumberFromURL(lastURL)
	rdmPage := rand.Intn(nbPageTotal)
	newURL := replacePageNumberInURL(lastURL, rdmPage)

	jsonData := getJSONData(newURL)
	nbAnimeInPage := getNbAnimeInPage_V2(jsonData)
	slug, titles, id := getAnime_3_V2_bis(jsonData, nbAnimeInPage-1)
	// image := getPictureAnime_v2(jsonData)
	image := getPictureAnime(id)

	return slug, titles, image

}

func getAnime_3_V2_bis(jsonDataMap map[string]interface{}, rdm int) (string, []string, string) {
	data := jsonDataMap["data"].([]interface{})
	anime := data[rdm].(map[string]interface{})
	attributes := anime["attributes"].(map[string]interface{})
	tmp_slug := attributes["slug"].(string)

	slug := regex.ReplaceAllString(tmp_slug, " ")
	slug = regex2.ReplaceAllString(slug, "")

	titles := attributes["titles"].(map[string]interface{})
	var titlesRes []string

	for i := range titles {
		if titles[i] != nil {

			title := titles[i].(string)
			title = regex.ReplaceAllString(title, " ")
			title = regex2.ReplaceAllString(title, "")
			titlesRes = append(titlesRes, title)
			// titlesRes = append(titlesRes, titles[i].(string))
		}
	}

	id := anime["id"].(string)

	return slug, titlesRes, id
}

// recupere le titre, la saison
func getAnime_4_V2(url string) (string, string) {
	lastURL := getLastURL(url)
	nbPageTotal := getPageNumberFromURL(lastURL)
	rdmPage := rand.Intn(nbPageTotal)
	newURL := replacePageNumberInURL(lastURL, rdmPage)

	jsonData := getJSONData(newURL)
	nbAnimeInPage := getNbAnimeInPage_V2(jsonData)
	return getAnime_4_V2_bis(jsonData, nbAnimeInPage-1)

}

func getAnime_4_V2_bis(jsonDataMap map[string]interface{}, rdm int) (string, string) {
	data := jsonDataMap["data"].([]interface{})
	anime := data[rdm].(map[string]interface{})
	attributes := anime["attributes"].(map[string]interface{})
	tmp_slug := attributes["slug"].(string)

	slug := regex.ReplaceAllString(tmp_slug, " ")
	slug = regex2.ReplaceAllString(slug, "")

	date, _ := attributes["startDate"].(string)
	startDate_anime, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	season := getSeason(startDate_anime)

	return slug, season
}

// recupere les titres, le synopsis
func getAnime_5_V2(url string) (string, []string, string) {
	lastURL := getLastURL(url)
	nbPageTotal := getPageNumberFromURL(lastURL)
	rdmPage := rand.Intn(nbPageTotal)
	newURL := replacePageNumberInURL(lastURL, rdmPage)

	jsonData := getJSONData(newURL)
	nbAnimeInPage := getNbAnimeInPage_V2(jsonData)
	return getAnime_5_V2_bis(jsonData, nbAnimeInPage-1)
}

func getAnime_5_V2_bis(jsonDataMap map[string]interface{}, rdm int) (string, []string, string) {
	data := jsonDataMap["data"].([]interface{})
	anime := data[rdm].(map[string]interface{})
	attributes := anime["attributes"].(map[string]interface{})
	synopsis := attributes["synopsis"].(string)
	tmp_slug := attributes["slug"].(string)

	slug := regex.ReplaceAllString(tmp_slug, " ")
	slug = regex2.ReplaceAllString(slug, "")

	titles := attributes["titles"].(map[string]interface{})
	var titlesRes []string

	for i := range titles {
		if titles[i] != nil {
			title := titles[i].(string)
			title = regex.ReplaceAllString(title, " ")
			title = regex2.ReplaceAllString(title, "")
			titlesRes = append(titlesRes, title)
			// titlesRes = append(titlesRes, titles[i].(string))
		}
	}

	return slug, titlesRes, synopsis
}

// recupere le titre, la saison, le synopsis
func getAnime_6_V2(url string) (string, string, string) {
	lastURL := getLastURL(url)
	nbPageTotal := getPageNumberFromURL(lastURL)
	rdmPage := rand.Intn(nbPageTotal)
	newURL := replacePageNumberInURL(lastURL, rdmPage)

	jsonData := getJSONData(newURL)
	nbAnimeInPage := getNbAnimeInPage_V2(jsonData)
	return getAnime_6_V2_bis(jsonData, nbAnimeInPage-1)
}

func getAnime_6_V2_bis(jsonDataMap map[string]interface{}, rdm int) (string, string, string) {
	data := jsonDataMap["data"].([]interface{})
	anime := data[rdm].(map[string]interface{})
	attributes := anime["attributes"].(map[string]interface{})
	synopsis := attributes["synopsis"].(string)
	tmp_slug := attributes["slug"].(string)

	slug := regex.ReplaceAllString(tmp_slug, " ")
	slug = regex2.ReplaceAllString(slug, "")

	date, _ := attributes["startDate"].(string)
	startDate_anime, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	season := getSeason(startDate_anime)

	return slug, season, synopsis
}

// recupere le titre et l'année de sortie
func getAnime_7_V2(url string) (string, int) {
	lastURL := getLastURL(url)
	nbPageTotal := getPageNumberFromURL(lastURL)
	rdmPage := rand.Intn(nbPageTotal)
	newURL := replacePageNumberInURL(lastURL, rdmPage)

	jsonData := getJSONData(newURL)
	nbAnimeInPage := getNbAnimeInPage_V2(jsonData)
	return getAnime_7_V2_bis(jsonData, nbAnimeInPage-1)
}

func getAnime_7_V2_bis(jsonDataMap map[string]interface{}, rdm int) (string, int) {
	data := jsonDataMap["data"].([]interface{})
	anime := data[rdm].(map[string]interface{})
	attributes := anime["attributes"].(map[string]interface{})
	tmp_slug := attributes["slug"].(string)

	slug := regex.ReplaceAllString(tmp_slug, " ")
	slug = regex2.ReplaceAllString(slug, "")

	date, _ := attributes["startDate"].(string)
	startDate_anime, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	year := startDate_anime.Year()

	return slug, year
}

// recupere le titre, synopsis, et l'année de sortie
func getAnime_8_V2(url string) (string, string, int) {
	lastURL := getLastURL(url)
	nbPageTotal := getPageNumberFromURL(lastURL)
	rdmPage := rand.Intn(nbPageTotal)
	newURL := replacePageNumberInURL(lastURL, rdmPage)

	jsonData := getJSONData(newURL)
	nbAnimeInPage := getNbAnimeInPage_V2(jsonData)
	return getAnime_8_V2_bis(jsonData, nbAnimeInPage-1)
}

func getAnime_8_V2_bis(jsonDataMap map[string]interface{}, rdm int) (string, string, int) {
	data := jsonDataMap["data"].([]interface{})
	anime := data[rdm].(map[string]interface{})
	attributes := anime["attributes"].(map[string]interface{})
	synopsis := attributes["synopsis"].(string)
	tmp_slug := attributes["slug"].(string)

	slug := regex.ReplaceAllString(tmp_slug, " ")
	slug = regex2.ReplaceAllString(slug, "")

	date, _ := attributes["startDate"].(string)
	startDate_anime, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	year := startDate_anime.Year()

	return slug, synopsis, year
}

func getAnime_9_V2(url string) (string, string, string) {
	lastURL := getLastURL(url)
	nbPageTotal := getPageNumberFromURL(lastURL)
	rdmPage := rand.Intn(nbPageTotal)
	newURL := replacePageNumberInURL(lastURL, rdmPage)

	jsonData := getJSONData(newURL)
	nbAnimeInPage := getNbAnimeInPage_V2(jsonData)
	return getAnime_9_V2_bis(jsonData, nbAnimeInPage-1)
}

func getAnime_9_V2_bis(jsonDataMap map[string]interface{}, rdm int) (string, string, string) {
	data := jsonDataMap["data"].([]interface{})
	anime := data[rdm].(map[string]interface{})
	attributes := anime["attributes"].(map[string]interface{})
	tmp_slug := attributes["slug"].(string)

	slug := regex.ReplaceAllString(tmp_slug, " ")
	slug = regex2.ReplaceAllString(slug, "")

	id := anime["id"].(string)
	image := getPictureAnime(id)
	// image := getPictureAnime_v2(jsonDataMap)
	return slug, id, image
}

func getAnime_10_V2(url string) (string, string, string) {
	lastURL := getLastURL(url)
	nbPageTotal := getPageNumberFromURL(lastURL)
	rdmPage := rand.Intn(nbPageTotal)
	newURL := replacePageNumberInURL(lastURL, rdmPage)

	jsonData := getJSONData(newURL)
	nbAnimeInPage := getNbAnimeInPage_V2(jsonData)
	return getAnime_10_V2_bis(jsonData, nbAnimeInPage-1)
}

func getAnime_10_V2_bis(jsonDataMap map[string]interface{}, rdm int) (string, string, string) {
	data := jsonDataMap["data"].([]interface{})
	anime := data[rdm].(map[string]interface{})
	attributes := anime["attributes"].(map[string]interface{})
	tmp_slug := attributes["slug"].(string)

	slug := regex.ReplaceAllString(tmp_slug, " ")
	slug = regex2.ReplaceAllString(slug, "")

	date, _ := attributes["startDate"].(string)
	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	season := getSeason(startDate)
	id := anime["id"].(string)
	image := getPictureAnime(id)
	// image := getPictureAnime_v2(jsonDataMap)

	return slug, season, image
}

func findTitleInSynopsis(synopsis string, title string) bool {
	synopsisLower := strings.ToLower(synopsis)
	titleLower := strings.ToLower(title)
	regex := regexp.MustCompile(titleLower)
	return regex.MatchString(synopsisLower)
}

func findTitlesInSynopsis(synopsis string, titles []string) bool {
	for _, title := range titles {
		if findTitleInSynopsis(synopsis, title) {
			return true
		}
	}
	return false
}
