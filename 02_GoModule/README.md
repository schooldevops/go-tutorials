# Private GitHub로 모듈 리포지토리 이용하기

- go module을 이용하면, 복잡한 패키지를 편하게 관리할 수 있다. 
- go package들을 관리하기 위해서는 module repository 에 모듈을 올려두고, 이를 'go get' 을 이용하여 모듈을 가져올 수 있다. 
- 리포지토리는 GitHub 에 private로 생성해두고, 이를 패키지 리포지토리로 이용할 수 있다. 
  
## GitHup 리포지토리 생성하기. 

- GitHub에서 다음과 같이 리포지토리를 생성하자. 

![github01](imgs/github_01.png)

  - 리포지토리 이름: go-module-repo
  - Description: Go Module Repository
  - Private: 비공개 모드로 설정한다. 

- 위와 같이 생성하자. 

![github02](imgs/github_02.png)

- 위와 같이 리포지토리 정보를 확인할 수 있다. 

## 모듈 생성하기. 

- 이제는 모듈을 생성할 차례이다. 

### 작업 디렉토리 생성하기. 

```go
$ mkdir go-module-repo
$ cd go-module-repo

```

- 위 디렉토리는 작업 디렉토리이다. 
- 모듈들을 한곳에 만들 예정이기 때문에 위와 같이 하나의 리포지토리에 라이브러리 모듈들을 저장 관리할 것이다. 

```go
Ξ GO/go-module-repo → git init
/Users/kido/Documents/06.PROGRAMMING_COURSES/GO/go-module-repo/.git/ 안의 빈 깃 저장소를 다시 초기화했습니다

Ξ GO/go-module-repo git:(main) ▶ echo "go module" > README.md
Ξ GO/go-module-repo git:(main) ▶ git add .

Ξ GO/go-module-repo git:(main) ▶ git commit -m "first commit"
[main (최상위-커밋) 60258e2] first commit
 1 file changed, 1 insertion(+)
 create mode 100644 README.md

Ξ GO/go-module-repo git:(main) ▶ git remote add origin https://github.com/schooldevops/go-module-repo.git

Ξ GO/go-module-repo git:(main) ▶ git push -u origin main
```

- 리포지토리에 최상위 모들을 푸시하였다. 
 
### go mod 설정하기 

```go
$ mkdir greeting
$ cd greeting

$ go mod init github.com/schooldevops/go-module-repo/greeting
```

![gomod01](imgs/go_mod_01.png)

- 위 명령을 실행하고 나면, 모듈파일이(go.mod) 생성이 된다. 
- 우리가 만들 모듈은 greeting 모듈이 된다. 

### greeting.go 모듈 만들기 

- 이제 greeting.go 파일을 생성하고, 모듈을 작성해 보자. 

```go
package greeting

import "fmt"

func Greeting(username string) string {
	return fmt.Sprintf("Hello %s. Have Good Day~", username)
}
```

- 단순하게 username 이라는 이름을 얻어서 인사를 하는 함수를 작성했다. 

### 리포지토리에 푸시하지. 

```go
$ git add .
$ git commit -m "add greeting module"
$ git push

$ git tag v0.0.1
$ git push --tag
```

- 위와 같이 우리가 작성한 코드를 리포지토리로 푸시 했다. 
- 그리고 첫번째 버변인 v0.0.1로 태깅을 한다. 
- git push --tag 를 통해서 태깅된 버젼을 리포지토리로 추가한다. 

## Go Module 사용하기 

- 이제는 private 리포지토리에서 go module을 get 하여 사용해 볼 것이다. 
- 아래와 같이 프로젝트 생성 및 모듈 설치를 해보자. 

### 어플리케이션 초기화 하기 

```go
$ mkdir go_module_test
$ cd go_module_test

$ go mod init github.com/schooldevops/test

```

- 디렉토리를 만들고, go 모듈을 지정하자. 
- 어플리케이션도 모듈로 관리되면 소스 관리가 편리해진다. 

### private 리포지토리 설정하기. 

- 다음은 go 환경에 private 리포지토리를 설정해 보자. 

```py
$ go env -w GOPRIVATE=github.com/schooldevops
```

- 위와 같이 github의 자신의 리포지토리를 private 로 지정해 준다. 
- GOPRIVATE 를 설정해주면 우리의 github 리포지토리가 private 리포지토리라고 환경에 등록해주게 된다. 

```py
$ go env -json 

{
        "AR": "ar",
        "CC": "clang",
        "CGO_CFLAGS": "-g -O2",
        "CGO_CPPFLAGS": "",
        "CGO_CXXFLAGS": "-g -O2",
        "CGO_ENABLED": "1",
        "CGO_FFLAGS": "-g -O2",
        "CGO_LDFLAGS": "-g -O2",
        "CXX": "clang++",
        "GCCGO": "gccgo",
        "GO111MODULE": "",
        "GOARCH": "amd64",
        "GOBIN": "",
        "GOCACHE": "/Users/kido/Library/Caches/go-build",
        "GOENV": "/Users/kido/Library/Application Support/go/env",
        "GOEXE": "",
        "GOFLAGS": "",
        "GOGCCFLAGS": "-fPIC -arch x86_64 -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/7w/k23qswcd04z8925pqjt296rw0000gp/T/go-build4039752654=/tmp/go-build -gno-record-gcc-switches -fno-common",
        "GOHOSTARCH": "amd64",
        "GOHOSTOS": "darwin",
        "GOINSECURE": "",
        "GOMOD": "/Users/kido/Documents/06.PROGRAMMING_COURSES/GO/go-tutorials/GoModule/go.mod",
        "GOMODCACHE": "/Users/kido/go/pkg/mod",
        "GONOPROXY": "github.com/schooldevops",
        "GONOSUMDB": "github.com/schooldevops",
        "GOOS": "darwin",
        "GOPATH": "/Users/kido/go",
        "GOPRIVATE": "github.com/schooldevops",
        "GOPROXY": "https://proxy.golang.org,direct",
        "GOROOT": "/usr/local/Cellar/go/1.16/libexec",
        "GOSUMDB": "sum.golang.org",
        "GOTMPDIR": "",
        "GOTOOLDIR": "/usr/local/Cellar/go/1.16/libexec/pkg/tool/darwin_amd64",
        "GOVCS": "",
        "GOVERSION": "go1.16",
        "PKG_CONFIG": "pkg-config"
}
```

