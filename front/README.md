# front

Kermesse Land project.

## Getting Started

This project is a starting point for a Flutter application.

A few resources to get you started if this is your first Flutter project:

- [Lab: Write your first Flutter app](https://docs.flutter.dev/get-started/codelab)
- [Cookbook: Useful Flutter samples](https://docs.flutter.dev/cookbook)

For help getting started with Flutter development, view the
[online documentation](https://docs.flutter.dev/), which offers tutorials,
samples, guidance on mobile development, and a full API reference.

## Prérequis

1. Flutter SDK : Assurez-vous que Flutter est installé sur votre machine. Si ce n'est pas déjà fait, suivez les instructions d'installation sur flutter.dev.
2. Dépendances : Ce projet utilise plusieurs packages Flutter, dont flutter_dotenv pour la gestion des variables d'environnement et flutter_stripe pour les paiements Stripe. Vous pouvez installer les dépendances avec la commande suivante :
bash
`flutter pub get`

## Fichier .env

Le projet utilise un fichier .env pour gérer certaines configurations sensibles comme les clés Stripe et l'URL de l'API.

1. Créez un fichier .env à la racine du projet à partir du fichier .env.example.
2. Ajoutez les informations suivantes :

`PROD_MODE=0
BASE_URL=https://kermesse-land-5d1bb3e5d576.herokuapp.com/
PUBLIC_STRIPE_KEY=sk_test_votre_cle_publique`

## Variables d'environnement
- PROD_MODE: Indique si l'application est en mode production (1) ou développement (0).
- BASE_URL: https://kermesse-land-5d1bb3e5d576.herokuapp.com/

## Démarrer l'application

Pour lancer l'application sur un simulateur ou un appareil physique, suivez les étapes ci-dessous :

1. Configurer les variables d'environnement : Assurez-vous que le fichier .env est bien configuré.
2. Lancer l'application :
`flutter run`

3. Mode Debug/Production :
   - En mode développement, gardez PROD_MODE=0.
   - Pour le mode production, changez la valeur en 1 et assurez-vous d'utiliser la clé Stripe et l'URL API correctes.
   
## Dépendances principales

- `flutter_dotenv`: Gestion des variables d'environnement.
- `flutter_stripe`: Intégration de Stripe pour les paiements.