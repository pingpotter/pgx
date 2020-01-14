\connect dloan_account

SET TIME ZONE 'UTC';

DROP TABLE IF EXISTS public.message_log;

CREATE TABLE public.message_log (
    job_id character varying(40) PRIMARY KEY
    -- message_type smallint NOT NULL,
    -- message jsonb,
    -- create_datetime_stamp timestamptz NOT NULL
);


ALTER TABLE public.message_log OWNER TO postgres;


-- CREATE INDEX idx_message_log_create_datetime_stamp 
-- on message_log (create_datetime_stamp);