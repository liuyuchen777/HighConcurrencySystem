/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 01:38:48
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-09 01:00:10
 * @Description: 
 * @FilePath: /spike_system/static/src/page/Main.js
 * @GitHub: https://github.com/liuyuchen777
 */

import React from 'react'
import ProductList from '../compoents/ProductList'
import Footer from '../compoents/Footer'
import Logout from '../compoents/Logout'
import ShowProducts from '../compoents/ShowProducts'

class Main extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      products: Object
    }
  }

  showHandler() {
    
  }

  render() {
    return (
      <div className="container">
        <center>
          <h1>Spike.com</h1>
          <h4>Buy Good Things</h4>
          <ProductList />
          <Logout />
          <ShowProducts />
          <Footer />
        </center>
      </div>
    )
  }
}

export default Main
