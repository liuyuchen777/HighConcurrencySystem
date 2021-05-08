/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 01:00:39
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 01:59:15
 * @Description: 
 * @FilePath: /frontend/src/page/Login.js
 * @GitHub: https://github.com/liuyuchen777
 */
import React from 'react'
import axios from 'axios'

class Login extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      username: "",
      password: ""
    }
    this.onSubmit = this.onSubmit.bind(this)
  }

  onSubmit(e) {
    e.preventDefault()

    console.log(this.state.username)
    console.log(this.state.password)

    // check
    if (this.state.username === "") {
      alert("Please enter username!")
    } else if (this.state.password === "") {
      alert("Please enter password!")
    }

    // communicate with backend to verify
    axios.post('/api/login', {
      username: this.state.username,
      password: this.state.password
    })
    .then((response) => {
      console.log(response)
    })
  }

  render() {
    return (
      <div className="container">
        <form className='add-form' onSubmit={this.onSubmit}>
            <div className='form-control'>
            <h2>Username: </h2>
            <input type='text' placeholder='username' 
            value={this.state.username} onChange={(e) => this.setState({username: e.target.value})}/>
            </div>

            <div className='form-control'>
            <h2>Password: </h2>
            <input type='password' placeholder='password' 
            value={this.state.password} onChange={(e) => this.setState({password: e.target.value})}/>
            </div>

            <input type='submit' value='Login' className='btn btn-block'/>
        </form>
      </div>
    )
  }
}

export default Login