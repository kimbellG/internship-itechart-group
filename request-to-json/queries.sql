/*
 *	Write a query to select any row with particular VehicleId.
 */

SELECT jsonb_pretty(value) FROM Vechicles,
	LATERAL jsonb_array_elements(Vechicles.data)
	WHERE value ->> 'VehicleId' = '6109270e31845c39b422cf0a';
	
/*
 *	Write a query to select rows which Metadata.CreatedOn is greater than particular date.	
 */

SELECT jsonb_pretty(value) FROM Vechicles,
	LATERAL jsonb_array_elements(Vechicles.data)
	WHERE (value -> 'Metadata' ->> 'CreatedOn')::date > current_date;

/*
 *	Write a query which groups all price histories for all rows by the PricedBy field and calculates average of OfferedPrice.	
 */

SELECT (el ->> 'VehicleId') AS id, AVG((prices ->> 'OfferedPrice')::float)  FROM Vechicles,
	LATERAL jsonb_array_elements(Vechicles.data) AS el,
	LATERAL jsonb_array_elements(el.value -> 'PriceHistory') AS prices
	GROUP BY el;

/*
 *	Write a query which groups all price histories for all rows by the PricedBy field and calculates average of OfferedPrice.	
 */

SELECT SUM((el ->> 'SoldPrice')::float) AS "All price" FROM Vechicles,
	LATERAL jsonb_array_elements(Vechicles.data) as el
	WHERE el ?& array['IsSold', 'SoldPrice'] AND (el ->> 'IsSold')::boolean;

/*
 *	Write a query to select all rows which have IsSold and SoldPrice fields.	
 */

SELECT jsonb_pretty(el) AS "vechicles on sale" FROM Vechicles,
	LATERAL jsonb_array_elements(Vechicles.data) as el
	WHERE el ?& array['IsSold', 'SoldPrice'];

