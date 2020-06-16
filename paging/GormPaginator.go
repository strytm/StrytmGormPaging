package paging

import (
	"github.com/jinzhu/gorm"
	"math"
)

func PagingGorm(skipTake SkipTakeModel, db *gorm.DB, totalCount uint, result interface{}, association string) (PagingModel, error) {

	if len(association) == 0 {
		if dbResult := db.Offset(skipTake.Skip).Limit(skipTake.Take).Find(result); dbResult.Error != nil {

			return PagingModel{}, dbResult.Error
		}
	} else {
		if dbResult := db.Offset(skipTake.Skip).Limit(skipTake.Take).Association(association).Find(result); dbResult.Error != nil {

			return PagingModel{}, dbResult.Error
		}
	}

	return InitPageModel(skipTake, totalCount), nil
}

func InitPageModel(skipTake SkipTakeModel, totalRecord uint) PagingModel {
	pageModel := PagingModel{}
	pageModel.Skip = skipTake.Skip
	pageModel.Take = skipTake.Take

	pageModel.TotalRecord = totalRecord

	pageModel.TotalPage = uint(math.Ceil(float64(pageModel.TotalRecord) / float64(skipTake.Take)))

	remainingRecord := pageModel.TotalRecord - skipTake.SkipTakeSum()

	if skipTake.SkipTakeSum() > pageModel.TotalRecord {
		remainingRecord = 0
	}

	remainingPage := math.Ceil(float64(remainingRecord) / float64(pageModel.Take))

	pageModel.Page = pageModel.TotalPage - uint(remainingPage)

	if pageModel.Page == 0 {
		pageModel.PrevPage = pageModel.Page
	} else if pageModel.Page > 0{
		pageModel.PrevPage = pageModel.Page -1
		}

	if pageModel.Page+1 >= pageModel.TotalPage {
		pageModel.NextPage = pageModel.TotalPage
	} else {
		pageModel.NextPage = pageModel.Page + 1
	}

	return pageModel
}
