#!/bin/sh

group_id=$1

rm tmp.txt
cat jd2.txt | while read line
do
	if [ -n "$line" ]; then
		echo $line | tr '\n' ',' >> tmp.txt
	else
		echo $line >> tmp.txt
	fi	
done

echo "\n" >> tmp.txt

cat tmp.txt | while read line
do
	if [ -n "$line" ]; then
		h_group_id=`echo $line | awk -F ',' '{if(NF > 5) print $4; else print $4}' | awk -F ' ' '{print $2}'`
		echo $h_group_id
		if [ "$group_id" = "$h_group_id" ]; then
			echo $line | awk -F ',' '{if(NF > 5) print $4",",$2",",$3",",$5",",$6",",$7",",$8; else print $4",",$2",",$3}' > jd5.txt
		fi
	fi
done

