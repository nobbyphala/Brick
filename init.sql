-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- public.disbursement definition
CREATE TABLE public.disbursement (
                                     id uuid DEFAULT uuid_generate_v4() NOT NULL,
                                     recipient_name varchar NOT NULL,
                                     recipient_account_number varchar NOT NULL,
                                     recipient_bank_code varchar NOT NULL,
                                     bank_transaction_id varchar NOT NULL,
                                     amount int8 NOT NULL,
                                     status int NULL,
                                     created_at timestamp NOT NULL,
                                     updated_at timestamp NOT NULL,
                                     CONSTRAINT disbursement_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX disbursement_bank_transaction_id_idx ON public.disbursement (bank_transaction_id);