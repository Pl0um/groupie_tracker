# Groupie Tracker

## Description

Groupie Tracker est une application web qui permet de suivre les informations sur les artistes musicaux. Elle affiche des détails tels que le nom, les membres, la date de création, le premier album, les lieux de concert et bien plus encore. Ce projet a été créé pour fournir une interface simple et intuitive pour accéder aux informations des artistes.

## Fonctionnalités

- Affichage des informations détaillées sur les artistes
- Recherche d'artistes par nom
- Navigation entre les pages d'accueil, de détails des artistes et de crédits
- Interface utilisateur responsive et attrayante

## Créateurs

- **Kéo** : [Pl0um](https://github.com/Pl0um)
- **Mael** : [CasualElf34](https://github.com/CasualElf34)
- **Lucas** : [Howexou](https://github.com/Howexou)

## Prérequis

- Go 1.22.2 ou supérieur
- Connexion Internet pour accéder à l'API des artistes

## Installation

1. Clonez le dépôt :
    ```sh
    git clone https://github.com/votre-utilisateur/groupie_tracker.git
    cd groupie_tracker
    ```

2. Installez les dépendances :
    ```sh
    go mod tidy
    ```

## Utilisation

1. Lancez le serveur :
    ```sh
    go run main.go
    ```

2. Ouvrez votre navigateur et accédez à :
    ```
    http://localhost:2026
    ```

## Structure du projet

- [main.go](http://_vscodecontentref_/0) : Point d'entrée principal de l'application
- [serveur](http://_vscodecontentref_/1) : Contient les fichiers de gestion des routes et des handlers
- [template](http://_vscodecontentref_/2) : Contient les fichiers HTML pour les différentes pages
- [css](http://_vscodecontentref_/3) : Contient les fichiers CSS pour le style des pages
- [js](http://_vscodecontentref_/4) : Contient les fichiers JavaScript pour les fonctionnalités dynamiques

## Routes

- `/` : Page d'accueil affichant la liste des artistes
- `/Groupie` : Page de détails d'un artiste
- `/Credit` : Page des crédits

## Exemple de données

Les données des artistes sont récupérées depuis l'API `https://groupietrackers.herokuapp.com/api/artists`.

## Contribution

Les contributions sont les bienvenues ! Si vous souhaitez contribuer, veuillez ouvrir une issue ou soumettre une pull request.

## Licence

Ce projet est sous licence MIT. Voir le fichier LICENSE pour plus de détails.