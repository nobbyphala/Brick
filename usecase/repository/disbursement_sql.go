package repository

const (
	queryInsertDisbursement = `
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
		id`

	queryUpdateDisbursement = `
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
		id = $7`

	querySelectByBankTransactionId = `
	SELECT
		*
	FROM
		disbursement
	WHERE
		bank_transaction_id = $1`
)
