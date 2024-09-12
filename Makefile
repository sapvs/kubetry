clt: svr
	docker build -t clt:0.2 --build-arg MOD=clt .

svr:
	docker build -t svr:0.1 --build-arg MOD=svr .


up: clt
	docker compose up

down:
	docker compose down