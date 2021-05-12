/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-09 00:28:06
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-09 00:29:19
 * @Description: 
 * @FilePath: /spike_system/static/src/compoents/Logout.js
 * @GitHub: https://github.com/liuyuchen777
 */
import React, { Component } from 'react'

export default class Logout extends Component {

  clickHandler() {
    console.log("I am Clicked!")
  }

  render() {
    return (
      <div>
        <center>
          <button className="btn" onClick={this.clickHandler}>Logout</button>
        </center>
      </div>
    )
  }
}
