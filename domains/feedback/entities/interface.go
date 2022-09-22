package feedbackentities

type IFeedBackRepo interface {
	Insert(feedBackEntity FeedBackEntity) error
	UpdateMentee(feedBackEntity FeedBackEntity) error
	FindAll(feedBackEntity FeedBackEntity) ([]FeedBackEntity, error)
}

type IFeedBackUseCase interface {
	Create(feedBackEntity FeedBackEntity) error
	GetAll(feedBackEntity FeedBackEntity) ([]FeedBackEntity, error)
}
