
------------------------------------------- institute table under base schema ---------------------------------------------------
CREATE TABLE base.institute
(
    institute_id character varying(20) COLLATE pg_catalog."default" NOT NULL,
    institute_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    institute_type character varying(100) COLLATE pg_catalog."default" NOT NULL,
    contact_person_salutation character varying(10) COLLATE pg_catalog."default" NOT NULL,
    contact_person_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    contact_person_email_id character varying(400) COLLATE pg_catalog."default",
    contact_person_phone_no character varying(200) COLLATE pg_catalog."default",
    web_url character varying(255) COLLATE pg_catalog."default",
    logo_path character varying(255) COLLATE pg_catalog."default",
    address_line1 character varying(500) COLLATE pg_catalog."default" NOT NULL,
    address_line2 character varying(500) COLLATE pg_catalog."default",
    address_line3 character varying(500) COLLATE pg_catalog."default",
    city character varying(50) COLLATE pg_catalog."default" NOT NULL,
    state character varying(50) COLLATE pg_catalog."default" NOT NULL,
    country character varying(50) COLLATE pg_catalog."default" NOT NULL,
    country_code character varying(50) COLLATE pg_catalog."default" NOT NULL,
    postal_code character varying(50) COLLATE pg_catalog."default" NOT NULL,
    tel_no character varying(50) COLLATE pg_catalog."default" NOT NULL,
    fax_no character varying(50) COLLATE pg_catalog."default",
    institute_start_date date NOT NULL,
    area_in_acres double precision,
    status character varying(20) COLLATE pg_catalog."default",
    created_by character varying(50) COLLATE pg_catalog."default",
    created_time timestamp with time zone default current_timestamp,
    updated_by character varying(50) COLLATE pg_catalog."default",
    updated_time timestamp with time zone default current_timestamp,
    CONSTRAINT institute_id_pkey PRIMARY KEY (institute_id),
    CONSTRAINT institute_name_uk UNIQUE (institute_name)
)


institute_type ==> affiliated/autonomous and it's details in separate table (parent table to institute)


/////////////////////////////////////////////////////////

INSERT INTO base.institute(
	institute_id, institute_name, institute_type, contact_person_salutation, contact_person_name, 
	contact_person_email_id, contact_person_phone_no, web_url, logo_path, address_line1, address_line2, 
	address_line3, city, state, country, country_code, postal_code, tel_no, fax_no, institute_start_date, 
	area_in_acres, status, created_by, created_time, updated_by, updated_time)
	VALUES ('inst2', 'inst_name2', 'engineering', 'Mr/Ms', 'name2', 'abc2@gmail.com', '9875432561', 'url2.xyz', 'logo/path/2', 
			'addr1', 'addr2', 'addr3', 'blore', 'ka', 'in', '123', '560222', '7767898977', '79999999', '10-05-1990', 
			55.5, 'active', 'admin', now(),'admin', now());
			
