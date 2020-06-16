package paging

import (
	"net/http"

	"github.com/mitchellh/mapstructure"
	"github.com/strytm/StrytmUtils/utils"
	"github.com/strytm/StrytmValidator/validator"
)

type SkipTakeModel struct {
	Skip uint
	Take uint
}

func (data *SkipTakeModel) Validation(validation *validator.Validation) {

	if data.Take > 200 {

		data.Take = 200
	}

	if data.Take == 0 {
		data.Take = 10
	}
}

func (data *SkipTakeModel) SkipTakeSum() uint {
	return data.Take + data.Skip
}

func (data *SkipTakeModel) BindData(c *http.Request) error {

	clearMapData := utils.GetAllFormRequestValue(c)

	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &data,
	}

	decoder, errNewDecoder := mapstructure.NewDecoder(config)
	if errNewDecoder != nil {
		return errNewDecoder
	}

	if err := decoder.Decode(clearMapData); err != nil {
		return err
	}

	return nil

}
