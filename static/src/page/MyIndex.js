/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 00:59:05
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 01:23:23
 * @Description: 
 * @FilePath: /frontend/src/MyIndex.js
 * @GitHub: https://github.com/liuyuchen777
 */

import React from 'react'
import { Link } from "react-router-dom"

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
      </div>
      </center>
    )
  }
}

export default Index