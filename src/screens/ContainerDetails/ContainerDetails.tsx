import React, { useState, useEffect } from 'react';
import style from './container-details.module.scss';
import { useParams } from "react-router-dom";
import ActionButtons from '../../components/ContainerDetails/ActionButtons/ActionButtons';
import LinePlot from '../../components/Charts/LinePlot/LinePlot';
import RatioPlot from '../../components/Charts/RatioPlot/RatioPlot';
import { lineDummyData, ratioDummyData } from './dummy-data';

function ContainerDetails() {
    const { id } = useParams();

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
            <div className={style.actionButons}>
                <ActionButtons />
            </div>
            <div className={style.containerStats}>
                {displayLineChart}
                {displayRatioPlot}
            </div>
        </div>
    )
}

export default ContainerDetails
