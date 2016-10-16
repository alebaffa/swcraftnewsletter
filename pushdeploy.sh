#!/bin/sh
echo "Commit..."
git add .
git commit -m "commit all files"

echo "remove private..."
rm private/private.go
cp private.go private/

echo "push to github..."
git push origin master

echo "push to heroku..."
git push heroku master