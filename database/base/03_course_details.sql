
------------------------------------------- course_details table under base schema ---------------------------------------------------

CREATE TABLE base.course
(
    course_id character varying(50) COLLATE pg_catalog."default" NOT NULL,
    institute_id character varying(20) COLLATE pg_catalog."default" NOT NULL,
    course_discipline_id character varying(100) COLLATE pg_catalog."default" NOT NULL,
    course_name character varying(200) COLLATE pg_catalog."default" NOT NULL,
    duration integer NOT NULL,
    description character varying(500) COLLATE pg_catalog."default" NOT NULL,
    added_date date NOT NULL,
    status character varying(20) COLLATE pg_catalog."default",
    created_by character varying(50) COLLATE pg_catalog."default",
    created_time timestamp with time zone,
    updated_by character varying(50) COLLATE pg_catalog."default",
    updated_time timestamp with time zone,
    CONSTRAINT course_id_pkey PRIMARY KEY (course_id),
    CONSTRAINT course_course_discipline_id_fkey FOREIGN KEY (course_discipline_id)
        REFERENCES base.graduation_program (graduation_program_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT course_institute_id_fkey FOREIGN KEY (institute_id)
        REFERENCES base.institute (institute_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);



student_admission_records(parent_table to student_table), marks_cards_records,previous_academic_records,scholarships,sports,library,hostel,transportation,participation



be
mtech
diploma
mba
