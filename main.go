package main

import (
	"context"
	"fmt"
	"log"

	"ent_sample/ent"
	"ent_sample/ent/user"

	entsql "entgo.io/ent/dialect/sql"
	"go.elastic.co/apm/module/apmsql/v2"
	_ "go.elastic.co/apm/module/apmsql/v2/pq"
)

func main() {
	db, err := apmsql.Open("postgres", "host=127.0.0.1 port=5432 user=test dbname=test password=test sslmode=disable")

	// opts := []ent.Option{
	// 	ent.Debug(),
	// }
	drv := entsql.OpenDB("postgresql", db)
	client := ent.NewClient(ent.Driver(drv))

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.TODO()

	// GenerateSchema(ctx, client)

	ctx = context.WithValue(ctx, "key", "val")
	createUser(ctx, client)
	user, err := getUser(ctx, client)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	fmt.Println(user.Edges.UserInfos)
}

func GenerateSchema(ctx context.Context, client *ent.Client) {
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func createUser(ctx context.Context, client *ent.Client) (*ent.User, error) {

	user, err := client.User.Create().SetName("sample").Save(ctx)

	client.UserInfo.Create().SetUsers(user).SetEtc("etc-1").ExecX(ctx)

	return user, err
}

func getUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	return client.User.Query().Order(ent.Desc(user.FieldID)).WithUserInfos().First(ctx)
}
