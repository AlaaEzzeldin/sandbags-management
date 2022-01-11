const {BASE_URL} = require("./src/api/config");

module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy: BASE_URL,
  }
}
