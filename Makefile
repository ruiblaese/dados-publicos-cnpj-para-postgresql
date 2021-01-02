build:
	docker build --pull --rm -f "Dockerfile" -t ruiblaese/cnpj:latest "."

push:
	docker push ruiblaese/cnpj:latest

clean:
	docker-compose down
	docker rm cnpj_app_1
	docker rm cnpj_db_1
	docker volume rm cnpj_dados

postgres-temp:
	docker run \
		--rm \
		--name teste-postgres \
		-p 5432:5432 \
		-e "POSTGRES_USER=postgres" \
		-e "POSTGRES_PASSWORD=postgres" \
		-e "POSTGRES_DB=empresas" \
		-e "sslmode=disable" \
		-e "TimeZone=America/Sao_Paulo" \
		postgres:13 \
