package rest_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nobbyphala/Brick/domain"
	"github.com/nobbyphala/Brick/external/validator"
	mock_usecase "github.com/nobbyphala/Brick/mock/usecase"
	"github.com/nobbyphala/Brick/usecase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDisbursementController_VerifyDisbursement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDisbursementUsecase := mock_usecase.NewMockDisbursement(ctrl)

	type fields struct {
		disbursementUsecase usecase.Disbursement
		validator           validator.Validator
	}
	type args struct {
		req interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       string
		wantStatus int
		mock       func()
	}{
		{
			name: "successfully verify disbursement",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: VerifyDisbursementRequest{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
				},
			},
			wantStatus: http.StatusOK,
			want:       `{"message":"disbursement successfully verified"}`,
			mock: func() {
				mockDisbursementUsecase.EXPECT().VerifyDisbursement(gomock.Any(), domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
				}).Return(nil)
			},
		},
		{
			name: "verify account not found",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: VerifyDisbursementRequest{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
				},
			},
			wantStatus: http.StatusOK,
			want: func() string {
				res := ErrorResponse{
					Message: "error bank account not found",
				}

				jsonByte, _ := json.Marshal(res)
				return string(jsonByte)
			}(),
			mock: func() {
				mockDisbursementUsecase.EXPECT().VerifyDisbursement(gomock.Any(), domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
				}).Return(errors.New("error bank account not found"))
			},
		},
		{
			name: "verify account blocked",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: VerifyDisbursementRequest{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
				},
			},
			wantStatus: http.StatusOK,
			want: func() string {
				res := ErrorResponse{
					Message: "error bank account is blocked",
				}

				jsonByte, _ := json.Marshal(res)
				return string(jsonByte)
			}(),
			mock: func() {
				mockDisbursementUsecase.EXPECT().VerifyDisbursement(gomock.Any(), domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
				}).Return(errors.New("error bank account is blocked"))
			},
		},
		{
			name: "verify error",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: VerifyDisbursementRequest{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
				},
			},
			wantStatus: http.StatusInternalServerError,
			want: func() string {
				res := ErrorResponse{
					Message: "error when trying verify disbursement",
				}

				jsonByte, _ := json.Marshal(res)
				return string(jsonByte)
			}(),
			mock: func() {
				mockDisbursementUsecase.EXPECT().VerifyDisbursement(gomock.Any(), domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
				}).Return(errors.New("error when trying verify disbursement"))
			},
		},
		{
			name: "invalid request validation",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: VerifyDisbursementRequest{
					RecipientName:          "",
					RecipientAccountNumber: "094578A",
					RecipientBankCode:      "",
					Amount:                 90000,
				},
			},
			wantStatus: http.StatusBadRequest,
			want: func() string {
				res := ValidationErrorResponse{
					Message: "invalid request",
					Errors: []ValidationErrorItem{
						{
							Field: "RecipientName",
							Error: "RecipientName must be greater than 1 character in length",
						},
						{
							Field: "RecipientAccountNumber",
							Error: "RecipientAccountNumber must be a valid numeric value",
						},
						{
							Field: "RecipientBankCode",
							Error: "RecipientBankCode must be at least 1 character in length",
						},
					},
				}

				jsonByte, _ := json.Marshal(res)
				return string(jsonByte)
			}(),
			mock: func() {

			},
		},
		{
			name: "invalid request schema",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: "",
			},
			wantStatus: http.StatusBadRequest,
			want: func() string {
				res := ErrorResponse{
					Message: "invalid request",
				}

				jsonByte, _ := json.Marshal(res)
				return string(jsonByte)
			}(),
			mock: func() {

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			controller := DisbursementController{
				disbursementUsecase: tt.fields.disbursementUsecase,
				validator:           tt.fields.validator,
			}

			router := gin.New()
			router.POST("/test", controller.VerifyDisbursement)

			requestBody, _ := json.Marshal(tt.args.req)

			req, err := http.NewRequest("POST", "/test", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			respRecorder := httptest.NewRecorder()

			router.ServeHTTP(respRecorder, req)

			assert.Equal(t, tt.wantStatus, respRecorder.Code)

			assert.Equal(t, tt.want, respRecorder.Body.String())
		})
	}
}

