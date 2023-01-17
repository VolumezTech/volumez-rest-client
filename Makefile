build-client:
	@echo "\n🛠️ Building vlz-api-client module..."
	rm -rf restapi ### remove the old model
	rm -rf client ### remove the old client
	swagger generate client -A volumez -m restapi/models -f restapi.json -c client