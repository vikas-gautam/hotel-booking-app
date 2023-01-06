POSTGRES_PASSWORD ?= password

run:
	docker run -itd -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -p 5432:5432 -v $(pwd)/data:/var/lib/postgresql/data --name postgresql postgres