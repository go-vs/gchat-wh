package main

import (
	gchat_wh "github.com/go-vs/gchat-wh"
	"github.com/go-vs/gchat-wh/cards"
	"github.com/go-vs/gchat-wh/messages"
	"github.com/google/uuid"
	"os"
)

func main() {
	config := gchat_wh.Config{
		SpaceID: os.Getenv("SPACE_ID"),
		Key:     os.Getenv("KEY"),
		Token:   os.Getenv("TOKEN"),
	}
	wh := gchat_wh.New(config)
	// new thread
	threadKey := uuid.NewString()
	msg := gchat_wh.Message("thread " + threadKey)
	msg.Thread = &messages.Thread{
		ThreadKey: threadKey,
	}
	msg.CardsV2 = []messages.CardWithID{
		{
			CardID: "id1",
			Card: &cards.Card{
				Header: cards.NewHeader("Header" + threadKey).
					SetSubtitle("Subtitle").
					SetImageType(cards.ImageTypeCircle).
					SetImageURL("https://www.w3schools.com/html/pic_trulli.jpg").
					SetImageAltText(""),

				Sections: []*cards.Section{
					cards.NewSection().SetHeader("Noti 1").
						SetCollapsible(false).
						SetUncollapsibleWidgetsCount(0).
						AddWidgets(cards.NewWidget().SetTextParagraph(
							cards.NewTextParagraph("Hello 1234"),
						)),

					cards.NewSection().SetHeader("Noti 2").
						AddWidgets(
							cards.NewWidget().SetDateTimePicker(
								cards.NewDateTimePicker().SetName("Pick").
									SetLabel("date").
									SetType(cards.DateTimePickerDateAndTime).
									SetValueMsEpoch("1672574400000").
									SetTimezoneOffsetDate(0).
									SetOnChangeAction(
										cards.NewAction(),
									),
							),
						).
						SetCollapsible(false).
						SetUncollapsibleWidgetsCount(0),
				},
				SectionDividerStyle: cards.DividerSolid,
			},
		},
	}

	msg, err := wh.Send(msg)
	if err != nil {
		panic(err)
	}
	// reply
	_, err = wh.Send(gchat_wh.Reply(msg.Thread.ThreadKey, "reply "+msg.Thread.ThreadKey))
	if err != nil {
		panic(err)
	}
}
