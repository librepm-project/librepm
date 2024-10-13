#!/bin/sh
for dir in $(go work edit -json | jq -r '.Use[]["DiskPath"]'); do (cd $dir && go mod tidy); done