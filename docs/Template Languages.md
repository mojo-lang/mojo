# Template Languages

[TOC]

## Languges

### [TT2](http://www.tt2.org/)

#### syntax

### [Django template language](https://docs.djangoproject.com/en/dev/ref/templates/language/)

This document explains the language syntax of the Django template system. If you’re looking for a more technical perspective on how it works and how to extend it, see [The Django template language: for Python programmers](https://docs.djangoproject.com/en/dev/ref/templates/api/).

Django’s template language is designed to strike a balance between power and ease. It’s designed to feel comfortable to those used to working with HTML. If you have any exposure to other text-based template languages, such as [Smarty](http://www.smarty.net/) or [Jinja2](http://jinja.pocoo.org/), you should feel right at home with Django’s templates.

### [Smarty](https://github.com/smarty-php/smarty/)

### [Jinja](http://jinja.pocoo.org/)

Jinja2 is a full featured template engine for Python. It has full unicode support, an optional integrated sandboxed execution environment, widely used and BSD licensed.

### [Mustache](http://mustache.github.io/)

Logic-less templates.

#### Synopsis

A typical Mustache template:

```
Hello {{name}}
You have just won {{value}} dollars!
{{#in_ca}}
Well, {{taxed_value}} dollars, after taxes.
{{/in_ca}}
```

Given the following hash:

```
{
  "name": "Chris",
  "value": 10000,
  "taxed_value": 10000 - (10000 * 0.4),
  "in_ca": true
}
```

Will produce the following:

```
Hello Chris
You have just won 10000 dollars!
Well, 6000.0 dollars, after taxes.
```

### [Mako](http://www.makotemplates.org/)

Mako is a template library written in Python. It provides a familiar, non-XML syntax which compiles into Python modules for maximum performance. Mako's syntax and API borrows from the best ideas of many others, including Django and Jinja2 templates, Cheetah, Myghty, and Genshi. Conceptually, Mako is an embedded Python (i.e. Python Server Page) language, which refines the familiar ideas of componentized layout and inheritance to produce one of the most straightforward and flexible models available, while also maintaining close ties to Python calling and scoping semantics.

Mako is used by [reddit.com](http://reddit.com/) where it delivers over [one billion page views per month](http://mashable.com/2011/02/02/reddit-surpasses-1-billion-monthly-pageviews/). It is the default template language included with the [Pylons and Pyramid](https://www.pylonsproject.org/) web frameworks.

### [Ctemplate](https://htmlpreview.github.io/?https://raw.githubusercontent.com/OlafvdSpek/ctemplate/master/doc/guide.html)

### [SSI](http://httpd.apache.org/docs/current/howto/ssi.html)

### [html_template](http://html-template.sourceforge.net/html_template.html)



## Implements

### LUA

#### [lemplate](https://github.com/openresty/lemplate) 

OpenResty/Lua template framework implementing Perl's TT2 templating language

### C++

#### [synth](https://github.com/ajg/synth)

A Powerful C++ Templating Framework

##### Implements

1. Django Template
2. SSI
3. HTML-Template

#### [Jinja2CppLight](https://github.com/hughperkins/Jinja2CppLight)

(very) lightweight version of Jinja2 for C++

Lightweight templating engine for C++, based on Jinja2

- no dependencies, everything you need to build is included
- templates follow Jinja2 syntax
- supports:
  - variable substitution
  - for loops
  - including nested for loops
  - if statements - partially: only if variable exists or not

##### License

Mozilla Public License

#### [mstch](https://github.com/no1msd/mstch)

mstch is a complete implementation of [{{mustache}}](http://mustache.github.io/) templates using modern C++. It's compliant with [specifications](https://github.com/mustache/spec)v1.1.3, including the lambda module.

#### [ctemplate](https://github.com/OlafvdSpek/ctemplate)

This library provides an easy to use and lightning fast text templating system to use with C++ programs.

It was originally called Google Templates, due to its origin as the template system used for Google search result pages. Now it has a more general name matching its community-owned nature.