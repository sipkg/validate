package messages_test

import (
	"testing"

	"github.com/sipkg/validate/messages"
)

func TestLang(t *testing.T) {
	testCases := []struct {
		desc       string
		lang       string
		original   string
		translated string
	}{
		{
			desc:       "default message",
			lang:       "en",
			original:   "is not a string",
			translated: "is not a string",
		},
		{
			desc:       "french",
			lang:       "fr",
			original:   "is not a string",
			translated: "n'est pas une chaîne de catactères",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			messages.ChangeLang(tc.lang)
			got := messages.Translate(tc.original)
			if tc.translated != got {
				t.Errorf("wait %s got %s", tc.translated, got)
			}
		})
	}
}
