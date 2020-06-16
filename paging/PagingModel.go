package paging

type PagingModel struct {
	TotalRecord uint `json:"total_record"`
	TotalPage   uint `json:"total_page"`
	Skip        uint `json:"skip"`
	Take        uint `json:"take"`
	Page        uint `json:"page"`
	PrevPage    uint `json:"prev_page"`
	NextPage    uint `json:"next_page"`
}
