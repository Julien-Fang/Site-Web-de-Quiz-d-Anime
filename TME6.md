Le site web sera dédié aux quiz sur les animés. Elle se distingue par sa variété de quiz qui peuvent être généraux ou spécifique à une catégorie, permettant aux utilisateurs de tester et d’affiner leurs connaissances. Épingler une question est possible si l'utilisateur souhaite découvrir un nouvel animé durant les quiz.

Les principales fonctionnalités :
- Quiz généraux et spécifiques : Les utilisateur peuvent choisir de participer à des quiz généraux couvrant une variété de thèmes d’animés différents, ou se concentrer sur des quiz spécifiques à un thème.
- Quiz d’entrainement : Des quiz d'entraînement conçus pour améliorer les compétences des utilisateurs en matière de connaissances sur les animés.
- Système de Classement : Un système de classement dynamique qui permet aux utilisateurs de se mesurer aux autres (et de suivre leur progression au fil du temps.)
- 

L’API choisit est : https://kitsu.docs.apiary.io/#
Cette API propose une grande base de données d’animée incluant les plus récents. On peut récupérer plusieurs informations, notamment : le nom de l’animé, un synopsis, les votes de ranking, son rang, etc …
Afin de récupérer les informations sur un animé, il faut avoir son id, cela suffit pour obtenir l’ensemble des informations.

Cas d’utilisation : 
- L’utilisateur se connecte, sélectionne un quiz général et obtient une note à la fin.
- L'utilisateur sélectionne le quiz d'entrainement, épingle un quiz, réponse aux questions et obtient une note.
- L'utilisateur sélectionne le classement, pour regarder son classement.

Base de donnée:
- Utilisateur : idUser, pseudo
- Pseudo : tableau de pseudo (pour eviter d'avoir des doublons)
- Classement : idUser, rang
- StatsUser : idUser, nombre de quiz lancé (généraux), nombre de quiz lancé (spécifique)
- Pin : idUser, liste[anime]

Les données sont mis à jour à chaque fois que l'utilisateur sélectionne un des quiz. Lors du quiz, sa table de StatsUser change car il effectue un quiz donc le nombre de quiz lancé augmente. Egalement, la table Classe et Pin peuvent changer si l'utilisateur obtient un meilleur score qu'un autre joueur et s'il épingle un animé durant son quiz.
L'API externe est appelée à chaque fois qu'un utilisateur clique sur un des quiz, par exemple s'il choisit d'être testé sur le thème action, le server fera une requête qui récupère tous les animés ayant le thème action.

Description du client : 
dans ce lien google doc : https://docs.google.com/document/d/1Pbr_hbi0apRdtExcw7mcpIjeIg68_S3joMOV0gwt4A0/edit


Description des requetes :
  Le serveur retourne une réponse JSON contenant son id, sa date de création, un synopsis, ses titres (dans differents pays), sa notation, etc ...
  La réponse JSON renvoyée est assez conséquente, il est préférable de voir par soi-même en mettant l'URL sur le navigateur. 
  Dans la réponse JSON, toutes les informations concernant 'One piece', y compris les détails de chaque saison, sont centralisées. Cela signifie qu'il n'est pas nécessaire de rechercher des informations spécifiques à chaque saison comme 'One piece saison 3', car ces informations sont déjà intégrées dans la structure de la réponse JSON

- Récupérer les informations concernant un animé a l'aide de son ID : https://kitsu.io/api/edge/anime/{id}
  Par exemple en mettant https://kitsu.io/api/edge/anime/1 , l'animé devrait etre : Cowboy Bebop.

- Récupérer les informations des animés en fonction du genre : https://kitsu.io/api/edge/anime?filter[genres]={genre}
  Par exemple : https://kitsu.io/api/edge/anime?filter[genres]=action,fantasy
  Le serveur renvoie les 10 premiers animés respectant ses genres. Et pour récupérer la suite des animés du même genre, s'il existe, il faut récupérer le lien indiqué dans "next" à la fin de la reponse JSON.

- Récupérer les informations d'un animé avec son titre : https://kitsu.io/api/edge/anime?filter[text]={titre}
  Par exemple : https://kitsu.io/api/edge/anime?filter[text]=my hero academia
  Il est également possible de le retrouver dans une autre langue si aucune faut d'orthographe n'est présent : https://kitsu.io/api/edge/anime?filter[text]=進撃の巨人 (correspond a "Attack on titan")

- Récupérer les informations d'un animé en fonction de sa note et du genre : https://kitsu.io/api/edge/anime?sort=-averageRating&filter[genres]={genre}
  Par exemple : https://kitsu.io/api/edge/anime?sort=-averageRating&filter[genres]=action,comedy
  Avec -averageRating permettant de prendre dans l'ordre décroissant.

- Récupérer le top 10 tendances des animés : https://kitsu.io/api/edge/trending/anime

 - 
