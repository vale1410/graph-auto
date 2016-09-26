#!/bin/zsh

time=3600

mkdir -p $3

for x in $1/*; 
do 
    for y in $2/*
    do 
        if [[ ! $x -ef $y ]]
        then 
            if [[ $x < $y ]]
            then 
                name=${x:e}-${y:e}
                echo running $x $y $time piping into $3/$name
                ./graph-iso-asp.sh $x $y $time | tee $3/$name
            fi 
        fi 
    done 
done
