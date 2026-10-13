---
layout: single
title: "HackerRank SQL Preparation: Japanese Cities' Names(MySQL)"
date: 2024-06-09 18:44:35 -3500
categories: sql, mysql, writeups, hackerrank
---

![HackerRank SQL Preparation: Japanese Cities' Names(MySQL)]({{ site.url }}{{ site.baseurl }}/assets/img/hackerrank-sql-preparation:-japanese-cities'-names(mysql)-00.png)
------------

**Problem Statement:**
Query the names of all Japanese cities in the **CITY** table. The `COUNTRYCODE` for Japan is `JPN`.

**Link:** [HackerRank - Japanese Cities Name](https://www.hackerrank.com/challenges/japanese-cities-name/problem)

**Solution:**

```sql
SELECT NAME FROM CITY WHERE COUNTRYCODE = 'JPN';
```

**Explanation:**

- `SELECT NAME`: This part of the query specifies that you want to retrieve the `NAME` column from the **CITY** table.
- `FROM CITY`: Indicates that you are selecting data from the **CITY** table.
- `WHERE COUNTRYCODE = 'JPN'`: This condition filters the rows to include only those cities where the `COUNTRYCODE` is 'JPN' (Japan).
