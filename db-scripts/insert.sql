INSERT INTO DrugType (value) VALUES
	('A'),
	('B'),
	('C');

INSERT INTO Status (value) VALUES
	('Screened'),
	('Randomized'),
	('EarlyCompleted'),
	('Completed');

INSERT INTO Gender (value) VALUES
	('Male'),
	('Female');

INSERT INTO clinics (name, address, phone_number, opendays, openup, closing) VALUES 
	('St. Mary Hospital', 'Witherspoon Street', '+375123452343', 
		'mon-fri', '08:00', '21:00'),
	('Cincinnati Hospital', 'Burnet Avenue', '+375178912312',
		'everyday', '08:00', '21:00'),
	('The Johns Hopkins Hospital', '401 N Broadway, Baltimore', '+375291332312',
		'mon-fri', '08:00', '21:00'),
	('Cleveland Clinic', '2050E 96th St, Cleveland', '+375291234321',
		'everyday', '08:00', '21:00'),
	('Mayo Clinic', '2nd Street Southwest, Rochester, MN', '+375293123212',
		'everyday', '08:00', '21:00'),
	('MD Anderson Center', '1515 Holcombe Blvd, Houston',
	       	'+375293223212', 'everyday', '08:00', '21:00');

INSERT INTO DrugUnits (name, drug_type, storage_life, cost, manual) SELECT
	'Азатиоприн', id, '1-10', 50.22, 'text text text. Big text' FROM DrugType WHERE value = 'A';
INSERT INTO DrugUnits (name, drug_type, storage_life, cost, manual) SELECT
	'Аллапинин', id, '2-2', 30, 'text text text. Big text' FROM DrugType WHERE value = 'A';
INSERT INTO DrugUnits (name, drug_type, storage_life, cost, manual) SELECT
	'Армин', id, '3-0', 25, 'text text text. Big text' FROM DrugType WHERE value = 'A';
INSERT INTO DrugUnits (name, drug_type, storage_life, cost, manual) SELECT
	'Антропин', id, '2-3', 11, 'text text text. Big text' FROM DrugType WHERE value = 'A';
INSERT INTO DrugUnits (name, drug_type, storage_life, cost, manual) SELECT
	'Бемегрид', id, '10-0', 100, 'text text text. Big text' FROM DrugType WHERE value = 'B';
INSERT INTO DrugUnits (name, drug_type, storage_life, cost, manual) SELECT
	'Бендазол', id, '0-2', 499, 'text text text. Big text' FROM DrugType WHERE value = 'B';
INSERT INTO DrugUnits (name, drug_type, storage_life, cost, manual) SELECT
	'Бензидамин', id, '2-3', 129, 'text text text. Big text' FROM DrugType WHERE value = 'B';
INSERT INTO DrugUnits (name, drug_type, storage_life, cost, manual) SELECT
	'Бенперидол', id, '12-11', 123, 'text text text. Big text' FROM DrugType WHERE value = 'B';
INSERT INTO DrugUnits (name, drug_type, storage_life, cost, manual) SELECT
	'Пантоган', id, '8-2', 223, 'text text text. Big text' FROM DrugType WHERE value = 'C';
INSERT INTO DrugUnits (name, drug_type, storage_life, cost, manual) SELECT
	'Пиритинол', id, '2-3', 12, 'text text text. Big text' FROM DrugType WHERE value = 'C';




INSERT INTO Patients (name, status, gender, drug_type)
	SELECT 'Rob Pike', s.id, g.id, d.id 
	FROM Status AS s, Gender AS g, DrugType AS d
	WHERE s.value = 'Screened' AND g.value = 'Male' AND d.value = 'A';

INSERT INTO Patients (name, status, gender, drug_type)
	SELECT 'Dennis Ritchie', s.id, g.id, d.id 
	FROM Status AS s, Gender AS g, DrugType AS d
	WHERE s.value = 'Randomized' AND g.value = 'Male' AND d.value = 'B';

INSERT INTO Patients (name, status, gender, drug_type)
	SELECT 'Steve Jobs', s.id, g.id, d.id 
	FROM Status AS s, Gender AS g, DrugType AS d
	WHERE s.value = 'EarlyCompleted' AND g.value = 'Male' AND d.value = 'C';

INSERT INTO Patients (name, status, gender, drug_type)
	SELECT 'Ken Thompson', s.id, g.id, d.id 
	FROM Status AS s, Gender AS g, DrugType AS d
	WHERE s.value = 'Completed' AND g.value = 'Male' AND d.value = 'B';

