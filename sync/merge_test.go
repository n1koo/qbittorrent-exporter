package sync

import (
	"testing"

	API "qbit-exp/api"
)

func TestMergeInfo_PartialUpdate(t *testing.T) {
	t.Parallel()

	existing := API.Info{
		Hash:     "abc123",
		Name:     "Original Name",
		State:    "downloading",
		Progress: 0.5,
		Dlspeed:  1000000,
		Upspeed:  500000,
		Size:     100000000,
	}

	// Delta only updates state and dlspeed
	newState := "seeding"
	newDlspeed := int64(0)
	delta := API.DeltaInfo{
		State:   &newState,
		Dlspeed: &newDlspeed,
	}

	result := MergeInfo(existing, delta)

	// Updated fields
	if result.State != "seeding" {
		t.Errorf("State: expected %q, got %q", "seeding", result.State)
	}
	if result.Dlspeed != 0 {
		t.Errorf("Dlspeed: expected 0, got %d", result.Dlspeed)
	}

	// Unchanged fields
	if result.Hash != "abc123" {
		t.Errorf("Hash: expected %q, got %q", "abc123", result.Hash)
	}
	if result.Name != "Original Name" {
		t.Errorf("Name: expected %q, got %q", "Original Name", result.Name)
	}
	if result.Progress != 0.5 {
		t.Errorf("Progress: expected 0.5, got %f", result.Progress)
	}
	if result.Upspeed != 500000 {
		t.Errorf("Upspeed: expected 500000, got %d", result.Upspeed)
	}
	if result.Size != 100000000 {
		t.Errorf("Size: expected 100000000, got %d", result.Size)
	}
}

func TestMergeInfo_EmptyDelta(t *testing.T) {
	t.Parallel()

	existing := API.Info{
		Hash:     "abc123",
		Name:     "Test",
		State:    "downloading",
		Progress: 0.75,
		Dlspeed:  2000000,
	}

	delta := API.DeltaInfo{} // All fields nil

	result := MergeInfo(existing, delta)

	// All fields should be unchanged
	if result.Hash != existing.Hash {
		t.Errorf("Hash changed unexpectedly")
	}
	if result.Name != existing.Name {
		t.Errorf("Name changed unexpectedly")
	}
	if result.State != existing.State {
		t.Errorf("State changed unexpectedly")
	}
	if result.Progress != existing.Progress {
		t.Errorf("Progress changed unexpectedly")
	}
	if result.Dlspeed != existing.Dlspeed {
		t.Errorf("Dlspeed changed unexpectedly")
	}
}

func TestMergeInfo_ZeroValues(t *testing.T) {
	t.Parallel()

	existing := API.Info{
		Hash:    "abc123",
		Dlspeed: 1000000,
		Upspeed: 500000,
		Eta:     3600,
	}

	// Setting fields to zero (not nil)
	zero := int64(0)
	delta := API.DeltaInfo{
		Dlspeed: &zero,
		Upspeed: &zero,
		Eta:     &zero,
	}

	result := MergeInfo(existing, delta)

	if result.Dlspeed != 0 {
		t.Errorf("Dlspeed: expected 0, got %d", result.Dlspeed)
	}
	if result.Upspeed != 0 {
		t.Errorf("Upspeed: expected 0, got %d", result.Upspeed)
	}
	if result.Eta != 0 {
		t.Errorf("Eta: expected 0, got %d", result.Eta)
	}
}

