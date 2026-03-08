---
title:  "The Case For Boring Tools"
date:   2026-03-08 00:00:00 -0500
categories: web
layout: single
classes:
  - landing
  - dark-theme
---
![Link text]({{ site.url }}{{ site.baseurl }}/assets/img/the-case-for-boring-tools.jpg)
------------
Many of us software developers spend time on personal projects. Mine was/is a simple personal tech blog. When I first built it, I wanted to stick to the tools I was most used to:

- React.js
- Some SSR metaframework (Next or Gatsby)
- A JavaScript build system (Webpack back then)

And that is what I did. I took a template for Ghost CMS and built a basic blog for fun.

After writing for a couple of months and having updated my site over 50 times, I began to get small problems.

# 1. Writing became boring

The way the CMS worked was by taking markdown content, adding some metadata to it like canonical links, metatags, etc., storing this info in a database (yes, this needed a server), and then triggering a build on the server-side rendered site. This meant that after the fun part of writing my little blog post, I had to go to an admin panel, log in, figure out all the metatags and attributes for my blog post, and fill out a massive form with a good amount of form validation (description cannot be longer than x, title is too short, etc.), just to trigger a build and wait a couple of minutes to see my blog post. I remember many times spending more time publishing the post than actually writing it.

# 2. This setup is not simple

If I go to Medium, Quora, or Substack, I can just take markdown, an image, and some tags and be ready to post a blog. This was not what my blog had, and I began to notice that my setup was too complex. Since I was probably used to complications at work, I did not see this until a couple of months of writing. Also, running this blog locally was annoying since dependencies were always broken for some reason, as most JavaScript projects are.

# 3. It costs more than needed

At first, this was not a problem because the golden years of free Heroku dynos were still a thing. But when they nuked their free tier, I noticed that I was running a server with a database when I just needed to serve static content.

So after a while, writing became boring. I lost my admin panel credentials and forgot about the site.

# Boring just works

One day I was browsing for web templates since I wanted a blog for non-tech content and stumbled upon Jekyll. It is a super basic static site generator that has:

- No database
- No admin panel
- No JS pipeline

It is true that it does not do a lot of the fancy SSR, IWR, SCG, and other permutations of letters I can't understand (I made some of those up, and you probably did not notice). But I could write a .md file and get a blog post in less than a minute. It took me longer to log in to my admin panel and find the "Create Post" button on the page than just having my site updated.

So now:

- To publish a blog post, I spend 90% writing and 10% making the post available online
- I understand how my simple setup works
- I don't have unnecessary costs

I even use Jekyll for some other projects that are non-blog sites, and it's so nice to have something simple that does not break and does not cause friction. I know there has to be a case for a more sophisticated setup but for now I like to keep it simple.

Boring but works.
