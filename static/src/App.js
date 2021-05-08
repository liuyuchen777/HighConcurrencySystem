/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 00:28:43
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 01:40:51
 * @Description: 
 * @FilePath: /frontend/src/App.js
 * @GitHub: https://github.com/liuyuchen777
 */

import React from 'react'
import { BrowserRouter as Router, Route } from "react-router-dom"

// self define compoents
import Index from './page/MyIndex'
import Login from './page/Login'

class App extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      data: "React"
    }
  }

  render() {
    return (
      <Router>
        <Route path='/' exact render={(props) => <Index {...props} />} />
        <Route path='/login' exact render={(props) => <Login {...props} />} />
      </Router>
    )
  }
}

export default App