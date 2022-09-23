package feedbackusecase

import (
	"immersive/config"
	entity "immersive/domains/feedback/entities"
	"immersive/exceptions"
	"immersive/utils/helpers"
	"time"
)

type feedbackUsecase struct {
	Repo entity.IFeedBackRepo
}

func New(repo entity.IFeedBackRepo) *feedbackUsecase {
	return &feedbackUsecase{
		Repo: repo,
	}
}

func (u *feedbackUsecase) Create(feedbackEntity entity.FeedBackEntity) error {

	fileExtension, err_image_extension := helpers.CheckFileExtension(feedbackEntity.FileName)
	if err_image_extension != nil {
		return exceptions.NewBadRequestError("image extension error")
	}

	err_image_size := helpers.CheckFileSize(feedbackEntity.FileSize, fileExtension)
	if err_image_size != nil {
		return exceptions.NewBadRequestError("image size error")
	}

	filename := time.Now().Format("2006-01-02 15:04:05") + "." + fileExtension

	if fileExtension != "pdf" {
		fileUrl, errUploadImg := helpers.UploadFileToS3(config.IMAGEDIR, filename, config.CONTENT_IMAGE, feedbackEntity.FileData)
		if errUploadImg != nil {
			return exceptions.NewInternalServerError(errUploadImg.Error())
		}
		feedbackEntity.Url = fileUrl
	} else {
		fileUrl, errUploadFile := helpers.UploadFileToS3(config.PDFDIR, filename, config.CONTENT_DOCUMENTS, feedbackEntity.FileData)
		if errUploadFile != nil {
			return exceptions.NewInternalServerError(errUploadFile.Error())
		}
		feedbackEntity.Url = fileUrl
	}

	err := u.Repo.Insert(feedbackEntity)
	if err != nil {
		return err
	}

	return u.Repo.UpdateMentee(feedbackEntity)
}

func (u *feedbackUsecase) GetAll(feedbackEntity entity.FeedBackEntity) ([]entity.FeedBackEntity, error) {
	return u.Repo.FindAll(feedbackEntity)
}
