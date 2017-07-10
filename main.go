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
	"log"
	"net/http"
	"os"
	"strings"

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
				} else if strings.Contains(message.Text, "愛你") || strings.Contains(message.Text, "愛妳") {
					//IP := event.ReplyToken
					out := fmt.Sprintf("I Love you, too")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(event.ReplyToken+" : "+out)).Do()
				} else if strings.Contains(message.Text, "你好") || strings.Contains(message.Text, "妳好") {
					out := fmt.Sprintf("你以為你是天線寶寶嗎?")
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do()
					//bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(message.OriginalContentURL+message.PreviewImageURL)).Do()
				} else if strings.Contains(message.Text, "幹") {
					out := fmt.Sprintf("尛")
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
				} else if strings.Contains(message.Text, "???") {
					originalContentURL := "https://www.hahatai.com/sites/default/files/u1031/8_2.jpg.pagespeed.ce.9yo2iS_Cxl.jpg"
					previewImageURL := "https://www.hahatai.com/sites/default/files/u1031/8_2.jpg.pagespeed.ce.9yo2iS_Cxl.jpg"
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(originalContentURL, previewImageURL)).Do()
				} else if strings.Contains(message.Text, "初音") {
					originalContentURL := "http://s65.youmaker.com/flv/2014/2-14/mp4909826563a3583eddc78f4a5da9cf3bbdeffecf79065.mp4"
					previewImageURL := "https://upload.wikimedia.org/wikipedia/zh/0/00/Miku_Hatsune.png"
					bot.ReplyMessage(event.ReplyToken, linebot.NewVideoMessage(originalContentURL, previewImageURL)).Do()
				} else if strings.Contains(message.Text, "1") {
					//out1 :=
					//	bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(out)).Do()
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
