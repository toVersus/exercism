#!/usr/bin/env bash

INPUT=$1

scrabble_score() {
    total=0
    for c in $(grep -o . <<<${INPUT^^} ); do
        case $c in
            [AEIOULNRST]) ((total++));;
            [DG])         ((total+=2));;
            [BCMP])       ((total+=3));;
            [FHVWY])      ((total+=4));;
            [K])          ((total+=5));;
            [JX])         ((total+=8));;
            *)            ((total+=10));;
        esac
    done

    echo $total
}

scrabble_score
