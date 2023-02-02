run:
	docker-compose up -d

down:
	docker-compose down

mailserver:
	docker run -itd -p 8025:8025 -p 1025:1025 --name mail-server mailhog/mailhog
