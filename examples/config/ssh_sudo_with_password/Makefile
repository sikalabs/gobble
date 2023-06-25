IMAGE = "ssh_sudo_with_password"

run:
	docker build --platform linux/amd64 -t $(IMAGE) .
	docker run -p 2223:22 --name $(IMAGE) -d $(IMAGE)