- 'go env -json' 을 통해서 환경 변수를 확인할 수 있다. 
- go 환경에 방금 지정한 값이 설정 되었음을 알 수 있다. 

### github 토큰 생성하기 

- 우리가 사용하는 go module repository는 private이므로 접근을 위한 보안 설정을 해야한다. 
- 이때 리포지토리에 접속하기 위한 token을 생성할 것이다. 

![go_mod_02](imgs/go_mod_02.png)

- github 에서 Settings를 클릭한다. 

![go_mod_03](imgs/go_mod_03.png)

- 좌측 메뉴에서 Developer settings를 선택한다. 

![go_mod_04](imgs/go_mod_04.png)

- 접속을 위한 토큰 생성을 위해 "Personal access tokens" 를 클릭한다. 
- 그리고 "Generate new token" 을 클릭한다. 

![go_mod_05](imgs/go_mod_05.png)

- 토큰 이름을 위와 같이 지정한다 
- 필요한 권한을 적절히 주자.
- 팔자의 케이스에서는 리포지토리에 대한 권한을 주었다. 

![go_mod_06](imgs/go_mod_06.png)

- 우리가 생성한 토큰을 위와 같이 확인할 수 있다. 
- 이를 복사해 두자. 

### Git 접근 설정하기. 

- github을 통해서 모듈을 가져올 것이기 때문에 git 설정을 다음과 같이 해 주어야한다. 
- user: github의 사용자 계정이름 
- personal_access_token: 좀전에 생성한 토큰
  - user/personal_access_token은 환경 변수로 지정해도 되고, 아래 플레이스 홀더에 채워 넣어도 된다. 

#### Git Hub 인경우 

```py
git config --global url."https://${user}:${personal_access_token}@github.com/schooldevops".insteadOf "https://github.com/schooldevops"
```

#### Git Lab 인 경우 

```py
git config --global url."https://oauth2:${personal_access_token}@privategitlab.com/schooldevops".insteadOf "https://privategitlab.com/schooldevops"

# or 

git config --global url."https://${user}:${personal_access_token}@privategitlab.com/schooldevops".insteadOf "https://privategitlab.com/schooldevops"

```

#### BitBucket 인 경우 

```py
git config --global url."https://${bitbucket_id}:${bitbucket_token}@privatebitbucket.com/schooldevops".insteadOf "https://privatebitbucket.com/schooldevops"
```

- 위와 같이 설정하면 다음과 같이 git config -l로 환경을 조회할 수 있다. 

```go
$ git config -l

...생략
filter.lfs.process=git-lfs filter-process
filter.lfs.required=true
url.https://schooldevops:ghp_9yR3HOsdfkjsdlfkjsflksdjflsdoci@github.com/schooldevops.insteadof=https://github.com/schooldevops
core.repositoryformatversion=0
core.filemode=true
core.bare=false
...생략
```

- 참고: 환경값이 잘못된 경우 수정을 하고자 한다면 다음 커맨드로 수정할 수 있다. 

```go
$ git config --global --edit
```

### 참고: SSH 접근 설정인경우 

- github 에서 SSH로 접근하는 방법에 대해서는 https://schooldevops.tistory.com/13 에서 도움을 얻을 수 있다. 

## 모듈 사용 프로그램 작성하기. 

- 이제 모듈을 사용하기 위한 프로그램을 작성해 보자. 

```go
$ mkdir greet-mod-test
$ cd greet-mod-test
```

```go
$ go mod init github.com/schooldevops/test

$ go get github.com/schooldevops/go-module-repo/greeting
```

- 위와 같이 모듈을 초기화 하고, 모듈을 'go get'을 활용하여 가져오자. 

```go
$ go get github.com/schooldevops/go-module-repo@v0.0.1

or 

$ go get -u github.com/schooldevops/go-module-repo
```

- 위와 같이 버젼을 @를 통해서 지정해줄 수 있다. 
- 혹은 최신 버젼을 가져오고자 한다면 'go get -u' 를 이용하여 최신 버젼을 가져올 수도 있다. 

### 모듈 사용하기. 

- 이제 드디어 코드를 작성할 차례이다. 
- main.go 파일을 만들고 다음과 같이 작성해보자. 

```go
package main

import (
	"fmt"

	"github.com/schooldevops/go-module-repo/greeting"
)

func main() {
	greet := greeting.Greeting("Schooldevops")
	fmt.Println(greet)
}

```

- 다음과 같이 실행해보자. 

```py
$ go run main.go                                        
Hello Schooldevops. Have Good Day~
```

- 원하는 모듈을 가져올 수 있게 되었다. 

## WrapUp

- 지금까지 github private 리포지토리를 이용하여 모듈을 푸시하고, 프로그램에서 생성된 모듈을 가져왔다. 
- 특정회사, 특정 프로젝트에서 사용하는 모듈을 private 리포지토리에서 생성하고, 해당 리포지토리 내부에 여러 모듈을 작성할 수 있음을 알게 되었다. 
