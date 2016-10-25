#!/bin/zsh

tmp1=$(mktemp)
tmp2=$(mktemp)

./graph $1 $2 > $tmp1 
./convert -f $tmp1 $4  > $tmp2
gringo3 model.lp $tmp2 | clasp --stat --configuration=many --time-limit $3
rm -fr $tmp1
rm -fr $tmp2
