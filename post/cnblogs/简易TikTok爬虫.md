## TikTok 简易爬虫实现

tiktok web 页面，为各种爬虫准备了一份数据，就是其页面源码中，一个 id 为 `SIGI_STATE` 的 script 里的 json 数据。实际上，tiktok web 页面使用 sigi 框架，并且配合 SSR 将 sigi 应用的 state 保存在了 dom 里，相当于 vue 的 data。这个 state 里包含了用户的相关信息，用户发布的视频等等信息。

所以需要做的就是拉取 web 页面，解析出这个 json, 并且获取感兴趣的字段。

第一步，访问 tk 页面。tk 是限制了访问区域的，比如国内以及想干的大部分 ip 都不能够访问。所以第一步就是需要有一台能够访问 tk 的机器。

第二步，在这台机器上使用 curl 访问 tk 主页

第三步，从 html 页面中解析出 json

第四步，从 JSON 中提取感兴趣的字段

### 代码实现

一：跳板机。因为 tk 对访问的区域敏感，所以准备了多个区域的多台机器备用。查询时，可以选择发起访问的区域

```go
// region: ip
var TkDestIpMap = map[string][string]{}
```

使用 ssh 工具，连接到指定区域，并执行命令。这里写一个简易的 ssh 工具

```go
// 现在太菜，后面补
```

连接上之后，就可以执行命令了。但是因为命令比较多，而且使用 shell 编写也比较麻烦，所以使用 golang 编写，再打包成可执行命令，然后只需要触发一下就可以了。



二：获取 HTML 并提取感兴趣的字段

编写一个 golang 命令行工具

```go
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var tkId = flag.String("t", "", "tk user id")

// 使用一下命令，将此 go 程序编译成可执行程序（这里编译后的可执行程序名为 fetch。 使用方式为 ./fetch -t <tk user id>）
// build: CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./fetch ./parseTkState.go
func main() {
	flag.Parse()

	sw := bytes.Buffer{}
	command := exec.Command("curl", "-s", "https://www.tiktok.com/@"+*tkId)
	command.Stdout = &sw
	err := command.Run()
	if err != nil {
		_ = fmt.Errorf("curl tk failed %v", err)
		return
	}

  // 使用 curl 获取到的 tk 主页 html
	html := sw.String()
	script, err := extraJsonInScript(html)
	if err != nil {
		_ = fmt.Errorf("extra json failed %v", err)
		return
	}
  // 将结果输出到 stdout，便于调用者获取
	fmt.Print(script)
}

// 从 html 中解析出含有用户信息的 json
var scriptsReg = regexp.MustCompile(`<script\s+.*?>(.*?)</script>`)
func extraJsonInScript(html string) (string, error) {
	ret := ""
	rets := scriptsReg.FindAllStringSubmatch(html, -1)
	for _, v := range rets {
		isStateScript := strings.Contains(v[0], `id="SIGI_STATE"`)
		if len(v) > 1 && isStateScript && json.Valid([]byte(v[1])) {
			state := TKState{}
			jsonStr := strings.Trim(v[1], " ")
			err := json.Unmarshal([]byte(jsonStr), &state)
			if err != nil {
				continue
			}
			userMap := state.UserModule.Users
			if userMap == nil || len(userMap) == 0 {
				return "", errors.New("用户不存在或账号已注销")
			}
			// 返回什么内容，由 TKState 决定
			stateStr, err := json.Marshal(state)
			if err != nil {
				continue
			}
			ret = string(stateStr)
			break
		}
	}
	return ret, nil
}

// 省略这个结构体的内容。具体内容可以手动把 tk 主页的 json 拉出来看，并且使用工具转换成结构体即可
type TKState struct {
}
```

有了 fetch 这个可执行程序，调用方就很简单了。

```go
sh := NewSSHHelper(sshConf)
cmd := fmt.Sprintf("./fetch -t %s", url.PathEscape(tiktokId))
Logger.Infof("show tk user cmd: %s", cmd)
sshRet, err := sh.RunCMD(cmd)
if err != nil {
  l.Logger.Errorf("curl failed, %v", err)
}
return sshRet
```

但是因为访问 tk 主页是个网络请求行为，所以不得不考虑超时问题。以下是处理超时的逻辑：

```go
sh := NewSSHHelper(sshConf)
// 接收 fetch 命令结构的 channel
retChan := make(chan string, 1)
// 异步执行，让 main 进入 select 流程 
go func() {
  cmd := fmt.Sprintf("./fetch -t %s", url.PathEscape(tiktokId))
  Logger.Infof("show tk user cmd: %s", cmd)
  sshRet, err := sh.RunCMD(cmd)
  if err != nil {
    l.Logger.Errorf("curl failed, %v", err)
    // 出错了写入空值，后面会判断
    retChan <- ""
  }
  retChan <- sshRet
}()

// 经典的 golang 超时控制结构
select {
  case <-time.After(45 * time.Second):
  	return "", status.New(codes.DeadlineExceeded, "超时").Err()
  case sshRet := <-retChan:
    if sshRet == "" {
      return "", status.New(codes.Unknown, "解析失败").Err()
    } else {
      return sshRet, nil
    }
}
```

至此，一个简易的 tk 爬虫便能够跑起来了。

但是需要注意的是，fetch 程序是运行在能够访问 tk 的机器上的，而 fetch 程序的调用者，需要通过 ssh 连接到这台机器上去触发。并且，因为有很多区域，每个区域都有一个主机，所以 fetch 程序的部署也是一个繁琐的事情。

一开始写了一个 shell 脚本，循环所有的机器列表，一个个通过 scp 把编译好的 fetch 程序部署上去。虽然也能用，但是由于机器数量巨大，机器分布在全球，访问时间长短不一，脚本又不能并行，所以就执行得很慢。

后经同事指点，了解了 ansible 这个工具。那是真好用。

所以现在的部署脚本就是：

```shell
# build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./fetch ./parseTkState.go
echo 'build done'
sleep 2

# deploy
ansible tk -i ./ansible.ini -m copy -a "src=./fetch dest=~/"
echo 'copy done'

ansible tk -i ./ansible.ini -m file -a "path=/root/fetch mode=0755"
echo 'deploy done'
```

ansible.ini 就是配置机器列表，大致长这样：

```ini
[tk]
hostname1 ansible_password=yyy
hostname2
...
[tk:vars]
ansible_connection=ssh
ansible_user=root
ansible_password=xxx
```

