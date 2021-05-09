/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 23:57:17
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 23:57:34
 * @Description: 
 * @FilePath: /spike_system/static/src/compoents/Footer.js
 * @GitHub: https://github.com/liuyuchen777
 */
import { Link } from 'react-router-dom'

function Footer() {
    return (
        <footer>
            <p>Author: @<a href="https://liuyuchen777.github.io">liuyuchen777</a></p>
            <p>Copyright &copy; 2021</p>
            <Link to='/About'>About</Link>
        </footer>
    )
}

export default Footer