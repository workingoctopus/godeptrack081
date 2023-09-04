package dtrack_test

import (
	"context"

	"github.com/google/uuid"

	"github.com/DependencyTrack/client-go"
)

// This example demonstrates how to fetch all findings for a given project.
func Example_fetchAllFindings() {
	client, _ := dtrack.NewClient("https://dtrack.example.com", dtrack.WithAPIKey("..."))
	projectUUID := uuid.MustParse("2d16089e-6d3a-437e-b334-f27eb2cbd7f4")

	_, err := dtrack.FetchAll(func(po dtrack.PageOptions) (dtrack.Page[dtrack.Finding], error) {
		return client.Finding.GetAll(context.TODO(), projectUUID, false, po)
	})
	if err != nil {
		panic(err)
	}
}
