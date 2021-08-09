/*
 *	Write a SQL function to count the required number of drug units to complete the rest patients' visits at the specified clinic
 */

CREATE FUNCTION count_drugs_for_clinic(clinic_id uuid) RETURNS TABLE(name varchar(250), num integer) AS $$
	WITH clinic_visits(id) AS (
		SELECT id FROM Visits WHERE clinic = clinic_id
	)
	SELECT name, COUNT(drug) FROM Visits 
		INNER JOIN DrugUnits ON DrugUnits.id = Visits.Drug
		WHERE Visits.id IN (SELECT * FROM clinic_visits)
		GROUP BY name, drug;
$$ LANGUAGE SQL;


/*
 *	Write a SQL function that will randomly choose the drug type and assign it to the specified patient.
 */
CREATE FUNCTION update_drugtype_of_patient(patient_id uuid) RETURNS void AS $$
	UPDATE Patients 
	SET drug_type=
		(SELECT id FROM DrugType OFFSET random() * (SELECT COUNT(*) FROM DrugType) LIMIT 1)
	WHERE id = patient_id;

$$ LANGUAGE SQL;

/*
 *	Write a SQL function that will randomly choose available drug unit of the specified drug type and assign it to the specified patient.
 */

CREATE FUNCTION choose_random_drug(drug_type_r integer, patient_r uuid) RETURNS void AS $$
	WITH drugUnit AS (
		SELECT id FROM DrugUnits WHERE drug_type = drug_type_r 
			OFFSET random() *
		       	(SELECT COUNT(*) FROM DrugUnits WHERE drug_type = drug_type_r)
		       	LIMIT 1

	), patient_clinic AS (
		SELECT clinic_id FROM PatientsOfClinics WHERE patient_id = patient_r LIMIT 1

	)
	INSERT INTO Visits (patient, clinic, drug, reason) VALUES 
	(patient_r, (SELECT * FROM patient_clinic), (SELECT * FROM drugUnit), 'random choose from system');

$$ LANGUAGE SQL;
