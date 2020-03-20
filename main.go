package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
)

func main() {
	// Defining the input
	if len(os.Args) != 3 {
		log.Fatal("usage:upload BUCKET_UR_FILE")

	}
	bucketURL := os.Args[1]
	file := os.Args[2]
	_, _ = bucketURL, file

	// Open a connection to the bucket
	b, err := blob.OpenBucket(context.Background(), bucketURL)
	if err != nil {
		log.Fatalf("Failed to setup bucket: %s", err)
	}
	defer b.close()

	// prepare the file for upload
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)

	}
	// Writing the file to the bucket
	w, err := b.NewWriter(ctx, file, nil)
	if err != nil {
		log.Fatalf("Failed to obtain writer: %s", err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("Failed to write to bucket: %s", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("Failed to close: %s", err)
	}
}
