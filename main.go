package main

import (
	"context"
	"fmt"
	"log"

	"ent_sample/ent"
	"ent_sample/ent/user"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=test dbname=test password=test sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx := context.TODO()
	ctx = context.WithValue(ctx, "key", "val")
	createUser(ctx, client)
	user, err := getUser(ctx, client)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	fmt.Println(user.Edges.UserInfos)
}

func createUser(ctx context.Context, client *ent.Client) (*ent.User, error) {

	user, err := client.User.Create().SetName("sample").Save(ctx)

	client.UserInfo.Create().SetUsers(user).SetEtc("etc-1").ExecX(ctx)

	return user, err
}

func getUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	return client.User.Query().Order(ent.Desc(user.FieldID)).WithUserInfos().First(ctx)
}
