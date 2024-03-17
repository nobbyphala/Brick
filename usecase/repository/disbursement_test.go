package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/nobbyphala/Brick/domain"
	"github.com/nobbyphala/Brick/external/database"
	"github.com/nobbyphala/Brick/mock"
	"github.com/nobbyphala/Brick/usecase/repository/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestNewDisbursement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockSQLDatabase(ctrl)

	type args struct {
		deps DisbursementDeps
	}
	tests := []struct {
		name string
		args args
		want *disbursementRepository
	}{
		{
			name: "new disbursement repository",
			args: args{
				deps: DisbursementDeps{DB: mockDB},
			},
			want: &disbursementRepository{
				db: mockDB,
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

func Test_disbursementRepository_GetByTransactionId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockSQLDatabase(ctrl)

	type fields struct {
		db database.SQLDatabase
	}
	type args struct {
		ctx               context.Context
		bankTransactionId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Disbursement
		wantErr error
		mock    func()
	}{
		{
			name: "get existing disbursement",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx:               context.TODO(),
				bankTransactionId: "txn-id-1",
			},
			wantErr: nil,
			want: &domain.Disbursement{
				Id:                     "disb-id-1",
				RecipientName:          "Nobby Phala",
				RecipientAccountNumber: "79823469",
				RecipientBankCode:      "Bank A",
				BankTransactionId:      "txn-id-1",
				Amount:                 1000000,
				Status:                 1,
			},
			mock: func() {
				mockDB.EXPECT().Get(gomock.Any(), gomock.Eq(&model.Disbursement{}), `
	SELECT
		*
	FROM
		disbursement
	WHERE
		bank_transaction_id = $1`, "txn-id-1").DoAndReturn(func(
					ctx context.Context,
					dest interface{},
					query string,
					args ...interface{},
				) error {
					res := dest.(*model.Disbursement)

					*res = model.Disbursement{
						Id:                     "disb-id-1",
						RecipientName:          "Nobby Phala",
						RecipientAccountNumber: "79823469",
						RecipientBankCode:      "Bank A",
						BankTransactionId:      "txn-id-1",
						Amount:                 1000000,
						Status:                 1,
					}

					return nil
				},
				)
			},
		},
		{
			name: "non existing record",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx:               context.TODO(),
				bankTransactionId: "txn-id-1",
			},
			wantErr: nil,
			want:    nil,
			mock: func() {
				mockDB.EXPECT().Get(gomock.Any(), gomock.Eq(&model.Disbursement{}), `
	SELECT
		*
	FROM
		disbursement
	WHERE
		bank_transaction_id = $1`, "txn-id-1").Return(sql.ErrNoRows)
			},
		},
		{
			name: "unknown error from driver",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx:               context.TODO(),
				bankTransactionId: "txn-id-1",
			},
			wantErr: errors.New("sql error"),
			want:    nil,
			mock: func() {
				mockDB.EXPECT().Get(gomock.Any(), gomock.Eq(&model.Disbursement{}), `
	SELECT
		*
	FROM
		disbursement
	WHERE
		bank_transaction_id = $1`, "txn-id-1").Return(errors.New("sql error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			disb := disbursementRepository{
				db: tt.fields.db,
			}
			got, err := disb.GetByTransactionId(tt.args.ctx, tt.args.bankTransactionId)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_disbursementRepository_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockSQLDatabase(ctrl)
	mockRow := mock.NewMockRow(ctrl)

	type fields struct {
		db database.SQLDatabase
	}
	type args struct {
		ctx          context.Context
		disbursement domain.Disbursement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
		mock    func()
	}{
		{
			name: "success insert",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx: context.TODO(),
				disbursement: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "79823469",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "txn-id-1",
					Amount:                 1000000,
					Status:                 1,
				},
			},
			wantErr: nil,
			want:    "disb-id-1",
			mock: func() {
				mockRow.EXPECT().Scan(gomock.Any()).DoAndReturn(func(dest ...any) error {
					res := dest[0].(*string)
					*res = "disb-id-1"
					return nil
				})

				mockDB.EXPECT().Query(gomock.Any(), `
	INSERT INTO
		disbursement
		(
		 recipient_name, 
		 recipient_account_number, 
		 recipient_bank_code, 
		 bank_transaction_id, 
		 amount,
		 status,
		 created_at,
		 updated_at
		 )
	VALUES
		($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	RETURNING
		id`, gomock.Any()).Return(mockRow)
			},
		},
		{
			name: "error insert from driver",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx: context.TODO(),
				disbursement: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "79823469",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "txn-id-1",
					Amount:                 1000000,
					Status:                 1,
				},
			},
			wantErr: errors.New("error scan"),
			want:    "",
			mock: func() {
				mockRow.EXPECT().Scan(gomock.Any()).DoAndReturn(func(dest ...any) error {
					return errors.New("error scan")
				})

				mockDB.EXPECT().Query(gomock.Any(), `
	INSERT INTO
		disbursement
		(
		 recipient_name, 
		 recipient_account_number, 
		 recipient_bank_code, 
		 bank_transaction_id, 
		 amount,
		 status,
		 created_at,
		 updated_at
		 )
	VALUES
		($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	RETURNING
		id`, gomock.Any()).Return(mockRow)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			disb := disbursementRepository{
				db: tt.fields.db,
			}
			got, err := disb.Insert(tt.args.ctx, tt.args.disbursement)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_disbursementRepository_UpdateById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockSQLDatabase(ctrl)
	mockResult := mock.NewMockResult(ctrl)

	type fields struct {
		db database.SQLDatabase
	}
	type args struct {
		ctx         context.Context
		id          string
		updatedData domain.Disbursement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
		mock    func()
	}{
		{
			name: "success update",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx: context.TODO(),
				id:  "disb-id-1",
				updatedData: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "45678",
					Amount:                 60000,
					Status:                 2,
				},
			},
			wantErr: nil,
			mock: func() {
				mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
				mockDB.EXPECT().Exec(gomock.Any(), `
	UPDATE
		disbursement
	SET
		recipient_name = $1, 
		recipient_account_number = $2, 
		recipient_bank_code = $3, 
		bank_transaction_id = $4, 
		amount = $5,
		status = $6,
		updated_at = CURRENT_TIMESTAMP
	WHERE
		id = $8`, gomock.Any()).Return(mockResult, nil)
			},
		},
		{
			name: "no row affected",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx: context.TODO(),
				id:  "disb-id-1",
				updatedData: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "45678",
					Amount:                 60000,
					Status:                 2,
				},
			},
			wantErr: errors.New("error expected there row be affected but got none"),
			mock: func() {
				mockResult.EXPECT().RowsAffected().Return(int64(0), nil)
				mockDB.EXPECT().Exec(gomock.Any(), `
	UPDATE
		disbursement
	SET
		recipient_name = $1, 
		recipient_account_number = $2, 
		recipient_bank_code = $3, 
		bank_transaction_id = $4, 
		amount = $5,
		status = $6,
		updated_at = CURRENT_TIMESTAMP
	WHERE
		id = $8`, gomock.Any()).Return(mockResult, nil)
			},
		},
		{
			name: "error when get affected row",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx: context.TODO(),
				id:  "disb-id-1",
				updatedData: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "45678",
					Amount:                 60000,
					Status:                 2,
				},
			},
			wantErr: errors.New("error row affected"),
			mock: func() {
				mockResult.EXPECT().RowsAffected().Return(int64(0), errors.New("error row affected"))
				mockDB.EXPECT().Exec(gomock.Any(), `
	UPDATE
		disbursement
	SET
		recipient_name = $1, 
		recipient_account_number = $2, 
		recipient_bank_code = $3, 
		bank_transaction_id = $4, 
		amount = $5,
		status = $6,
		updated_at = CURRENT_TIMESTAMP
	WHERE
		id = $8`, gomock.Any()).Return(mockResult, nil)
			},
		},
		{
			name: "error exec the query",
			fields: fields{
				db: mockDB,
			},
			args: args{
				ctx: context.TODO(),
				id:  "disb-id-1",
				updatedData: domain.Disbursement{
					RecipientName:          "Nobby Phala",
					RecipientAccountNumber: "6789567",
					RecipientBankCode:      "Bank A",
					BankTransactionId:      "45678",
					Amount:                 60000,
					Status:                 2,
				},
			},
			wantErr: errors.New("error exec query"),
			mock: func() {
				mockDB.EXPECT().Exec(gomock.Any(), `
	UPDATE
		disbursement
	SET
		recipient_name = $1, 
		recipient_account_number = $2, 
		recipient_bank_code = $3, 
		bank_transaction_id = $4, 
		amount = $5,
		status = $6,
		updated_at = CURRENT_TIMESTAMP
	WHERE
		id = $8`, gomock.Any()).Return(mockResult, errors.New("error exec query"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			disb := disbursementRepository{
				db: tt.fields.db,
			}
			err := disb.UpdateById(tt.args.ctx, tt.args.id, tt.args.updatedData)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
