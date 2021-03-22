package gopherav

import "testing"

func TestOpen(t *testing.T) {
	f, err := Open("testsrc2.mp4", nil)
	if err != nil {
		t.Fatalf("Got error opening file: %s", err)
	}
	defer f.Close();

	if err := f.InitStreamInfo(nil); err != nil {
		t.Fatalf("InitStreamInfo() got error: %s", err)
	}
}
