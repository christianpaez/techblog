---
layout: single
title: "HackerRank SQL Preparation: Weather Observation Station 3(MySQL)"
date: 2024-08-27 02:11:49 -4900
categories: sql, hackerrank, writeups, mysql
---
![HackerRank SQL Preparation: Weather Observation Station 3(MySQL)]({{ site.url }}{{ site.baseurl }}/assets/img/hackerrank-sql-preparation:-weather-observation-station-3(mysql)-00.png)
------------
**Problem Statement:**
Query a list of CITY names from **STATION** for cities that have an even ID number. Print the results in any order, but exclude duplicates from the answer.

**Link:** [HackerRank - Weather Observation Station 3](https://www.hackerrank.com/challenges/weather-observation-station-3/problem?isFullScreen=true)

**Solution:**

```sql
SELECT DISTINCT CITY FROM STATION WHERE ID % 2 = 0;
```

**Explanation:**

- `SELECT DISTINCT CITY`: This part of the query specifies that you want to retrieve the `CITY` column from the **STATION** table, and `DISTINCT` ensures that each city name is unique, eliminating any duplicates from the result.
- `FROM STATION`: Indicates that you are selecting data from the **STATION** table.
- `WHERE ID % 2 = 0`: This condition filters the rows to include only those cities where the `ID` is an even number. The modulo operation (`% 2 = 0`) checks if the `ID` is divisible by 2, meaning it's even.

This query will return a list of unique city names from the **STATION** table where the city has an even `ID` number. The result will exclude any duplicate city names.
