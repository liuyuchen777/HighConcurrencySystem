/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 23:57:46
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-09 00:02:30
 * @Description: 
 * @FilePath: /spike_system/static/src/page/About.js
 * @GitHub: https://github.com/liuyuchen777
 */
import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import Footer from '../compoents/Footer'

export default class About extends Component {

  render() {
    return (
      <div className='container'>
        <center>
            <h4>Version 1.0.0</h4>
            <Link to="/">Go Back</Link>
            <Footer />
        </center>
      </div>
    )
  }
}
