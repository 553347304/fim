// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type ImageRequest struct {
}

type ImageResponse struct {
	Url string `json:"url"`
}

type ImageShowRequest struct {
	ImageType string `path:"image_type"`
	ImageName string `path:"image_name"`
}
