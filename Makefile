# Сборка
# $ docker build -t prod-server:local -f DockerFile .

run:
	docker run -p 3000:3333 prod-server:local