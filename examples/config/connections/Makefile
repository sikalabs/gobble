all:
	make down
	make up
	sleep 3
	make test
	make see
	make down

up:
	docker compose up -d --build

test:
	gobble run

see:
	docker compose exec ssh_with_password_sudo_with_password cat /now
	docker compose exec ssh_without_password_sudo_with_password cat /now

down:
	docker compose down
