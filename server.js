'use strict';

const express = require('express');

// Constants
const PORT = 3000;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/', (req, res) => {
  res.send('updated version PART 3');
});
app.get('/api', (req, res) => {
  res.send('api e...o');
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT} ...`);