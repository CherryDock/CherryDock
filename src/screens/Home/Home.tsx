import React from 'react';
import style from './home.module.scss';
import LinePlot from '../../components/Charts/LinePlot/LinePlot';
import RatioPlot from '../../components/Charts/RatioPlot/RatioPlot';
import { lineDummyData, ratioDummyData } from './../ContainerDetails/dummy-data';

/*
    token = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjY2ODkxNDUsInVzZXIiOiIifQ._nUdbYxWsBI-clzFr16M30D4zyvUu7-SRsAfVisZvLg
*/

function Home() {
    const displayLineChart = lineDummyData.map(data => {
        return (
            <div key={data.id} className={style.plot}>
                <LinePlot heightScreenRatio={0.35} data={data.data} title={data.title} labels={data.labels} />
            </div>
        )
    });

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
                    {displayLineChart}
                </div>
                <div className={style.ratioStats}>
                    {displayRatioPlot}
                </div>
            </div>
        </div>
    )
}

export default Home
