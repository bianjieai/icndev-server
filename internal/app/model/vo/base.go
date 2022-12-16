package vo

type BaseVO interface {
	Validate() error
}

type PageVO struct {
	Page  int  `json:"page" form:"page" binding:"required"`
	Size  int  `json:"size" form:"size" binding:"required"`
	Total bool `json:"total" form:"total"`
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BaseResponseWithPage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
	Total   int         `json:"total,omitempty"`
}

type Page struct {
	Page  int  `json:"page" form:"page" binding:"required"`
	Size  int  `json:"size" form:"size" binding:"required"`
	Total bool `json:"total" form:"total"`
}
