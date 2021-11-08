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
	pid INT PRIMARY KEY,
	name VARCHAR NOT NULL
    );

    CREATE UNLOGGED TABLE IF NOT EXISTS category (
	catId INT PRIMARY KEY,
	name VARCHAR NOT NULL,
	typeId INT
    );

    CREATE UNLOGGED TABLE IF NOT EXISTS transaction (
	tranId INT PRIMARY KEY,
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
        INSERT INTO person(pid, name)
        VALUES(1,'kenzou'),
        (2,'personwithlongname');
    END IF;

    IF NOT EXISTS (
        SELECT * 
        FROM category
    )
    THEN 
        INSERT INTO category(catId, name, typeId)
        VALUES(1,'Food',1),
        (2,'Bills',1);
    END IF;

    -- TRUNCATE transaction table to prepare for bulk loading
    TRUNCATE TABLE transaction;
END;
$body$
language 'plpgsql';


-- Toast is used to store extra from row that doesnt fit 8kb
-- Table Size Including Toast
-- SELECT pg_total_relation_size('person'); 