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
import { makeLinePlotItem } from '../../utils/containers-data-processing';
import { GlobalContainers } from '../../interfaces/data.interface';
import { fetchGlobalContainers } from '../../utils/api-fetch';

function Home() {
    const [realTimeStats, setRealTimeStats] = useState<GlobalContainers>();
    const [linePlotData, setLinePlotData] = useState<LinePlotData[]>([]);
    const realTimeLimit = 10;

    /**
     * Init linePlotData
     */
    function initLinePlotData() {
        fetchGlobalContainers('/api/monitor/stats')
            .then(data => {
                if (data != undefined) {
                    let id = 1;
                    linePlotConf.forEach(plotConf => {
                        const dataItem = [makeLinePlotItem(data, plotConf.kpiName)!];
                        const plotData: LinePlotData = {
                            id, title: plotConf.title, data: dataItem
                        }
                        id++;
                        setLinePlotData(emptyArray => [...emptyArray, plotData]);
                    })
                }
            });
    }

    /**
     * Shift all the elements of an Array one position to the left in order
     * to delete the first element and make space for the new one
     * @param array - Generic Array
     */
    function shiftArrayLeft<T>(array: T[]): T[] {
        /* Delete the last item of the array by slicing it (0, n-1) */
        array = array.slice(1, array.length);
        return array;
    }

    /**
     * Update linePlotData
     */
    function updateLinePlotData() {
        /* API call */
        const rtStatsRoute = apiConf.routes.find(route => route.name === 'GET-ALL-RT-CTN-STATS')?.route!;
        //fetchDataApi<GlobalContainers>(rtStatsRoute)
        fetchGlobalContainers('/api/monitor/stats')
            .then(data => {
                if (data != undefined) {
                    let updatedData = linePlotData.map(plotData => {
                        const plotTitle = plotData.title;
                        const kpiName = linePlotConf.find(plotConf => plotConf.title === plotTitle)?.kpiName!;
                        const dataItem = makeLinePlotItem(data, kpiName);

                        if (plotData.data.length > realTimeLimit) {
                            console.log(plotData.data.length);
                            plotData.data = shiftArrayLeft<LinePlotItem>(plotData.data);
                            console.log(plotData.data.length);
                        }

                        if (dataItem != undefined)
                            plotData.data.push(dataItem)

                        return plotData;
                    });
                    setLinePlotData(updatedData);
                }
            });
    }

    useEffect(() => {
        initLinePlotData();
    }, [])

    useEffect(() => {
        const scheduledUpdate = setTimeout(() => {
            updateLinePlotData();
        }, 2000);
        return () => clearTimeout(scheduledUpdate);
    }, [linePlotData])


    const displayLineChart = linePlotData.map(kpi => {
        return (
            <div key={kpi.id} className={style.plot}>
                <LinePlot
                    heightScreenRatio={0.35}
                    title={kpi.title}
                    data={kpi.data.map(item => item.value)}
                    labels={kpi.data.map(item => item.label)} />
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
                {/* <div className={style.ratioStats}>
                    {displayRatioPlot}
                </div> */}
            </div>
        </div>
    )
}

export default Home;
