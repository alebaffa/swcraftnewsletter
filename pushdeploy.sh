#!/bin/sh
echo "replace private..."
cp private.go private/

echo "Commit..."
git add .
git commit -m "commit all files"

echo "push to github..."
git push origin master