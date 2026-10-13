---
layout: single
title: "HackerRank SQL Preparation: Select All(MySQL)"
date: 2024-06-07 10:27:53 -5300
categories: sql, writeups, hackerrank, mysql
---

![HackerRank SQL Preparation: Select All(MySQL)]({{ site.url }}{{ site.baseurl }}/assets/img/hackerrank-sql-preparation:-select-all(mysql)-00.png)
------------
**Problem Statement:**
Query all columns (attributes) for every row in the **CITY** table.

---

**Link:** [HackerRank - Select All SQL](https://www.hackerrank.com/challenges/select-all-sql/problem)

**Solution:**

```sql
SELECT * FROM CITY;
```

**Explanation:**

- `SELECT *`: The asterisk (*) is a wildcard character in SQL that means "all columns." This part of the query specifies that you want to retrieve all columns from the table.
- `FROM CITY`: Indicates that you are selecting data from the **CITY** table.

This query will return all the columns for every row in the **CITY** table. It's useful when you need to inspect or analyze all data available in the table without any specific filtering or conditions.
