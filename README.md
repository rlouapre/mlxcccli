mlxcccli
========

MarkLogic cli in Go

Execute XQuery script on XCC Server
- Provide configuration in xcc.yml
- Or provide it as parameters

Required parameters:
- host
- port
- authentication: basic or digest supported
- username
- password
- filename: xquery script to be executed
