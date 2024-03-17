package usecase

import (
	"context"
	"errors"
	"github.com/nobbyphala/Brick/domain"
	"github.com/nobbyphala/Brick/external/database"
	"github.com/nobbyphala/Brick/mock"
	mock_api "github.com/nobbyphala/Brick/mock/api"
	mock_repository "github.com/nobbyphala/Brick/mock/repository"
	"github.com/nobbyphala/Brick/usecase/api"
	"github.com/nobbyphala/Brick/usecase/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestNewDisbursement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankApi := mock_api.NewMockBank(ctrl)
	mockDisbursementRepo := mock_repository.NewMockDisbursement(ctrl)

	type args struct {
		deps DisbursementDeps
	}
	tests := []struct {
		name string
		args args
		want *disbursementUsecase
	}{
		{
			name: "new disbursement usecase",
			args: args{deps: DisbursementDeps{
				BankApi:                mockBankApi,
				DisbursementRepository: mockDisbursementRepo,
			}},
			want: &disbursementUsecase{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDisbursement(tt.args.deps)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_disbursementUsecase_Disburse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankApi := mock_api.NewMockBank(ctrl)
	mockDisbursementRepo := mock_repository.NewMockDisbursement(ctrl)

	type fields struct {
		bankApi                api.Bank
		disbursementRepository repository.Disbursement
	}
	type args struct {
		ctx          context.Context
		disbursement domain.Disbursement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Disbursement
		wantErr error
		mock    func()
	}{
		{
			name: "success disburse",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				ctx: context.TODO(),
				disbursement: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					Amount:                 60000,
				},
			},
			want: domain.Disbursement{
				Id:                     "disb-id-1",
				RecipientName:          "Nobby Phala",
				RecipientAccountNumber: "6789567",
				RecipientBankCode:      "Bank A",
				BankTransactionId:      "txn-id-1",
				Amount:                 60000,
				Status:                 1,
			},
			wantErr: nil,
			mock: func() {
				mockBankApi.EXPECT().VerifyAccount(gomock.Any(), api.VerifyAccountRequest{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
				}).Return(api.VerifyAccountResponse{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
					AccountStatus:       "status: account verified",
				}, nil)
				mockBankApi.EXPECT().TransferMoney(gomock.Any(), api.TransferRequest{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
					DestinationBankCode: "Bank A",
					Amount:              60000,
				}).Return(api.TransferResponse{
					TransactionId:       "txn-id-1",
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
					DestinationBankCode: "BANK A",
					Amount:              60000,
					TransferStatus:      "COMPLETED",
				}, nil)
				mockDisbursementRepo.EXPECT().Insert(gomock.Any(), domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "txn-id-1",
					Amount:                 60000,
					Status:                 1,
				}).Return("disb-id-1", nil)
			},
		},
		{
			name: "error when insert disbursement to database",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				ctx: context.TODO(),
				disbursement: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					Amount:                 60000,
				},
			},
			want:    domain.Disbursement{},
			wantErr: errors.New("error when try to disburse"),
			mock: func() {
				mockBankApi.EXPECT().VerifyAccount(gomock.Any(), api.VerifyAccountRequest{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
				}).Return(api.VerifyAccountResponse{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
					AccountStatus:       "status: account verified",
				}, nil)
				mockBankApi.EXPECT().TransferMoney(gomock.Any(), api.TransferRequest{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
					DestinationBankCode: "Bank A",
					Amount:              60000,
				}).Return(api.TransferResponse{
					TransactionId:       "txn-id-1",
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
					DestinationBankCode: "BANK A",
					Amount:              60000,
					TransferStatus:      "COMPLETED",
				}, nil)
				mockDisbursementRepo.EXPECT().Insert(gomock.Any(), domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "txn-id-1",
					Amount:                 60000,
					Status:                 1,
				}).Return("", errors.New("error insert"))
			},
		},
		{
			name: "error when transfer money to bank partner",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				ctx: context.TODO(),
				disbursement: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					Amount:                 60000,
				},
			},
			want:    domain.Disbursement{},
			wantErr: errors.New("temporary bank network error"),
			mock: func() {
				mockBankApi.EXPECT().VerifyAccount(gomock.Any(), api.VerifyAccountRequest{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
				}).Return(api.VerifyAccountResponse{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
					AccountStatus:       "status: account verified",
				}, nil)
				mockBankApi.EXPECT().TransferMoney(gomock.Any(), api.TransferRequest{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
					DestinationBankCode: "Bank A",
					Amount:              60000,
				}).Return(api.TransferResponse{
					TransactionId:       "txn-id-1",
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
					DestinationBankCode: "BANK A",
					Amount:              60000,
					TransferStatus:      "COMPLETED",
				}, errors.New("error api transfer money"))
			},
		},
		{
			name: "destination account not verified",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				ctx: context.TODO(),
				disbursement: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					Amount:                 60000,
				},
			},
			want:    domain.Disbursement{},
			wantErr: errors.New("error bank account not found"),
			mock: func() {
				mockBankApi.EXPECT().VerifyAccount(gomock.Any(), api.VerifyAccountRequest{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
				}).Return(api.VerifyAccountResponse{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "6789567",
					AccountStatus:       "status: account not found",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			disb := disbursementUsecase{
				bankApi:                tt.fields.bankApi,
				disbursementRepository: tt.fields.disbursementRepository,
			}
			got, err := disb.Disburse(tt.args.ctx, tt.args.disbursement)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_disbursementUsecase_ProcessBankCallback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSQL := mock.NewMockSQLDatabase(ctrl)
	mockBankApi := mock_api.NewMockBank(ctrl)
	mockDisbursementRepo := mock_repository.NewMockDisbursement(ctrl)
	mockUtilRepo := mock_repository.NewMockUtils(ctrl)

	type fields struct {
		bankApi                api.Bank
		disbursementRepository repository.Disbursement
		utilsRepository        repository.Utils
	}
	type args struct {
		ctx          context.Context
		bankCallback BankCallbackData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
		mock    func()
	}{
		{
			name: "disbursement status updated",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
				utilsRepository:        mockUtilRepo,
			},
			args: args{
				ctx: context.TODO(),
				bankCallback: BankCallbackData{
					TransactionId: "txn-id-1",
					Status:        "COMPLETED",
				},
			},
			wantErr: nil,
			mock: func() {
				mockDisbursementRepo.EXPECT().WithTx(gomock.Any()).Return(mockDisbursementRepo).Times(2)
				mockDisbursementRepo.EXPECT().GetByTransactionId(gomock.Any(), "txn-id-1").Return(&domain.Disbursement{
					Id:                     "disb-id-1",
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "txn-id-1",
					Amount:                 60000,
					Status:                 1,
				}, nil)
				mockDisbursementRepo.EXPECT().UpdateById(gomock.Any(), "disb-id-1", domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "txn-id-1",
					Amount:                 60000,
					Status:                 2,
				}).Return(nil)
				mockUtilRepo.EXPECT().RunWithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, handler func(Tx database.SQLDatabase) error) error {
					return handler(mockSQL)
				})
			},
		},
		{
			name: "error when update status",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
				utilsRepository:        mockUtilRepo,
			},
			args: args{
				ctx: context.TODO(),
				bankCallback: BankCallbackData{
					TransactionId: "txn-id-1",
					Status:        "COMPLETED",
				},
			},
			wantErr: errors.New("error when updated disbursement status"),
			mock: func() {
				mockDisbursementRepo.EXPECT().WithTx(gomock.Any()).Return(mockDisbursementRepo).Times(2)
				mockDisbursementRepo.EXPECT().GetByTransactionId(gomock.Any(), "txn-id-1").Return(&domain.Disbursement{
					Id:                     "disb-id-1",
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "txn-id-1",
					Amount:                 60000,
					Status:                 1,
				}, nil)
				mockDisbursementRepo.EXPECT().UpdateById(gomock.Any(), "disb-id-1", domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "txn-id-1",
					Amount:                 60000,
					Status:                 2,
				}).Return(errors.New("error update database"))
				mockUtilRepo.EXPECT().RunWithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, handler func(Tx database.SQLDatabase) error) error {
					return handler(mockSQL)
				})
			},
		},
		{
			name: "error disbursement status not PENDING",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
				utilsRepository:        mockUtilRepo,
			},
			args: args{
				ctx: context.TODO(),
				bankCallback: BankCallbackData{
					TransactionId: "txn-id-1",
					Status:        "COMPLETED",
				},
			},
			wantErr: errors.New("error invalid disbursement status"),
			mock: func() {
				mockDisbursementRepo.EXPECT().WithTx(gomock.Any()).Return(mockDisbursementRepo).Times(1)
				mockUtilRepo.EXPECT().RunWithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, handler func(Tx database.SQLDatabase) error) error {
					return handler(mockSQL)
				})
				mockDisbursementRepo.EXPECT().GetByTransactionId(gomock.Any(), "txn-id-1").Return(&domain.Disbursement{
					Id:                     "disb-id-1",
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "txn-id-1",
					Amount:                 60000,
					Status:                 2,
				}, nil)
			},
		},
		{
			name: "error disbursement not found",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
				utilsRepository:        mockUtilRepo,
			},
			args: args{
				ctx: context.TODO(),
				bankCallback: BankCallbackData{
					TransactionId: "txn-id-1",
					Status:        "COMPLETED",
				},
			},
			wantErr: errors.New("error disbursement not found"),
			mock: func() {
				mockDisbursementRepo.EXPECT().WithTx(gomock.Any()).Return(mockDisbursementRepo).Times(1)
				mockUtilRepo.EXPECT().RunWithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, handler func(Tx database.SQLDatabase) error) error {
					return handler(mockSQL)
				})
				mockDisbursementRepo.EXPECT().GetByTransactionId(gomock.Any(), "txn-id-1").Return(nil, nil)
			},
		},
		{
			name: "error when get existing disbursement from database",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
				utilsRepository:        mockUtilRepo,
			},
			args: args{
				ctx: context.TODO(),
				bankCallback: BankCallbackData{
					TransactionId: "txn-id-1",
					Status:        "COMPLETED",
				},
			},
			wantErr: errors.New("error when processing bank callback"),
			mock: func() {
				mockDisbursementRepo.EXPECT().WithTx(gomock.Any()).Return(mockDisbursementRepo)
				mockUtilRepo.EXPECT().RunWithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, handler func(Tx database.SQLDatabase) error) error {
					return handler(mockSQL)
				})
				mockDisbursementRepo.EXPECT().GetByTransactionId(gomock.Any(), "txn-id-1").Return(nil, errors.New("error get from database"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			disb := disbursementUsecase{
				bankApi:                tt.fields.bankApi,
				disbursementRepository: tt.fields.disbursementRepository,
				utilsRepository:        tt.fields.utilsRepository,
			}
			err := disb.ProcessBankCallback(tt.args.ctx, tt.args.bankCallback)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_disbursementUsecase_VerifyDisbursement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankApi := mock_api.NewMockBank(ctrl)
	mockDisbursementRepo := mock_repository.NewMockDisbursement(ctrl)

	type fields struct {
		bankApi                api.Bank
		disbursementRepository repository.Disbursement
	}
	type args struct {
		ctx          context.Context
		disbursement domain.Disbursement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
		mock    func()
	}{
		{
			name: "disbursement verified",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				ctx: context.TODO(),
				disbursement: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "98765",
				},
			},
			wantErr: nil,
			mock: func() {
				mockBankApi.EXPECT().VerifyAccount(gomock.Any(), api.VerifyAccountRequest{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "98765",
				}).Return(api.VerifyAccountResponse{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "98765",
					AccountStatus:       "status: account verified",
				}, nil)
			},
		},
		{
			name: "disbursement bank account not found",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				ctx: context.TODO(),
				disbursement: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "98765",
				},
			},
			wantErr: errors.New("error bank account not found"),
			mock: func() {
				mockBankApi.EXPECT().VerifyAccount(gomock.Any(), api.VerifyAccountRequest{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "98765",
				}).Return(api.VerifyAccountResponse{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "98765",
					AccountStatus:       "status: account not found",
				}, nil)
			},
		},
		{
			name: "disbursement bank account blocked",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				ctx: context.TODO(),
				disbursement: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "98765",
				},
			},
			wantErr: errors.New("error bank account is blocked"),
			mock: func() {
				mockBankApi.EXPECT().VerifyAccount(gomock.Any(), api.VerifyAccountRequest{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "98765",
				}).Return(api.VerifyAccountResponse{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "98765",
					AccountStatus:       "status: account blocked",
				}, nil)
			},
		},
		{
			name: "disbursement bank account unknown status",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				ctx: context.TODO(),
				disbursement: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "98765",
				},
			},
			wantErr: errors.New("error bank account not found"),
			mock: func() {
				mockBankApi.EXPECT().VerifyAccount(gomock.Any(), api.VerifyAccountRequest{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "98765",
				}).Return(api.VerifyAccountResponse{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "98765",
					AccountStatus:       "status: unknown",
				}, nil)
			},
		},
		{
			name: "error from bank partner",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				ctx: context.TODO(),
				disbursement: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "98765",
				},
			},
			wantErr: errors.New("error when trying verify disbursement"),
			mock: func() {
				mockBankApi.EXPECT().VerifyAccount(gomock.Any(), api.VerifyAccountRequest{
					AccountHolderName:   "Nobby Phala",
					AccountHolderNumber: "98765",
				}).Return(api.VerifyAccountResponse{}, errors.New("error api"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			disb := disbursementUsecase{
				bankApi:                tt.fields.bankApi,
				disbursementRepository: tt.fields.disbursementRepository,
			}
			err := disb.VerifyDisbursement(tt.args.ctx, tt.args.disbursement)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_disbursementUsecase_mapTransferStatusToDisbursementStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankApi := mock_api.NewMockBank(ctrl)
	mockDisbursementRepo := mock_repository.NewMockDisbursement(ctrl)

	type fields struct {
		bankApi                api.Bank
		disbursementRepository repository.Disbursement
	}
	type args struct {
		transferStatus api.TransferStatus
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   domain.DisbursementStatus
	}{
		{
			name: "transfer status COMPLETED",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				transferStatus: api.TransferStatusCompleted,
			},
			want: domain.DisbursementStatusCompleted,
		},
		{
			name: "transfer status ACCEPTED",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				transferStatus: api.TransferStatusAccepted,
			},
			want: domain.DisbursementStatusPending,
		},
		{
			name: "transfer status FAILED",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				transferStatus: api.TransferStatusFailed,
			},
			want: domain.DisbursementStatusFailed,
		},
		{
			name: "transfer status REJECTED",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				transferStatus: api.TransferStatusRejected,
			},
			want: domain.DisbursementStatusRejected,
		},
		{
			name: "transfer status unknown",
			fields: fields{
				bankApi:                mockBankApi,
				disbursementRepository: mockDisbursementRepo,
			},
			args: args{
				transferStatus: api.TransferStatus("random status"),
			},
			want: domain.DisbursementStatusUnknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disb := disbursementUsecase{
				bankApi:                tt.fields.bankApi,
				disbursementRepository: tt.fields.disbursementRepository,
			}
			got := disb.mapTransferStatusToDisbursementStatus(tt.args.transferStatus)
			assert.Equal(t, tt.want, got)
		})
	}
}
