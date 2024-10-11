package internals

type repo struct {
	url    string
	branch string
}

var repoStore = map[string]repo{
	"simple-web": {
		url:    "https://bitbucket.org/blackcoders2019/simple_web.git",
		branch: "master",
	},
	"mvn-wildfly": {
		url:    "https://bitbucket.org/blackcoders2019/simple_web.git",
		branch: "features/wildfly",
	},
}

func getRepoURL(repo string) *repo {
	r, ok := repoStore[repo]
	if !ok {
		return nil
	}

	return &r
}
