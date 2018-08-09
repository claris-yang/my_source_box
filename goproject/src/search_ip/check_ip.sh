#/bin/bash
tmp_ip1=""
tmp_ip2=""
cat mi_china.txt | while read line
do
	echo $line
	startIp=`echo "$line" | awk -F ',' '{print $1}'`
	endIp=`echo "$line" | awk -F ',' '{print $2}'`
	
	if [ $startIp -gt $endIp ]
	then
		echo "失败1, $startIp|$endIp"		
	fi

	if [ -z $tmp_ip2 ]
	then
		echo $tmp_ip2
		if [ $tmp_ip2 -gt $startIp ]
		then
			echo "失败2, $tmp_ip2|$startIp|"
		fi
	fi

	tmp_ip1=$startIp
	tmp_ip2=$endIp

done

