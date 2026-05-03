# Database Transactions

## DB Locks
https://wiki.postgresql.org/wiki/Lock_Monitoring

```sql
-- simulate lock - session1 - different order of account id updates!!!
-- update account 2 before account 3
BEGIN;
UPDATE accounts SET balance = balance - 10 WHERE id = 2;
UPDATE accounts SET balance = balance + 10 WHERE id = 3;
ROLLBACK;

-- simulate lock - session2 - different order of account id updates!!!
-- update account 3 before account 2
BEGIN;
UPDATE accounts SET balance = balance - 10 WHERE id = 3;
UPDATE accounts SET balance = balance + 10 WHERE id = 2;
ROLLBACK;

-- view locks
SELECT a.application_name
     , l.relation::regclass
     , l.transactionid
     , l.mode
     , l.GRANTED
     , a.usename
	 , a.pid
     , a.query
     , a.query_start
     , age(now(), a.query_start) AS "age"
  FROM pg_stat_activity a
  JOIN pg_locks l ON l.pid = a.pid
 WHERE a.application_name = 'psql'
 ORDER BY a.query_start;

show transaction isolation level;                   -- read committed

begin;
set transaction isolation level read uncommitted;   -- same as read committed in pg
set transaction isolation level read committed;
set transaction isolation level repeatable read;
set transaction isolation level serializable;
show transaction isolation level; 
```

## Isolation Level
### Read Phenomena
- **Dirty Read**: a transaction reads data written by another uncommitted transaction
- **Non-Repeatable Read**: a transaction reads the same row twice and sees different values because it has been modified by other committed transactions.
- **Phantom Read**: a transaction re-executes a query to find rows that satisfy a condition and sees a different set of rows, due to changes by other committed transactions.
- **Serialization Anomaly**: the result of a group of concurrent committed transactions is impossible to achieve if we try to run them sequentially in any order without overlapping.

<br>

<pre>
 -low--1----------2-----------3------------4--high-
      Read      Read     Repeatable   Serializable
   Uncommitted Committed    Read
</pre>

| | Read Committed | Read Committed | Repeatable Read | Serializable |
|-| -------------- | -------------- | --------------- | ------------ |
| Dirty Read            |  Y  |  N  |        N        |       N      | 
| Non-Repeatable Read   |  Y  |  Y  |        N        |       N      | 
| Phantom Read          |  Y  |  Y  |        N        |       N      | 
| Serialization Anomaly |  Y  |  Y  |        Y        |       N      | 
