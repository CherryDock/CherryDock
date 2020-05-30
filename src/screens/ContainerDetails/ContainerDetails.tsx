import React, { useState, useEffect } from 'react';
import style from './container-details.module.scss';
import { useParams } from "react-router-dom";
import ActionButtons from '../../components/ContainerDetails/ActionButtons/ActionButtons';
import LinePlot from '../../components/Charts/LinePlot/LinePlot';
import RatioPlot from '../../components/Charts/RatioPlot/RatioPlot';
import { lineDummyData, ratioDummyData } from './dummy-data';
import LinePlotsManager from '../../components/ChartsManagers/LinePlotsManager/LinePlotsManager';
import { fetchSingleContainer } from '../../utils/api-fetch';
import { makeSingleCtnItem } from '../../utils/containers-data-processing';

function ContainerDetails() {
    const { id } = useParams();

    return (
        <div className={style.container}>
            <div className={style.actionButons}>
                <ActionButtons />
            </div>
            <div className={style.containerStats}>
                <div className={style.lineStats}>
                    <LinePlotsManager
                        realTimeLimit={10}
                        fetchData={fetchSingleContainer}
                        makePlotItem={makeSingleCtnItem}
                    />
                </div>
            </div>
        </div>
    )
}

export default ContainerDetails
