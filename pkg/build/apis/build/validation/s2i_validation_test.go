package validation

import (
	"net/url"
	"testing"

	_ "github.com/openshift/openshift-apiserver/pkg/build/apis/build/install"
)

func TestParseGitURL(t *testing.T) {
	testCases := map[string]URL{
		"ssh://git@github.com:example/my-example-repo": URL{
			URL: url.URL{
				Scheme:     "",
				Opaque:     "",
				User:       url.User("git"),
				Host:       "github.com",
				Path:       "example/my-example-repo",
				RawPath:    "",
				ForceQuery: false,
				RawQuery:   "",
				Fragment:   ""},
			Type: URLTypeSCP,
		},
		"http://github.com/example/my-example-repo": URL{
			URL: url.URL{
				Scheme:     "http",
				Opaque:     "",
				User:       nil,
				Host:       "github.com",
				Path:       "/example/my-example-repo",
				RawPath:    "",
				ForceQuery: false,
				RawQuery:   "",
				Fragment:   ""},
			Type: URLTypeURL,
		},
		"https://github.com/example/my-example-repo": URL{
			URL: url.URL{
				Scheme:     "https",
				Opaque:     "",
				User:       nil,
				Host:       "github.com",
				Path:       "/example/my-example-repo",
				RawPath:    "",
				ForceQuery: false,
				RawQuery:   "",
				Fragment:   ""},
			Type: URLTypeURL,
		},
		"/example/my-example-repo": URL{
			URL: url.URL{
				Scheme:     "",
				Opaque:     "",
				User:       nil,
				Host:       "",
				Path:       "/example/my-example-repo",
				RawPath:    "",
				ForceQuery: false,
				RawQuery:   "",
				Fragment:   ""},
			Type: URLTypeLocal,
		},
		"file:///example/my-example-repo": URL{
			URL: url.URL{
				Scheme:     "file",
				Opaque:     "",
				User:       nil,
				Host:       "",
				Path:       "/example/my-example-repo",
				RawPath:    "",
				ForceQuery: false,
				RawQuery:   "",
				Fragment:   ""},
			Type: URLTypeLocal,
		},
		"~/example/my-example-repo": URL{
			URL: url.URL{
				Scheme:     "",
				Opaque:     "",
				User:       nil,
				Host:       "",
				Path:       "~/example/my-example-repo",
				RawPath:    "",
				ForceQuery: false,
				RawQuery:   "",
				Fragment:   ""},
			Type: URLTypeLocal,
		},
	}

	for testURL, wanted := range testCases {
		url, err := parseGitURL(testURL)
		if err != nil {
			t.Errorf("error occurred parsing git url %s: %#v", testURL, err.Error())
		}
		if !compareURLs(*url, wanted) {
			t.Errorf("git url %q was not parsed correctly wanted: %#v, got %#v", testURL, wanted, *url)
		}

	}

}

func compareURLs(v1, v2 URL) bool {

	if v1.Type != v2.Type {
		return false
	}

	if v1.URL.Scheme != v2.URL.Scheme {
		return false
	}

	if v1.URL.Opaque != v2.URL.Opaque {
		return false
	}

	if v1.URL.Host != v2.URL.Host {
		return false
	}

	if v1.URL.Path != v2.URL.Path {
		return false
	}

	if v1.URL.RawPath != v2.URL.RawPath {
		return false
	}

	if v1.URL.ForceQuery != v2.URL.ForceQuery {
		return false
	}

	if v1.URL.RawQuery != v2.URL.RawQuery {
		return false
	}

	if v1.URL.Fragment != v2.URL.Fragment {
		return false
	}

	if v1.URL.User.Username() != v2.URL.User.Username() {
		return false
	}

	v1Password, v1PasswordSet := v1.URL.User.Password()
	v2Password, v2PasswordSet := v2.URL.User.Password()

	if v1PasswordSet != v2PasswordSet {
		return false
	}

	if v1Password != v2Password {
		return false
	}

	return true
}
