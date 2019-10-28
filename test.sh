#!/bin/sh

export URL=http://127.0.0.1:8080

echo "- List:"
curl $URL/list
echo ""

echo "- Add todo:"
curl -X POST $URL/add/Do%20This
echo ""

echo "- Add another todo:"
curl -X POST $URL/add/Do%20That
echo ""

echo "- List:"
curl $URL/list
echo ""

echo "- Delete first one:"
curl -X DELETE $URL/remove/1
echo ""

echo "- List:"
curl $URL/list
echo ""

echo "- Delete non-esisting one:"
curl -X DELETE $URL/remove/10
echo ""
