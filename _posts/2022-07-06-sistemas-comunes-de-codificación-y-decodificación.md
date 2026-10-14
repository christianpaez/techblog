---
layout: single
title: "Sistemas comunes de codificación y decodificación"
date: 2022-07-06 20:54:27 -2700
categories: webdev, spanish
---

![Sistemas comunes de codificación y decodificación]({{ site.url }}{{ site.baseurl }}/assets/img/sistemas-comunes-de-codificación-y-decodificación-00.jpg)
------------
Es común en informática no poder presentar o almacenar datos en texto plano por motivos de seguridad. Para ello, es muy común utilizar sistemas de codificación; esto consiste en tomar un dato como una letra o palabra y convertirlo en símbolos que parecen ilegibles. La decodificación es el proceso opuesto, tomar los símbolos codificados y convertirlos en algo que sea legible por humanos. Veamos algunos de los sistemas de codificación y decodificación más comunes utilizados en la actualidad.

Base64
======

La idea detrás de esta codificación es usar 64 caracteres que son comunes en todos los sistemas de codificación y también se puedan imprimir, las variaciones base64 más comunes usan caracteres alfanuméricos (`A-Z`, `a-z`, `0-9`) para los primeros 62 valores y alguna combinación de un signo más (`+`), una barra diagonal (`/`) para los últimos 2 caracteres y quizás un signo igual (`=`) para el relleno.

Ejemplo
-------

Original: `hola mundo`

Base64: `aGVsbG8gd29ybGQ=`

Hexadecimal (Base 16)
=====================

En este sistema, los datos se codifican en secuencias de 4 bits utilizando 16 símbolos del juego de caracteres ASCII, los caracteres más utilizados son las letras `A` a `F` (a veces minúsculas `a-f`) y los números arábigos o dígitos `0-9`.

Ejemplo
-------

Original: `hola mundo`

Hexadecimal (Base16): `68656c6c6f20776f726c640a`

ROT13 (Cifrado César)
=====================

Mejor definido como un cambio de sustitución simple que consiste en tomar una letra del alfabeto y reemplazarla con una letra 13 posiciones hacia abajo en el alfabeto (se pueden usar otros números de posiciones); esto también se conoce como el Cifrado César, ya que fue utilizado por Julio César en su correspondencia privada en la antigua Roma.

Ejemplo
-------

Original: `hola mundo`

ROT13: `uryyb jbeyq`

Estas codificaciones se usan comúnmente en desarrollo web, criptografía y encriptación de correo electrónico. Esperamos que esta publicación pueda proporcionar una comprensión básica de estas codificaciones y cómo transforman los datos.

Mira esta publicación en Art of Code: https://es.artofcode.tech/sistemas-comunes-de-codificacion-y-decodificacion/
