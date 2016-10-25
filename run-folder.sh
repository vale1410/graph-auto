#!/bin/zsh

go build convert.go
tmp1=/tmp/test1.lp
tmp2=/tmp/test2.lp
tmp3=/tmp/test3.lp

for x in $1/*; 
do 
    ./convert -f $x -di > $tmp1
    ./convert -f $x -ddi > $tmp2
    ./convert -f $x -dd2 > $tmp3
    echo $(basename $x) $(gringo3 model.lp $tmp1 | clasp --stat --configuration=many --time-limit 180 | grep 'SATI\|CPU')  $(gringo3 model.lp $tmp2 | clasp --stat --configuration=many --time-limit 180 | grep 'SATI\|CPU')   $(gringo3 model.lp $tmp3 | clasp --stat --configuration=many --time-limit 180 | grep 'SATI\|CPU')  

done 



