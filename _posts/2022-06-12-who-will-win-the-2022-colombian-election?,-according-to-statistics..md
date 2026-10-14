---
layout: single
title: "Who will win the 2022 Colombian Election?, according to statistics."
date: 2022-06-12 18:45:05 -0500
categories: datascience
---

![Who will win the 2022 Colombian Election?, according to statistics.]({{ site.url }}{{ site.baseurl }}/assets/img/who-will-win-the-2022-colombian-election?,-according-to-statistics.-00.jpg)
------------
In about a week, A new President will be elected to lead the Republic of Colombia, since this has become a popular topic in the media, many news outlets have run opinion polls on presidential candidates; what does this data tell us about a possible winner?.

### Does polling data provide useful insights?

The short answer is, yes. Thanks to public polling data, we can build a statistical model that gives us the probability of a candidate winning the second round of the election. However, there are some caveats:

### Polling Variability

When we build our prediction model, we need to take into account possible variability in the collected data; this can happen for many reasons:

-   People lie in polls.
-   People's opinions change over time.
-   Voting and polled population may differ.
-   Pollster Bias.

Despite these problems, we can still construct an effective model to predict a winner.

### How was this model built, despite polling data issues?

All of these flaws found in polling data are interpreted as *variability* in our model, which works perfectly because we want to predict who will win, not precisely the number or percentage of votes, previous election data is also available to use as reference in order to tell what is the expected behavior of poll data vs election results and how bigger is the percentage of votes obtained by the election winner.

### Check out our current forecast here: <https://artofcode.tech/2022-colombian-election-forecast/>
&nbsp;

### How to interpret this model?

Currently, it's a tossup. One candidate has a higher probability of winning, but a victory is not certain, in statistics we are often certain about an event if we see a 95% probability of it happening, this is not the case here; similar scenarios have been seen in other elections like the USA 2016 Election, a model very similar to this one predicted a win for Hillary Clinton with a 71% probability, however, she lost the election; it is true that she had a higher percentage of winning but unless we see a probability of over 90% of winning, it's hard to be certain about a clear victory.

### Conclusion

Statistics and Data Science are very useful to predict results given polling data, however, it seems that many popular news outlets have problems when interpreting these numbers since it is becoming a popular sentiment among the general public that polls do not provide useful information or are rigged, hopefully here you can a find a more realistic assessment of the meaning of opinion polling results and we invite you to check our current forecast, we update it as often as we can.

### References:

Forecast: <https://artofcode.tech/2022-colombian-election-forecast/>
Github: <https://github.com/christianpaez/colombian-election-analysis-2022>
