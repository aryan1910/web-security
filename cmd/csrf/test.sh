curl -c cookies.txt http://localhost:8080/
curl -b cookies.txt \                     
     -X POST http://localhost:8080/submit \
     -d "name=Alice&_csrf=$(grep -oP '_csrf=\K[^;]+' cookies.txt)"