---
layout: single
title: "Subdominios vs. Hosts virtuales"
date: 2022-07-24 18:21:23 -2300
categories: spanish, webdev
---

![Subdominios vs. Hosts virtuales]({{ site.url }}{{ site.baseurl }}/assets/img/subdominios-vs.-hosts-virtuales-00.jpg)
------------
Es importante entender la diferencia entre subdominios y VHosts porque puede afectar la forma en que se accede a su sitio web y cómo los motores de búsqueda indexan su sitio web. Si no está utilizando el tipo de alojamiento correcto, también puede afectar la velocidad y el rendimiento de su sitio web e incluso ser vulnerable a las vulnerabilidades de seguridad cibernética como el secuestro de DNS, las vulnerabilidades del certificado SSL y las secuencias de comandos entre sitios.

Subdominio
==========

Un subdominio es un dominio de segundo nivel que forma parte de un dominio más grande. Por ejemplo, si tiene un sitio web en `example.com`, puede crear un subdominio en `subdomain.example.com`. Un subdominio se puede usar para crear un sitio web separado, o se puede usar para apuntar a un sitio web o directorio diferente en el mismo servidor.

Host Virtual
============

Un host virtual (VH) es un servicio de alojamiento de Internet que permite a las organizaciones alojar sus sitios web en un único servidor. Un VH se puede usar para alojar varios sitios web, cada uno con su propio nombre de dominio, o se puede usar para alojar varios sitios web que comparten el mismo nombre de dominio. Los VHosts pueden o no tener registros de DNS públicos, por lo que para acceder a su sitio, es posible que deba cambiar los nombres y las direcciones de los hosts en su servidor local (generalmente ubicado en `etc/hosts`) o usar el encabezado `Host:` de un solicitud HTTP.

---


En conclusión, es importante comprender la diferencia entre subdominios y VHosts para garantizar que su sitio web sea accesible y funcione de manera óptima.

Revisa esta publicación en Art Of Code: https://es.artofcode.tech/subdominios-vs-hosts-virtuales/
