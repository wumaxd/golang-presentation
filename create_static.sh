present &
background_present=$!

echo "$background_present"

rm ./build/*.html
sleep 2

curl -s http://127.0.0.1:3999/static/notes.js > build/static/notes.js
curl -s http://127.0.0.1:3999/static/slides.js | sed -E 's/(\/static\/)/.\1/' > build/static/slides.js
curl -s http://127.0.0.1:3999/static/styles.css > build/static/styles.css
curl -s http://127.0.0.1:3999/static/notes.css > build/static/notes.css

page_array=()

for page in $(cat "page_list.json" | jq -r '.pages[]'); do
  name=$(echo "$page" | sed 's/\.[^.]*$//')
  curl -s http://127.0.0.1:3999/$page | sed -E 's/(\/static\/)/.\1/' | sed -E '/play\.js/d' > build/$name.html
  page_array+=("$name")
done

cp -r ./images ./build

cp index_template.md ./build/index.md

echo "" >> ./build/index.md
echo "" >> ./build/index.md
echo "| Link |" >> ./build/index.md
echo "| ---- |" >> ./build/index.md

for entry in "${page_array[@]}"; do
  echo "| [$entry](./""$entry"".html) |" >> ./build/index.md
done

kill $background_present
