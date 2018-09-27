'use strict'

const HtmlWebpackPlugin = require('html-webpack-plugin')
const CopyWebpackPlugin = require('copy-webpack-plugin')
const { VueLoaderPlugin } = require('vue-loader')
const MiniCssExtractPlugin = require('mini-css-extract-plugin')

const utils = require('./utils')

const vuePkg = [
  'vue',
  'vue-router',
  'buefy',
];

module.exports = {
  entry: {
    index: utils.resolve('src/index.js'), 
    vue: vuePkg,
  },

  output: {
    path: utils.resolve('public'),
    filename: '[name].js',
  },

  resolve: {
    extensions: ['.js', '.vue', '.json'],
    alias: {
    }
  },

  optimization: {
    splitChunks: {
      chunks: 'all',
      minSize: 30000,
      minChunks: 1,
      maxAsyncRequests: 5,
      maxInitialRequests: 3,
      automaticNameDelimiter: '-',
      cacheGroups: {
        vendor: {
          test: /[\\/]node_modules[\\/]/,
          priority: -10,
          enforce: true,
        },
        default: {
          minSize: 0,
          minChunks: 2,
          priority: -20,
          reuseExistingChunk: true
        }
      }
    }
  },

  module: {
    rules: [
      {
        test: /\.vue$/,
        exclude: /node_modules/, // 排除不处理的目录
        use: 'vue-loader'
      },
      {
        test: /\.js$/,
        exclude: /node_modules/, // 排除不处理的目录
        use: {
          loader: 'babel-loader',
          options: {
            compact: 'false'
          }
        }
      },
      {
        test: /\.css/,
        use: [MiniCssExtractPlugin.loader, "css-loader", {
          loader: "postcss-loader",
          options: {
            plugins: () => [require('autoprefixer')]
          }
        }]
      },
      {
        test: /\.less$/,
        use: ['style-loader', 'css-loader', 'less-loader'] // 编译顺序从右往左
      },

      {
        test: /(\.scss|\.sass)$/,
        exclude: /node_modules/,
        use: ['style-loader', 'css-loader',   'sass-loader']
      },
      // {
      //   test: /\.scss$/,
      //   use: [
      //     "style-loader", // creates style nodes from JS strings
      //     "css-loader", // translates CSS into CommonJS
      //     "sass-loader" // compiles Sass to CSS, using Node Sass by default
      //   ]
      // },
      {
        test: /\.(png|jpe?g|gif|svg)(\?.*)?$/,
        use: {
          loader: 'url-loader',
          options: {
            limit: 10000,
            name: utils.assetsPath('img/[name].[hash:7].[ext]')
          }
        }
      },
      {
        test: /\.(mp4|webm|ogg|mp3|wav|flac|aac)(\?.*)?$/,
        use: {
          loader: 'url-loader',
          options: {
            limit: 10000,
            name: utils.assetsPath('media/[name].[hash:7].[ext]')
          }
        }
      },
      {
        test: /\.(woff2?|eot|ttf|otf)(\?.*)?$/,
        use: {
          loader: 'url-loader',
          options: {
            limit: 10000,
            name: utils.assetsPath('fonts/[name].[hash:7].[ext]')
          }
        }
      }
    ]
  },

  plugins: [

    new MiniCssExtractPlugin({
      filename: "[name].css"
    }),

    new HtmlWebpackPlugin({
      template: utils.resolve('src/assets/ejs/index.ejs'),
      filename: utils.resolve('public/index.html'),
      hash: true,
      inject: true
    }),

    new VueLoaderPlugin(),

    new CopyWebpackPlugin([{
      from: utils.resolve('src/assets/imgs'),
      to: utils.resolve('public/assets/img')
    }])
  ]
}
