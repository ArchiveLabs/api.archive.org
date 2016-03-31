# Archive.org API requirements

Will Fitzgerald, will@archive.org  
Last change: March 11, 2016  

## General goals

We want to provide a set of well-supported, consistent, well-documented APIs for internal and external consumption, including a transition path for current API use. 

This document is primarily about API calls over HTTP.

## General 

1. In general terms, the ideas behind [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) should guide API development
2. The APIs should be federated from a single host (e.g. api.archive.org)
3. APIs should be well-documented:
  - Parameters and their types
  - Return values and their types
  - Location of parameter data (query parameters, etc)
4. APIs and their documentation should reside on Internet Archive hardware
5. APIs should be versioned, and support for versioned levels should be clear and allow for eventual sun-setting.
6. The archive.org stable endpoints will not change. (e.g. http://archive.org/details/current, http://www.archive.org/stream/aliceinwonderlan00carriala#page/23). There may be API endpoint version equivalents, however (and we need to handle the possibility of confusion).
7. The APIs should be discoverable.
8. The APIs should be well-reviewed and tested.
9. APIs should be “as simple as possible” but for consumers and our developers (for example, one function per API, composable, consistency in naming etc)

## Data Formats

### Request
(Kenji notes: There was a debate on request format for availability API v2. form-urlencoded? mixing URL parameters and POST body? JSON? JSONL? PHP-style array encoding? CSV? how to identify request format?)

### Response

1. For non-streaming APIs, Nominal and Error values returned should be in object format. (Object here means a set of unordered key value pairs, where the key is a string, and the value is either another object, or a “plain type” (number, string, array, boolean, or a null object), or an array of objects or plain types. Arrays elements should be of all the same type. Objects that are semantically dates or timestamps should be either strings in ISO 8601 format, or the number of seconds since the Unix epoch.
2. The above statement implies nominal and error values should not be returned as arrays or plain types
3. Streaming APIs should return a stream of objects (note: this is different sense from the archive.org/stream endpoint). 
4. Nominal and Error values should be encoded as JSON. Specific cases may provide data in other object formats (e.g. XML).
5. Responses should also return an appropriate HTTP status code  

### Security 
1. APIs should have security appropriate to the service provided (authorization/authentication)
2. APIs should be rate-limitable; in particular, it should be possible to prevent DDOS-like events by revoking access

### Routing and Parameters

1. Routes should not expose filetypes (e.g., .php)
2. Routes should expose their version number 
3. Routes should follow the following format: /{general_service}/{version}/{specific_service}[{specific_object}]  
4. Version will expose major version only (with, perhaps, exceptional cases for major plus minor version)
5. When applicable, the following parameters should be used:
   - format, the format of the return type; defaults to json
   - callback, the Javascript or Javascript-like callback to be used (JSONP).
   - q, a query parameter
   - start, the location in a stream of data from which to start (for example, the nth item in a set of search results; zero-based.)
   - count, the number of results requested
   - item_id, for item ids (but, generally, the item id should be part of the route)
   - mediatype for media types.
   - whoami, to give callers who can’t change their user-agent an easy way to identify themselves

Examples (to be filled out):

    /search/1.0/scrape?q=William+BIllings&size=100
    /item/1.0/metadata/commute


## Discoverability of the API

The plan of record for an API should be a file in YAML or JSON format as defined by the Open API Initiative (OAI), 
also known as “Swagger.”

For a particular general service, the endpoint /{service}/{version} should return the plan of record JSON 
format for all specific services.

## Testing

Unit testing the endpoints can ensure contract behavior is being met.

An IA-maintained reference library can help bootstrap other language libraries being developed in the open-source world.  
(A system like Swagger/OpenAPI may be able to auto-generate bindings.)

An API review for each endpoint can ensure consistency across all APIs before its released to the wild, 
where it’s more difficult to revise.
