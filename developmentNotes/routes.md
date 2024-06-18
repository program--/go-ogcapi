
# Routes Outline
High level overview of routes that will be implemented. Each route has additional YAML requirements that will not be specified here.
## */* Landing page(home)
- links to
    - api definitions
    - conformance endpoint
    - collections endpoint
## */api-definition*
- defintion describing capability of the server
# */conformance* Conformance
- list of what api should conform(ex: json, html, geojson...)
# */collections*
- must support GET
    - essentially a list of all collections
    - lots of meta data requirements
    - extent property to list spatial and temporal requirements
# */collections/{collectionId}* feature collection
# */collections/{collectionId}/items
- parameters:
    - limit
    - bbox
    - datetime(can be interval)
    - filtering(for "simple" filterable characteristics)
- **unsure what the diff between this and no item route is**
# */collections/{collectionId}/items/{featureId}*

