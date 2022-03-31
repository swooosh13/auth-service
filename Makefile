
run-app: build-app
    @docker-compose up -d --force-recreate authapp
build-app:
    @docker-compose build --no-cache authapp
