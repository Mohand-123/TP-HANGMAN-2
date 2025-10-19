🛒 TP – Rendre un site web dynamique avec Go
📌 Description
Ce projet est un TP réalisé dans le cadre du cours de programmation web avec Golang.
L’objectif est de mettre en pratique la gestion des routes, des templates HTML, et des formulaires afin de rendre un site e-commerce dynamique.
Le site simule une boutique en ligne de style streetwear, avec une page d’accueil listant les produits, une page de détails pour chaque article, et un formulaire (en cours) pour ajouter de nouveaux produits.

🚀 Fonctionnalités implémentées
✅ Challenge 1 – Affichage de la liste des articles
- Page d’accueil avec :
- En-tête (logo + menu de navigation)
- Liste de 5 produits minimum
- Affichage : image, nom, prix, réduction (optionnelle)
- Données stockées dans une variable globale
- Utilisation de templates Go avec boucles, variables et conditions
✅ Challenge 2 – Afficher l’article sélectionné
- Page de détails pour chaque produit :
- Nom, description, image, prix, stock, réduction (optionnelle)
- Gestion des erreurs si l’article n’existe pas
- Bouton "Voir le produit" ajouté sur la page d’accueil pour accéder aux détails
⏳ Challenge 3 – Ajouter un produit
- En cours de développement
- Objectif : créer un formulaire permettant d’ajouter un produit à la liste
- Redirection prévue vers la page de détails du produit ajouté

🛠️ Technologies utilisées
- Go (Golang) – gestion des routes et logique serveur
- HTML / CSS – templates et mise en forme
- net/http – serveur web
- text/template – moteur de templates Go

▶️ Lancer le projet
- Cloner le dépôt :
git clone https://github.com/ton-compte/ton-repo.git
cd ton-repo
- Lancer le serveur Go :
go run main.go
- Ouvrir le navigateur à l’adresse :
http://localhost:8000



