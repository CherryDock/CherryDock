import React from 'react'
import NavBar from './../components/NavBar/NavBar';
import SideBar from './../components/SideBar/SideBar';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
} from "react-router-dom";
import style from './app-routing.module.scss';
import Containers from '../screens/Containers/Containers';

function AppRouting() {
    return (
        <div className={style.layout}>
            <div className={style.sidebar}>
                <SideBar />
            </div>
            <div className={style.appScreen}>
                <NavBar />
                <Containers />
            </div>
        </div>
    )
}

export default AppRouting
