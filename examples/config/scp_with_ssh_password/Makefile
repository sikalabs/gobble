IMAGE = "scp_with_ssh_password"

run:
	docker build --platform linux/amd64 -t $(IMAGE) .
	docker run -p 2222:22 --name $(IMAGE) -d $(IMAGE)
	sleep 3
	@echo
	gobble run
	@echo
	sshpass -p demo ssh ssh://demo@127.0.0.1:2222 cat /hello.txt
	sshpass -p demo ssh ssh://demo@127.0.0.1:2222 cat /template.txt
	@echo
	@make cleanup

cleanup:
	docker rm -f $(IMAGE)
