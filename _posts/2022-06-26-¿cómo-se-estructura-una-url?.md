---
layout: single
title: "¿Cómo se estructura una URL?"
date: 2022-06-26 19:05:21 -2100
categories: webdev, programming, beginners
---

![¿Cómo se estructura una URL?]({{ site.url }}{{ site.baseurl }}/assets/img/¿cómo-se-estructura-una-url?-00.jpg)
------------
Una gran parte del software moderno está relacionado con las tecnologías web, principalmente realizado por un servidor web que interpreta una URL (Localizador Uniforme de Recursos) para localizar un recurso. En este breve escrito, aprenderemos cómo se estructura e interpreta una URL.

### Ejemplo de URL

[`https://example.com:443/path/file.html?key1=value1&key2=value2#SomeLocationInDocument`](https://example.com:443/path/file.html?key1=value1&key2=value2#SomeLocationInDocument)

Una URL como esta se divide en varios componentes:

Esquema (Requerido)
-------------------

![]({{ site.url }}{{ site.baseurl }}/assets/img/¿cómo-se-estructura-una-url?-01.png)

Indica el protocolo que debe usarse para solicitar el recurso, puede ser https (Con SSL/Seguro) o http (Sin SSL/No seguro).

Usuario (Opcional)
------------------

![]({{ site.url }}{{ site.baseurl }}/assets/img/¿cómo-se-estructura-una-url?-02.png)

Utilizado para la autenticación básica HTTP, actualmente en desuso ya que es propenso a fallas de seguridad, se escribe en la forma `usuario: contraseña` seguida de un signo `@`.

Host (Requerido)
----------------

![](https://res-2.cloudinary.com/hbig9cpoo/image/upload/q_auto/v1/ghost-blog-images/URL-2-.png)

Esta es la ubicación del recurso, podría ser una dirección IP (por ejemplo, 192.0.2.146) o un nombre de host (por ejemplo, ejemplo.com). un nombre de host puede incluir un subdominio seguido de un punto `.` antes del nombre de dominio (es decir, subdominio.ejemplo.com) y debe incluir un dominio de nivel superior (por ejemplo, .com).

Puerto (Opcional)
-----------------

![](https://res-3.cloudinary.com/hbig9cpoo/image/upload/q_auto/v1/ghost-blog-images/URL-3-.png)

Este número indica el número de puerto del servidor web al que queremos conectarnos, indicado después del host y precedido por dos puntos `:` el valor predeterminado es 443 para https y 80 para http.

Ruta/Directorio (Opcional)
--------------------------

![](https://res-5.cloudinary.com/hbig9cpoo/image/upload/q_auto/v1/ghost-blog-images/URL-4-.png)

Apunta al recurso al que queremos acceder, puede ser un archivo o una carpeta; todas las aplicaciones web tienen una ruta predeterminada, generalmente `index.html`.

Cadena de consulta (opcional)
-----------------------------

![](https://res-1.cloudinary.com/hbig9cpoo/image/upload/q_auto/v1/ghost-blog-images/URL-5-.png)

Estos son parámetros adicionales enviados al servidor web, comienza con un signo de interrogación `?` seguido de uno o varios pares de valores clave en el formato `llave=valor`.

Fragmento/Anclaje (Opcional)
----------------------------

![](https://res-2.cloudinary.com/hbig9cpoo/image/upload/q_auto/v1/ghost-blog-images/URL-8-.png)

Esto es utilizado por las secciones de localización dentro del documento que muestra el navegador; este valor solo se procesa del lado del cliente.

Como puede ver, suceden muchas cosas en una cadena de URL, esperamos que este breve escrito pueda proporcionar información útil sobre lo que significan estos valores.

Revisa esta publicación en Art Of Code: https://es.artofcode.tech/como-se-estructura-una-url/
