
------------------------------------------- department table under base schema ---------------------------------------------------


CREATE TABLE base.department
(
    department_id character varying(20) COLLATE pg_catalog."default" NOT NULL,
    institute_id character varying(20) COLLATE pg_catalog."default" NOT NULL,
    department_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    description character varying(500) COLLATE pg_catalog."default" NOT NULL,
    status character varying(20) COLLATE pg_catalog."default",
    created_by character varying(50) COLLATE pg_catalog."default",
    created_time timestamp with time zone,
    updated_by character varying(50) COLLATE pg_catalog."default",
    updated_time timestamp with time zone,
    CONSTRAINT department_id_pkey PRIMARY KEY (department_id),
    CONSTRAINT department_name_uk UNIQUE (department_name),
    CONSTRAINT department_institute_id_fkey FOREIGN KEY (institute_id)
        REFERENCES base.institute (institute_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);


activities
budget
contributions
achievements
