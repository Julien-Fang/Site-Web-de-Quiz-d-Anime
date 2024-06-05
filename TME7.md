

- services/ressources:
    - Pour l'utilisateur : http://localhost:8000/user/...
        - login : pour se connecter
        - signup : pour se créer un compte
        - logout : pour se deconnecter
        -  : sessionExpired

- Methode GET :
    - getTrendAnim() []string : Obtenir les animes tendances 
    - getRecentAnime() []string : Obtenir les animés récents à l'aide d'une fonction récursive getRecentAnimePage pour parcourir les pages jusqu'à récuperer les animés les plus récents à partir d'aujourd'hui
    - getRecentAnimePage(url string, cpt int) []string : Parcours les pages des animés à l'aide de url et récupère les 10 premiers avec cpt
    - getAnimeTitle(animeID string) (string) : Obtenir le titre d'un animé grâce à l'id
    - getAnimePerSeasonWithAverageRating(season string) []string : Obtenir les 10 premiers animés en fonction d'une saison et de leur note moyenne
    - getAnimeByTitle(title string) string : Obtenir l'id de l'animé par son titre
    - getAnimeBySeason(season string) []string : Obtenir les 10 premiers animés d'une saison (classé par popularityRank, pas sûr)
    - getAnimeByGenre(genre []string) []string : Obtenir les 10 premiers animés d'une/des genre(s) (classé par popularityRank, pas sûr)
    - getAnimeByAverageRating() []string : Obtenir les 10 premiers animés ayant les meilleures notes
    - generateGeneralQuiz(idJoueur string) QuizCollection
