package posts_service

import "google.golang.org/protobuf/types/known/timestamppb"

type Post struct {
	Txt string
	TimeUpdated timestamppb.Timestamp
}
