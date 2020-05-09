import React from 'react';
import style from './nav-bar.module.scss';
import cx from 'classnames'

function NavBar() {
    const imgPath = require('../../assets/img/profile_pic_white.jpg');
    return (
        <header className={style.container}>
            <span className={style.currentMenu}>Containers List</span>
            <div className={style.account}>
                <img className={style.avatar} src={imgPath} alt="INF" />
                <span className={style.username}>Admin</span>
            </div>
            <div className={style.logout}>
                <i className={cx("fas fa-sign-out-alt", style.icon)}></i>
                <span className={style.label}>Log out</span>
            </div>
        </header>
    )
}

export default NavBar;
