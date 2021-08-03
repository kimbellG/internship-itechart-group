/*
	Write a query that will choose patients with only one visit.
*/
SELECT Patients.* 
	FROM (SELECT patient AS p
		FROM Visits
		GROUP BY Visits.patient
		HAVING COUNT(Visits.Patient) = 1
	) t INNER JOIN Patients ON Patients.id = t.p;

/*
	Write a query that will choose patients who completed all visits.
*/
SELECT * FROM Patients INNER JOIN Status ON Patients.status = Status.id
	WHERE value = 'Completed';

/*
	Write a query that will choose clinics with patients who performed at least one visit.
*/
SELECT Patients.* 
	FROM (SELECT patient AS p
		FROM Visits
		GROUP BY Visits.patient
		HAVING COUNT(Visits.Patient) > 0
	) t INNER JOIN Patients ON Patients.id = t.p;


/*
	Write a query that will choose clinics with no patients.
*/
SELECT * FROM clinics WHERE id NOT IN (SELECT DISTINCT clinic_id FROM PatientsOfClinics);


/*
	Write a query that will choose clinics with patients who completed all visits.
*/
SELECT * FROM clinics WHERE id IN 
	(SELECT DISTINCT clinic_id FROM PatientsOfClinics 
		INNER JOIN Patients ON Patients.id = PatientsOfClinics.patient_id
		INNER JOIN Status ON Patients.status = Status.id
		WHERE Status.value = 'Completed'
	);

/*
	Write a query that will show the number of visits performed at each clinic.
*/
SELECT clinics.name, COUNT(Visits.clinic) FROM Visits
	RIGHT JOIN clinics ON clinics.id = Visits.clinic
	GROUP BY Visits.clinic, clinics.name;

/*
	Write a query that will show the number of drugs of each type.
*/
SELECT value, COUNT(drug_type) FROM DrugUnits
	RIGHT JOIN DrugType ON DrugUnits.drug_type = DrugType.id
	GROUP BY drug_type, value;

/*
	Write a query that will show the number of drugs of what type were dispensed.
*/
SELECT value, COUNT(drug) FROM Visits 
	INNER JOIN DrugUnits ON Visits.drug = DrugUnits.id
	RIGHT JOIN DrugType ON DrugUnits.drug_type = DrugType.id
	GROUP BY drug, value;

/*
	Write a query that will show the number of patients of each status.
*/
SELECT value, COUNT(Patients.status) FROM Patients
	RIGHT JOIN Status ON Patients.status = Status.id
	GROUP BY Patients.status, value;

/*
	Write a query that will select drug units of the drug type that was not dispensed for any patient.
*/
WITH NotDispensedType(id) AS (
	SELECT DrugType.id FROM Visits
	 INNER JOIN DrugUnits ON Visits.drug = DrugUnits.id
	 RIGHT JOIN DrugType ON DrugUnits.drug_type = DrugType.id
	 GROUP BY drug, DrugType.id 
	 HAVING COUNT(drug) = 0

)
SELECT value AS type, * FROM DrugUnits
	INNER JOIN DrugType ON DrugType.id = DrugUnits.drug_type
	WHERE DrugType.id IN (SELECT * FROM NotDispensedType);


/*
	Write a query that will add Expiration Date to the drug units and fill that column.
*/
ALTER TABLE DrugUnits ADD COLUMN IF NOT EXISTS exp_date timestamp;
ALTER TABLE DrugUnits ALTER exp_date SET DEFAULT date_in + storage_life;

/*
 *	UPDATE DrugUnits SET exp_date = date_in + storage_life;
 */

/*
	Write a query that will add Expiration Date to the drug units and fill that column.
*/
SELECT * FROM DrugUnits WHERE exp_date < current_timestamp;
