CREATE TABLE IF NOT EXISTS Vechicles (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	data jsonb NOT NULL,
	modifiedOn date NOT NULL DEFAULT current_date
);
