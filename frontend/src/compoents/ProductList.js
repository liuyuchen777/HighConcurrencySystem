/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 23:55:40
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-09 00:30:53
 * @Description: 
 * @FilePath: /spike_system/static/src/compoents/ProductList.js
 * @GitHub: https://github.com/liuyuchen777
 */
import React, { Component } from 'react'
import Product from './Product'

export default class ProductList extends Component {
  render() {
    const numbers = Array.from({length: this.props.products.length}, (_, k) => k)
    const listItems = numbers.map((number) => 
        <Product key={number.id} product={this.props.products[number]}/>
    )

    return (
      <div className='sub-container'>
        <ul>{listItems}</ul>
      </div>
    )
  }
}
