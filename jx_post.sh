#/bin/sh
echo 'add changes'
git add .
echo 'commit and add comments'
git commit -m 'post blog and update README.md(commit by script)'
echo 'push to master'
git push origin master
echo 'Done'
