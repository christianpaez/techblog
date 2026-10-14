---
layout: single
title: "Common Encoding and Decoding systems"
date: 2022-07-06 20:05:57 -5700
categories: webdev
---

![Common Encoding and Decoding systems]({{ site.url }}{{ site.baseurl }}/assets/img/common-encoding-and-decoding-systems-00.jpg)
------------
Sometimes data cannot be stored or presented in plain text due to security reasons. For this, it is very common to use encoding systems; this consists of taking a piece of data like a letter or word and converting it into symbols that looks unreadable. Decoding is the opposite process, taking the encoded symbols and converting them back into something that's human readable. Let's see some of the most common encoding and decoding systems used nowadays.

Base64
======

The idea behind this encoding is to use 64 characters that are common across encoding systems and are also printable, most common base64 variations use alphanumeric characters (`A-Z`, `a-z`, `0-9`) for the first 62 values and some combination of a plus sign (`+`), a forward slash (`/`)for the last 2 characters and perhaps an equal sign (`=`) for padding.

Example
-------

Original: `hello world`

Base64: `aGVsbG8gd29ybGQ=`

Hex (Base 16)
=============

In this system, data is encoded in 4-bit sequences using 16 symbols from the ASCII character set, most commonly used characters are letters `A` to `F` (sometimes lowercase `a-f`) and the Arabic numerals or digits `0-9`.

Example
-------

Original: `hello world`

Hex (Base16): `68656c6c6f20776f726c640a`

ROT13 (Caesar Cypher)
=====================

More of a simple substitution shift that consists of taking an alphabet letter and replacing it with a letter 13 positions down the alphabet (other numbers of positions can be used); this is also known as the Caesar Cypher since it was used by Julius Caesar in his private correspondence in ancient Rome.

Example
-------

Original: `hello world`

ROT13: `uryyb jbeyq`

These encodings are commonly used in web development, cryptography, and email encryption. We hope this post can provide a basic understanding of these encodings and how they transform data.

Check out this post on Art of Code: https://artofcode.tech/common-encoding-and-decoding-systems/
