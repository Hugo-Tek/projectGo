# projectGo
    Projet qui consiste à optimiser le parcours des transpalettes dans un entrepôt afin de maximiser le remplissage de camions de transports partant d’un entrepôt.
    Module GO - EPITECH
## Installer go(lang)

avec [homebrew](http://mxcl.github.io/homebrew/):

```Shell
sudo brew install go
```

avec [apt](http://packages.qa.debian.org/a/apt.html)-get:

```Shell
sudo apt-get install golang
```

[installer Golang manuellement](https://golang.org/doc/install)

### Compilation de notre projet

Pour exécuter notre programme il suffit d'écrire ```go run main.go``` suivi du nom du fichier que vous voulez passer en paramètre par exemple ```file.txt```.

### Description du fichier en entrée

La première ligne contiendra 3 nombre séparés par un espace :
* nombre de cases en largeur de l’entrepôt
* nombre de cases en longueur de l’entrepôt
* nombre de tours à simuler

Les n lignes suivantes décriront lescolisprésentsdans l’entrepôt.Chaque ligne contiendra 3 éléments :
* le nom du **colis** (une chaîne de caractères sans espace)
* la position du **colis** (2 nombres séparés par des espaces)
* la couleur du **colis** (une chaîne de caractères valant GREEN,YELLOW ou BLUE, pouvant être en majuscule ou en minuscule)

Les k lignes suivantes décriront les **transpalettes** présents dans l’**entrepôt**.
Chaque ligne contiendra 2 éléments :
* le nom du **transpalette** (chaîne de caractères sansespace)
* la position du **transpalette** (2 nombres séparés pardes espaces)

Enfin les p dernières lignes du fichier seront composées des éléments suivants :
* le nom du **camion**(chaîne de caractères sans espace)
* la position de l’**aire de dépôt** (2 nombres séparés par des espaces)
* la charge maximale du **camion** (1 nombre)
* le nombre de tours avant que le **camion** ne soit de nouveau disponible


Voici un exemple de fichier avec 4 **colis**, 1 **transpalette**, 1 **camion** et un **entrepôt** de forme carré :

```Shell
5 5 1000 
colis_a_livrer 2 1 green 
paquet 2 2 BLUE 
deadpool 0 3 yellow 
colère_DU_dragon 4 1 green
transpalette_1 0 0 
camion_b 3 4 4000 5
```


### Organisations et packages

À la racine nous avons le ```main.go``` avec le fichier pour l'entrée ```file.txt```.
Ensuite nous avons un dossier ```src``` où se trouve tout le code,
séparé en 3 parties les **components** dans un dossier ```components```, la **gestion d'erreur**  dans un dossier ```errors```, et les fichiers principaux à la racine de ```src```.

### 