.PHONY: docker-up
docker-up:
	docker-compose -f docker-compose.yaml up --build

.PHONY: docker-down
docker-down: ## Stop docker containers and clear artefacts.
	docker-compose -f docker-compose.yaml down
	docker system prune 

.PHONY: bundle
bundle: ## bundles the submission for... submission
	git bundle create guestlist.bundle --all

.PHONY: integration-test
integration-test:
	docker-compose -f docker-compose.test.yaml up -d
# sleep 30s to make sure the mysql container started
# possible to move the go test stage to docker too so it can keep retry until the mysql started
	sleep 30
	go test ./... --tags integration
	docker-compose -f docker-compose.test.yaml down

.PHONY: unit-test
unit-test:
	go test ./...