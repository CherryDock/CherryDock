import React from 'react';
import style from './home.module.scss';
import ActionButtons from '../../components/Containers/ActionButtons/ActionButtons';

function Home() {
    return (
        <div className={style.container}>
            <h3 className={style.screenName}>Containers</h3>
            <ActionButtons />
        </div>
    )
}

export default Home
