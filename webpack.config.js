/* eslint-env node */
'use strict';

var path = require('path');

var babelOptions = {
  optional: ['runtime', 'es7.asyncFunctions']
};

module.exports = {
  entry: './assets/js/index.js',
  output: {
    path: path.join(__dirname, 'static', 'js'),
    filename: 'index.js'
  },
  module: {
    loaders: [
      {
        test: /\.js$/,
        loader: 'babel?' + JSON.stringify(babelOptions)
      }
    ]
  }
};
