package handlers

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ONSdigital/dp-api-clients-go/dataset"
	"github.com/ONSdigital/dp-api-clients-go/filter"
	dprequest "github.com/ONSdigital/dp-net/request"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

const mockUserAuthToken = "Foo"
const mockServiceAuthToken = ""
const mockCollectionID = "Bar"
const mockFilterID = ""
const batchSize = 100

func TestUpdateTime(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ctx := gomock.Any()
	Convey("test update time function", t, func() {
		Convey("the redirect is successful when the user has selected via the list time-selection", func() {
			options := []string{"Aug-11", "Aug-12"}
			mockClient := NewMockFilterClient(mockCtrl)
			mockClient.EXPECT().RemoveDimension(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, mockFilterID, "time").Return(nil)
			mockClient.EXPECT().AddDimension(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, mockFilterID, "time").Return(nil)
			mockClient.EXPECT().SetDimensionValues(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, mockFilterID, "time", options).Return(nil) // Might not be able to use gomock.Any() here
			formData := "latest-option=Nov-17&latest-month=November&latest-year=2017&month-single=Select&year-single=Select&start-month=Select&start-year=Select&end-month=Select&end-year=Select&time-selection=list&months=August&start-year-grouped=2011&end-year-grouped=2012&save-and-return=Save+and+return"
			w := callTimeUpdate(formData, mockClient, nil)
			So(w.Code, ShouldEqual, 302)
		})

		Convey("the redirect is successful when the user has selected via the latest", func() {
			option := "Jul-20"
			mockClient := NewMockFilterClient(mockCtrl)
			mockClient.EXPECT().RemoveDimension(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, mockFilterID, "time").Return(nil)
			mockClient.EXPECT().AddDimension(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, mockFilterID, "time").Return(nil)
			mockClient.EXPECT().AddDimensionValue(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, mockFilterID, "time", option).Return(nil) // Might not be able to use gomock.Any() here
			formData := "time-selection=latest&latest-option=Jul-20&latest-month=July&latest-year=2020&first-year=1988&first-month=January&month-single=Select&year-single=Select&start-month=Select&start-year=Select&end-month=Select&end-year=Select&months=February&start-year-grouped=2000&end-year-grouped=2002&save-and-return=Save+and+return"
			w := callTimeUpdate(formData, mockClient, nil)
			So(w.Code, ShouldEqual, 302)
		})

		Convey("the redirect is successful when the user has selected via the single", func() {
			option := "May-19"
			mockClient := NewMockFilterClient(mockCtrl)
			mockClient.EXPECT().RemoveDimension(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, mockFilterID, "time").Return(nil)
			mockClient.EXPECT().AddDimension(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, mockFilterID, "time").Return(nil)
			mockClient.EXPECT().AddDimensionValue(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, mockFilterID, "time", option).Return(nil) // Might not be able to use gomock.Any() here
			formData := "latest-option=Jul-20&latest-month=July&latest-year=2020&first-year=1988&first-month=January&time-selection=single&month-single=May&year-single=2019&start-month=Select&start-year=Select&end-month=Select&end-year=Select&start-year-grouped=Select&end-year-grouped=Select&save-and-return=Save+and+return"
			w := callTimeUpdate(formData, mockClient, nil)
			So(w.Code, ShouldEqual, 302)
		})

		Convey("the redirect is successful when the user has selected via the range and the option labels are correctly converted into filter option values", func() {
			expectedFilterModel := filter.Model{
				Links: filter.Links{
					Version: filter.Link{
						HRef: "http://localhost:23200/v1/datasets/abcde/editions/2017/versions/1",
					},
				},
			}
			datasetOptions := dataset.Options{
				Items: []dataset.Option{
					{Label: "Jan-00", Option: "Jan-00"},
					{Label: "Feb-00", Option: "Feb-00"},
					{Label: "Mar-00", Option: "Mar-00"},
				},
			}
			filterOptions := []string{"Jan-00", "Feb-00", "Mar-00"}
			mockFilterClient := NewMockFilterClient(mockCtrl)
			mockDatasetClient := NewMockDatasetClient(mockCtrl)
			mockFilterClient.EXPECT().RemoveDimension(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, mockFilterID, "time").Return(nil)
			mockFilterClient.EXPECT().AddDimension(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, mockFilterID, "time").Return(nil)
			mockFilterClient.EXPECT().GetJobState(ctx, mockUserAuthToken, mockServiceAuthToken, "", mockCollectionID, mockFilterID).Return(expectedFilterModel, nil)
			mockDatasetClient.EXPECT().GetOptions(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, "abcde", "2017", "1", "time").Return(datasetOptions, nil)
			mockFilterClient.EXPECT().SetDimensionValues(ctx, mockUserAuthToken, mockServiceAuthToken, mockCollectionID, mockFilterID, "time", filterOptions).Return(nil) // Might not be able to use gomock.Any() here
			formData := "latest-option=Jul-20&latest-month=July&latest-year=2020&first-year=1988&first-month=January&month-single=Select&year-single=Select&time-selection=range&start-month=January&start-year=2000&end-month=March&end-year=2000&start-year-grouped=Select&end-year-grouped=Select&save-and-return=Save+and+return"
			w := callTimeUpdate(formData, mockFilterClient, mockDatasetClient)
			So(w.Code, ShouldEqual, 302)
		})
	})
}

func callTimeUpdate(formData string, mockFilterClient *MockFilterClient, mockDatasetClient *MockDatasetClient) *httptest.ResponseRecorder {
	target := fmt.Sprintf("/filters/%s/dimensions/time/update", mockFilterID)
	reader := strings.NewReader(formData)
	req := httptest.NewRequest("POST", target, reader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add(dprequest.FlorenceHeaderKey, mockUserAuthToken)
	req.Header.Add(dprequest.CollectionIDHeaderKey, mockCollectionID)
	w := httptest.NewRecorder()
	f := NewFilter(nil, mockFilterClient, mockDatasetClient, nil, nil, nil, mockServiceAuthToken, "", "/v1", false, batchSize)
	f.UpdateTime().ServeHTTP(w, req)
	return w
}
