# notes
## definitions
**feature**: "abstractions of real-world phenomena" (https://www.w3.org/TR/sdw-bp/#spatial-things-features-and-geometry)
    - essentialy a view that contains everything of interest

**OGC API**: the family of standards for geospatial Web APIs that in ISO is published as "Geospatial API;"

**OGC API - Features**: the multipart standard for features that in ISO is published as ISO 19168 / "Geographic Information - Geospatial API for Features;" and

**OGC API - Features - Part 1**: Core: this document that in ISO is published as ISO 19168-1 / "Geographic Information - Geospatial API for Features - Part 1: Core."



# Assorted info
- Standard expects everything to be specified with http GET(other methods will be referenced later)
- query operations are to retrieve features from data store based upon provided filter criteria


# Assorted reminders
- Each feature should only be a part of one collection

# req and recc
- http 1.1
    - req method HEAD for anything that supports GET
- OPTIONS (if cors is implemented)
    - If the server is intended to be accessed from the browser, cross-origin requests SHOULD be supported. Note that support can also be added in a proxy layer on top of the server.

- status codes https://docs.ogc.org/is/17-069r4/17-069r4.html#status_codes
    - 400 response anytime unspecified query parameters are given
- entity tags for web caching
- html is reccomended alongside geojson for:
    - search engine crawling
    - allow users to view data set via web browser
- Unless the client explicitly requests a different coordinate reference system, all spatial geometries SHALL be in the coordinate reference system http://www.opengis.net/def/crs/OGC/1.3/CRS84 (WGS 84 longitude/latitude) for geometries without height information and http://www.opengis.net/def/crs/OGC/0/CRS84h (WGS 84 longitude/latitude plus ellipsoidal height) for geometries with height information.
- links in payload should be included as header
- all 200's must support:
    - html (only if we want to meet html req)
    - geo json for resources included feature content (only if we want to meet geojson req class)
    - json for resources that do not include feature content (only if we want to meet geojson req class)


# Routes
High level overview of the routes and their implementaion
# */* Landing page
- meet yaml requirements
- links to
    - api definitions
    - conformance endpoint
    - collections endpoint
# */api-definition*
- defintion describing capability of the server
# */conformance* Conformance
- list of what api should conform(ex: json, html, geojson...)
# */collections*
- return list of all collections
- must support GET
- lots of meta data requirements
- extent property to list spatial and temporal requirements
# */collections/{collectionId}* feature collection
- A local identifier for the collection that is unique for the dataset;
- A list of coordinate reference systems in which geometries may be returned by the server: the first CRS is the default CRS (in the Core, the default is always WGS 84 with axis order longitude/latitude);
- An optional title and description for the collection;
- An optional extent that can be used to provide an indication of the spatial and temporal extent of the collection - typically derived from the data;
- An optional indicator about the type of the items in the collection (the default value, if the indicator is not provided, is 'feature').
## */collections/{collectionId}/features* Features

# */collections/{collectionId}/items
- accesing *features* should return all features in a collection
- parameters:
    - limit
    - bbox
    - datetime(can be interval)
    - filtering(for "simple" filterable characteristics)
- should return metadata such as numberMatched...
# */collections/{collectionId}/items/{featureId}*



Open API has additional requirements

https://docs.ogc.org/is/17-069r4/17-069r4.html
