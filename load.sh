#!/bin/sh

# Input is taken from http://parltrack.euwiki.org/dumps/eurlex.json.xz
#
# Since every document in this json is on its own line, just remove
# everything before the first doc and everything after the last one
# (that should be a ']' on its own) and save it as eurlex2.json.
#
# After that this script takes documents and pushes them on your local
# couchdb instance.

cat eurlex2.json | sed 's/^,$//' | while read line
do
  if [ -n "$line" ]; then
    echo $line | curl -XPOST -H 'Content-Type: application/json' http://localhost:5984/test_jq -d @-
  fi
done
