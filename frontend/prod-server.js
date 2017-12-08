var express = require('express')
var history = require('connect-history-api-fallback');

var app = express()
app.use(history());

app.use(express.static(__dirname + '/dist'));

var server = app.listen(process.env.PORT || 3000, function () {
  var host = server.address().address
  var port = server.address().port
  console.log('App listening at http://%s:%s', host, port)
})
