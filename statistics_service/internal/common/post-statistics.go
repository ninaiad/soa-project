package common

type PostStatistics struct {
	PostId     uint64 `ch:"post"`
	AuthorId   uint64 `ch:"author"`
	TotalLikes uint64 `ch:"total_likes"`
	TotalViews uint64 `ch:"total_views"`
}

type UserStatistics struct {
	AuthorId   uint64 `ch:"author"`
	TotalLikes uint64 `ch:"total_likes"`
	TotalViews uint64 `ch:"total_views"`
}
