import React from 'react';
import style from './side-bar.module.scss';
import cx from 'classnames'

const SideBar = () => {
    return (
        <div className={style.container}>
            <span className={style.appName}>Cherry Doc</span>
            <nav>
                <ul className={style.links}>
                    <li className={style.link}>
                        <i className={cx("fas fa-tachometer-alt", style.icon)}></i>
                        <a className={style.label} href="">Home</a>
                    </li>
                    <li className={style.link}>
                        <i className={cx("fas fa-box", style.icon)}></i>
                        <a className={style.label} href="">Containers</a>
                    </li>
                </ul>
            </nav>
        </div>
    );
}

export default SideBar;
