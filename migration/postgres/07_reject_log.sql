\connect dloan_account

SET TIME ZONE 'UTC';

DROP TABLE IF EXISTS public."reject_log";

CREATE TABLE IF NOT EXISTS public."reject_log"
(
    "date" date NOT NULL,
    topic character varying(40) NOT NULL,
    chrono_sequence character varying(40) NOT NULL,
    sequence integer NOT NULL,
    job_id character varying(40) NOT NULL,
    header_service character varying(255) NOT NULL,
    header_service_version character varying(255) NOT NULL,
    "timestamp" timestamp with time zone NOT NULL,
    message text NOT NULL,
    error_code integer NOT NULL,
    error_category character varying(100) NOT NULL,
    error_description character varying(250) NOT NULL,
    error_caused_by text,
    PRIMARY KEY ("date", topic, chrono_sequence, sequence)
);

ALTER TABLE public."reject_log" OWNER TO postgres;