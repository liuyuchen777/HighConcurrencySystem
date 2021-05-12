/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 01:27:15
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 01:27:45
 * @Description: 
 * @FilePath: /frontend/src/setupProxy.js
 * @GitHub: https://github.com/liuyuchen777
 */

const {createProxyMiddleware} = require('http-proxy-middleware')

module.exports = function (app) {
    // proxy第一个参数为要代理的路由
    // 第二参数中target为代理后的请求网址，changeOrigin是否改变请求头，其他参数请看官网
    app.use(createProxyMiddleware('/api', {
        target: 'http://192.168.231.129:9090',
        changeOrigin: true
    }))
}