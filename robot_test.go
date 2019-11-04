package dingtalk_group_robot

import (
	"log"
	"testing"
)

func TestSendText(t *testing.T) {
	message := NewText().SetContent("我就是我, 是不一样的烟火").AtAll() //or .AtMobiles([]string{"xxxxxxx"})
	b, err := New().SetAccessToken("59cfcfec07375399df981a930fca7ce84d8a94d6c686b9518082e6c0cebff8e9").Send(message)
	log.Println(b, err)
}

func TestSendLink(t *testing.T) {
	message := NewLink().SetContent(
		"国产飞碟型武装直升机亮相", //title
		"10月10日，第五届中国天津国际直升机博览会在天津滨海新区空港经济区中航直升机产业基地内举行。外形酷似飞碟的“超级大白鲨”武装直升机在会上对外展出。", //text
		"https://news.ifeng.com/c/7qeji8ll3sm", //messageUrl
		"http://x0.ifengimg.com/res/2019/C63EEBD9080A3EE7039E7233C193718181ECE61B_size109_w600_h450.jpeg", //picUrl
	)
	b, err := New().SetAccessToken("59cfcfec07375399df981a930fca7ce84d8a94d6c686b9518082e6c0cebff8e9").Send(message)
	log.Println(b, err)
}

func TestSendMarkDown(t *testing.T) {
	// 注意：被@人的手机号(在text内容里要有@手机号)
	message := NewMarkDown().SetContent(
		"杭州天气", //title
		"#### 杭州天气@156xxxx8827\n"+
			"> 9度，西北风1级，空气良89，相对温度73%\n\n"+
			"> ![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)\n"+
			"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n", //text
	).AtMobiles([]string{"156xxxx8827"})
	b, err := New().SetAccessToken("59cfcfec07375399df981a930fca7ce84d8a94d6c686b9518082e6c0cebff8e9").Send(message)
	log.Println(b, err)
}

func TestSendActionCard(t *testing.T) {
	message := NewActionCard().SetContent(
		"乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身", //title
		"![screenshot](@lADOpwk3K80C0M0FoA) ### 乔布斯 20 年前想打造的苹果咖啡厅 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划", //text
	//单个按钮
	// ).AddBtn("阅读全文", "https://www.dingtalk.com/")
	//多个按钮
	).AddBtns([][]string{
		[]string{"内容不错", "https://www.dingtalk.com/"},
		[]string{"不感兴趣", "https://www.dingtalk.com/"},
	}).HideAvatar().BtnOrientation()

	b, err := New().SetAccessToken("59cfcfec07375399df981a930fca7ce84d8a94d6c686b9518082e6c0cebff8e9").Send(message)
	log.Println(b, err)
}

func TestSendFeedCard(t *testing.T) {
	message := NewFeedCard().AddCards([][]string{
		[]string{"为互利合作注入新动力 为互联互通开辟新空间", "https://news.ifeng.com/c/7qeVEtaOZFo", "https://x0.ifengimg.com/ucms/2019_41/87832BA4D47AF7DF6FBCAF2ED750E966BC107D63_w900_h460.jpg"},
		[]string{"世界银行继续把中国列为营商环境改善度最高国家之一", "https://news.ifeng.com/c/7qedFUMM6BE", "https://x0.ifengimg.com/ucms/2019_41/53EFB3AF5F8DCD54EA74BD06CED980BB284603B6_w600_h450.jpg"},
	}).AddCard("国产飞碟型武装直升机亮相", "https://news.ifeng.com/c/7qeji8ll3sm", "http://x0.ifengimg.com/res/2019/C63EEBD9080A3EE7039E7233C193718181ECE61B_size109_w600_h450.jpeg")

	b, err := New().SetAccessToken("59cfcfec07375399df981a930fca7ce84d8a94d6c686b9518082e6c0cebff8e9").Send(message)
	log.Println(b, err)
}

func TestSendLinkWithSign(t *testing.T) {
	message := NewLink().SetContent(
		"国产飞碟型武装直升机亮相", //title
		"10月10日，第五届中国天津国际直升机博览会在天津滨海新区空港经济区中航直升机产业基地内举行。外形酷似飞碟的“超级大白鲨”武装直升机在会上对外展出。", //text
		"https://news.ifeng.com/c/7qeji8ll3sm", //messageUrl
		"http://x0.ifengimg.com/res/2019/C63EEBD9080A3EE7039E7233C193718181ECE61B_size109_w600_h450.jpeg", //picUrl
	)
	b, err := New().SetAccessToken("feb3fad6dc8575b5a0e95fb2576ac86c61d9760039a91ae44809607cf7b42172").AddSign("SEC6761eafa77c228432873c1e4006f264d40bd3adf08fdb65d784a53d11cf8174d").Send(message)
	log.Println(b, err)
}
