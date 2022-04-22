
------------------------------------------- student_admission table under base schema ---------------------------------------------------

CREATE TABLE base.student_admission
(
    admission_id character varying(50) COLLATE pg_catalog."default" NOT NULL,
    institute_id character varying(20) COLLATE pg_catalog."default" NOT NULL,
    course_id character varying(50) COLLATE pg_catalog."default" NOT NULL,
    department_id character varying(50) COLLATE pg_catalog."default" NOT NULL,
    student_salutation character varying(10) COLLATE pg_catalog."default" NOT NULL,
    first_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    middle_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    last_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    dob date,
    age character varying(10) COLLATE pg_catalog."default" NOT NULL,
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
    CONSTRAINT admission_id_pkey PRIMARY KEY (admission_id),
    CONSTRAINT student_admission_course_id_fkey FOREIGN KEY (course_id)
        REFERENCES base.course (course_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT student_admission_department_id_fkey FOREIGN KEY (department_id)
        REFERENCES base.department (department_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT student_admission_institute_id_fkey FOREIGN KEY (institute_id)
        REFERENCES base.institute (institute_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);



student_admission_records(parent_table to student_table), marks_cards_records,previous_academic_records,scholarships,sports,library,hostel,transportation,participation