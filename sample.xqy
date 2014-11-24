xquery version '1.0-ml';
declare variable $XML := element request {
  attribute uuid { sem:uuid-string() },
  attribute name { "hello" },
  element message {
    xdmp:random()
  }
};
$XML