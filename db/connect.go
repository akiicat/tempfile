package db

import (
  "context"
  "fmt"

  "cloud.google.com/go/datastore"
)

type datastoreDB struct {
	client *datastore.Client
}

//// Ensure datastoreDB conforms to the FileDatabase interface.
var _ FileDatabase = &datastoreDB{}

// newDatastoreDB creates a new FileDatabase backed by Cloud Datastore.
// See the datastore and google packages for details on creating a suitable Client:
// https://godoc.org/cloud.google.com/go/datastore
func NewDatastoreDB(projectID string) (FileDatabase, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	// Verify that we can communicate and authenticate with the datastore service.
	t, err := client.NewTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	if err := t.Rollback(); err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	return &datastoreDB{
		client: client,
	}, nil
}

func (db *datastoreDB) ListFilesByToken(token string) (*File, error) {
	ctx := context.Background()
	if token == "" {
		return nil, fmt.Errorf("parameter token is empty")
	}

	files := make([]*File, 0)
	q := datastore.NewQuery("File").
		Filter("Token =", token).
		Order("-CreatedAt").
    Limit(1)

	keys, err := db.client.GetAll(ctx, q, &files)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list files: %v", err)
	}

  if len(keys) == 0 {
		return nil, fmt.Errorf("Token Not Found: %s", token)
  }

	return files[0], nil
}

// AddFile saves a given book, assigning it a new ID.
func (db *datastoreDB) AddFile(b *File) (id int64, err error) {
	ctx := context.Background()
	k := datastore.IncompleteKey("File", nil)
	k, err = db.client.Put(ctx, k, b)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put File: %v", err)
	}
	return k.ID, nil
}


