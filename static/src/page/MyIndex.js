/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 00:59:05
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-09 00:01:47
 * @Description: 
 * @FilePath: /spike_system/static/src/page/MyIndex.js
 * @GitHub: https://github.com/liuyuchen777
 */

import React from 'react'
import { Link } from "react-router-dom"
import Footer from '../compoents/Footer'

class Index extends React.Component {
  constructor(props) {
    super(props)
    this.state = {

    }
  }

  render() {
    return (
      <center>
      <div className="container">
          <h1 className="header">Welcome to Spike.com!</h1>
        <Link to="/login">
        <button to='/login' className="btn btn-fix" style={{ width: "8em" }}>Login</button>
        </Link>
        <br />
        <button className="btn btn-fix" style={{ width: "8em" }}>Register</button>
        <Footer />
      </div>
      </center>
    )
  }
}

export default Index