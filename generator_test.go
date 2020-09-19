package main

import "testing"

var tests = []struct {
	id      int
	url     string
	correct bool
}{
	{
		id:      -1,
		url:     "",
		correct: false,
	},
	{
		id:      0,
		url:     string(charSet[0]),
		correct: true,
	},
	{
		id:      162,
		url:     string(charSet[162%len(charSet)]) + string(charSet[(162/len(charSet))%len(charSet)]),
		correct: true,
	},
}

func TestGenerateShortURL(t *testing.T) {
	for i, test := range tests {
		url, err := GenerateShortURL(test.id)
		if test.correct {
			if err != nil {
				t.Errorf("Incorrect behavior on %v with num:%v. Got error %w ,but expected correct work", test, i, err)
			} else {
				if url != test.url {
					t.Errorf("Incorrect behavior on %v with num:%v. Got incorrect answer: %v", test, i, url)
				}
			}
		} else {
			if err == nil {
				t.Errorf("Incorrect behavior on %v with num:%v. Expected error,but got nil", test, i)
			}
		}
	}
}
