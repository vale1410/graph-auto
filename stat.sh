#!/bin/zsh


echo size ';' pairing ';' result ';' total ';'  search ';' clauses ';' conflicts ';' 
ls CHH/B/CHH_cc\(*  | sed 's/.*_.*_//g' | sed 's/\..*//g' | uniq | while read size ; 
do
    for x in Results-*-$size/*; 
    do 
        echo $size ';' $(basename $x) ';' $(cat $x | grep 'SATIS' ) ';' \
        $(cat $x | grep 'Solving: ' | sed 's/Time *: //g' | sed 's/ (.*//g' )  ';' \
        $(cat $x | grep 'Solving:' | sed 's/.*g: //g' | sed 's/1st.*//g') ';' \
        $(cat $x  | grep 'Rules' | sed 's/.* : //g' | sed 's/ *(.*//g' ) ';' \
        $(cat $x  | grep 'Conflicts' | sed 's/.* : //g' | sed 's/ *(.*//g' ) ';' 
    done 
done
