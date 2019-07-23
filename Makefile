

run:
	@docker build  -t cookie-setter .
	@docker run   --rm -d --network host cookie-setter