#!/bin/bash

find -P src '*' -type f -perm -o=x -exec ok '{}' \;