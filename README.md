# Ting
###*A powerful, minimialist approach to content management.*
***
###What is Ting?

Ting is an open source alternative to the 'Enterprise CMS'. It is made up of a series of small, scalable open source components.

###Who is it aimed at?
If you're a non-programmer or a small non-technical company, solutions like Wordpress are fine. Plugins can be installed quickly and without technical knowledge. Content can be managed effectively. You're generally not going to have many problems, at least early on.

The enterprise CMS is a different beast. It's not working. It doesn't work for clients and it certainly doesn't work for developers. Modern web frameworks have evolved far beyond what's possible with a monolithic enterprise CMS. Deployments become a nightmare. Build scripts become unfathomable. Databases end up littered with legacy fields, or even legacy *tables*. As the codebase grows architectural principles go out the window, and the project ends up as a mass of techincal debt where the developers are terrified to change any part of the code. While CMS projects can technically be checked into version control, there are so many caveats that most of the benefits are lost.

###The cons of current solutions.
The internet is saturated in CMS solutions, and many software houses still prefer to write their own. The problems that can arise from this are too many to list, but here are the main culprits:

* *Fully-featured CMS solutions are bulky and slow.*
* *Clients may only use a small portion of the features offered by their CMS.*
* *Clients often require custom functionality which is very time-consuming to add into an existing CMS.*
* *The client often has the power to break their own website.*
* *Time consuming GUI-based changes made on the dev server have to be painstakingly replicated on staging and live.*
* *No database migrations - Databases are repeatedly dumped and imported across builds.*
* *CMS solutions generally take a 'command-line free' approach. This applies not only to managing content, but often to managing databases, deployments, plugins, styling, and other factors the client should have nothing to do with.*
* *DevOps engineers often have nightmares about enterprise CMS solutions.*
* *The benefits of using a version control system are lost.*
* *Because it's one codebase, there is an incredible amount of cognitive load in even attempting to fix issues or accomodate new requirements.*

###How Ting is different.
The idea behind Ting is that content management should be kept isolated from the rest of the system. It aims to be like [prismic.io](http://prismic.io), but self-hosted, open source, and importantly, fully compatible with modern DevOps/PaaS/IaaS software like [Flynn](http://flynn.io), [Deis](http://deis.io), [CoreOS](http://coreos.com), [Docker](http://docker.io), etc.

At the core of Ting is a RESTful API layer backed by a NoSQL database resource (Currently MongoDB only). Ting defines repositories, which in turn define the types of content used in that project. We use document-oriented database backends because a piece of content is essentially a document.

###Todo:

* Design the components.
* Design the primitives.
* Design the migration system.
* Actually write some code.