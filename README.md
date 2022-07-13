#Drive

a simple net-disk project backend

a trial of building web server with [ent, logrus, golang-migrate, viper, gin, ...]

## Config
must have a file named "config.yaml" in project's root directory
```yaml
db:
  driver: postgres
  host: your_host
  port: your_port
  username: your_username
  password: your_password
  database: your_dbname

redis:
  host: your_host
  port: your_port
  password: your_password_if_exist

minio:
  host: your_host:your_port
  access_key: your_access_key
  secret_key: your_secret_key
```

## Ent
after change in ./ent/schema
run go generate
```bash
# pwd: xxx/<project>
go generate ./ent
```

## Migration
```
# try to generate new migration files
err = migrate.NamedDiff(ctx, url, "update", opts...)

# try to migrate, but often error due to no change
m, err := golangMigrate.New("file://migrations", url)
m.Up()
```
