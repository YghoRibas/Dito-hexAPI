CREATE TABLE ADDRESS (
	address_id SERIAL NOT NULL,
	street varchar NOT NULL,
	address_number int NOT NULL,
	country varchar NOT NULL,
	postal_code int NOT NULL,
	PRIMARY KEY (address_id)
);

CREATE TABLE PERSON (
	person_id SERIAL NOT NULL,
	person_name varchar NOT NULL,
	person_age int8 NOT NULL,
	person_address_id int NOT NULL,
	gender char,
	PRIMARY KEY (person_id),
	FOREIGN KEY (person_address_id) REFERENCES ADDRESS(address_id)
);

CREATE TABLE DEPENDENT (
	dependent_id SERIAL NOT NULL,
	person_id SERIAL NOT NULL,
	dependent_name varchar NOT NULL,
	CONSTRAINT fk_dependent FOREIGN KEY (person_id) REFERENCES PERSON(person_id) ON DELETE CASCADE,
	PRIMARY KEY (person_id, dependent_id)
);

INSERT INTO ADDRESS (street, address_number, country, postal_code) VALUES ('Rua. Miguel Pereira Lacerda', 845, 'Brazil', 31950095);
INSERT INTO ADDRESS (street, address_number, country, postal_code) VALUES ('Av. Joaquim Jose Diniz', 315, 'Brazil', 31910520);
INSERT INTO ADDRESS (street, address_number, country, postal_code) VALUES ('Rua. Guajajaras', 30817, 'Brazil', 30180000);

INSERT INTO PERSON (person_name, person_age, person_address_id, gender)
INSERT INTO PERSON VALUES (DEFAULT, 'Caio Ribeiro', 29, 4, 'M');
INSERT INTO PERSON VALUES (DEFAULT, 'Maria de Almeida Ribeiro', 27, 4, 'F');
INSERT INTO PERSON VALUES (DEFAULT, 'Alberto Carlos', 53, 3, 'M');
INSERT INTO PERSON VALUES (DEFAULT, 'Maria Costa', 27, 2, NULL);

INSERT INTO dependent (person_id, dependent_name) VALUES ( 7, 'Gabriel Costa');
INSERT INTO dependent (person_id, dependent_name) VALUES ( 6, 'Alberto Jr');
INSERT INTO dependent (person_id, dependent_name) VALUES ( 4, 'Fernando Ribeiro');
INSERT INTO dependent (person_id, dependent_name) VALUES ( 4, 'Vinicius Ribeiro');

SELECT person.person_name, person_age, person.gender, 
address.street, address.address_number, address.country, address.postal_code, 
dependent.dependent_name  
FROM PERSON INNER JOIN ADDRESS ON PERSON.person_address_id = ADDRESS.address_id
INNER JOIN DEPENDENT ON PERSON.person_id = DEPENDENT.person_id; 