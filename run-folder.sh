#!/bin/zsh

for x in $1/*; 
do 
    echo $x
    gringo3 model.lp $x | clasp --stat --configuration=many -t 10 --time-limit 60 > $2/$(basename $x)
done 



