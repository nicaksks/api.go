package database

import "context"

func Delete(collection string, data any) error {
	client, ctx := StartConnect()
	defer client.Disconnect(ctx)

	c := client.Database("Anime").Collection(collection)

	_, err := c.DeleteOne(context.Background(), data)

	return err
}
