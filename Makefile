build:
	docker buildx build -t esilva-printer .

clean:
	docker container rm esilva-printer-container-test --force

attach:
	docker exec -u root -it esilva-printer-container zsh

attach-test:
	docker exec -u root -it esilva-printer-container-test zsh

test: clean
	docker run \
		--name esilva-printer-container-test \
		-p 9099:9099 \
		-it esilva-printer zsh

debug: clean test
