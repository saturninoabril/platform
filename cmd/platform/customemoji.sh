#!/bin/bash

declare -a files=("cat001.jpeg"
                 "cat002.jpeg"
                 "cat003.jpeg"
                 "cat004.jpeg"
                 "cat005.jpeg"
                 "cat006.jpeg"
                 "cat007.jpeg"
                 "cat008.jpeg"
                 "cat009.jpeg"
                 "cat010.jpeg"
                 "catgif001.gif"
                 "catgif002.gif"
                 "catgif003.gif"
                 "catgif004.gif"
                 "catgif005.gif"
                 "catgif006.gif"
                 "catgif007.gif"
                 "catgif008.gif"
                 "catgif009.gif"
                 "catgif010.gif"
                )

for i in {6..4000}
do
    NAME=$i
    echo $NAME
    TOKEN="k5h7e9yadt8pbmd8y56yd4xo8w"
    echo $TOKEN
    IMAGE="@custom_emojis/${files[$((i % 20))]}"
    echo $IMAGE
    FILENAME="${files[$((i % 20))]}"
    echo $FILENAME
    EMOJI='{"name":"'"$NAME"'","creator_id":"s7ofurh883nfdm8u45sgkhdxfy"}'
    echo $EMOJI
    # curl -i -X POST -H "Authorization: Bearer $TOKEN" -H "Content-Type: multipart/form-data" -F "image=$IMAGE" -F "filename=$FILENAME" -F "emoji=$EMOJI" http://localhost:8065/api/v4/emoji
    curl -i -X POST -H "Authorization: Bearer $TOKEN" -H "Content-Type: multipart/form-data" -F "image=$IMAGE" -F "filename=$FILENAME" -F "emoji=$EMOJI" -k https://ci-linux-mysql-prev.mattermost.com/api/v4/emoji
done

# 1. Run below command first then get the token and user_id from the cookies.txt

# curl -c cookies.txt -d '{"login_id":"saturn","password":"saturn"}' http://localhost:8065/api/v4/users/login
# curl -c cookies.txt -d '{"login_id":"xxx","password":"xxx"}' -k https://ci-linux-mysql-prev.mattermost.com/api/v4/users/login
