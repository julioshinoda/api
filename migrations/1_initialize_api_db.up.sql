-- public.accounts definition

-- Drop table

-- DROP TABLE public.accounts;

CREATE TABLE if not exists accounts (
	account_id serial NOT NULL,
	document_number varchar(14) NOT NULL,
	CONSTRAINT accounts_pk PRIMARY KEY (account_id),
	CONSTRAINT accounts_un UNIQUE (document_number)
);


-- public.operations_types definition

-- Drop table

-- DROP TABLE public.operations_types;

CREATE TABLE if not exists operations_types (
	operationtype_id serial NOT NULL,
	description varchar NOT NULL,
	CONSTRAINT operations_types_pk PRIMARY KEY (operationtype_id)
);

-- public.transactions definition

-- Drop table

-- DROP TABLE public.transactions;

CREATE TABLE if not exists  transactions (
	transaction_id serial NOT NULL,
	account_id int4 NOT NULL,
	operationtype_id int4 NULL,
	amount float8 NOT NULL,
	event_date timestamp NULL,
	CONSTRAINT transactions_pk PRIMARY KEY (transaction_id)
);