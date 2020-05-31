import React, { useState, useEffect } from 'react';
import style from './home.module.scss';
import { fetchGlobalContainers } from '../../utils/api-fetch';
import LinePlotsManager from '../../components/ChartsManagers/LinePlotsManager/LinePlotsManager';
import RatioChartsManager from '../../components/ChartsManagers/RatioChartsManager/RatioChartsManager';
import { makeGlobalCtnLineItem, makeGlobalCtnRatioItem } from '../../utils/containers-data-processing';

function Home() {
    
    return (
        <div className={style.container}>
            <div className={style.containerStats}>
                <div className={style.lineStats}>
                    <LinePlotsManager
                        realTimeLimit={10}
                        fetchData={fetchGlobalContainers}
                        makePlotItem={makeGlobalCtnLineItem}
                    />
                </div>
                <div className={style.ratioStats}>
                    <RatioChartsManager
                        realTimeLimit={10}
                        fetchData={fetchGlobalContainers}
                        makePlotItem={makeGlobalCtnRatioItem}
                    />
                </div>
            </div>
        </div>
    )
}

export default Home;
