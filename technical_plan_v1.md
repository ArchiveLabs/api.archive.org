# Technical Plan, Version 1.0 #

## Goals

The goals of the first version of `api.archive.org` are to:

- Gather all existing and emerging Interet Archive API endpoints into one location
- Gather API endpoint _documentation_ into one location
- Provide authentication/authorization consistent with current practices
- Create an technical architecture that is well tested and ready for grown
- Learn additional requirements for Version 2.0 of an API

## General Plan and resourcing

Version 1.0 will be delivered in a series of half-month sprints starting 
April 1 and going through May 31.

Will Fitzgerald (WF) of the Internet Archive will lead the effort, and spend half
time (0.5 FTE) on it. Michael (Mek) Karpeles will consult with approximately 
(x.x FTE). 

Other staff and voluteers will advise and work as needed and appropriate.

As much as possible, progress and design will happen in the open at the [api repository](ttps://github.com/ArchiveLabs/api.archive.org/)

## Sprint 1: Information Gathering (April 1-15)

Tasks:

- Compare PHP/Python/Go/Kong endpoints on:
  - Ease of integration
  - latency
  - number of concurrent sessions
  - memory use
  - CPU use
- Report and recommend V1 endpoint architecture
- Compare documentation methods, specifically Swagger/Raml/Slide/Others on:
  - ease of use
  - design and beauty
  - discoverabilty
- Report and recommend documentation method
- Make initial list of APIs
  
## Sprint 2: Implementing 1/3 of APIs, Investigating AUTH (April 16-30)

Tasks:

- Implement "archive labs" APIs
- Implement Scraping API, metadata read API (including documentation)
- Others?
- Investigate using OAuth 2.0 within authentication/authorization needs
- Report on OAuth

## Sprint 3: Implementing OpenLibrary Dash API with AUTH, other APIs (May 1-15)

Tasks:

- Implement OpenLibrary Dash API with AUTH
- Implement as many APIs left as possible (including documentation)
- Plan Sprint 4

## Sprint 4: Final cleanup and further planning (May 16-31)

Tasks:
 - Final cleanup, finish APIs
 - Further planning towards 2.0
 - Marketing and communication


