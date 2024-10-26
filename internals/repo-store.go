package internals

import "log"

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
	"standalone-persistence": {
		url:    "https://bitbucket.org/blackcoders2019/simple_web.git",
		branch: "features/standalone/persistence",
	},
}

func getRepoURL(repo string) *repo {
	r, ok := repoStore[repo]
	if !ok {
		log.Fatalf("Project [%s] does not exists", repo)
	}

	return &r
}
