curl -X POST http://localhost:8080/apply -d '{"fullName":"Student A","mathScore":5,"informaticsScore":5,"englishScore":5}' -H "Content-Type: application/json"
curl -X POST http://localhost:8080/apply -d '{"fullName":"Student B","mathScore":4,"informaticsScore":5,"englishScore":5}' -H "Content-Type: application/json"
curl -X POST http://localhost:8080/apply -d '{"fullName":"Student C","mathScore":2,"informaticsScore":4,"englishScore":3}' -H "Content-Type: application/json"

curl http://localhost:8080/admitted
