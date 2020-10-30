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

Mongo Atlas (0$/mois) en fonction de son utilisation 

3. Localisation du service Cloud 

Nous utiliserons AWS afin de profiter d’une facilité d’utilisation et de documentation. 
N’ayant pas de visibilité sur la quantité de donnée que nous utiliserons, AWS nous donne l’avantage de pouvoir anticipé une bonne évolutivité de notre application grâce à l’Auto Scaling tout en ne payant seulement les ressources utilisées.
En terme de fiabilité, AWS a l’avantage d’être au cœur de l’organisation d’infrastructure de grands groupe tel que Netflix, garantissant la sécurité de nos opérations.

4. Solution mongoDB

La solution MongoDb Atlas a été retenue du fait de sa simplicité de connexion avec le service cloud d'Amazon et de la diffusion des données à travers le monde. De plus, elle dispose d'une version gratuite.
De plus avec Atlas Global Cluster, nous pouvons repartir les clusters sur plusieurs pays

5. Les limites

Infrastructure a faible ressources

6. Votre local (PC) est un cloud privé