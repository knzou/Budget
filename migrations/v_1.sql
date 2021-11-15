DO
$body$
BEGIN
-- advantage of using PL/pgsql is that you can group a block of computation and a series of queries - power of precedural language and easy of use of sql
-- Command to Use: psql -h 127.0.0.1 -d kenzou -f migrations/v_1.sql -c "\copy transaction FROM '/Users/kenzou/Desktop/Budget/migrations/transaction.csv' delimiter ',' csv header"
    IF NOT EXISTS (
        SELECT usename
        FROM pg_catalog.pg_user
        WHERE usename = 'test_user'
    )
    THEN
        -- TODO: SHOULD BE PULL FROM CONFIG
        CREATE ROLE test_user WITH SUPERUSER CREATEDB LOGIN PASSWORD '123';
    END IF;

    -- Tables
    CREATE UNLOGGED TABLE IF NOT EXISTS person(
	pid SERIAL PRIMARY KEY,
	name VARCHAR NOT NULL
    );

    CREATE UNLOGGED TABLE IF NOT EXISTS category (
	catId SERIAL PRIMARY KEY,
	name VARCHAR NOT NULL,
	typeId INT
    );

    CREATE UNLOGGED TABLE IF NOT EXISTS transaction (
	tranId SERIAL PRIMARY KEY,
	catId INT REFERENCES category (catId),
    transDate DATE NOT NULL DEFAULT CURRENT_DATE,
	amount INT
    );

    -- HARDCODE Data
    IF NOT EXISTS (
        SELECT * 
        FROM person
    )
    THEN 
        INSERT INTO person(name)
        VALUES('kenzou'),
        ('Georgeanna Alexandria'),
        ('Rodney Jed'),
        ('Ken Dotty'),
        ('Connell Vance'),
        ('Cairo Ruth'),
        ('Rod Jedi'),
        ('Ken Geordie'),
        ('Iggy Aimee'),
        ('Rona Nydia'),
        ('Jeri Raelyn'),
        ('Ormonde Karina'),
        ('Drogo Burt'),
        ('Paulie Bryan'),
        ('Eddy Rickey'),
        ('Marjory Aria'),
        ('Topsy Grover'),
        ('Harmony Ibbie'),
        ('Buster Alf'),
        ('Kaydence Lew'),
        ('Barry Jayna'),
        ('Lorinda Natasha'),
        ('Dawn Sophia'),
        ('Windsor Lyle'),
        ('Myron Edythe'),
        ('Jaime Derryl'),
        ('Carlyle Yancy'),
        ('Issac Rachel'),
        ('Yancy Maisy'),
        ('Terry Wilda');
    END IF;

    IF NOT EXISTS (
        SELECT * 
        FROM category
    )
    THEN 
        INSERT INTO category(name, typeId)
        VALUES('Food',1),
        ('Bills',1);
    END IF;

    -- TRUNCATE transaction table to prepare for bulk loading
    TRUNCATE TABLE transaction;
END;
$body$
language 'plpgsql';


-- Toast is used to store extra from row that doesnt fit 8kb
-- Table Size Including Toast
-- SELECT pg_total_relation_size('person'); 