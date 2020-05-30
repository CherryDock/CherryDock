import React, { useState, useEffect } from 'react';
import style from './home.module.scss';
import LinePlot from '../../components/Charts/LinePlot/LinePlot';
import RatioPlot from '../../components/Charts/RatioPlot/RatioPlot';
import { lineDummyData, ratioDummyData } from './../ContainerDetails/dummy-data';
import { fetchDataApi } from '../../utils/api-fetch';
import apiConf from '../../conf/api.conf';
import { linePlotConf } from '../../conf/charts.conf';
import Axios from 'axios';
import { LinePlotProps, LinePlotData, LinePlotItem } from '../../interfaces/charts.interface';
import { makeGlobalCtnItem } from '../../utils/containers-data-processing';
import { GlobalContainers } from '../../interfaces/data.interface';
import { fetchGlobalContainers } from '../../utils/api-fetch';
import LinePlotsManager from '../../components/ChartsManagers/LinePlotsManager/LinePlotsManager';

function Home() {
    const [realTimeStats, setRealTimeStats] = useState<GlobalContainers>();
    const [linePlotData, setLinePlotData] = useState<LinePlotData[]>([]);
    const realTimeLimit = 10;

    
    const displayRatioPlot = ratioDummyData.map(data => {
        return (
            <div key={data.id} className={style.plot}>
                <RatioPlot heightScreenRatio={0.35} data={data.data} title={data.title} />
            </div>
        )
    });

    return (
        <div className={style.container}>
            <div className={style.containerStats}>
                <div className={style.lineStats}>
                    <LinePlotsManager 
                        realTimeLimit={10}
                        fetchData={fetchGlobalContainers}
                        makePlotItem={makeGlobalCtnItem}
                        />
                </div>
            </div>
        </div>
    )
}

export default Home;
