xquery version '1.0-ml';
declare namespace html = "http://www.w3.org/1999/xhtml";
(
  xdmp:get-server-field-names(),
  xdmp:estimate(xdmp:directory("/"))
)
