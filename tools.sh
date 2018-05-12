#!/bin/bash

COMMENT=$1

if [ -z $COMMENT ]; then
    echo "必须加上注释"
    exit 1
fi

git add .;
git commit -m "$1";
git push origin master;
