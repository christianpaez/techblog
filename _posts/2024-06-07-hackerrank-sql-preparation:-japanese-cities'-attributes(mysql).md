---
layout: single
title: "HackerRank SQL Preparation: Japanese Cities' Attributes(MySQL)"
date: 2024-06-07 10:27:22 -2200
categories: sql, writeups, hackerrank, mysql
---

![HackerRank SQL Preparation: Japanese Cities' Attributes(MySQL)]({{ site.url }}{{ site.baseurl }}/assets/img/hackerrank-sql-preparation:-japanese-cities'-attributes(mysql)-00.jpeg)
------------
**Problem Statement:**
Query all attributes of every Japanese city in the **CITY** table. The `COUNTRYCODE` for Japan is `JPN`.

**Link:** [HackerRank - Japanese Cities Attributes](https://www.hackerrank.com/challenges/japanese-cities-attributes/problem)

**Solution:**

```sql
SELECT * FROM CITY WHERE COUNTRYCODE = 'JPN';
```

**Explanation:**

- `SELECT *`: The asterisk (*) is a wildcard character in SQL that means "all columns." This part of the query specifies that you want to retrieve all columns from the table.
- `FROM CITY`: Indicates that you are selecting data from the **CITY** table.
- `WHERE COUNTRYCODE = 'JPN'`: This condition filters the rows to include only the cities where the `COUNTRYCODE` is 'JPN' (Japan).

This query will return all the columns for every row in the **CITY** table where the `COUNTRYCODE` is 'JPN'. It's useful when you need to retrieve all details about every city in Japan.

By running this query, you can see the complete dataset of Japanese cities stored in the **CITY** table, including information such as city names, populations, districts, and other relevant attributes.
