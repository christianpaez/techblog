---
layout: single
title: "HackerRank SQL preparation; Revising the Select Query II(MySQL)"
date: 2024-04-03 22:01:45 -4500
categories: sql, writeups, hackerrank
---

![HackerRank SQL preparation; Revising the Select Query II(MySQL)]({{ site.url }}{{ site.baseurl }}/assets/img/hackerrank-sql-preparation;-revising-the-select-query-ii(mysql)-00.png)
------------
**Problem Statement:**
Query the `NAME` field for American cities in the **CITY** table with populations larger than 120,000. The `CountryCode` for America is `USA`.

**Link:**[HackerRank - Revising the SELECT Query 2](https://www.hackerrank.com/challenges/revising-the-select-query-2/problem)

**Solution:**

```sql
SELECT NAME FROM CITY WHERE POPULATION > 120000 AND COUNTRYCODE = 'USA';
```

**Explanation:**

- `SELECT NAME`: This part of the query specifies that you want to retrieve the `NAME` field.
- `FROM CITY`: Indicates that you are selecting data from the **CITY** table.
- `WHERE POPULATION > 120000`: This condition filters the rows and includes only those cities where the population is larger than 120,000.
- `AND COUNTRYCODE = 'USA'`: This additional condition ensures that only American cities (USA) are included in the result, based on the `CountryCode` column.
