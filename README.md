# Go Api Skeleton
Skeleton for a go rest api.

- Set configuration for database
```yaml
// config.yaml
database:
  host: "localhost"
  password: "nopass"
  name: "todo_ent"
  user: "root"
  port: "3308"
```
- Run the schema migrations
```shell script
go run main.go database migrate --force
```
- Run the server
```shell script
go run main.go serve -p 8000
``` 