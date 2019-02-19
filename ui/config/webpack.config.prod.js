const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
const webpack = require('webpack');

module.exports = {
  mode: 'production',
  entry: [
    './src/index.js'
  ],
  output: {
    pathinfo: true,
    filename: 'assets/[name].[chunkhash].js',
    publicPath: '/beta/'
  },
  resolve: {
    extensions: ['.js', '.jsx']
  },
  optimization: {
    minimizer: [
      new UglifyJsPlugin()
    ],
    splitChunks: {
      chunks: 'all'
    }
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        use: 'babel-loader'
      },
      {
        test: /\.css$/,
        exclude: /(index.css|global.css|node_modules)/,
        use: [
          {
            loader: MiniCssExtractPlugin.loader
          },
          {
            loader: 'css-loader',
            options: {
              importLoaders: 1,
              localIdentName: "[name]__[local]___[hash:base64:5]",
              modules: true,
            },
          }
        ]
      },
      {
        test: /\.css$/,
        include: /(index.css|global.css|node_modules)/,
        use: [
          {
            loader: MiniCssExtractPlugin.loader
          },
          {
            loader: 'css-loader',
            options: {
              importLoaders: 1
            },
          }
        ]
      }
    ]
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: 'assets/[name].[chunkhash].css'
    }),
    new HtmlWebpackPlugin({
      inject: true,
      template: './public/index.html'
    }),
    new webpack.DefinePlugin({
      API_URL: JSON.stringify('http://paste.click')
    })
  ]
};
