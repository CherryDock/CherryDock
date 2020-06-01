import React, { useState, useEffect } from 'react';
import style from './line-plots-manager.module.scss';
import LinePlot from '../../Charts/LinePlot/LinePlot';
import { linePlotConf } from '../../../conf/charts.conf';
import { LinePlotData, ChartItem } from '../../../interfaces/charts.interface';

interface LinePlotsManagerProps<T> {
    realTimeLimit: number;
    fetchData: () => Promise<T>;
    makePlotItem: (data: T, kpiName: string) => ChartItem | undefined;
}

function LinePlotsManager<T extends object>({ realTimeLimit, fetchData, makePlotItem }: LinePlotsManagerProps<T>) {

    const [linePlotData, setLinePlotData] = useState<LinePlotData[]>([]);
    /**
     * Init linePlotData
     */
    function initLinePlotData() {
        fetchData()
            .then(data => {
                if (data != undefined) {
                    let id = 1;
                    linePlotConf.forEach(plotConf => {

                        const dataItem = [makePlotItem(data, plotConf.kpiName)!];
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
        //const rtStatsRoute = apiConf.routes.find(route => route.name === 'GET-ALL-RT-CTN-STATS') ?.route!;
        //fetchDataApi<GlobalContainers>(rtStatsRoute)
        fetchData()
            .then(data => {
                if (data != undefined) {
                    let updatedData = linePlotData.map(plotData => {
                        const plotTitle = plotData.title;
                        const kpiName = linePlotConf.find(plotConf => plotConf.title === plotTitle)!.kpiName;
                        const dataItem = makePlotItem(data, kpiName);

                        if (plotData.data.length > realTimeLimit) {
                            plotData.data = shiftArrayLeft<ChartItem>(plotData.data);
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
                    heightScreenRatio={0.40}
                    title={kpi.title}
                    data={kpi.data.map(item => item.value)}
                    labels={kpi.data.map(item => item.label)} />
            </div>
        )
    });

    return (
        <div className={style.plotManager}>
            {displayLineChart}
        </div>
    )
}

export default LinePlotsManager
