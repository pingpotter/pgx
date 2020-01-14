\connect dloan_account

SET TIME ZONE 'UTC';

DROP TABLE IF EXISTS public.opened_account_journal;

CREATE TABLE public.opened_account_journal (
    uid character varying(100),
    account_number bigint
    -- is_sent boolean NOT NULL,
    -- create_datetime_stamp timestamp with time zone NOT NULL,
    -- update_datetime_stamp timestamp with time zone NOT NULL,
    -- PRIMARY KEY (uid,account_number)
);

ALTER TABLE public.opened_account_journal OWNER TO postgres;

-- CREATE INDEX idx_opened_account_journal_is_sent ON public.opened_account_journal USING btree (is_sent);

-- CREATE TRIGGER set_timestamp BEFORE UPDATE ON public.opened_account_journal FOR EACH ROW EXECUTE PROCEDURE public.trigger_set_timestamp();