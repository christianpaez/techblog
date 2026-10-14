---
layout: single
title: "¿Quién ganará las elecciones presidenciales de Colombia de 2022?, según la estadística."
date: 2022-06-12 18:42:11 -1100
categories: datascience
---

![¿Quién ganará las elecciones presidenciales de Colombia de 2022?, según la estadística.]({{ site.url }}{{ site.baseurl }}/assets/img/¿quién-ganará-las-elecciones-presidenciales-de-colombia-de-2022?,-según-la-estadística.-00.jpg)
------------
En menos de una semana, se elegirá un nuevo presidente para dirigir la República de Colombia, dado que este se ha convertido en un tema popular en redes sociales, muchos medios de comunicación han realizado encuestas de opinión sobre los candidatos presidenciales; ¿Qué nos dicen estos datos sobre un posible ganador?.

### ¿Los datos de las encuestas proporcionan información útil?

La respuesta corta es sí. Gracias a los datos de encuestas públicas, podemos construir un modelo estadístico que nos da la probabilidad de que un candidato gane la segunda vuelta de las elecciones. Sin embargo, hay algunos detalles que debemos tener en cuenta:

### Variabilidad de sondeo

Cuando construimos nuestro modelo de predicción, debemos tener en cuenta la posible variabilidad en los datos recopilados; esto puede suceder por muchas razones:

-   La gente miente en las encuestas.
-   Las opiniones de las personas cambian con el tiempo.
-   La población votante y encuestada puede diferir.
-   Parcialidad del encuestador ó fuente encuestada.

A pesar de estos problemas, todavía podemos construir un modelo eficaz para predecir un ganador.

### ¿Cómo se construyó este modelo, a pesar de los problemas de datos de encuestas?

Todas estas fallas encontradas en los datos de las encuestas se interpretan como variabilidad en nuestro modelo, el cual funciona perfectamente porque queremos predecir quién ganará, no precisamente el número o porcentaje de votos, los datos de elecciones anteriores también están disponibles para usar como referencia para analizar cuál es el comportamiento esperado de los datos de las encuestas frente a los resultados de las elecciones y qué tan grande es el porcentaje de votos obtenidos por el ganador de las elecciones.

### Consulte nuestro pronóstico actual aquí: <https://es.artofcode.tech/pronostico-elecciones-presidenciales-de-colombia-de-2022/>  
&nbsp; 
  

### ¿Cómo interpretar este modelo?

Actualmente, es un empate. Un candidato tiene una mayor probabilidad de ganar, pero una victoria no es segura, en las estadísticas a menudo estamos seguros de un evento si vemos un 95% de probabilidad de que suceda, este no es el caso aquí; escenarios similares se han visto en otras elecciones como la del año 2016 en Estados Unidos, un modelo muy similar a este predijo una victoria de Hillary Clinton con un 71% de probabilidad, sin embargo, perdió las elecciones; es cierto que tenía un mayor porcentaje de obtener la victoria, pero a menos que veamos una probabilidad de más del 90% de ganar, es difícil estar seguro de una victoria clara.

### Conclusión

La estadística y la ciencia de datos son muy útiles para predecir resultados de elecciones con base en datos de encuestas de opinión, sin embargo, parece que muchos medios de comunicación populares tienen problemas al interpretar estos números, ya que se está convirtiendo en una idea popular entre el público en general que las encuestas no brindan información útil o están manipuladas. Esperamos que aquí pueda encontrar una evaluación más realista del significado de los resultados de las encuestas de opinión y lo invitamos a consultar nuestro pronóstico actual, lo actualizamos con la mayor frecuencia posible.

### Referencias:

Pronóstico: <https://es.artofcode.tech/pronostico-elecciones-presidenciales-de-colombia-de-2022/>
Github: <https://github.com/christianpaez/colombian-election-analysis-2022>
