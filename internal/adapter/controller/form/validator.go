package form

import (
	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"gopkg.in/go-playground/validator.v9"
)

type formValidator struct{}

const (
	isListLiveVideoSortValidateName string = "is_list_live_video_sort"
	isVideoPlatformValidateName     string = "is_video_platform"
)

func Validate(form interface{}) error {
	validate := validator.New()

	formValidator := &formValidator{}
	err := formValidator.registerValidations(validate)
	if err != nil {
		return err
	}

	return formValidator.validateStruct(validate, form)
}

func (fv *formValidator) validateStruct(validate *validator.Validate, form interface{}) error {
	err := validate.Struct(form)
	if err == nil {
		return nil
	}

	/*
		for _, err := range err.(validator.ValidationErrors) {
			errMessage := fv.validatorFieldMessage(err.Field())
			// ここでエラーメッセージを新しく配列として作る
		}
	*/
	return err
}

func (fv *formValidator) validatorFieldMessage(fieldName string) string {
	switch fieldName {
	case "SearchKeywords":
		return "入力された検索キーワードが正しくありません。"
	case "Platform":
		return "配信プラットフォームが正しくありません。"
	case "Sort":
		return "ソートキーが正しくありません。"
	case "Page":
		return "現在ページが正しくありません。"
	case "Limit":
		return "取得する件数が正しくありません。"
	}

	return "validation error"
}

func (fv *formValidator) registerValidations(validate *validator.Validate) error {
	err := validate.RegisterValidation(isListLiveVideoSortValidateName, fv.isListLiveVideoSort)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation(isVideoPlatformValidateName, fv.isVideoPlatform)
	if err != nil {
		return err
	}

	return nil
}

func (fv *formValidator) isListLiveVideoSort(fl validator.FieldLevel) bool {
	value := fl.Field().Int()
	switch pb.ListLiveVideosRequest_Sort(value) {
	case pb.ListLiveVideosRequest_SORT_UNKNOWN,
		pb.ListLiveVideosRequest_SORT_VIEWER_ASC,
		pb.ListLiveVideosRequest_SORT_VIEWER_DESC,
		pb.ListLiveVideosRequest_SORT_STARTED_DATETIME_ASC,
		pb.ListLiveVideosRequest_SORT_STARTED_DATETIME_DESC:

		return true
	}

	return false
}

func (fv *formValidator) isVideoPlatform(fl validator.FieldLevel) bool {
	value := fl.Field().Int()
	switch pb.VideoPlatform(value) {
	case pb.VideoPlatform_VIDEO_PLATFORM_UNKNOWN,
		pb.VideoPlatform_TWITCH,
		pb.VideoPlatform_YOUTUBE:

		return true
	}

	return false
}
