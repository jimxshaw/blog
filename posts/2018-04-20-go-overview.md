---
layout: post
title:  "Overview of Go (Golang)"
date:   2018-04-20 04:20:00 -0500
categories: go
---

<span style="display:block;text-align:center;">![Go Gopher][gogopher]</span>

Go was created by a team inside Google back in 2009. Three big languages that were used at Google at that time were C/C++, Java and Python. The creators of Go saw that each of those three languages was powerful but had limitations. For example, Python was easy to pickup and learn but ran into trouble when running apps at the scale a massive company like Google needed. Java started out simple but its type system grew increasingly complex. C/C++ was wickedly fast but it also had a complex type system and it suffered from notoriously slow compile-times. The three languages were also created at a time when multi-threaded applications were rare. Concurrency was added to them at a later date. Google's often highly concurrent and parallel environments posed challenges for them. Go was the answer to many of the concerns Python, Java and C/C++ had.

### Go Features

Go is strong and statically typed. Strong typing means that the type of a variable cannot change over time. When you declare a variable x to hold a boolean value then it will always hold a boolean value. Statically typing means that those variables must be declared at compile-time. 

Go has a strong and passionate community supporting the language. A language lives or dies based on its community support as it's the community who spreads the word and makes sure new developers wanting to learn the language have as an easy time as possible to learn it. Go certainly has a fervent community.

Go benefits from simplicity and simplicity ought to be recognized as a language feature. Simplicity as in the creators consciously left out features that are popular in other languages because they felt they would add complexity to Go.

Go has very fast compile-times and is a garbage collected language. Developers don't need to manage their own memory. While it's true that Go isn't as fast as a non-garbage collected language like C++, Go's releases over time have tried to address the pause when garbage collecting happens.   

Go has concurreny baked into the base language so there's no need to import another library for this feature.

A Go app compiles down to a single self-contained binary, which makes versioning headaches obsolete. 

### Resources

Be sure to check out the official Go [website][godocs] for more information.

[gogopher]: /assets/images/2018-04-20-go-overview/go-gopher.png

[godocs]: https://golang.org/