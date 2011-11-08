#/usr/bin/sh

rm -rv ./static
cp -rv ../static ./
echo "var tourMode = \"appengine\";" > ./static/mode.js
~/bin/google_appengine/appcfg.py update ./
