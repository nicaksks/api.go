package database

import "context"

func Insert(collection string, data any) error {
	client, ctx := StartConnect()
	defer client.Disconnect(ctx)

	c := client.Database("Anime").Collection(collection)

	_, err := c.InsertOne(context.Background(), data)

	return err
}
