#!/bin/sh

rm tmp.txt
cat jd2.txt | while read line
do
	if [ -n "$line" ]; then
		echo $line | tr '\n' ',' >> tmp.txt
	else
		echo $line >> tmp.txt
	fi	
done

cat tmp.txt | while read line
do
	echo $line | awk -F ',' '{if(NF > 5) print $4",",$2",",$3",",$5",",$6",",$7",",$8; else print $4",",$2",",$3}' >> jd5.txt
done

rm tmp.txt
