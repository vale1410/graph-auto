#!/bin/zsh

go build convert.go
tmp=/tmp/test.lp

for x in $1/*; 
do 
    ./convert -f $x -di > $tmp
    echo $(basename $x) $(gringo3 model.lp $tmp | clasp --stat --configuration=many --time-limit 60 | grep SATI) 
done 



