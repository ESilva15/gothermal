build:
	docker buildx build -t thprinter .

clean:
	docker container rm thprinter-container-test --force

attach:
	docker exec -u root -it thprinter-container zsh

attach-test:
	docker exec -u root -it thprinter-container-test zsh

test: clean
	docker run \
		--name thprinter-container-test \
		-p 9099:9099 \
		-it thprinter zsh

debug: clean test
