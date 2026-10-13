---
layout: single
title: "HackerRank SQL Preparation: Weather Observation Station 1(MySQL)"
date: 2024-06-09 18:41:32 -3200
categories: sql, writeups, hackerrank, mysql
---

![HackerRank SQL Preparation: Weather Observation Station 1(MySQL)]({{ site.url }}{{ site.baseurl }}/assets/img/hackerrank-sql-preparation:-weather-observation-station-1(mysql)-00.png)
------------

**Problem Statement:**
Query a list of CITY and STATE from the **STATION** table.

**Link:** [HackerRank - Weather Observation Station 1](https://www.hackerrank.com/challenges/weather-observation-station-1/problem)

**Solution:**

```sql
SELECT CITY, STATE FROM STATION;

```

**Explanation:**

- `SELECT CITY, STATE`: This part of the query specifies that you want to retrieve the `CITY` and `STATE` columns from the **STATION** table.
- `FROM STATION`: Indicates that you are selecting data from the **STATION** table.
