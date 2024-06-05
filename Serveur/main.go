package main

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	conf "pc3r_projet/config"
	log "pc3r_projet/services/log"
	pin "pc3r_projet/services/pin"
	quiz "pc3r_projet/services/quiz"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongoDB() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://Jm:26012001L!n@pc3r-anime.xpht143.mongodb.net/?retryWrites=true&w=majority&appName=PC3R-Anime").SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

func main() {
	client := connectMongoDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	conf.SetClient(client)
	conf.SetUserCollection(client.Database("Projet_PC3R").Collection("User"))

	r := mux.NewRouter()

	// Serve static files from the build directory
	staticDir := "../client/build"
	absStaticDir, err := filepath.Abs(staticDir)
	fmt.Println(absStaticDir)
	if err != nil {
		fmt.Println("Failed to get absolute path")
	}
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(absStaticDir+"/static"))))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(absStaticDir)))

	// Handle API routes
	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API endpoint")
	}).Methods("GET")

	fmt.Println("Serveur démarré sur http://localhost:8000")
	err = http.ListenAndServe(":8000", router())
	if err != nil {
		fmt.Println("Serveur arrêté")
	}
}

func router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/user/login", log.LoginHandler)
	r.HandleFunc("/user/signup", log.CreateUserHandler)
	r.HandleFunc("/user/logout", log.LogoutHandler)
	r.HandleFunc("/user/sessionExpired", log.SessionExpiredHandler)
	r.HandleFunc("/user", log.GetUserHandler)

	r.HandleFunc("/anime/trend", quiz.GetTrendAnimeHandler)
	r.HandleFunc("/anime/recent", quiz.GetRecentAnimeHandler)
	r.HandleFunc("/anime/generalQuiz", quiz.PostGeneralQuizHandler)
	r.HandleFunc("/anime/genreQuiz", quiz.PostGenreQuizHandler)
	r.HandleFunc("/anime/pictureQuiz", quiz.PictureQuizHandler)
	r.HandleFunc("/anime/synopsisQuiz", quiz.SynopsisQuizHandler)
	r.HandleFunc("/anime/answer", quiz.VerifyAnswerHandler)
	r.HandleFunc("/stats/generalQuiz", quiz.GetTop5QuizGeneral)
	r.HandleFunc("/stats/genreQuiz", quiz.GetTop5QuizGenre)
	r.HandleFunc("/stats/pictureQuiz", quiz.GetTop5QuizPicture)
	r.HandleFunc("/stats/synopsisQuiz", quiz.GetTop5QuizSynopsis)
	r.HandleFunc("/quiz/unfinishedQuiz", quiz.UnfinishedQuiz)

	r.HandleFunc("/pin/find", pin.FindPin)
	r.HandleFunc("/pin/create", pin.CreatePin)
	r.HandleFunc("/pin/delete", pin.DeletePin)
	r.HandleFunc("/pin/getAllPin", pin.GetAllPin)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
	})

	return c.Handler(r)
}
