# TP2 – Pilotes de Formule 1

Programme en Go (ligne de commande) qui interroge une API publique de F1 et affiche les informations d'un pilote à partir de son **numéro**, ou de **tous les pilotes** de la saison.

Pour chaque pilote, il affiche : numéro, nom complet, nationalité et code (ex. `HAM`).

## Prérequis

- [Go] 1.27.1 ou plus
- Une connexion Internet (les données viennent de `https://f1-companion-api-ba5k.onrender.com/drivers`)

## Lancer le programme

Depuis le dossier du projet (celui qui contient `go.mod`) :

```bash
go run main.go
 
```

## Utilisation

Le programme demande un numéro de pilote :

- un **entier positif** (ex. `44`) → affiche le pilote correspondant ;
- `all` (ou `ALL`) → affiche tous les pilotes.

Une saisie invalide (texte, nombre négatif) est refusée et la question est reposée.

### Exemples

> Les données ci-dessous sont indicatives : elles dépendent de ce que renvoie l'API.

Un pilote :

```text
sélectionnez un numéro de pilote : 44
vous avez choisi:  44
---   pilote n. 44    ---
nom complet:  Lewis Hamilton
nationalité :  British
code :  HAM
```

Tous les pilotes :

```text
sélectionnez un numéro de pilote : all

vous avez choisi d'afficher les informations de tous les pilotes.
---   pilote n. 44    ---
nom complet:  Lewis Hamilton
nationalité :  British
code :  HAM
---   pilote n. 16    ---
nom complet:  Charles Leclerc
nationalité :  Monegasque
code :  LEC
...
```

Numéro inconnu :

```text
sélectionnez un numéro de pilote : 99
vous avez choisi:  99
le nombre 99 ne correspond à aucun pilote.
```

> Si l'API est injoignable ou met plus de 10 s à répondre, le programme affiche les erreurs rencontrées, puis `impossible de récupérer la liste des pilotes.` Dans ce cas, vérifiez votre connexion et relancez-le.

## Structure du projet

| Fichier | Rôle |
|---|---|
| `main.go` | Saisie utilisateur et enchaînement des étapes |
| `API/appel.go` | Requête HTTP GET vers l'API (timeout 10 s) |
| `Json/JSON.go` | Décodage du JSON en liste de `Driver` |
| `fetch/fetchNumber.go` | Sélection d'un pilote par numéro (ou de tous) |
| `affichage_pilote/afficher_pilote.go` | Affichage des pilotes dans le terminal |
