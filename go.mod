module go-blog

go 1.16

replace (
	local.pkg/ctrl => ./ctrl
	local.pkg/db => ./db
	local.pkg/routes => ./routes
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/joho/godotenv v1.3.0
	local.pkg/ctrl v1.0.0 // indirect
	local.pkg/db v1.0.0 // indirect
	local.pkg/routes v1.0.0
)
