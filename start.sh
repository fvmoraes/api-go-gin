#!/bin/bash

#Docker objects cleaner
docker stop $(docker ps -qa) && docker system prune -af --volumes
docker compose up

#Curl messages to initial populate database
req=$1
count=0

for ((count=1;count<=req;count++)); do
    curl --request GET --url http://172.33.0.4:9000/api/v1/person/ --header 'Content-Type: application/json' >> /dev/null
    curl --request GET --url http://172.33.0.4:9000/api/v1/person/$count --header 'Content-Type: application/json' >> /dev/null
    curl --request POST --url http://172.33.0.4:9000/api/v1/person/ --header 'Content-Type: application/json' --data '{"name": "User","job": "Job","age":25,"photo": "https://aws.s3.bucket/PHOTO.jpg"}' >> /dev/null
    if [ $(($count % 5)) -eq 0  ]; then
        curl --request DELETE --url http://172.33.0.4:9000/api/v1/person/$count --header 'Content-Type: application/json' >> /dev/null
        curl --request PUT --url http://172.33.0.4:9000/api/v1/person/ --header 'Content-Type: application/json' --data '{"name": "NewUser","job": "NewJob","age": 30,"photo": "https://aws.s3.bucket/PHOTO$count.jpg"}' >> /dev/null
    fi
        if [  $(($count % 15)) -eq 0 ]; then
        curl --request DELETE --url http://172.33.0.4:9000/api/v1/person/id --header 'Content-Type: application/json' >> /dev/null
        curl --request PUT --url http://172.33.0.4:9000/api/v1/person/ --header 'Content-Type: application/json' --data '{"name": "User$count""job": "Job$count""age": $count"photo": "https://aws.s3.bucket/PHOTO$count.jpg"}' >> /dev/null
        curl --request GET --url http://172.33.0.4:9000/api/v1/server/ --header 'Content-Type: application/json' >> /dev/null
        curl --request GET --url http://172.33.0.4:9000/api/v1/nothing/ --header 'Content-Type: application/json' >> /dev/null
    fi 
done