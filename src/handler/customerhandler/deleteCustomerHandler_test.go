package customerhandler

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	"github.com/overusevery/golang-echo-practice2/src/domain/usecase/customerusecase"
	mock_repository "github.com/overusevery/golang-echo-practice2/src/repository/mock"
	"github.com/overusevery/golang-echo-practice2/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteCustomerHandler_DeleteCustomer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		e, close, m := setupMock(t)
		defer close()

		m.EXPECT().GetCustomer(context.Background(), gomock.Eq("1")).Return(forceNewCustomer(
			"1",
			"山田 太郎",
			"東京都練馬区豊玉北2-13-1",
			"176-0013",
			"03-1234-5678",
			"個人",
			"日本",
			time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
			1,
		), nil)
		m.EXPECT().DeleteCustomer(gomock.Any(),
			entity.DeletedCustomer(*forceNewCustomer(
				"1",
				"山田 太郎",
				"東京都練馬区豊玉北2-13-1",
				"176-0013",
				"03-1234-5678",
				"個人",
				"日本",
				time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
				1,
			)))
		res := testutil.DELETE(e, "/customer/1")
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("not found", func(t *testing.T) {
		e, close, m := setupMock(t)
		defer close()

		m.EXPECT().GetCustomer(context.Background(), gomock.Eq("1")).Return(nil, repository.ErrCustomerNotFound)
		res := testutil.DELETE(e, "/customer/1")
		assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	})
	t.Run("internal server error", func(t *testing.T) {
		e, close, m := setupMock(t)
		defer close()

		m.EXPECT().GetCustomer(context.Background(), gomock.Eq("1")).Return(forceNewCustomer(
			"1",
			"山田 太郎",
			"東京都練馬区豊玉北2-13-1",
			"176-0013",
			"03-1234-5678",
			"個人",
			"日本",
			time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
			1,
		), nil)
		m.EXPECT().DeleteCustomer(gomock.Any(), gomock.Any()).Return(errors.New("some error"))
		res := testutil.DELETE(e, "/customer/1")
		assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	})
}

func setupMock(t *testing.T) (*echo.Echo, func(), *mock_repository.MockCustomerRepository) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	close := func() {
		ctrl.Finish()
	}
	m := mock_repository.NewMockCustomerRepository(ctrl)
	h := NewDeleteCustomerHandler(*customerusecase.NewDeleteCustomerUseCase(m))
	h.RegisterRouter(e)
	return e, close, m
}
