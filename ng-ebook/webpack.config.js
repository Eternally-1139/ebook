var path = require('./');

module.exports = {
  entry: {
    main: './app/main.ts'
  },
  output: {
    path: path.resolve(__dirname, 'dist'),
    publicPath: '/dist/',
    filename: 'bundle.js'
  },
  module: {
    loaders: [
      {
        test: /\.ts$/,
        loader: 'ts'
      }
    ]
  },
  // require 文件时可省略后缀 .js 和 .ts
  resolve: {
    extensions: ['', '.js', '.ts']
  },
  // 配置 webpack-dev-server
  devServer:{
    historyApiFallback: true,
    hot: true,
    inline: true,
    port: 9090 // 修改端口，一般默认是8080
  }
};
