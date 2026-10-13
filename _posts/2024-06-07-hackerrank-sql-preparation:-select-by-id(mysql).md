---
layout: single
title: "HackerRank SQL Preparation: Select by ID(MySQL)"
date: 2024-06-07 10:27:39 -3900
categories: sql, writeups, hackerrank, mysql
---

![HackerRank SQL Preparation: Select by ID(MySQL)]({{ site.url }}{{ site.baseurl }}/assets/img/hackerrank-sql-preparation:-select-by-id(mysql)-00.png)
------------
**Problem Statement:**
Query all columns for a city in the **CITY** table with the ID 1661.

---

**Link:** [HackerRank - Select by ID](https://www.hackerrank.com/challenges/select-by-id/problem)

**Solution:**

```sql
SELECT * FROM CITY WHERE ID = 1661;
```

**Explanation:**

- `SELECT *`: The asterisk (*) is a wildcard character in SQL that means "all columns." This part of the query specifies that you want to retrieve all columns from the table.
- `FROM CITY`: Indicates that you are selecting data from the **CITY** table.
- `WHERE ID = 1661`: This condition filters the rows to include only the city with the specific `ID` of 1661.

This query will return all the columns for the row in the **CITY** table where the `ID` is 1661. It's useful when you need to retrieve all details about a specific city identified by its unique ID.

By running this query, you can see all the attributes (such as city name, country code, population, etc.) for the city that has the ID of 1661.
