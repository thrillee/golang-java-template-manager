package internals

var repoStore = map[string]string{
	"simple-web": "https://bitbucket.org/blackcoders2019/simple_web.git",
}

func getRepoURL(repo string) string {
	url, ok := repoStore[repo]
	if !ok {
		return ""
	}

	return url
}
