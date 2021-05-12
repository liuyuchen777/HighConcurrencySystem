/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-09 00:22:44
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-09 00:30:14
 * @Description: 
 * @FilePath: /spike_system/static/src/compoents/Product.js
 * @GitHub: https://github.com/liuyuchen777
 */
import React, { Component } from 'react'

export default class Product extends Component {
  constructor(props) {
    super(props)
    console.log(props.product)
  }

  render() {
    return (
      <div className="product-container">
        <h2 align="left">{this.props.product.Name}</h2>
      </div>
    )
  }
}
