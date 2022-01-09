# Concurrent Web Scrapper

A Web Scrapper that crawls a given web page, as well as all the sub-pages linked from that page (part of the same domain), their referred sub-pages, and so on recursively.
All the images from these pages are extracted and stored in a given folder., and stores them in a given folder without duplicates and without processing the same URL twice (without cycles).
Duplicate URLs and Images are not processed twice.

The program receives as command line arguments:

1) Name of the folder to store images to (creates a folder if it is non-existant);

2) URL of the start webpage;

3) MaxRoutines number - the maximal number of goroutines that should be run concurrently at any given moment;

5) MaxResults - the maximum number of web pages processed;
