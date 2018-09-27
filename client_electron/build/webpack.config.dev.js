'use strict'

// const webpack = require('webpack')
const merge = require('webpack-merge')
const baseConfig = require('./webpack.config.base')
// const ExtractTextPlugin = require("extract-text-webpack-plugin");

module.exports = merge(baseConfig, {
  mode: 'development',

  module: {
  
  },

  plugins: [  
  ]
})
