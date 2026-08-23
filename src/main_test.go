package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestCreateStars(t *testing.T) {
    stars := createStars()
    if len(stars) != starCount {
        t.Fatalf("expected %d stars, got %d", starCount, len(stars))
    }
    for i, s := range stars {
        if s.z < 0 || s.z > float32(screenWidth) {
            t.Fatalf("star %d z out of range: %v", i, s.z)
        }
    }
}



func FuzzNewApp(f *testing.F) {
	

	
}
