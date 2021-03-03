package dingding
import (
	"github.com/CodyGuo/dingtalk"
	"github.com/CodyGuo/dingtalk/pkg/robot"
	"github.com/CodyGuo/glog"
	"io/ioutil"
)

var (
	WebHookSecret = "SEC4fbbdfa7308df998ac09e8e93aef3b5460dfacf2976c9d14664ee032838f28fd"
	WebHookUrl    = "https://oapi.dingtalk.com/robot/send?access_token=2a70c3a6e989a2e7d784b24ab114d199773012177454e7abd2c1d9765e206d8f"
	DDingTalk     *dingtalk.DingTalk
)

func init() {
	initDingTalk()
}

// initDingTalk
func initDingTalk() {
	glog.SetFlags(glog.LglogFlags)
	webHook := WebHookUrl
	secret := WebHookSecret
	DDingTalk = dingtalk.New(webHook, dingtalk.WithSecret(secret))
}

func NewDingDing(webHook, secret string) *dingtalk.DingTalk {
	return dingtalk.New(webHook, dingtalk.WithSecret(secret))
}

// printResult
func printResult(dt *dingtalk.DingTalk) {
	response, err := dt.GetResponse()
	if err != nil {
		glog.Fatal(err)
	}
	reqBody, err := response.Request.GetBody()
	if err != nil {
		glog.Fatal(err)
	}
	reqData, err := ioutil.ReadAll(reqBody)
	if err != nil {
		glog.Fatal(err)
	}
	glog.Infof("发送消息成功, message: %s", reqData)
}

// NormalText 发送普通内容
func NormalText(dt *dingtalk.DingTalk, text string, members []string) {
	textContent := text
	// 要@的用户手机号
	atMobiles := robot.SendWithAtMobiles(members)
	if err := dt.RobotSendText(textContent, atMobiles); err != nil {
		glog.Fatal(err)
	}
	printResult(dt)
}

// NormalText 发送link类型
func NormalLink(dt *dingtalk.DingTalk, linkTitle, linkText, linkMessageURL, linkPicURL string) {
	if err := dt.RobotSendLink(linkTitle, linkText, linkMessageURL, linkPicURL); err != nil {
		glog.Fatal(err)
	}
	printResult(dt)
}

// NormalText 发送markdown类型
func NormalMarkdown(dt *dingtalk.DingTalk, markdownTitle, markdownText string) {
	if err := dt.RobotSendMarkdown(markdownTitle, markdownText); err != nil {
		glog.Fatal(err)
	}
	printResult(dt)
}

// NormalActionCard 整体跳转ActionCard类型
func NormalActionCard(dt *dingtalk.DingTalk, actionCardTitle, actionCardText, actionCardSingleTitle, actionCardSingleURL, actionCardBtnOrientation string) {
	if err := dt.RobotSendEntiretyActionCard(actionCardTitle,
		actionCardText,
		actionCardSingleTitle,
		actionCardSingleURL,
		actionCardBtnOrientation); err != nil {
		glog.Fatal(err)
	}
	printResult(dt)
}

// SingleActionCard 整体跳转ActionCard类型
func SingleActionCard(dt *dingtalk.DingTalk, actionCardSingleURL, actionCardTitle, actionCardText, actionCardBtnOrientation string) {
	// 独立跳转ActionCard类型
	btns := map[string]string{
		"内容不错": actionCardSingleURL,
		"不感兴趣": actionCardSingleURL,
	}
	if err := dt.RobotSendIndependentActionCard(actionCardTitle,
		actionCardText,
		actionCardBtnOrientation,
		btns); err != nil {
		glog.Fatal(err)
	}
	printResult(dt)
}

// NormalFreeCard
func NormalFreeCard(dt *dingtalk.DingTalk, linkTitle, linkMessageURL, linkPicURL string) {
	// FeedCard类型
	link1 := robot.FeedCardLink{
		Title:      linkTitle,
		MessageURL: linkMessageURL,
		PicURL:     linkPicURL,
	}
	link2 := robot.FeedCardLink{
		Title:      linkTitle + "2",
		MessageURL: linkMessageURL,
		PicURL:     linkPicURL,
	}
	links := []robot.FeedCardLink{link1, link2}
	if err := dt.RobotSendFeedCard(links); err != nil {
		glog.Fatal(err)
	}
	printResult(dt)
}
