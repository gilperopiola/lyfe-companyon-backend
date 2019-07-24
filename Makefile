all: run
run:
	go run server.go router.go auth.go auth_controller.go users.go --env=$(env)