package form

import (
	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/go-playground/validator.v9"
)

type formValidator struct{}

type validationError struct {
	field       string
	description string
}

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

	validationErrors := []validationError{}
	for _, err := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, fv.validatorFieldMessage(err.Field()))
	}

	return fv.generateBadRequestError(validationErrors)
}

func (fv *formValidator) generateBadRequestError(validationErrors []validationError) error {
	fieldViolations := make([]*errdetails.BadRequest_FieldViolation, 0, len(validationErrors))
	for _, validationError := range validationErrors {
		fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{
			Field:       validationError.field,
			Description: validationError.description,
		})
	}

	st := status.New(codes.InvalidArgument, "リクエストが不正です。")
	v := &errdetails.BadRequest{
		FieldViolations: fieldViolations,
	}

	s, _ := st.WithDetails(v)
	return s.Err()
}

func (fv *formValidator) validatorFieldMessage(fieldName string) validationError {
	vError := validationError{}
	switch fieldName {
	case "SearchKeywords":
		vError.description = "入力された検索キーワードが正しくありません。"
	case "Platform":
		vError.description = "配信プラットフォームが正しくありません。"
	case "Sort":
		vError.description = "ソートキーが正しくありません。"
	case "Page":
		vError.description = "現在ページが正しくありません。"
	case "Limit":
		vError.description = "取得する件数が正しくありません。"
	}

	vError.field = fieldName
	return vError
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
