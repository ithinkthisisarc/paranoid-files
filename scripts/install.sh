cp "$1" "$1_cp"
sudo mv "./$1" /usr/bin
mv "$1_cp" "$1"
