---
layout: post
title:  ".NET Core: The Big Picture"
date:   2017-06-12 16:10:34 -0500
categories: dotnet
---

![dotnet][dotnet]

The new-ish [.NET Core][dotnetcore] is a modular verion of the full [.NET Framework][fulldotnet]. Microsoft designed .NET Core to be portable across multiple platform in order to maximize code sharing and code re-use. It's a subset of the full .NET Framework that provides key functionalities to create the application we want and be able to reuse that app's code wherever, regardless of platform. 

Our choice of platform makes no difference as .NET Core features are universal across all of them. For web apps, we could choose Windows, Linux or MacOS. Non-web app platforms include desktops, mobile and other devices. 

.NET Core is called modular because it's divided into many small specialized assemblies on Nuget as opposed to a single massive assembly with the full .NET Framework. This allows us to individually select the packages we want to incorporate in our app while bypassing everything else. Our app's size is kept to a minimum this way.

Finally, .NET Core implements the new [.NET Standard][dotnetstandard]. The purpose of .NET Standard is to have an official agreement of .NET apis that are available across all platforms of .NET. It's essentially one giant library to rule them all. Code sharing between multiple .NET platforms becomes easier for those platforms that implement this standard. The old way of using Portable Class Libraries jumping from platform to platform should fade.

### ASP.NET Core

#### Overview

ASP.NET Core is an open-source, cross-platform framework for building modern, cloud-based internet connected applications. In addition to web apps, this framework can be used for mobile apps and [Internet-Of-Things][iot] apps as well. 

Being open-souce, we can peruse ASP.NET Core's source code at [github.com/aspnet][aspnetcoresource] and freely contribute to that project. 

ASP.NET Core is cross-platform so we can write our apps on Windows, MacOS or Linux. A variety of IDEs and text editors are available are on those platforms from the ever reliable Visual Studio to Visual Studio Code to [Project Rider by JetBrains][projectrider]. 

Microsoft created ASP.NET Core from scratch. It's not an update to ASP.NET 4 as that framework was based on the System.Web namespace, which meant having many excess functionalities. The Core team decided to forgo big and bulky for granular and lean. As with .NET Core, ASP.NET Core is also based on modular Nuget packages. We only import the packages we need for our app. Less excess results in tighter security, reduced maintenance and improved performance.

#### Frameworks

ASP.NET Core apps can target either the full .NET Framework or .NET Core. If we choose full .NET then we'll have all the dependencies we're used to having. Since .NET Core is newer, not all the assemblies are available yet. There's the possibility that .NET Core will never port over certain assemblies. However, if the app wants to maximize modularity, have a small memory footprint and be cross-platform then nothing beats targeting .NET Core. Unless we utilize .NET features to it's fullest, ASP.NET Core apps should choose .NET Core as their framework.

[dotnet]: https://weblog.west-wind.com/images/2016/ASP.NET%20Core%20Overview/NetPlatformOverviewTomorrow.png
[dotnetcore]: https://docs.microsoft.com/en-us/dotnet/core/
[fulldotnet]: https://docs.microsoft.com/en-us/dotnet/framework/
[dotnetstandard]: https://docs.microsoft.com/en-us/dotnet/standard/library
[iot]: https://en.wikipedia.org/wiki/Internet_of_things
[aspnetcoresource]: https://github.com/aspnet
[projectrider]: https://www.jetbrains.com/rider/