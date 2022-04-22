
------------------------------------------- graduation_program table under base schema ---------------------------------------------------
CREATE TABLE base.graduation_program
(
    graduation_program_id character varying(50) COLLATE pg_catalog."default" NOT NULL,
    graduation_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    graduation_type_added_date date NOT NULL,
    institute_id character varying(20) COLLATE pg_catalog."default" NOT NULL,
    status character varying(20) COLLATE pg_catalog."default",
    created_by character varying(50) COLLATE pg_catalog."default",
    created_time timestamp with time zone,
    updated_by character varying(50) COLLATE pg_catalog."default",
    updated_time timestamp with time zone,
    CONSTRAINT graduation_program_id_pkey PRIMARY KEY (graduation_program_id),
    CONSTRAINT graduation_program_institute_id_fkey FOREIGN KEY (institute_id)
        REFERENCES base.institute (institute_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);
