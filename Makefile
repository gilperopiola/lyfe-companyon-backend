all: run
run:
	go run server.go router.go common.go auth.go auth_controller.go users.go users_controller.go users_model.go tags.go tags_controller.go tags_model.go tasks.go tasks_controller.go tasks_model.go --env=$(env)