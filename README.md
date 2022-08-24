# The Problem

Given the attached text file as an argument, your program will read the file, and output the 20 most frequently used words in the file in order, along with their frequency. <br />
The output should be the same to that of the following bash program: <br />

#!/usr/bin/env bash <br />

cat $1 | tr -cs 'a-zA-Z' '[\n*]' | grep -v "^$" | tr '[:upper:]' '[:lower:]'| sort | uniq -c | sort -nr | head -20 <br />