func TestDisbursementController_Disburse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDisbursementUsecase := mock_usecase.NewMockDisbursement(ctrl)

	type fields struct {
		disbursementUsecase usecase.Disbursement
		validator           validator.Validator
	}
	type args struct {
		req interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       string
		wantStatus int
		mock       func()
	}{
		{
			name: "successfully Disburse",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: DisburseRequest{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
				},
			},
			wantStatus: http.StatusOK,
			want: func() string {
				res := DisbursementResponse{
					Id:                     "disb-id-1",
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
					Status:                 "PENDING",
				}

				jsonByte, _ := json.Marshal(res)
				return string(jsonByte)
			}(),
			mock: func() {
				mockDisbursementUsecase.EXPECT().Disburse(gomock.Any(), domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
				}).Return(domain.Disbursement{
					Id:                     "disb-id-1",
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
					Status:                 1,
				}, nil)
			},
		},
		{
			name: "error when disbursing",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: DisburseRequest{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
				},
			},
			wantStatus: http.StatusInternalServerError,
			want: func() string {
				res := ErrorResponse{Message: "error when try to disburse"}

				jsonByte, _ := json.Marshal(res)
				return string(jsonByte)
			}(),
			mock: func() {
				mockDisbursementUsecase.EXPECT().Disburse(gomock.Any(), domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "94578",
					RecipientBankCode:      "BANK A",
					Amount:                 90000,
				}).Return(domain.Disbursement{}, errors.New("error when try to disburse"))
			},
		},
		{
			name: "invalid request validation",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: DisburseRequest{
					RecipientName:          "",
					RecipientAccountNumber: "094578A",
					RecipientBankCode:      "",
					Amount:                 90000,
				},
			},
			wantStatus: http.StatusBadRequest,
			want: func() string {
				res := ValidationErrorResponse{
					Message: "invalid request",
					Errors: []ValidationErrorItem{
						{
							Field: "RecipientName",
							Error: "RecipientName must be greater than 1 character in length",
						},
						{
							Field: "RecipientAccountNumber",
							Error: "RecipientAccountNumber must be a valid numeric value",
						},
						{
							Field: "RecipientBankCode",
							Error: "RecipientBankCode must be at least 1 character in length",
						},
					},
				}

				jsonByte, _ := json.Marshal(res)
				return string(jsonByte)
			}(),
			mock: func() {

			},
		},
		{
			name: "invalid request schema",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: "",
			},
			wantStatus: http.StatusBadRequest,
			want: func() string {
				res := ErrorResponse{
					Message: "invalid request",
				}

				jsonByte, _ := json.Marshal(res)
				return string(jsonByte)
			}(),
			mock: func() {

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			controller := DisbursementController{
				disbursementUsecase: tt.fields.disbursementUsecase,
				validator:           tt.fields.validator,
			}

			router := gin.New()
			router.POST("/test", controller.Disburse)

			requestBody, _ := json.Marshal(tt.args.req)

			req, err := http.NewRequest("POST", "/test", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			respRecorder := httptest.NewRecorder()

			router.ServeHTTP(respRecorder, req)

			assert.Equal(t, tt.wantStatus, respRecorder.Code)

			assert.Equal(t, tt.want, respRecorder.Body.String())
		})
	}
}

func TestDisbursementController_HandleBankCallback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDisbursementUsecase := mock_usecase.NewMockDisbursement(ctrl)

	type fields struct {
		disbursementUsecase usecase.Disbursement
		validator           validator.Validator
	}
	type args struct {
		req interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       string
		wantStatus int
		mock       func()
	}{
		{
			name: "successfully Disburse",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: BankTransferCallbackRequest{
					TransactionId: "txn-id-1",
					Status:        "COMPLETED",
				},
			},
			wantStatus: http.StatusOK,
			want:       "",
			mock: func() {
				mockDisbursementUsecase.EXPECT().ProcessBankCallback(gomock.Any(), usecase.BankCallbackData{
					TransactionId: "txn-id-1",
					Status:        "COMPLETED",
				}).Return(nil)
			},
		},
		{
			name: "error when process callback",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: BankTransferCallbackRequest{
					TransactionId: "txn-id-1",
					Status:        "COMPLETED",
				},
			},
			wantStatus: http.StatusInternalServerError,
			want: func() string {
				res := ErrorResponse{Message: "error when processing bank callback"}

				jsonByte, _ := json.Marshal(res)
				return string(jsonByte)
			}(),
			mock: func() {
				mockDisbursementUsecase.EXPECT().ProcessBankCallback(gomock.Any(), usecase.BankCallbackData{
					TransactionId: "txn-id-1",
					Status:        "COMPLETED",
				}).Return(errors.New("error when processing bank callback"))
			},
		},
		{
			name: "invalid request validation",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: BankTransferCallbackRequest{
					TransactionId: "",
					Status:        "",
				},
			},
			wantStatus: http.StatusBadRequest,
			want: func() string {
				res := ValidationErrorResponse{
					Message: "invalid request",
					Errors: []ValidationErrorItem{
						{
							Field: "TransactionId",
							Error: "TransactionId must be at least 1 character in length",
						},
						{
							Field: "Status",
							Error: "Status must be at least 1 character in length",
						},
					},
				}

				jsonByte, _ := json.Marshal(res)
				return string(jsonByte)
			}(),
			mock: func() {

			},
		},
		{
			name: "invalid request schema",
			fields: fields{
				disbursementUsecase: mockDisbursementUsecase,
				validator:           validator.NewValidator(),
			},
			args: args{
				req: "",
			},
			wantStatus: http.StatusBadRequest,
			want: func() string {
				res := ErrorResponse{
					Message: "invalid request",
				}

				jsonByte, _ := json.Marshal(res)
				return string(jsonByte)
			}(),
			mock: func() {

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			controller := DisbursementController{
				disbursementUsecase: tt.fields.disbursementUsecase,
				validator:           tt.fields.validator,
			}

			router := gin.New()
			router.POST("/test", controller.HandleBankCallback)

			requestBody, _ := json.Marshal(tt.args.req)

			req, err := http.NewRequest("POST", "/test", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			respRecorder := httptest.NewRecorder()

			router.ServeHTTP(respRecorder, req)

			assert.Equal(t, tt.wantStatus, respRecorder.Code)

			assert.Equal(t, tt.want, respRecorder.Body.String())
		})
	}
}
