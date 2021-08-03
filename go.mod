module go-blog

go 1.13

replace (
	local.pkg/routes => ./routes
	local.pkg/ctrl => ./ctrl
	local.pkg/db => ./db
)

require (
	github.com/go-gosql-driver/mysql v1.5.0
	github.com/google/logger v1.1.0
	github.com/joho/godotenv v1.3.0
	local.pkg/routes v1.0.0
	local.pkg/ctrl v1.0.0
	local.pkg/db v1.0.0
)
