Ce code  en ligne de commande, écrit en Go , lit un fichier texte, applique une série de règles de correction, puis écrit le résultat dans un fichier texte de sortie.

Fonctionnalités

Le code applique, ligne par ligne et dans cet ordre :
1.	Les balises de transformation : (hex), (bin), (up), (low), (cap)
2.	La correction de la ponctuation
3.	La correction des apostrophes
4.	La correction des articles a / an
Les retours à la ligne du texte d'origine sont conservés.

Structure du projet
.
├── go.mod           module Go
├── main.go          lecture des arguments, lecture et écriture des fichiers
├── rules.go         Process et toutes les règles de correction
├── rules_test.go    tests unitaires
└── entree.txt       fichier d'exemple

les tests

go test ./... donne ok
go test -v ./... donne ok

Le fichier entree.txt sert comme exemple d'utilisation des balises Étape par Étape.


Le fichier exemple_complet.txt comporte toutes les balises dans un texte unique ci-dessous. 

Entrée :
Simply add 1E (hex) and 10 (bin) and you will see the result is 32 .
this is so exciting (up, 2) , but i should stop SHOUTING (low) and welcome to the brooklyn bridge (cap, 3) .
It was a amazing day...we waited a hour , then BAMM !! it was over , wasn't it ?
As he said : ' I am the happiest men in the world ' , and I agree .
There is a unusual story here (up, 2) and a honest answer (cap) !

Sortie :
Simply add 30 and 2 and you will see the result is 32.
this is SO EXCITING, but i should stop shouting and welcome to The Brooklyn Bridge.
It was an amazing day... we waited an hour, then BAMM!! it was over, wasn't it?
As Elton John said: 'I am the happiest men in the world', and I agree.
There is an unusual STORY HERE and an honest Answer!
