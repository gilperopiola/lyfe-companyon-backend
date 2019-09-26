all: run
run:
	go run server.go router.go common.go cron.go auth.go auth_controller.go users.go users_controller.go users_model.go tags.go tags_controller.go tags_model.go tasks.go tasks_controller.go tasks_model.go --env=$(env)

get:
	go get github.com/gin-gonic/gin \
	go get github.com/gin-contrib/cors \
	go get github.com/dgrijalva/jwt-go \
	go get github.com/go-sql-driver/mysql \
	go get github.com/spf13/viper \