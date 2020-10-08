// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"

	//"net/url"
	"math/rand"
	"os"
	"strings"

	//"strconv"
	"time"

	//"github.com/JustinBeckwith/go-yelp/yelp"
	//"github.com/guregu/null"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	//var err2 error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)

}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	//events, err2 := strings.Contains(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	//for _, result := range events.Results {
	//content := result.Content()
	//}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {

			switch message := event.Message.(type) {

			case *linebot.TextMessage:
				//----------------回聲範例---------------------
				/*if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" TextOK!")).Do(); err != nil {
					//發送訊息的格式
					log.Print(err)
				}*/
				//----------------------------------------------------------------------
				//----------------關鍵字回復--------------------
				if strings.Contains(message.Text, "/help") || strings.Contains(message.Text, "/HELP") {
					out := fmt.Sprintf("HELP")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
					/*
						} else if strings.Contains(string(event.Source.UserID), "U54c79813d639c4ba3cffdb3d0987c99d"){
							out := fmt.Sprintf("yoyoyo")
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
						} else if strings.Contains(message.Text, "test"){
							bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage(event.Source.UserID)).Do()
					*/
				} else if strings.Contains(message.Text, "speed") {
					replytoken := event.ReplyToken
					start := time.Now()
					bot.ReplyMessage(replytoken, linebot.NewTextMessage("...")).Do()
					end := time.Now()
					result := fmt.Sprintf("%f [sec]", (end.Sub(start)).Seconds())
					_, err := bot.PushMessage(event.Source.GroupID, linebot.NewTextMessage(result)).Do()
					if err != nil {
						_, err := bot.PushMessage(event.Source.RoomID, linebot.NewTextMessage(result)).Do()
						if err != nil {
							_, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(result)).Do()
							if err != nil {
								log.Fatal(err)
							}
						}
					}
				} else if strings.Contains(message.Text, "吃飯") {
					out := fmt.Sprintf("https://www.foodpanda.com.tw")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				} else if strings.Contains(message.Text, "吃") {
					rand.Seed(time.Now().UnixNano()) // Try changing this number!
					answers := []string{
						"吃飯",
						"吃麵",
						"吃水餃",
						"吃炸雞",
						"吃PIZZA",
						"吃手手",
					}
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(answers[rand.Intn(len(answers))])).Do()
				} else if strings.Contains(message.Text, "bye") || strings.Contains(message.Text, "哭") {
					bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("3", "187")).Do()
				} else if strings.Contains(message.Text, "now") {
					now := time.Now()
					out := now.Format(time.RFC3339)
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				} else if strings.Contains(message.Text, "到了") {
					//IP := event.ReplyToken
					out := fmt.Sprintf("胡了! 大三元、混一色、對對胡、北風、無字無花、平胡...106台!!")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				} else if strings.Contains(message.Text, "愛你") || strings.Contains(message.Text, "愛妳") {
					//IP := event.ReplyToken

					out := fmt.Sprintf("I Love You, too $")

					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
					//.AddEmoji(linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "221"))
				} else if strings.Contains(message.Text, "你好") || strings.Contains(message.Text, "妳好") {
					out := fmt.Sprintf("你以為你是天線寶寶嗎?")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
					//bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.OriginalContentURL+message.PreviewImageURL)).Do()
				} else if strings.Contains(message.Text, "幹") {
					out := fmt.Sprintf("尛")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				} else if strings.Contains(message.Text, "功課") || strings.Contains(message.Text, "作業") {
					originalContentURL := "https://gss0.baidu.com/-fo3dSag_xI4khGko9WTAnF6hhy/zhidao/wh%3D600%2C800/sign=d4257a7f07d162d985bb6a1a21ef85d4/d50735fae6cd7b89487c0e38022442a7d9330e7c.jpg"
					previewImageURL := "https://gss0.baidu.com/-fo3dSag_xI4khGko9WTAnF6hhy/zhidao/wh%3D600%2C800/sign=d4257a7f07d162d985bb6a1a21ef85d4/d50735fae6cd7b89487c0e38022442a7d9330e7c.jpg"
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(originalContentURL, previewImageURL)).Do()
				} else if strings.Contains(message.Text, "謝謝書凱") {
					originalContentURL := "https://hkchen.tech/images/shukae.png"
					previewImageURL := "https://hkchen.tech/images/shukae.png"
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(originalContentURL, previewImageURL)).Do()
				} else if strings.Contains(message.Text, "Elton Worksheet") {
					out := fmt.Sprintf("https://drive.google.com/file/d/1Mv7Gjt1ib7m3YaZQajq4sfW9ET_5JHeR/view?usp=sharing")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				} else if strings.Contains(message.Text, "成員分類") {	
					out := fmt.Sprintf("有害的:工作愛做不做，愛抱怨，和自己不相關的事不願意做 \n平庸的:安排做什麼就做什麼，沒安排就不做，等著事情過來的人 \n極好的:發自內心做事，做事有條理、認真負責，真心為公司操心的 \n")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				} else if strings.Contains(message.Text, "成員資料") {
					out := fmt.Sprintf("姓名:周義淳 \n性別:男 \n年次:民國79年 \n工作年資:6年 \n學歷:黎明技術學院 資訊工程系 \n專長:LINUX \n嗜好:看電影、打電動 \n優點:樂觀、負責 \n缺點:本人說太月半 \n工作目標:守護運維的和平 \n\n姓名:張智綱 \n性別:男 \n年次:民國80年 \n工作年資:5年 \n學歷:龍華科技大學 資訊管理系 \n專長:LINUX \n嗜好:聽音樂 \n優點:細心、認真 \n缺點:嗓門大 \n工作目標:抵禦APP TEAM的騷擾")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				} else if strings.Contains(message.Text, "王道") {
					if strings.Contains(message.Text, "環境") {
						out := fmt.Sprintf("172.30.5.89 RbESTMTPBWB01 \n 10.88.20.112 RbESTMTPAWB01 \n 10.88.20.113 RbESTMTPCV01 \n 10.88.20.115 RbESTMTTAP01")
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
					} else if strings.Contains(message.Text, "電話") {
						out := fmt.Sprintf("(02) 8178-3177 \n 梁韡峻 ext.12325 \n 翁誠鴻 ext.12332 \n 黃雋幃 ext.12371")
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
					}

				} else if strings.Contains(message.Text, "幾點了") {
					out := fmt.Sprintf("不知道")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				} else if strings.Contains(message.Text, "c8763") || strings.Contains(message.Text, "C8763") || strings.Contains(message.Text, "星爆氣流斬") {
					out := fmt.Sprintf("噓")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				} else if strings.Contains(message.Text, "姆咪姆咪") {
					out := fmt.Sprintf("心動動")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				} else if strings.Contains(message.Text, "好了") || strings.Contains(message.Text, "好惹") || strings.Contains(message.Text, "水喔") {
					originalContentURL := "https://www.cool-gif.com/media/gif/48/18/4818634f4c87146039e774c2bc752be1/cool.gif"
					previewImageURL := "https://www.cool-gif.com/media/gif/48/18/4818634f4c87146039e774c2bc752be1/cool.gif"
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(originalContentURL, previewImageURL)).Do()
				} else if strings.Contains(message.Text, "唱歌") {
					originalContentURL := "https://i.imgur.com/kJ3KoVs.png"
					previewImageURL := "https://i.imgur.com/kJ3KoVs.png"
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(originalContentURL, previewImageURL)).Do()
				} else if strings.Contains(message.Text, "大大") {
					originalContentURL := "https://pic.pimg.tw/ash1118/1445931964-184045615.jpg"
					previewImageURL := "https://pic.pimg.tw/ash1118/1445931964-184045615.jpg"
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(originalContentURL, previewImageURL)).Do()
				} else if strings.Contains(message.Text, "...") {
					originalContentURL := "https://pic.pimg.tw/chenwei389/1448377600-433737766_n.jpg"
					previewImageURL := "https://pic.pimg.tw/chenwei389/1448377600-433737766_n.jpg"
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(originalContentURL, previewImageURL)).Do()
				} else if strings.Contains(message.Text, "表示:") {
					originalContentURL := "https://i.ytimg.com/vi/4pKMRvrW69s/hqdefault.jpg"
					previewImageURL := "https://i.ytimg.com/vi/4pKMRvrW69s/hqdefault.jpg"
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(originalContentURL, previewImageURL)).Do()
				} else if strings.Contains(message.Text, "???") {
					originalContentURL := "https://vignette.wikia.nocookie.net/evchk/images/e/ec/2471912.jpg/revision/latest/scale-to-width-down/1000?cb=20171012125530"
					previewImageURL := "https://vignette.wikia.nocookie.net/evchk/images/e/ec/2471912.jpg/revision/latest/scale-to-width-down/1000?cb=20171012125530"
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(originalContentURL, previewImageURL)).Do()
				} else if strings.Contains(message.Text, "一個人") || strings.Contains(message.Text, "沒人一起") || strings.Contains(message.Text, "沒人要一起") {
					originalContentURL := "https://www.youtube.com/watch?v=gELloR2QiYM"
					previewImageURL := "https://i.ytimg.com/vi/gELloR2QiYM/maxresdefault.jpg"
					bot.ReplyMessage(event.ReplyToken, linebot.NewVideoMessage(originalContentURL, previewImageURL)).Do()
				} else if strings.Contains(message.Text, "初音") {
					originalContentURL := "https://s65.youmaker.com/flv/2014/2-14/mp4909826563a3583eddc78f4a5da9cf3bbdeffecf79065.mp4"
					previewImageURL := "https://upload.wikimedia.org/wikipedia/zh/0/00/Miku_Hatsune.png"
					bot.ReplyMessage(event.ReplyToken, linebot.NewVideoMessage(originalContentURL, previewImageURL)).Do()
				} else if strings.Contains(message.Text, "去哪") || strings.Contains(message.Text, "好玩") {
					bot.ReplyMessage(event.ReplyToken, linebot.NewLocationMessage(
						"Disney Resort",
						"〒279-0031 千葉県浦安市舞浜１−１",
						35.632211,
						139.881234)).Do()
				} else if strings.Contains(message.Text, "全頻廣播") {

					//IP := event.ReplyToken //飲茶
					//IP[1] = "b4c929b7ceec4e21912e6e16304ff0ee" //台南吃吃吃
					//IP := []string{"6c65c8b36882491faa32493bfeba736", "b4c929b7ceec4e21912e6e16304ff0ee"}
					//IPP := IP[2]
					//bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("*注意*")).Do()
					//IP := event.ReplyToken
					bot.PushMessage(event.ReplyToken, linebot.NewTextMessage("hello")).Do()
				}

				/* else if strings.Contains(message.Text, "rdrrJC") {
					//type:= "image",
					OURL := fmt.Sprintf("https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png")
					PURL := fmt.Sprintf("https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png")
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(OURL+PURL)).Do()
				}*/
			/*else if strings.Contains(message.Text, "rdrrJC") {
				message.OriginalContentURL = "https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png"
				bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.OriginalContentURL+"ImageOK!")).Do()
			}*/
			/*else {      //回聲功能
				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Lucy:"+message.Text+" Aye")).Do() //message.ID
			}*/
			//----------------------------------------------------------------------

			case *linebot.LocationMessage:
				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這裡好玩嗎?")).Do()

			case *linebot.ImageMessage:
				//if message.ID == "RS232.jpg" {
				//out := fmt.Sprintf("這是圖片")
				//bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out+message.OriginalContentURL+message.PreviewImageURL)).Do()
				//}
				//out := fmt.Sprintf("這是圖片")
				//bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				/*if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.ID+":"+message.OriginalContentURL+message.PreviewImageURL+"ImageOK!")).Do(); err != nil {
					log.Print(err)
				}*/
				//message.OriginalContentURL = "https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png"
				//message.PreviewImageURL = "https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png"
				/*if bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.OriginalContentURL+message.PreviewImageURL+"ImageOK!")).Do() {
					out := fmt.Sprintf("這是圖片")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
				}*/

				/*if strings.Contains(message.Text, "rdrrJC") {
				//type:= "image",
				OURL := fmt.Sprintf("https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png")
				PURL := fmt.Sprintf("https://2.bp.blogspot.com/-qb9ZYC7-dAg/Ts0hjukIZnI/AAAAAAAACq0/1xX9ujTZa5g/s1600/JackieChan.png")
				bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(OURL+PURL+"ImageOK!")).Do()
				}*/

				//--------------------------------------------------------------
				/*
					case *linebot.ImageMessage:
						if _, err2 = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.ID+":"+message.OriginalContentURL+message.PreviewImageURL+"ImageOK!")).Do(); err2 != nil {
							log.Print(err2)
						}
				*/
			}
			//--------------------------------------------------------------- + message.PreviewImageURL
		}
	}
	/*
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.ImageMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.ID+":/n"+message.OriginalContentURL+"/n"+message.PreviewImageURL+"/n OK!")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	*/

}
