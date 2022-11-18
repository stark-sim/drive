package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"drive/config"
	"drive/db"
	"drive/logger"
	"drive/services"
	"drive/tools"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"

	"drive/ent"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 初始化日志
	logger.InitLog()

	// 初始化配置文件
	err := config.Init()
	if err != nil {
		return
	}

	// 初始化通用工具：例如雪花 ID 生成器
	err = tools.Init()
	if err != nil {
		return
	}

	// 数据库准备工作：例如迁移
	db.Init()

	// 业务服务层初始化工作：例如建立数据库连接供各个业务使用
	services.Init()

	client := db.NewDBClient()

	//// Create ent.Client and run the schema migration.
	//client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	//if err != nil {
	//	log.Fatal("opening ent client", err)
	//}
	//if err := client.Schema.Create(
	//	context.Background(),
	//	migrate.WithGlobalUniqueID(true),
	//); err != nil {
	//	log.Fatal("opening ent client", err)
	//}

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(ent.NewSchema(client))
	// Transaction, optional for Isolation Levels
	srv.Use(entgql.Transactioner{
		TxOpener: entgql.TxOpenerFunc(func(ctx context.Context) (context.Context, driver.Tx, error) {
			tx, err := client.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
			if err != nil {
				return nil, nil, err
			}
			ctx = ent.NewTxContext(ctx, tx)
			ctx = ent.NewContext(ctx, tx.Client())
			return ctx, tx, nil
		}),
	})
	http.Handle("/",
		playground.Handler("Test", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
