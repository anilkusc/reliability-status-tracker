const proxy = require('http-proxy-middleware');

module.exports = function(app) {
  app.use(proxy('/backend/', { target: 'http://localhost:80' }));
  app.use(proxy('/ws/status/', { target: 'ws://localhost:80/ws/status/', ws: true }));
};