.PHONY: checkhash gethash createhash

checkhash:
	curl -X POST -H "Content-Type: text/plain" -d "Hello, world!" http://localhost:8080/checkhash

gethash:
	curl -X POST -H "Content-Type: text/plain" -d "Hello, world!" http://localhost:8080/gethash

createhash:
	curl -X POST -H "Content-Type: text/plain" -d "Hello, world!" http://localhost:8080/createhash
