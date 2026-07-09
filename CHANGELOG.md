# v1.0.0

* Initial commit

# v1.1.0

* Ignoring links without any title
* CORS enabled

# v1.2.0

* Replaced HTML parsing with RSS parsing for news
* Replaced Geziyor with http.Client + goquery for faster parsing
* Added gzip compression for response payloads
* Optimized HTTP client connection pooling
* Added request timeout middleware
* Improved date comparison