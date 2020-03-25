#/bin/sh
echo '1. add all changes'
git add .

echo '2. commit and add comments'
read -p 'enter commit message(no -):'
git commit -m "$REPLY"

echo '3. push to master'
git push origin master

echo '4. Done'