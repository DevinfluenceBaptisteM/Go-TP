TP cloud - Groupe Antoine MILOUX et Baptiste MABILLE

1. Jusification de l'utilisation du CI/CD

Dans le but de garantir une automatisation et une surveillance continue tout au long du cycle de vie de l'application, nous avons decidé d'utiliser CodeFresh. Notre choix c'est porté vers cette plateforme car elle est classé n°1 des plateforme Devops pour application kubernetes et qu'elle s'intègre facilement à des outils comme Slack, Git et solutions cloud. De plus, dans le cadre de se projet, elle nous permet de créer facilement et en quelques clic des pipelines. 

2. Detail du cout mensuel

AWS (0$/mois)

CodeFresh :

version gratuite
 - 3 utilisateurs
 - Deploiement SaaS
 - Build illimité
 - Dossiers illimité
 - Pipeline
 - Dépôt de Helm privé
 - Conservation des données pendant 1 mois

 version payante (à partir de 34$/mois)
 - Jusqu'à 10 utilisateurs
 - Tableau de bord Kubernetes
 - Tableau de bord de pipeline
 - Débogage du pipeline en direct
 - Tableaux Kanban de libération de pipeline
 - Conservation des données pendant 6 mois

Mongo Atlas (0$/mois)

3. Localisation du service Cloud 

L'utilisation de AWS nous permet de stocker en Europe et en asie du fait de la localisation de leurs centres de données.

4. Solution mongoDB

La solution MongoDb Atlas a été retenue du fait de sa simplicité de connexion avec un service cloud et de la diffusion des données à travers le monde. De plus, elle dispose d'une version gratuite.

5. Les limites

6. Votre local (PC) est un cloud privé

I/ Test unitaires :
    1. POST -> OK
    2. GET -> OK
    3. PUT -> xxx
    4. DELETE -> xxx

II/ Mise en place de MongoDB
    OK

III/ Mise en place d'un CI/CD
    ???

IV/ Automatisation des tests
    xxx

V/ Dépot de code
    xxx

VI/ Plug de la stack CI/CD et Github
    xxx

VII/ Création d'un cloud client dans un cloud provider
    OK

VIII/ Déploiement d’un espace (VM free-tier type linux)
    xxx

IIX/ Configuration des accès à l’espace entre l’outil de CI/CD et la VM
    xxx

IX/ Si test ok déploiement dans la VM 
    xxx
    
IIX/ Réaliser plusieurs pull request en créant des nouvelles route
    xxx