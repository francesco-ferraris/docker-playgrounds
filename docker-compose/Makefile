.PHONY: run
run:
	@docker compose up

.PHONY: clean
clean:
	@docker compose down

.PHONY: run-and-watch
run-and-watch:
	@docker compose watch

.DEFAULT_GOAL := run-and-watch