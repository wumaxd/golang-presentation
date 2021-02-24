present &
background_present=$!

echo "$background_present"

rm ./docs/*.html
sleep 2

curl -s http://127.0.0.1:3999/static/notes.js > docs/static/notes.js
curl -s http://127.0.0.1:3999/static/slides.js | sed -E 's/(\/static\/)/.\1/' > docs/static/slides.js
curl -s http://127.0.0.1:3999/static/styles.css > docs/static/styles.css
curl -s http://127.0.0.1:3999/static/notes.css > docs/static/notes.css

page_array=()

for page in $(cat "page_list.json" | jq -r '.pages[]'); do
  name=$(echo "$page" | sed 's/\.[^.]*$//')
  curl -s http://127.0.0.1:3999/$page | sed -E 's/(\/static\/)/.\1/' | sed -E '/play\.js/d' > docs/$name.html
  page_array+=("$name")
done

cp -r ./images ./docs

cp index_template.md ./docs/index.md

echo "" >> ./docs/index.md
echo "" >> ./docs/index.md

for entry in "${page_array[@]}"; do
  echo "- [$entry](./""$entry"".html)" >> ./docs/index.md
done

kill $background_present
