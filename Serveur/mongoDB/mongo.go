package mongoDB

import (
	"context"
	"fmt"

	conf "pc3r_projet/config"
	ds "pc3r_projet/dataStructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func FindUserByID(userid string) (ds.User, error) {
	objectID, _ := primitive.ObjectIDFromHex(userid)
	filtre := bson.M{"_id": objectID}
	var user ds.User
	err := conf.GetUserCollection().FindOne(context.Background(), filtre).Decode(&user)
	if err != nil {
		return ds.User{}, err
	}

	return user, nil
}

func FindUserByLogin(login string) (ds.User, error) {
	var user ds.User
	err := conf.GetUserCollection().FindOne(context.Background(), bson.M{"connexion.login": login}).Decode(&user)
	if err != nil {
		return ds.User{}, err
	}

	return user, nil
}

func IsLoginTaken(login string) bool {
	var result ds.User

	err := conf.GetUserCollection().FindOne(context.Background(), bson.M{"connexion.login": login}).Decode(&result)
	switch err {
	case nil: // Un utilisateur avec ce login existe déjà
		return true
	case mongo.ErrNoDocuments: // Aucun utilisateur avec ce login trouvé
		return false
	default: // Une erreur inattendue s'est produite
		fmt.Println("Erreur lors de la recherche de l'utilisateur:", err)
		return true // On suppose que le login est déjà pris par défaut
	}
}

func IsPseudoTaken(pseudo string) bool {
	var result ds.User

	err := conf.GetUserCollection().FindOne(context.Background(), bson.M{"pseudo": pseudo}).Decode(&result)
	switch err {
	case nil:
		return true
	case mongo.ErrNoDocuments:
		return false
	default:
		fmt.Println("Erreur lors de la recherche de l'utilisateur:", err)
		return true
	}
}

func InsertUser(user ds.User) error {
	user.Connexion.Password = HachedPassword(user.Connexion.Password)
	user.Connexion.Connected = false
	_, err := conf.GetUserCollection().InsertOne(context.Background(), user)
	return err
}

func HachedPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	hashedPasswordString := string(hashedPassword)
	return hashedPasswordString
}

// Passer de connecter à déconnecté et vice versa
func ToggleConnectionStatus(userid string, newEtat bool) (*mongo.UpdateResult, error) {
	objectID, _ := primitive.ObjectIDFromHex(userid)
	filter := bson.M{"_id": objectID}
	var user ds.User

	conf.GetUserCollection().FindOne(context.Background(), filter).Decode(&user)
	update := bson.M{"$set": bson.M{"connexion.connected": newEtat}}
	return conf.GetUserCollection().UpdateOne(context.Background(), filter, update)
}

func InsertStats(userid string) error {
	collectionStat := conf.GetClient().Database("Projet_PC3R").Collection("statsUser")
	userS := ds.StatsUser{
		IdJoueur:          userid,
		NbQuizGeneral:     0,
		NoteTotalGeneral:  0,
		NbQuizGenre:       0,
		NoteTotalGenre:    0,
		NbQuizPicture:     0,
		NoteTotalPicture:  0,
		NbQuizSynopsis:    0,
		NoteTotalSynopsis: 0}

	_, err := collectionStat.InsertOne(context.Background(), userS)

	return err
}