func TestMergeInfo_AllFields(t *testing.T) {
	t.Parallel()

	existing := API.Info{
		Hash:              "abc123",
		AddedOn:           1000,
		AmountLeft:        50000,
		Category:          "old-cat",
		Comment:           "old comment",
		CompletionOn:      0,
		Dlspeed:           1000,
		Downloaded:        25000,
		DownloadedSession: 10000,
		Eta:               3600,
		MaxRatio:          1.0,
		Name:              "Old Name",
		NumLeechs:         5,
		NumSeeds:          10,
		Progress:          0.5,
		Ratio:             0.5,
		SavePath:          "/old/path",
		Size:              100000,
		State:             "downloading",
		Tags:              "old,tags",
		Tracker:           "http://old.tracker",
		TimeActive:        7200,
		Uploaded:          12500,
		UploadedSession:   5000,
		Upspeed:           500,
	}

	// Create delta with all fields set
	addedOn := int64(2000)
	amountLeft := int64(0)
	category := "new-cat"
	comment := "new comment"
	completionOn := int64(3000)
	dlspeed := int64(0)
	downloaded := int64(100000)
	downloadedSession := int64(50000)
	eta := int64(0)
	maxRatio := 2.0
	name := "New Name"
	numLeechs := int64(0)
	numSeeds := int64(20)
	progress := 1.0
	ratio := 1.0
	savePath := "/new/path"
	size := int64(100000)
	state := "seeding"
	tags := "new,tags"
	tracker := "http://new.tracker"
	timeActive := int64(14400)
	uploaded := int64(100000)
	uploadedSession := int64(50000)
	upspeed := int64(1000)

	delta := API.DeltaInfo{
		AddedOn:           &addedOn,
		AmountLeft:        &amountLeft,
		Category:          &category,
		Comment:           &comment,
		CompletionOn:      &completionOn,
		Dlspeed:           &dlspeed,
		Downloaded:        &downloaded,
		DownloadedSession: &downloadedSession,
		Eta:               &eta,
		MaxRatio:          &maxRatio,
		Name:              &name,
		NumLeechs:         &numLeechs,
		NumSeeds:          &numSeeds,
		Progress:          &progress,
		Ratio:             &ratio,
		SavePath:          &savePath,
		Size:              &size,
		State:             &state,
		Tags:              &tags,
		Tracker:           &tracker,
		TimeActive:        &timeActive,
		Uploaded:          &uploaded,
		UploadedSession:   &uploadedSession,
		Upspeed:           &upspeed,
	}

	result := MergeInfo(existing, delta)

	// Hash should be unchanged (not in delta)
	if result.Hash != "abc123" {
		t.Errorf("Hash should not change")
	}

	// All other fields should be updated
	if result.AddedOn != 2000 {
		t.Errorf("AddedOn: expected 2000, got %d", result.AddedOn)
	}
	if result.AmountLeft != 0 {
		t.Errorf("AmountLeft: expected 0, got %d", result.AmountLeft)
	}
	if result.Category != "new-cat" {
		t.Errorf("Category: expected %q, got %q", "new-cat", result.Category)
	}
	if result.Comment != "new comment" {
		t.Errorf("Comment: expected %q, got %q", "new comment", result.Comment)
	}
	if result.CompletionOn != 3000 {
		t.Errorf("CompletionOn: expected 3000, got %d", result.CompletionOn)
	}
	if result.Dlspeed != 0 {
		t.Errorf("Dlspeed: expected 0, got %d", result.Dlspeed)
	}
	if result.Downloaded != 100000 {
		t.Errorf("Downloaded: expected 100000, got %d", result.Downloaded)
	}
	if result.DownloadedSession != 50000 {
		t.Errorf("DownloadedSession: expected 50000, got %d", result.DownloadedSession)
	}
	if result.Eta != 0 {
		t.Errorf("Eta: expected 0, got %d", result.Eta)
	}
	if result.MaxRatio != 2.0 {
		t.Errorf("MaxRatio: expected 2.0, got %f", result.MaxRatio)
	}
	if result.Name != "New Name" {
		t.Errorf("Name: expected %q, got %q", "New Name", result.Name)
	}
	if result.NumLeechs != 0 {
		t.Errorf("NumLeechs: expected 0, got %d", result.NumLeechs)
	}
	if result.NumSeeds != 20 {
		t.Errorf("NumSeeds: expected 20, got %d", result.NumSeeds)
	}
	if result.Progress != 1.0 {
		t.Errorf("Progress: expected 1.0, got %f", result.Progress)
	}
	if result.Ratio != 1.0 {
		t.Errorf("Ratio: expected 1.0, got %f", result.Ratio)
	}
	if result.SavePath != "/new/path" {
		t.Errorf("SavePath: expected %q, got %q", "/new/path", result.SavePath)
	}
	if result.Size != 100000 {
		t.Errorf("Size: expected 100000, got %d", result.Size)
	}
	if result.State != "seeding" {
		t.Errorf("State: expected %q, got %q", "seeding", result.State)
	}
	if result.Tags != "new,tags" {
		t.Errorf("Tags: expected %q, got %q", "new,tags", result.Tags)
	}
	if result.Tracker != "http://new.tracker" {
		t.Errorf("Tracker: expected %q, got %q", "http://new.tracker", result.Tracker)
	}
	if result.TimeActive != 14400 {
		t.Errorf("TimeActive: expected 14400, got %d", result.TimeActive)
	}
	if result.Uploaded != 100000 {
		t.Errorf("Uploaded: expected 100000, got %d", result.Uploaded)
	}
	if result.UploadedSession != 50000 {
		t.Errorf("UploadedSession: expected 50000, got %d", result.UploadedSession)
	}
	if result.Upspeed != 1000 {
		t.Errorf("Upspeed: expected 1000, got %d", result.Upspeed)
	}
}

func TestMergeInfo_DoesNotMutateOriginal(t *testing.T) {
	t.Parallel()

	existing := API.Info{
		Hash:    "abc123",
		Name:    "Original",
		Dlspeed: 1000,
	}

	newName := "Updated"
	delta := API.DeltaInfo{
		Name: &newName,
	}

	result := MergeInfo(existing, delta)

	// Result should have new name
	if result.Name != "Updated" {
		t.Errorf("Result name: expected %q, got %q", "Updated", result.Name)
	}

	// Original should be unchanged (Go passes struct by value, so this is automatic)
	if existing.Name != "Original" {
		t.Errorf("Original name was mutated: got %q", existing.Name)
	}
}
