
------------------------------------------- parent_or_guardian_info table under base schema ---------------------------------------------------

CREATE TABLE base.parent_or_guardian_info
(
    parent_id character varying(50) COLLATE pg_catalog."default" NOT NULL,
    admission_id character varying(50) COLLATE pg_catalog."default" NOT NULL,
    parent_or_guardian_salutation character varying(50) COLLATE pg_catalog."default" NOT NULL,
    parent_or_guardian_first_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    parent_or_guardian_first_middle character varying(100) COLLATE pg_catalog."default" NOT NULL,
    parent_or_guardian_first_last character varying(100) COLLATE pg_catalog."default" NOT NULL,
    parent_or_guardian_phone_no character varying(50) COLLATE pg_catalog."default" NOT NULL,
    parent_or_guardian_email character varying(50) COLLATE pg_catalog."default" NOT NULL,
    gender character varying(10) COLLATE pg_catalog."default" NOT NULL,
    phone_number character varying(50) COLLATE pg_catalog."default" NOT NULL,
    email_id character varying(50) COLLATE pg_catalog."default" NOT NULL,
    address character varying(200) COLLATE pg_catalog."default" NOT NULL,
    city character varying(50) COLLATE pg_catalog."default" NOT NULL,
    state character varying(50) COLLATE pg_catalog."default" NOT NULL,
    country character varying(50) COLLATE pg_catalog."default" NOT NULL,
    postal_code character varying(10) COLLATE pg_catalog."default" NOT NULL,
    status character varying(20) COLLATE pg_catalog."default",
    created_by character varying(50) COLLATE pg_catalog."default",
    created_time timestamp with time zone,
    updated_by character varying(50) COLLATE pg_catalog."default",
    updated_time timestamp with time zone,
    CONSTRAINT parent_id_pkey PRIMARY KEY (parent_id),
    CONSTRAINT parent_or_guardian_info_admission_id_fkey FOREIGN KEY (admission_id)
        REFERENCES base.student_admission (admission_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

student_admission_records(parent_table to student_table), marks_cards_records,previous_academic_records,scholarships,sports,library,hostel,transportation,participation