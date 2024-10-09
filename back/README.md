- La liste des binaires et les fonctionnalités fournies
- Les procédure de compilation et dʼexécution du ou des binaires
- Les procédures de lancement des tests unitaires
- Un exemple de configuration pour lancement sur poste de travail local

# Démarrer l'application
`go run main.go`

accès en local: localhost:8080

lien du back: https://kermesse-land-5d1bb3e5d576.herokuapp.com/

### La liste des binaires et les fonctionnalités fournies

### Les procédure de compilation et dʼexécution du ou des binaires

### Les procédures de lancement des tests unitaires

### Un exemple de configuration pour lancement sur poste de travail local 

## Prérequis

1. Go : Assurez-vous que Go est installé sur votre machine. Si ce n'est pas déjà fait, téléchargez-le depuis golang[.org.](https://golang.org/dl/)![](Aspose.Words.e08f853a-ea3d-4cce-956a-28c62c5d1e97.001.png)
2. Docker : Assurez-vous que Docker est installé. Vous pouvez télécharger Docker Desktop depuis docker.com.

## Fichier .env

Le back-end utilise un fichier .env pour stocker les configurations sensibles comme les informations de connexion à la base de données et les clés Stripe.

## Créer un fichier .env

Créez un fichier .env à la racine du projet avec le contenu suivant : makefile

Copier le code DB\_URL=postgres://user:password@localhost:5432/dbname STRIPE\_SECRET\_KEY=sk\_test\_votre\_cle\_secrete PORT=8080

## Variables d'environnement

- DB\_URL: L'URL de connexion à votre base de données PostgreSQL locale.
- STRIPE\_SECRET\_KEY: La clé secrète Stripe pour les paiements.
- PORT: Le port sur lequel le serveur Go écoutera (par défaut 8080).

## Lancer le projet localement avec Docker

Le projet inclut un fichier Dockerfile pour faciliter l'exécution du back-end en local via Docker. Suivez les étapes ci-dessous pour tester le projet en local.

Dockerfile à modifier

Modifiez votre fichier Dockerfile pour faciliter le test en local : Dockerfile

Copier le code

- Dockerfile
- Utilisation de l'image Go officielle `FROM golang:1.20-alpine`
- Définir le répertoire de travail `WORKDIR /app`
- Copier les fichiers de go.mod et go.sum pour installer les dépendances `COPY go.mod go.sum ./`
- Télécharger les dépendances Go RUN go mod download
- Copier le reste du code de l'application `COPY . .`
- Construire l'application Go `RUN go build -o main .`
- Exposer le port sur lequel l'application écoutera `EXPOSE 8080`
- Lancer l'application CMD ["./main"]

Lancer l'application en local avec Docker

1\. Build de l'image Docker :

`docker build -t go-backend .`

2\ Exécuter un conteneur Docker :

`docker run --env-file .env -p 8080:8080 go-backend`

3\ Accéder à l'application : L'API sera accessible à l'adresse suivante :

http://localhost:8080.

## Tester les paiements Stripe en local

Assurez-vous que le secret Stripe dans votre fichier .env est une clé de test. Stripe permet de tester les paiements en utilisant des cartes de test comme 4242 4242 4242 4242.
