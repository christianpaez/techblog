---
layout: single
title: "¿Cómo funcionan la decodificación y la codificación de las URL?"
date: 2022-09-12 20:27:49 -4900
categories: webdev, software, spanish
---

![¿Cómo funcionan la decodificación y la codificación de las URL?]({{ site.url }}{{ site.baseurl }}/assets/img/¿cómo-funcionan-la-decodificación-y-la-codificación-de-las-url?-00.jpg)
------------
Cada vez que visitas un sitio web, tu navegador codifica y decodifica las URLs detrás de cámaras. Esto se hace para evitar la transmisión de datos no válidos o inseguros; veamos cómo funciona la decodificación de URL a nivel básico y por qué es importante.

Codificar una URL es el proceso de tomar una cadena como esta:

`https://example.com?param=value#someAnchor`

Y aplicar algunas reglas de transformación para que se parezca a esto:

`https%3A%2F%2Fexample.com%3Fparam%3Dvalue%23someAnchor`

Esta transformación es importante ya que las URLs deben seguir el estándar de sintaxis URI para enviar sólo datos válidos y seguros. La decodificación y la codificación de las URLs suelen ser gestionadas por el navegador o el servidor web que estés utilizando. Sin embargo, hay ocasiones en las que puedes necesitar codificar o decodificar una URL tú mismo, así que vamos a ver cómo funciona a nivel básico.

### Conceptos básicos de decodificación/codificación de URLs

Como hemos mencionado anteriormente, hay que seguir la norma de sintaxis URI. Sólo se permite un subconjunto especial de caracteres ASCII en las URL. Este consiste en los conocidos símbolos alfanuméricos, y algunos caracteres reservados para su uso como caracteres de control.

![Tomada de {{ site.url }}{{ site.baseurl }}/assets/img/¿cómo-funcionan-la-decodificación-y-la-codificación-de-las-url?-01.png)

### Métodos de codificación y decodificación

El sistema de codificación más común es el llamado codificación porcentual, que consiste en tomar un signo `%` y añadir la representación hexadecimal del valor ASCII de este carácter. Por ejemplo, el signo `@` se representa con `#40`, ya que 40 es el valor de este carácter en la tabla ASCII.

A veces también se utiliza el método base64. Con este método, cada carácter se representa por un conjunto de seis caracteres. Por ejemplo, el carácter `#` se representaría como `iVBORw0K`; puede obtener más información sobre base64 en nuestro post anterior sobre Sistemas comunes de decodificación/codificación: <https://es.artofcode.tech/sistemas-comunes-de-codificacion-y-decodificacion/>

Esperamos que este post le haya ayudado a entender un poco más cómo funcionan las tecnologías web modernas y los navegadores web detrás de cámaras.

Mira este post en Art Of Code: https://es.artofcode.tech/como-funcionan-la-decodificacion-y-la-codificacion-de-las-url/
Crédito de la imagen: Richy Great en Unsplash.
