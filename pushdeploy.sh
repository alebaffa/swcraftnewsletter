#!/bin/sh
echo "Commit..."
git add .
git commit -m "commit all files"

echo "push to heroku..."
git push heroku master

echo "replace private..."
cp private.go private/

echo "push to github..."
git push origin master