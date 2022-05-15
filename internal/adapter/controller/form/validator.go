package form

import (
	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"gopkg.in/go-playground/validator.v9"
)

type formValidator struct{}

const (
	isListLiveVideoSortValidateName string = "is_list_live_video_sort"
)

func Validate(form interface{}) error {
	validate := validator.New()

	formValidator := &formValidator{}
	err := formValidator.registerValidations(validate)
	if err != nil {
		return err
	}

	err = validate.Struct(form)
	return err
}

func (fv *formValidator) registerValidations(validate *validator.Validate) error {
	err := validate.RegisterValidation(isListLiveVideoSortValidateName, fv.isListLiveVideoSort)
	if err != nil {
		return err
	}

	return nil
}

func (fv *formValidator) isListLiveVideoSort(fl validator.FieldLevel) bool {
	value := fl.Field().Int()
	switch pb.ListLiveVideosRequest_Sort(value) {
	case pb.ListLiveVideosRequest_SORT_UNKNOWN,
		pb.ListLiveVideosRequest_SORT_PLATFORM,
		pb.ListLiveVideosRequest_SORT_VIEWER_ASC,
		pb.ListLiveVideosRequest_SORT_VIEWER_DESC,
		pb.ListLiveVideosRequest_SORT_STARTED_DATETIME_ASC,
		pb.ListLiveVideosRequest_SORT_STARTED_DATETIME_DESC:

		return true
	}

	return false
}
