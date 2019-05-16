// Copyright © 2019 Andreas Fritzler <andreas.fritzler@gmail.com>
//
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

package types

const (
	ButtonsType      = "buttons"
	TextType         = "text"
	CardType         = "card"
	QuickRepliesType = "quickReplies"
	CarouselType     = "carousel"
	ListType         = "list"
)

const (
	RequestErrorMessage = "Looks like there was a hick-up in my though process. Could you please try again?"
	HelloWorldMessage   = "hello world!"
)

// TextMessage defines a response of type text message.
// Example:
// {
// 	"type": "text",
//	"delay": 2,
//	"content": "MY_TEXT",
// }
type TextMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Delay   int    `json:"delay,omitempty"`
}

// QuickReplies defines a response of type text message.
// Example:
// {
// 	"type": "quickReplies",
// 	"content": {
// 		"title": "TITLE",
// 		"buttons": [
// 			{
// 				"title": "BUTTON_TITLE",
// 				"value": "BUTTON_VALUE"
// 			}
// 		]
// 	}
// }
type QuickReplies struct {
	Type    string              `json:"type"`
	Content QuickRepliesContent `json:"content"`
}

// QuickRepliesContent defines a subtype of the QuickReplies type.
type QuickRepliesContent struct {
	Title   string                `json:"title"`
	Buttons []QuickRepliesButtons `json:"buttons"`
}

// QuickRepliesButtons defines a subtype of the QuickRepliesContent type.
type QuickRepliesButtons struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

// Card defines a response of type card.
// Example:
// {
//     "type": "card",
//     "content": {
//       "title": "CARD_TITLE",
//       "subtitle": "CARD_SUBTITLE",
//       "imageUrl": "IMAGE_URL",
//       "buttons": [
//         {
//           "title": "BUTTON_TITLE",
//           "type": "BUTTON_TYPE",
//           "value": "BUTTON_VALUE"
//         }
//       ]
//     }
// }
type Card struct {
	Type    string      `json:"type"`
	Content CardContent `json:"content"`
}

// CardContent defines a subtype of the Card type.
type CardContent struct {
	Title    string   `json:"title"`
	SubTitle string   `json:"subtitle,omitempty"`
	ImageURL string   `json:"imageUrl,omitempty"`
	Buttons  []Button `json:"buttons"`
}

// Button defines a subtype for buttons.
type Button struct {
	Title string `json:"title"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Buttons defines a response of type buttons.
// Example:
// {
//     "type": "buttons",
//     "content": {
//       "title": "BUTTON_TITLE",
//       "buttons": [
//         {
//           "title": "BUTTON_TITLE",
//           "type": "BUTTON_TYPE",
//           "value": "BUTTON_VALUE"
//         }
//       ]
//     }
// }
type Buttons struct {
	Type    string         `json:"type"`
	Content ButtonsContent `json:"content"`
}

// ButtonsContent defines a subtype of the ButtonsType.
type ButtonsContent struct {
	Title   string   `json:"title"`
	Buttons []Button `json:"buttons"`
}

// Carousel defines a response of type buttons.
// Example:
// {
//     "type": "carousel",
//     "content": [
//       {
//         "title": "CARD_1_TITLE",
//         "subtitle": "CARD_1_SUBTITLE",
//         "imageUrl": "IMAGE_URL",
//         "buttons": [
//           {
//             "title": "BUTTON_1_TITLE",
//             "type": "BUTTON_1_TYPE",
//             "value": "BUTTON_1_VALUE"
//           }
//         ]
//       }
//     ]
// }
type Carousel struct {
	Type    string        `json:"type"`
	Content []CardContent `json:"content"`
}

// List defines a response of type buttons.
// Example:
// {
//     "type": "list",
//     "content": {
//       "elements": [
//         {
//           "title": "ELEM_1_TITLE",
//           "imageUrl": "IMAGE_URL",
//           "subtitle": "ELEM_1_SUBTITLE",
//           "buttons": [
//             {
//               "title": "BUTTON_1_TITLE",
//               "type": "BUTTON_TYPE",
//               "value": "BUTTON_1_VALUE"
//             }
//           ]
//         }
//       ],
//       "buttons": [
//         {
//           "title": "BUTTON_1_TITLE",
//           "type": "BUTTON_TYPE",
//           "value": "BUTTON_1_VALUE"
//         }
//       ]
//     }
// }
type List struct {
	Type    string      `json:"type"`
	Content ListContent `json:"content"`
}

// ListContent defines a subtype of the List type.
type ListContent struct {
	Elements []CardContent `json:"elements"`
	Buttons  []Button      `json:"buttons"`
}

// Request is a definition of the request type.
// Example:
// {
//
// }
type Request struct {
	IndexJsonUrl string `json:"indexJsonUrl"`
	Query        string `json:"query"`
	ResponseType string `json:"responseType"`
}

// Conversation is a subtype of the Request type.
type Conversation struct {
	Memory Memory `json:"memory"`
}

// Memory is a subtype of the Conversation type.
type Memory struct {
	Query string `json:"query"`
}

// Replies
type Replies struct {
	Replies []interface{} `json:"replies"`
}

// Reply
type Reply struct {
}

// DataArray
type DataArray []DataResponse

// DataResponse
type DataResponse struct {
	URI         string   `json:"uri"`
	Title       string   `json:"title"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
}
