#!/bin/bash

apk add --no-cache curl jq git zip
apt-get update -qq \
    && apt-get install -y -q --no-install-recommends \
        curl \
        jq

EVENT_DATA=$(cat $GITHUB_EVENT_PATH)
echo $EVENT_DATA | jq .
UPLOAD_URL=$(echo $EVENT_DATA | jq -r .release.upload_url)
UPLOAD_URL=${UPLOAD_URL/\{?name,label\}/}
RELEASE_NAME=$(echo $EVENT_DATA | jq -r .release.tag_name)
PROJECT_NAME=$(basename $GITHUB_REPOSITORY)
NAME="${NAME:-${PROJECT_NAME}_${RELEASE_NAME}}_${GOOS}_${GOARCH}"

EXT=''
if [ $GOOS == 'windows' ]; then
EXT='.exe'
fi
FILE_LIST="${PROJECT_NAME}${EXT}"

echo $FILE_LIST

FILE_LIST="${FILE_LIST} ${EXTRA_FILES}"

FILE_LIST=`echo "${FILE_LIST}" | awk '{$1=$1};1'`


if [ $GOOS == 'windows' ]; then
ARCHIVE=tmp.zip
zip -9r $ARCHIVE ${FILE_LIST}
else
ARCHIVE=tmp.tgz
tar cvfz $ARCHIVE ${FILE_LIST}
fi

CHECKSUM=$(md5sum ${ARCHIVE} | cut -d ' ' -f 1)

curl \
  -X POST \
  --data-binary @${ARCHIVE} \
  -H 'Content-Type: application/octet-stream' \
  -H "Authorization: Bearer ${GITHUB_TOKEN}" \
  "${UPLOAD_URL}?name=${NAME}.${ARCHIVE/tmp./}"

curl \
  -X POST \
  --data $CHECKSUM \
  -H 'Content-Type: text/plain' \
  -H "Authorization: Bearer ${GITHUB_TOKEN}" \
  "${UPLOAD_URL}?name=${NAME}_checksum.txt"
