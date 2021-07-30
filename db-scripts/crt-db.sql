CREATE TABLE IF NOT EXISTS clinics (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name varchar(250) NOT NULL UNIQUE,
	address varchar(200) NOT NULL UNIQUE,
	phone_number varchar(13) NOT NULL CHECK(phone_number LIKE '+375_________') UNIQUE,
	opendays varchar(20) NOT NULL,
	openup time NOT NULL,
	closing time NOT NULL
);


CREATE TABLE IF NOT EXISTS DrugType (
	id serial PRIMARY KEY,
	value varchar(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS DrugUnits (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name varchar(250) NOT NULL UNIQUE,
	drug_type integer NOT NULL REFERENCES DrugType,
	date_in date NOT NULL DEFAULT current_date,
	storage_life interval,
	cost money NOT NULL,
	manual text
);

CREATE TABLE IF NOT EXISTS Status (
	id serial PRIMARY KEY,
	value text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Gender (
	id serial PRIMARY KEY,
	value text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Patients (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name varchar(200) NOT NULL,
	status integer NOT NULL REFERENCES Status,
	gender integer NOT NULL REFERENCES Gender,
	drug_type integer REFERENCES DrugType NOT NULL,
	start_date timestamp NOT NULL DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS Visits (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	patient uuid NOT NULL REFERENCES Patients,
	clinic uuid NOT NULL REFERENCES clinics,
	drug uuid NOT NULL REFERENCES DrugUnits,
	visit_date timestamp NOT NULL DEFAULT current_timestamp,
	reason text NOT NULL
);

CREATE TYPE Roles AS ENUM ('sponsor', 'investigator', 'manager');

CREATE TABLE IF NOT EXISTS Users (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	email varchar(100) NOT NULL UNIQUE CHECK (email LIKE '%@%.%'),
	password varchar(30) NOT NULL,
	role Roles NOT NULL
);

CREATE TABLE IF NOT EXISTS DrugsOfClinics (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	clinic_id uuid NOT NULL REFERENCES clinics,
	drug_id uuid NOT NULL REFERENCES DrugUnits
);

CREATE TABLE IF NOT EXISTS PatientsOfClinics (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	clinic_id uuid NOT NULL REFERENCES clinics,
	patient_id uuid NOT NULL REFERENCES Patients
);	



