package controllers

import (
	"encoding/json"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	m2 "restArchitecture/mikail/App/mock"
	"restArchitecture/mikail/App/models"
	"testing"
)

func TestProductControllerRepo_InsertProductData(t *testing.T) {
	Convey("Insert Product Controllers", t, func() {
		Convey("Positive Scenario", func() {
			uc := NewControllers("test")
			cl := m2.ClientMock{}
			rd := m2.RedisMock{}

			cl.On("InsertProductList", mock.Anything).Return(nil)

			uc.setMock(&cl, &rd, nil)

			err := uc.InsertProductData(&models.InsertProductRequest{})
			So(err, ShouldBeNil)
		})
		Convey("Negative scenario", func() {
			Convey("error not nil", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}

				cl.On("InsertProductList", mock.Anything).Return(errors.New("err"))

				uc.setMock(&cl, &rd, nil)
				err := uc.InsertProductData(&models.InsertProductRequest{})
				So(err, ShouldNotBeNil)
			})
			Convey("Init error not nil", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}

				cl.On("InsertProductList", mock.Anything).Return(errors.New("err"))

				uc.setMock(&cl, &rd, errors.New("err"))
				err := uc.InsertProductData(&models.InsertProductRequest{})
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestProductControllerRepo_SortProductData(t *testing.T) {
	Convey("Sort Product Data", t, func() {
		Convey("Positive Scenario", func() {
			Convey("exist in redis", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}
				var res []*models.SortProductResponse
				res = append(res, &models.SortProductResponse{ProductName: "21"})
				resp, _ := json.Marshal([]*models.SortProductResponse{})
				cl.On("SortPriceProductList", mock.Anything).Return(res, nil)
				rd.On("Get", mock.Anything).Return(string(resp), true, nil)
				rd.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				uc.setMock(&cl, &rd, nil)
				res, err := uc.SortProductData("price", "asc")

				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
			Convey("Positive Scenario price", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}
				var res []*models.SortProductResponse
				res = append(res, &models.SortProductResponse{ProductName: "21"})

				cl.On("SortPriceProductList", mock.Anything).Return(res, nil)
				rd.On("Get", mock.Anything).Return("", false, nil)
				rd.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				uc.setMock(&cl, &rd, nil)
				res, err := uc.SortProductData("price", "asc")

				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
			Convey("Positive Scenario new", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}
				var res []*models.SortProductResponse
				res = append(res, &models.SortProductResponse{ProductName: "21"})

				cl.On("SortDateProductList", mock.Anything).Return(res, nil)
				rd.On("Get", mock.Anything).Return("", false, nil)
				rd.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				uc.setMock(&cl, &rd, nil)
				res, err := uc.SortProductData("new", "asc")

				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
			Convey("Positive Scenario name", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}
				var res []*models.SortProductResponse
				res = append(res, &models.SortProductResponse{ProductName: "21"})

				cl.On("SortNameProductList", mock.Anything).Return(res, nil)
				rd.On("Get", mock.Anything).Return("", false, nil)
				rd.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				uc.setMock(&cl, &rd, nil)
				res, err := uc.SortProductData("name", "asc")

				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("Get Error", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}
				var res []*models.SortProductResponse
				res = append(res, &models.SortProductResponse{ProductName: "21"})

				cl.On("SortPriceProductList", mock.Anything).Return(res, nil)
				rd.On("Get", mock.Anything).Return("", false, errors.New("err"))
				rd.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				uc.setMock(&cl, &rd, nil)
				res, err := uc.SortProductData("price", "asc")

				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
			Convey("SortPrice Error", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}

				cl.On("SortPriceProductList", mock.Anything).Return(nil, errors.New("err"))
				rd.On("Get", mock.Anything).Return("", false, nil)
				rd.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				uc.setMock(&cl, &rd, nil)
				res, err := uc.SortProductData("price", "asc")

				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
			Convey("Set Error", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}
				var res []*models.SortProductResponse
				res = append(res, &models.SortProductResponse{ProductName: "21"})

				cl.On("SortPriceProductList", mock.Anything).Return(res, nil)
				rd.On("Get", mock.Anything).Return("", false, nil)
				rd.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("err"))

				uc.setMock(&cl, &rd, nil)
				res, err := uc.SortProductData("price", "asc")

				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
			Convey("init error", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}
				var res []*models.SortProductResponse
				res = append(res, &models.SortProductResponse{ProductName: "21"})

				cl.On("SortPriceProductList", mock.Anything).Return(res, nil)
				rd.On("Get", mock.Anything).Return("", false, nil)
				rd.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				uc.setMock(&cl, &rd, errors.New("err"))
				res, err := uc.SortProductData("price", "asc")

				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
			Convey("order key not asc nor desc", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}
				var res []*models.SortProductResponse
				res = append(res, &models.SortProductResponse{ProductName: "21"})

				cl.On("SortPriceProductList", mock.Anything).Return(res, nil)
				rd.On("Get", mock.Anything).Return("", false, nil)
				rd.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				uc.setMock(&cl, &rd, nil)
				res, err := uc.SortProductData("price", "ppp")

				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
			Convey("error Scenario new", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}

				cl.On("SortDateProductList", mock.Anything).Return(nil, errors.New("err"))
				rd.On("Get", mock.Anything).Return("", false, nil)
				rd.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				uc.setMock(&cl, &rd, nil)
				res, err := uc.SortProductData("new", "asc")

				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
			Convey("error Scenario name", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}
				var res []*models.SortProductResponse
				res = append(res, &models.SortProductResponse{ProductName: "21"})

				cl.On("SortNameProductList", mock.Anything).Return(nil, errors.New("err"))
				rd.On("Get", mock.Anything).Return("", false, nil)
				rd.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				uc.setMock(&cl, &rd, nil)
				res, err := uc.SortProductData("name", "asc")

				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
			Convey("error Scenario key", func() {
				uc := NewControllers("test")
				cl := m2.ClientMock{}
				rd := m2.RedisMock{}
				var res []*models.SortProductResponse
				res = append(res, &models.SortProductResponse{ProductName: "21"})

				rd.On("Get", mock.Anything).Return("", false, nil)
				rd.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				uc.setMock(&cl, &rd, nil)
				res, err := uc.SortProductData("keyyy", "asc")

				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})
	})

}
