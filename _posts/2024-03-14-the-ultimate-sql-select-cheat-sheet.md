---
layout: single
title: "The Ultimate SQL SELECT Cheat Sheet"
date: 2024-03-14 20:28:48 -4800
categories: sql, cheatsheet
---

![The Ultimate SQL SELECT Cheat Sheet]({{ site.url }}{{ site.baseurl }}/assets/img/the-ultimate-sql-select-cheat-sheet-00.png)
------------
Have you ever found yourself staring blankly at your SQL code, trying to recall the syntax for a SELECT query? Don't worry; you're not alone. SQL, with its vast array of commands and clauses, can be daunting to remember every detail. Here I will share a useful cheat sheet to write the most common SELECT queries.

### The Power of a Cheat Sheet

Memorizing every SQL command and its syntax is not easy, that's where cheat sheets come to the rescue. Instead of flipping through documentation a cheat sheet provides quick and easy access to essential commands.

**Deciphering the SELECT Query**

Let's check the anatomy of a SELECT query. At its core, this query retrieves data from one or more tables in a database. Here's a breakdown of its components:

- **SELECT:** This keyword specifies the columns you want to retrieve from the database. You can either select specific columns or use the wildcard (*) to fetch all columns.
- **FROM:** Here, you specify the table or tables from which you want to retrieve data.
- **DISTINCT (optional)**: This optional keyword removes duplicate rows from the result set.
- **WHERE(optional):** This optional clause filters rows based on specified conditions. It allows you to narrow down your results to only those that meet certain criteria.
- **GROUP BY(optional):** When you want to group your results based on the values of one or more columns, you use this clause. It's commonly used in conjunction with aggregate functions.
- **HAVING(optional):** Similar to the WHERE clause, HAVING filters grouped rows based on specified conditions. It comes into play after the GROUP BY clause.
- **ORDER BY (optional)**: This optional clause sorts the result set based on specified columns or expressions.

The base for a SELECT query will look like this:

```sql
-- [Mandatory] SELECT: This keyword specifies the columns you want to retrieve from the database.
SELECT 
    column1, -- Specify columns to retrieve or * to fetch all columns
    column2, 
    ...
    -- [Optional] DISTINCT: This keyword removes duplicate rows from the result set, ensuring only unique rows are returned.
    -- Tipically used immediately after the SELECT keyword.
    -- For example:
    -- SELECT DISTINCT column1, column2 FROM table_name;
-- [Mandatory] FROM: Here, you specify the table or tables from which you want to retrieve data.
FROM 
    table1, -- [Mandatory] Specify at least one table
    table2, 
    ...;
    
-- [Optional] WHERE: This clause filters rows based on specified conditions.
-- For example:
-- SELECT * FROM employees WHERE salary > 50000;
    
-- [Optional] GROUP BY: When you want to group your results based on the values of one or more columns,
-- you use this clause. It's commonly used in conjunction with aggregate functions.
-- For example:
-- SELECT department_id, COUNT(*) FROM employees GROUP BY department_id;

-- [Optional] HAVING: Similar to the WHERE clause, HAVING filters grouped rows based on specified conditions.
-- It comes into play after the GROUP BY clause.
-- For example:
-- SELECT department_id, COUNT(*) FROM employees GROUP BY department_id HAVING COUNT(*) > 3;

-- [Optional] ORDER BY: This clause sorts the result set based on specified columns or expressions.
-- It's typically used to arrange the data in ascending or descending order.
-- For example:
-- SELECT employee_name, salary FROM employees ORDER BY salary DESC;
```

### Some examples

This query retrieves all columns from the employees table without applying any filtering criteria.

```sql
-- Basic Example without WHERE Clause
SELECT *
FROM employees;
```

This query uses the `DISTINCT` keyword to retrieve unique values of the `department_id` column from the employees table.

```sql
-- Example with DISTINCT
-- This query retrieves distinct values of the department_id column from the employees table.
SELECT DISTINCT department_id
FROM employees;
```

This query filters employee records based on the salary column, only returning those with a salary greater than $50,000.

```sql
-- Example with WHERE Clause
-- This query retrieves employee records with a salary greater than $50,000.
SELECT employee_name, salary
FROM employees
WHERE salary > 50000;

```

This query groups employee records by department and calculates the count of employees in each department.

```sql
-- Example with GROUP BY Clause
-- This query counts the number of employees in each department.
SELECT department_id, COUNT(*) AS employee_count
FROM employees
GROUP BY department_id;

```

This query first groups employee records by department and then filters the grouped results to only include departments with more than three employees.

```sql
-- Example with HAVING Clause
-- This query finds departments with more than three employees.
SELECT department_id, COUNT(*) AS employee_count
FROM employees
GROUP BY department_id
HAVING COUNT(*) > 3;

```

This example retrieves employee records and sorts them by salary in descending order using the `ORDER BY` clause.

```sql
-- Example with ORDER BY Clause
-- This query retrieves employee records sorted by salary in descending order.
SELECT employee_name, salary
FROM employees
ORDER BY salary DESC;

```

Happy querying!
