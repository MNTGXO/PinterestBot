package pinterest

import (
 "fmt"
 "strings"

 "github.com/PaulSonOfLars/gotgbot/v2"
 "github.com/PaulSonOfLars/gotgbot/v2/ext"
 "github.com/Mishel-07/PinterestBot/api"
)

func FindImageInline(b *gotgbot.Bot, ctx *ext.Context) error {
 query := ctx.InlineQuery.Query
 if query == "" {
  _, err := ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{
   gotgbot.InlineQueryResultArticle{
    Title: "No Query Provided",
    InputMessageContent: &gotgbot.InputTextMessageContent{
     MessageText: "Please provide a query.",
    },
   },
  }, nil)
  return err
 }

 quotequery := strings.Replace(query, " ", "+", -1)
 urls, err := api.SearchPinterest(quotequery)
 if err != nil {
  fmt.Println(err)
  _, err = ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{
   gotgbot.InlineQueryResultArticle{
    Title: "Image not found",
    InputMessageContent: &gotgbot.InputTextMessageContent{
     MessageText: "Image not found for your query.",
    },
   },
  }, nil)
  return err
 }

 media := make([]gotgbot.InlineQueryResult, 0)
 for _, item := range urls.Data {
  if item.URL != "" {
   media = append(media, gotgbot.InlineQueryResultPhoto{
    PhotoURL: item.URL,
    Title:    "Found Image",
    ThumbURL: item.URL,
   })
  } else {
   fmt.Println("Skipped empty URL")
  }
 }

 if len(media) == 0 {
  _, err := ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{
   gotgbot.InlineQueryResultArticle{
    Title: "No Images Found",
    InputMessageContent: &gotgbot.InputTextMessageContent{
     MessageText: "No images found for your query.",
    },
   },
  }, nil)
  return err
 }

 _, err = ctx.InlineQuery.Answer(b, media, nil)
 return err
}
