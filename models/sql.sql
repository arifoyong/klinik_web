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
