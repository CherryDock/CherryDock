import React from 'react';
import style from './side-bar.module.scss';
import cx from 'classnames';
import { Link } from "react-router-dom";

const SideBar = () => {
    const brandLogo = require('../../assets/img/brand-image.png');
    return (
        <div className={style.container}>
            <div className={style.brand}>
                <img className={style.brandImg} src={brandLogo} alt=""/>
                <span className={style.appName}>Cherry Doc</span>
            </div>
            <nav>
                <ul className={style.links}>
                    <li className={style.link}>
                        <i className={cx("fas fa-tachometer-alt", style.icon)}></i>
                        <Link className={style.label} to='/'>Home</Link>
                    </li>
                    <li className={style.link}>
                        <i className={cx("fas fa-box", style.icon)}></i>
                        <Link className={style.label} to='/containers'>Containers</Link>
                    </li>
                </ul>
            </nav>
        </div>
    );
}

export default SideBar;
