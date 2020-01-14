\connect dloan_account

SET TIME ZONE 'UTC';

DROP TABLE IF EXISTS public.product;

CREATE TABLE public.product (
    product_name character varying(30) PRIMARY KEY,
    product_group character varying(10) NOT NULL,
    product_type character varying(10) NOT NULL,
    product_subtype character varying(10) NOT NULL,
    is_revolving boolean NOT NULL,
    transaction_plan_name character varying(30) NOT NULL,
    interest_plan_name character varying(30) NOT NULL,
    payment_plan_name character varying(30) NOT NULL,
    penalty_plan_name character varying(30) NOT NULL,
    grace_plan_name character varying(30) NOT NULL,
    description character varying(100) NOT NULL
);


ALTER TABLE public.product OWNER TO postgres;
-- TODO check config transaction_plan_name
INSERT INTO public.product (product_name,product_group,product_type,product_subtype,is_revolving,transaction_plan_name,interest_plan_name,payment_plan_name,penalty_plan_name,grace_plan_name,description) VALUES ('7200120090001','ret','per','rev',true,'txnlazada','intlazada','pmtlazada','penalty28','gracelazada','K Loan Lazada');
INSERT INTO public.product (product_name,product_group,product_type,product_subtype,is_revolving,transaction_plan_name,interest_plan_name,payment_plan_name,penalty_plan_name,grace_plan_name,description) VALUES ('7200120090002','ret','per','rev',true,'txnlazada','intlazada','pmtlazada2','penalty28','gracelazada','K Loan Lazada 2');
INSERT INTO public.product (product_name,product_group,product_type,product_subtype,is_revolving,transaction_plan_name,interest_plan_name,payment_plan_name,penalty_plan_name,grace_plan_name,description) VALUES ('7200120090003','ret','per','rev',true,'txnlazada','intlazada','pmtlazada3','penalty28','gracelazada2','K Loan Lazada 3');