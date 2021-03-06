# dingtalk_group_robot
push message or notification to dingtalk group

### 支持钉钉群组机器人所有消息格式

- [x] Text
- [x] Link
- [x] markdown
- [x] actionCard
- [x] feedCard
- [x] 增加签名调用支持 2019-11-04

###  运行测试用例

需要在robot_test.go中填入自己的 accesstoken。如何获取accesstoken见[官方文档](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)

``` golang
go test -v -count=1 *.go
```

###  如何调用
``` golang
package main

import robot "github.com/irebit/dingding_bot"

func main() {

	//Text
	textMessage := robot.NewText().SetContent("国产飞碟型武装直升机亮相")

	robot.New().SetAccessToken("59cfcfec07375399df981a930fca7ce84d8a94d6c686b9518082e6c0cebff8e9").Send(textMessage)

	//Link 适用于旧的群组机器人
	linkMessage := robot.NewLink().SetContent(
		"国产飞碟型武装直升机亮相", //title
		"10月10日，第五届中国天津国际直升机博览会在天津滨海新区空港经济区中航直升机产业基地内举行。外形酷似飞碟的“超级大白鲨”武装直升机在会上对外展出。", //text
		"https://news.ifeng.com/c/7qeji8ll3sm", //messageUrl
		"http://x0.ifengimg.com/res/2019/C63EEBD9080A3EE7039E7233C193718181ECE61B_size109_w600_h450.jpeg", //picUrl

	)
	robot.New().Send(linkMessage)


	// link with sign  适用于新申请的群组机器人
	linkMessage := robot.NewLink().SetContent(
		"国产飞碟型武装直升机亮相", //title
		"10月10日，第五届中国天津国际直升机博览会在天津滨海新区空港经济区中航直升机产业基地内举行。外形酷似飞碟的“超级大白鲨”武装直升机在会上对外展出。", //text
		"https://news.ifeng.com/c/7qeji8ll3sm", //messageUrl
		"http://x0.ifengimg.com/res/2019/C63EEBD9080A3EE7039E7233C193718181ECE61B_size109_w600_h450.jpeg", //picUrl
	)

	robot.New().SetAccessToken("feb3fad6dc8575b5a0e95fb2576ac86c61d9760039a91ae44809607cf7b42172").AddSign("SEC6761eafa77c228432873c1e4006f264d40bd3adf08fdb65d784a53d11cf8174d").Send(textMessage)

}
```

### 效果如下：
![示例图片](https://github.com/irebit/dingtalk_group_robot/raw/master/example.jpg)



更多调用示例见robot_test.go
