package repository

const (
	queryInsertDisbursement = `
	INSERT INTO
		disbursement
		(
		 recipient_name, 
		 recipient_account_number, 
		 recipient_bank_code, 
		 transfer_channel, 
		 bank_transaction_id, 
		 amount,
		 status,
		 created_at,
		 updated_at
		 )
	VALUES
		($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	RETURNING
		id`

	queryUpdateDisbursement = `
	UPDATE
		disbursement
	SET
		recipient_name = $1, 
		recipient_account_number = $2, 
		recipient_bank_code = $3, 
		transfer_channel = $4, 
		bank_transaction_id = $5, 
		amount = $6,
		status = $7
	WHERE
		id = $8`

	querySelectByBankTransactionId = `
	SELECT
		*
	FROM
		disbursement
	WHERE
		bank_transaction_id = $1`
)