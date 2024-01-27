IMAGE = "ssh_sudo_with_password"

run:
	docker build --platform linux/amd64 -t $(IMAGE) .
	docker run -p 2222:22 --name $(IMAGE) -d $(IMAGE)
	sleep 3
	@echo
	gobble run
	@echo
	@make cleanup

cleanup:
	docker rm -f $(IMAGE)
