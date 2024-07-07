# kleo-back

Kleo is a personal financial management app for Android written in go and kotlin.

## Vocabulary

* Transaction / transaction : Une entrée ou une sortie d'argent sur un compte en banque (peut être positif ou négatif, ne peux être nul). Est caractérisé par un montant (dans un certaine devise), le compte en banque (unique) sur lequel la transaction à lieux, une date, une date originale (peut être différente de la date), un libellé (ou label). Une transaction, lors du pointage, est assignée à une catégorie, elle correspond donc à un certain type de mouvement.
* Amount / Montant : Montant en devise d'une transaction (par défaut en euro)
* Category / catégorie : une catégorie permet de classer les transactions. Elle est caractérisée par un libellé (ou label), un type de mouvement (flow type), et un type de catégorie
* Category type / type de catégorie : Classification permettant de subdiviser les catégories en différents types. Par exemple, différents types de dépenses : obligatoires fixes, obligatoires variables, facultatives.
* Flow type / type de mouvement : permet de qualifier le type de mouvement d'argent d'une transaction : dépense, revenu, épargne
