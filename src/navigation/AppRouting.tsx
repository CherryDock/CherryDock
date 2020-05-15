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
import Home from '../screens/Home/Home';
import ContainerDetails from '../screens/ContainerDetails/ContainerDetails';

function AppRouting() {
    return (
        <Router>
            <div className={style.layout}>
                <div className={style.sidebar}>
                    <SideBar />
                </div>
                <div className={style.appScreen}>
                    <NavBar />
                    <Switch>
                        <Route exact path="/">
                            <Home />
                        </Route>
                        <Route path="/containers">
                            <Containers />
                        </Route>
                        <Route path="/container-details/:id">
                            <ContainerDetails />
                        </Route>
                    </Switch>
                </div>
            </div>
        </Router>
    )
}

export default AppRouting
