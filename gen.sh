#!/bin/zsh

ls CHH/B/CHH_cc\(*  | sed 's/.*_.*_//g' | sed 's/\..*//g' | uniq | while read size ; 
do echo $size ; 
    CHHB=CHH-$size-B
    CHHC=CHH-$size-C
    rm -fr $CHHB
    rm -fr $CHHC
    mkdir $CHHB
    mkdir $CHHC
    ls CHH/B/*_$size\.* |sort -R |tail -n3 | sort | while read file; 
    do  
        cp $file $CHHB
    done
    ls CHH/C/*_$size\.* |sort -R |tail -n3 | sort | while read file; 
    do  
        cp $file $CHHC
    done
    
    RBB=Results-BB-$size
    RCC=Results-CC-$size
    RBC=Results-BC-$size
    mkdir -p $RBB
    mkdir -p $RCC
    mkdir -p $RBC
    ./pair.sh $CHHB $CHHB $RBB
    ./pair.sh $CHHC $CHHC $RCC
    ./pair.sh $CHHB $CHHC $RBC

done 


