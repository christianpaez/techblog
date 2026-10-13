---
layout: single
title: "The Ultimate SQL JOIN Cheat Sheet"
date: 2024-09-10 15:36:46 -4600
categories: sql, cheatsheet
---

![The Ultimate SQL JOIN Cheat Sheet]({{ site.url }}{{ site.baseurl }}/assets/img/the-ultimate-sql-join-cheat-sheet-00.png)
------------
If you're working with relational databases, you'll often need to combine data from multiple tables. SQL `JOINs` help you achieve this.

### Understanding INNER JOIN

An `INNER JOIN` retrieves rows where there is a match between the tables based on the specified join condition. If no match is found, no rows are returned for that combination.

The syntax for an `INNER JOIN` looks like this:

```sql
-- [Mandatory] SELECT: Specifies the columns you want to retrieve.
SELECT
    -- [Optional] DISTINCT: Ensures the result set contains only unique rows by removing duplicates.
    -- Place DISTINCT immediately after SELECT if you want only unique rows.
    -- DISTINCT 
    column1, -- Specify the first column you want to retrieve
    column2, -- Specify the second column, and so on...
    ...

-- [Mandatory] FROM: The table from which you're selecting data.
FROM
    table1

-- [Mandatory] INNER JOIN: Joins two tables and retrieves rows that have matching values in both tables.
INNER JOIN
    table2 -- The second table you're joining with the first table

-- [Mandatory] ON: Specifies the condition that defines how the tables are related. 
-- Only rows where the columns match will be included in the result set.
ON
    table1.column = table2.column

-- [Optional] WHERE: Filters the rows returned by the join based on a specific condition.
-- Only rows that meet this condition will be included in the final result.
-- WHERE condition

-- [Optional] ORDER BY: Sorts the result set by the specified columns.
-- You can choose to sort in ascending (ASC) or descending (DESC) order.
-- ORDER BY column1 [ASC|DESC], column2 [ASC|DESC];
```

### Understanding OUTER JOINs

While `INNER JOINs` return only matching rows between two tables, `OUTER JOINs` go further by also including rows that do not have a match in one of the tables. There are three types of `OUTER JOINs`:

1. **LEFT OUTER JOIN** (or simply `LEFT JOIN`): returns all rows from the left table and the matching rows from the right table. If no match is found, the result will contain `NULL` for the columns from the right table.
2. **RIGHT OUTER JOIN** (or simply `RIGHT JOIN`): works in the opposite way of a `LEFT JOIN`. It returns all rows from the right table and the matching rows from the left table. If no match is found, the result will contain `NULL` for the columns from the left table.
3. **FULL OUTER JOIN:** returns all rows when there is a match in either the left or right table. If there is no match, the result will contain `NULL` for the columns from the table without a match.

Each of these joins is useful depending on the specific data retrieval needs.

---

The syntax for an `OUTER JOIN` looks like this:

```sql
-- [Mandatory] SELECT: Specifies the columns you want to retrieve.
SELECT
    -- [Optional] DISTINCT: Ensures the result set contains only unique rows by removing duplicates.
    -- Place DISTINCT immediately after SELECT if you want only unique rows.
    -- DISTINCT 
    column1, -- Specify the first column you want to retrieve
    column2, -- Specify the second column, and so on...
    ...

-- [Mandatory] FROM: The table from which you're selecting data.
FROM
    table1

-- [Mandatory] OUTER JOIN: Includes rows from the specified table even if there is no match in the other table.
-- You can use LEFT OUTER JOIN, RIGHT OUTER JOIN, or FULL OUTER JOIN based on your needs.
-- LEFT OUTER JOIN: Includes all rows from the left table and matched rows from the right table.
-- RIGHT OUTER JOIN: Includes all rows from the right table and matched rows from the left table.
-- FULL OUTER JOIN: Includes all rows from both tables, with NULLs where there is no match.
-- Example: LEFT OUTER JOIN
    LEFT OUTER JOIN
    table2 -- The table you're joining with the first table

-- [Mandatory] ON: Specifies the condition that defines how the tables are related. 
-- Only rows where the columns match will be included in the result set, or all rows from one table depending on the join type.
ON
    table1.column = table2.column

-- [Optional] WHERE: Filters the rows returned by the join based on a specific condition.
-- Only rows that meet this condition will be included in the final result.
-- WHERE condition

-- [Optional] ORDER BY: Sorts the result set by the specified columns.
-- You can choose to sort in ascending (ASC) or descending (DESC) order.
-- ORDER BY column1 [ASC|DESC], column2 [ASC|DESC];
```

### Example of INNER JOIN

To retrieve employee names and their associated department names, where there is a matching department for each employee, you can use an `INNER JOIN`:

```sql
SELECT
    employees.employee_name,
    departments.department_name
FROM
    employees
INNER JOIN
    departments
ON
    employees.department_id = departments.department_id;
```

This query will return only the rows where there is a match between the `department_id` columns in both the `employees` and `departments` tables. If an employee does not have a corresponding department or a department has no employees, those rows will not appear in the result set.

### Example of LEFT OUTER JOIN

Let’s say we have an `employees` table and a `departments` table, and we want to retrieve all employees, even if they are not assigned to a department:

```sql
SELECT
    employees.employee_name,
    departments.department_name
FROM
    employees
LEFT OUTER JOIN
    departments
ON
    employees.department_id = departments.department_id;
```

In this example, the query will return all rows from the `employees` table, and if an employee does not belong to a department, the `department_name` will show as `NULL`.

### Example of RIGHT OUTER JOIN

Let’s modify our previous example to retrieve all departments, even if they don’t have any employees assigned:

```sql
SELECT
    employees.employee_name,
    departments.department_name
FROM
    employees
RIGHT OUTER JOIN
    departments
ON
    employees.department_id = departments.department_id;
```

This query will return all rows from the `departments` table, and if a department doesn’t have any employees, the `employee_name` will show as `NULL`.

### Example of FULL OUTER JOIN

To retrieve all employees and all departments, including those that don’t have a match in the other table, you can use a `FULL OUTER JOIN`:

```sql
SELECT
    employees.employee_name,
    departments.department_name
FROM
    employees
FULL OUTER JOIN
    departments
ON
    employees.department_id = departments.department_id;
```

This query will return all rows from both tables. If an employee is not assigned to a department, `department_name` will be `NULL`. If a department has no employees, `employee_name` will be `NULL`.

### Example: LEFT OUTER JOIN with WHERE and ORDER BY

Here’s a query that retrieves all employees (even those without a department) and filters the results to only include employees from specific departments. The result is sorted in ascending order by employee name.

```sql
SELECT
    DISTINCT employees.employee_name,
    departments.department_name
FROM
    employees

LEFT OUTER JOIN
    departments
ON
    employees.department_id = departments.department_id
WHERE
    departments.department_name IN ('Sales', 'HR')

ORDER BY
    employees.employee_name ASC;
```

Happy querying!
