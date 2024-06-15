package statistics

type Post struct {
	PostId   int64  `ch:"post_id"`
	AuthorId int64  `ch:"author_id"`
	NumLikes uint64 `ch:"num_likes"`
	NumViews uint64 `ch:"num_views"`
}

type User struct {
	Id       int64  `ch:"author_id"`
	NumLikes uint64 `ch:"num_likes"`
	NumViews uint64 `ch:"num_views"`
}
