DROP TABLE patients;
CREATE TABLE patients (
  id 	serial primary key,
  firstname 	varchar(50),
  lastname 	varchar(50) NOT NULL,
  ic 	varchar(20) NOT NULL,
  dob 	date NOT NULL,
  email 	varchar(50) NOT NULL, 
  phone 	varchar(10) NOT NULL,
  address 	varchar(100)   );


CREATE TABLE visits (
  id	serial primary key,
  date	timestamp with time zone,
  patient_id	integer REFERENCES patients(id),
  problems	varchar(300),
  diagnosis	varchar(3000,
  prescription_id integer 
);


CREATE TABLE drugs (
  drug_id serial primary key,
  name varchar(50) not null,
  unit_price decimal not null
);