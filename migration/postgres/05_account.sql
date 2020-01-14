\connect dloan_account

SET TIME ZONE 'UTC';

DROP TABLE IF EXISTS public.account;

CREATE TABLE public.account (
    account_number bigint PRIMARY KEY
    -- product_name character varying(30) NOT NULL,
    -- customer_number character varying(20) NOT NULL,
    -- customer_type character varying(4) NOT NULL,
    -- account_name character varying(800) NOT NULL,
    -- account_branch integer NOT NULL,
    -- response_unit integer,
    -- credit_term_number smallint NOT NULL,
    -- credit_term_unit character varying(20) NOT NULL,
    -- disbursement_account character varying(20),
    -- open_account_jobid character varying(40) NOT NULL,
    -- open_date date NOT NULL,
    -- open_datetime_stamp timestamptz NOT NULL,
    -- application_id character varying(50),
    -- closed_date date,
    -- maturity_date date NOT NULL
); 

ALTER TABLE public.account OWNER TO postgres;