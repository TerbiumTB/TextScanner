start:
	docker-compose build
	$(MAKE) run

stop:
	docker-compose down

run:
	docker-compose up -d
restart:
	$(MAKE) stop
	$(MAKE) start