INSERT INTO Patients (name, status, gender, drug_type)
	SELECT 'Barbara Liskov', s.id, g.id, d.id 
	FROM Status AS s, Gender AS g, DrugType AS d
	WHERE s.value = 'Screened' AND g.value = 'Female' AND d.value = 'A';

INSERT INTO Patients (name, status, gender, drug_type)
	SELECT 'Bill Gates', s.id, g.id, d.id 
	FROM Status AS s, Gender AS g, DrugType AS d
	WHERE s.value = 'Completed' AND g.value = 'Male' AND d.value = 'C';



	
INSERT INTO PatientsOfClinics (clinic_id, patient_id)
	SELECT c.id, p.id FROM clinics AS c, Patients AS p
	WHERE c.name = 'St. Mary Hospital' AND p.name = 'Steve Jobs';

INSERT INTO PatientsOfClinics (clinic_id, patient_id)
	SELECT c.id, p.id FROM clinics AS c, Patients AS p
	WHERE c.name = 'Cincinnati Hospital' AND p.name = 'Rob Pike';

INSERT INTO PatientsOfClinics (clinic_id, patient_id)
	SELECT c.id, p.id FROM clinics AS c, Patients AS p
	WHERE c.name = 'The Johns Hopkins Hospital' AND p.name = 'Ken Thompson';

INSERT INTO PatientsOfClinics (clinic_id, patient_id)
	SELECT c.id, p.id FROM clinics AS c, Patients AS p
	WHERE c.name = 'St. Mary Hospital' AND p.name = 'Bill Gates';




INSERT INTO DrugsOfClinics (clinic_id, drug_id)
	SELECT c.id, d.id FROM clinics AS c, DrugUnits AS d
	WHERE c.name = 'Cincinnati Hospital' AND d.name = 'Антропин';

INSERT INTO DrugsOfClinics (clinic_id, drug_id)
	SELECT c.id, d.id FROM clinics AS c, DrugUnits AS d
	WHERE c.name = 'St. Mary Hospital' AND d.name = 'Бендазол';

INSERT INTO DrugsOfClinics (clinic_id, drug_id)
	SELECT c.id, d.id FROM clinics AS c, DrugUnits AS d
	WHERE c.name = 'Mayo Clinic' AND d.name = 'Пантоган';

INSERT INTO DrugsOfClinics (clinic_id, drug_id)
	SELECT c.id, d.id FROM clinics AS c, DrugUnits AS d
	WHERE c.name = 'MD Anderson Center' AND d.name = 'Армин';




INSERT INTO Visits (patient, clinic, drug, reason) 
	SELECT p.id, c.id, d.id, 'Some reason' 
	FROM Patients AS p, clinics AS c, DrugUnits As d
	WHERE c.name = 'St. Mary Hospital' AND p.name = 'Steve Jobs' AND d.name = 'Бендазол';

INSERT INTO Visits (patient, clinic, drug, reason) 
	SELECT p.id, c.id, d.id, 'Some reason' 
	FROM Patients AS p, clinics AS c, DrugUnits As d
	WHERE c.name = 'Cincinnati Hospital' AND p.name = 'Rob Pike' and d.name = 'Антропин';

INSERT INTO Visits (patient, clinic, drug, reason) 
	SELECT p.id, c.id, d.id, 'Some reason' 
	FROM Patients AS p, clinics AS c, DrugUnits As d
	WHERE c.name = 'The Johns Hopkins Hospital' AND p.name = 'Ken Thompson' AND d.name = 'Пиритинол';

INSERT INTO Visits (patient, clinic, drug, reason) 
	SELECT p.id, c.id, d.id, 'Some reason' 
	FROM Patients AS p, clinics AS c, DrugUnits As d
	WHERE c.name = 'St. Mary Hospital' AND p.name = 'Bill Gates' AND d.name = 'Бендазол';

INSERT INTO Visits (patient, clinic, drug, reason) 
	SELECT p.id, c.id, d.id, 'Some reason' 
	FROM Patients AS p, clinics AS c, DrugUnits As d
	WHERE c.name = 'The Johns Hopkins Hospital' AND p.name = 'Ken Thompson' AND d.name = 'Пиритинол';

INSERT INTO Visits (patient, clinic, drug, reason) 
	SELECT p.id, c.id, d.id, 'Some reason' 
	FROM Patients AS p, clinics AS c, DrugUnits As d
	WHERE c.name = 'St. Mary Hospital' AND p.name = 'Bill Gates' AND d.name = 'Бендазол';

