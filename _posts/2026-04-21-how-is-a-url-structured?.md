---
layout: single
title: "How is a URL structured?"
date: 2026-04-21 05:07:10 -1000
categories: webdev
---

![How is a URL structured?]({{ site.url }}{{ site.baseurl }}/assets/img/how-is-a-url-structured?-00.jpg)
------------
A huge part of modern software is related to web technologies, mostly done by a web server interpreting a URL (Uniform Resource Locator) to locate a resource. In this short writing we will learn how a URL is structured and interpreted.

## URL Example

https://example.com:443/path/file.html?key1=value1&key2=value2#SomeLocationInDocument

A URL like this one is divided into multiple components:

## Scheme (Required)

![Image description]({{ site.url }}{{ site.baseurl }}/assets/img/how-is-a-url-structured?-01.png)

Indicates the protocol that has to be used to request the resource, can be either https (With SSL/Secure) or http (Without SSL/Unsecure).

## User (Optional)

![Image description]({{ site.url }}{{ site.baseurl }}/assets/img/how-is-a-url-structured?-02.png)

Used for HTTP Basic Authentication, currently deprecated since it is prone to security flaws, is written in the form user:password followed by an @ sign.

## Host (Required)

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/g5b4k97ustt12uu1fabr.png)

This is the resource location, could be an IP address (e.g., 192.0.2.146) or a host name (e.g., example.com). a host name can include a subdomain followed by a dot . before the domain name (i.e., subdomain.example.com) and has to include a top-level domain (e.g., .com).

## Port (Optional)

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/gh6av21bybnwkmzrv4da.png)

This number indicates the web server port number we want to connect to, denoted after the host and preceded by colon : defaults to 443 for https and 80 to http.

## Path/Directory (Optional)

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/ogpytebzaq14qa475q8w.png)

Points to the resource we want to access, can be a file or a folder; all web applications have a default path, usually index.html.

## Query String (Optional)

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/7l9pmxi99rp69ptcnnvj.png)

These are extra parameters sent to the web server, starts with a question mark ?
followed by one or many key value pairs in the form key=value.

## Fragment/Anchor (Optional)

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/jh7p2999mk53ogf3x8lg.png)

This is used by the locate sections within the document displayed by the browser; this value is only processed on the client.

As you can see, there is a lot going on in a URL string, we hope this short writing can provide useful info as to what these values mean.

Check this post on Art Of Code: https://artofcode.tech/how-is-a-url-structured/
