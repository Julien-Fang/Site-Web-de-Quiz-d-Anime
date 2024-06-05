Nous avons utilisé React, Bootstrap pour la partie client.
Le serveur s'est réalisé en Go avec la partie base de donnée avec MongoDB.

comment elles sont utilisées ????????????????????????

Ecrans du site : 
- Page de connexion/inscription
- Page d'accueil
- Page de quiz

Evenements :
- Créer un compte
- Se connecter
- Se déconnecter
- Chargement des animés tendances
- Chargement des animés récemment sortis
- Indication des quiz non terminés et pouvoir le continuer
- Affichage d'un classement en fonction du quiz, un leaderboard différent pour chaque quiz
- Lancer un quiz aux choix
- Epingler un animé durant un quiz
- Supprimer un épingle

Enchainement des callback :
- Dans la page d'accueil :
   - Récuperer les 10 premiers animés du tendance
   - Récupérer les animés sortis récemments
   - S'il y a un quiz non terminé, alors, affichage des quiz non terminés
   - S'il n y a aucun quiz non terminé, affichage des animés épinglés
- Lorsque l'utilisateur sélectionne l'un des types de quiz disponibles : "General", "Genre", "Picture", "Synopsis". La génération du quiz correspondant démarre uniquement s'il n'a pas déjà un quiz non terminé du même type.
- Durant un quiz :
   - Récupérer l'information de l'animé deja épinglé ou non
   - Epingler ou désépingler un animé
   - Récupérer les informations pour afficher un classement



Liste des appels AJAX serveur :
- http://localhost:3000/user/login : Pouvoir se connecter
- http://localhost:3000/user/signup : Pouvoir s'inscrire
- http://localhost:3000/user/logout : Se deconnecter
- http://localhost:3000/user/sessionExpired :
- http://localhost:3000/user : 
- http://localhost:3000/anime/trend : Obtenir les dernières animés les plus populaires
- http://localhost:3000/anime/recent : Obtenir les dernières animés sortis
- http://localhost:3000/anime/generalQuiz : Générer un quiz général
- http://localhost:3000/anime/genreQuiz : Générer un quiz où il faut deviner le genre
- http://localhost:3000/anime/pictureQuiz : Générer un quiz contenant des images comme indice
- http://localhost:3000/anime/synopsisQuiz : Générer un quiz contenant des synopsis comme indice
- http://localhost:3000/anime/answer : 
- http://localhost:3000/stats/generalQuiz : Obtenir le top 5 du classement du quiz général
- http://localhost:3000/stats/genreQuiz : Obtenir le top 5 du classement du quiz genre (où il faut trouver le genre de l'animé)
- http://localhost:3000/stats/pictureQuiz : Obtenir le top 5 du classement du quiz d'image
- http://localhost:3000/stats/synopsisQuiz : Obtenir le top 5 du classement du quiz de synopsis
- http://localhost:3000/quiz/unfinishedQuiz : Récuperer les quiz non terminés
- http://localhost:3000/pin/find : Trouver les animés épinglés par l'utilisateur
- http://localhost:3000/pin/create : Créer d'un épingle
- http://localhost:3000/pin/delete : Suppression d'un épingle
- http://localhost:3000/pin/getAllPin : Récupérer tous les épingles de l'utilisateur
