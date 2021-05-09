/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 00:28:43
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-09 00:32:17
 * @Description: 
 * @FilePath: /spike_system/static/src/App.js
 * @GitHub: https://github.com/liuyuchen777
 */

import React from 'react'
import { BrowserRouter as Router, Route } from "react-router-dom"

// self define compoents
import Index from './page/MyIndex'
import Login from './page/Login'
import Register from './page/Register'
import Main from './page/Main'
import About from './page/About'

class App extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      data: "React",
      username: "",
    }
  }

  render() {
    return (
      <Router>
        <Route path='/' exact render={(props) => <Index {...props} />} />
        <Route path='/main' exact render={(props) => <Main {...props} />} />
        <Route path='/login' exact render={(props) => <Login {...props} />} />
        <Route path='/register' exact render={(props) => <Register {...props} />} />
        <Route path='/about' exact render={(props) => <About {...props} />} />
      </Router>
    )
  }
}

export default App