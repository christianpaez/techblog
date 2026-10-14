---
layout: single
title: "How does URL decoding and encoding work?"
date: 2022-09-12 20:27:59 -5900
categories: webdev, software
---

![How does URL decoding and encoding work?]({{ site.url }}{{ site.baseurl }}/assets/img/how-does-url-decoding-and-encoding-work?-00.jpg)
------------

Every time you visit a website, your browser is encoding and decoding URLs under the hood. This is done in order to avoid transmitting invalid or unsafe data; let's see how URL decoding works on a basic level and why it is important.

Encoding a URL is the process of taking a string like this one: `https://example.com?param=value#someAnchor`

And apply some transformation rules to make it look like this one:

`https%3A%2F%2Fexample.com%3Fparam%3Dvalue%23someAnchor`

This transformation is important since URLs must follow the URI syntax standard to send only valid and safe data. URL decoding and encoding are usually handled by the web browser or web server that you are using. However, there are times when you might need to encode or decode a URL yourself, so let's see how it works on a basic level.

### URL decoding/encoding basics

As we mentioned earlier, the URI syntax standard must be followed. Only a special subset of ASCII characters is allowed in URL's. This consists of the familiar alphanumeric symbols, and some reserved characters for use as control characters.

![Taken from {{ site.url }}{{ site.baseurl }}/assets/img/how-does-url-decoding-and-encoding-work?-01.png)

### Encoding and Decoding Methods

The most common encoding system is called Percent-encoding, which consists of taking a `%` sign and appending the hexadecimal representation of the ASCII value of this character. For example, the `@` sign is represented by `#40`, since 40 is the value of this character in the ASCII table.

The "base64" method is also sometimes used. With this method, each character is represented by a set of six characters. For example, the character `#` would be represented as `iVBORw0K`; you can learn more about base64 in our previous post on Common decoding/encoding systems: <https://artofcode.tech/common-encoding-and-decoding-systems/>

We hope that this post has helped you understand a bit more about how the modern web technologies and web browsers work under the hood.

Check out this post in Art Of Code: https://artofcode.tech/how-does-url-decoding-and-encoding-work/
Image credit: Image from Richy Great on Unsplash.
