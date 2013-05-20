GLIP - GitHub Language Index of Popularity
==========================================

GLIP compares the relative popularity of programming languages by calculating a 
popularity score. Simply speaking, a programming language's popularity score is 
the sum of stars of a programming language's 20 most watched projects on 
GitHub.

This program scrapes GitHub (there's not API for that kind of stuff, yet) for 
that information and generates a table for comparison.


Example
-------

    go run glip-scrape/scrape.go | go run glip-gen/gen.go -t glip-gen/html-chart.tmpl > chart.html


Author
------

Andreas Krennmair <ak@synflood.at>


License
-------

See the file LICENSE.md for licensing terms.
