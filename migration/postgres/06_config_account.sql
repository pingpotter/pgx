\connect dloan_account

SET TIME ZONE 'UTC';

DROP TABLE IF EXISTS public.config_account;

CREATE TABLE public.config_account (
    key_value character varying(40) PRIMARY KEY,
    account_number bigint
); 


ALTER TABLE public.config_account OWNER TO postgres;

INSERT INTO public.config_account (key_value,account_number) VALUES ('init-account-number',60000000000);