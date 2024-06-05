package dataStructure

type UserLogin struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	Connected bool   `json:"connected"`
}

type User struct {
	ID        string    `bson:"_id,omitempty"`
	Connexion UserLogin `json:"user"`
	Pseudo    string    `json:"pseudo"`
}

type Quiz struct {
	Question        string   `json:"question"`
	Anime           string   `json:"anime"`
	ReponsePossible []string `json:"reponsepossible"`
	Image           string   `json:"image"`
	RepJoueur       string   `json:"repjoueur"`
	Valide          bool     `json:"valide"`
	BonneReponse    string   `json:"bonnereponse"`
}

type QuizCollection struct {
	IdJoueur        string `json:"idjoueur"`
	QuizCollection  []Quiz `json:"quizcollection"`
	Mark            int    `json:"mark"` //initialisé à 0
	Finished        bool   `json:"finished"`
	Number_question int    `json:"number_question"` // savoir à quelle question est le joueur
}

type StatsUser struct {
	IdJoueur          string `json:"idjoueur"`
	NbQuizGeneral     int    `json:"quizgeneral"`
	NoteTotalGeneral  int    `json:"moyennegeneral"`
	NbQuizGenre       int    `json:"quizgenre"`
	NoteTotalGenre    int    `json:"moyennegenre"`
	NbQuizPicture     int    `json:"quizpicture"`
	NoteTotalPicture  int    `json:"moyennepicture"`
	NbQuizSynopsis    int    `json:"quizsynopsis"`
	NoteTotalSynopsis int    `json:"moyennesynopsis"`
}

type UserAnswer struct {
	Answer string `json:"answer"`
	Quiz   string `json:"quiz"`
}

type Pin struct {
	IdJoueur string   `json:"idjoueur"`
	Slugs    []string `json:"slugs"`
}
