package services

import (
	"github.com/jatinsaini25/gin-http-poc/entity"
	"github.com/jatinsaini25/gin-http-poc/repository"
)

type VideoService interface {
	FindAll() []entity.Video
	Save(video entity.Video) entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (v *videoService) FindAll() []entity.Video {
	return v.videoRepository.FindAll()
}

func (v *videoService) Save(data entity.Video) entity.Video {
	v.videoRepository.Save(data)
	return data
}

func (v *videoService) Update(video entity.Video) {
	v.videoRepository.Update(video)
}

func (v *videoService) Delete(video entity.Video) {
	v.videoRepository.Delete(video)
}
