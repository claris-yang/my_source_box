#!/bin/bash  
  
s='{"rv":0,flag:1,"url":"http://www.jinhill.com","msg":"test"}' 
  
  
parse_json(){    
value=`echo $1 | sed 's/.*"$2":[^,}]*[^,}]*.*/\1/'`  
echo $value | sed 's/\"//g'  
} 
  
  
echo $s  
  
value=$(parse_json $s "url")  
  
echo $value  
