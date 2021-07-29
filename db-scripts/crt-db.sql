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
	value varchar(50) NOT NULL UNIQUE,
);

CREATE TABLE IF NOT EXISTS DrugUnits (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name varchar(250) NOT NULL UNIQUE,
	drug_type integer REFERENCES DrugType,
	data_in date NOT NULL DEFAULT current_date,
	storage_life interval,
	location varchar(100) NOT NULL,
	cost money NOT NULL,
	manual text
);
