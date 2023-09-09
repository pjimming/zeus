new_app:
	goctl api new $(app) --style goZero

api:
	goctl api go -api $(app)/$(app).api -dir $(app) --home template
	goctl api format --dir $(app)/apis -declare
