import React, { useState, useEffect } from 'react';
import style from './ratio-charts-manager.module.scss';
import { RatioPlotData, ChartItem } from '../../../interfaces/charts.interface';
import { ratioPlotConf } from '../../../conf/charts.conf';
import { GlobalContainers } from '../../../interfaces/data.interface';
import RatioPlot from '../../Charts/RatioPlot/RatioPlot';

interface RatioPlotsManagerProps {
    realTimeLimit: number;
    fetchData: () => Promise<GlobalContainers>;
    makePlotItem: (data: GlobalContainers, cntId: string, kpiName: string, addId: boolean) => ChartItem | undefined;
}

function RatioChartsManager({ realTimeLimit, fetchData, makePlotItem }: RatioPlotsManagerProps) {

    const [ratioPlotData, setratioPlotData] = useState<RatioPlotData[]>([]);

    /**
     * Init linePlotData
     */
    function initRatioPlotData() {
        fetchData()
            .then(data => {
                if (data != undefined) {
                    console.log("Data: ", data);
                    let id = 1;
                    ratioPlotConf.forEach(plotConf => {
                        const kpiData = data.Containers.map(cnt => {
                            const dataItem = makePlotItem(data, cnt.Id, plotConf.kpiName, true)!;
                            return dataItem;
                        });

                        const plotData: RatioPlotData = {
                            id,
                            title: plotConf.title,
                            data: kpiData
                        }
                        id++;
                        setratioPlotData(emptyArray => [...emptyArray, plotData]);
                    });
                }
            });
    }

    useEffect(() => {
        initRatioPlotData();
    }, [])

    const displayRatioChart = ratioPlotData.map(kpi => {
        return (
            <div key={kpi.id} className={style.plot}>
                <RatioPlot
                    heightScreenRatio={0.45}
                    title={kpi.title}
                    data={kpi}
                />
            </div>
        )
    });
    return (
        <div className={style.plotManager}>
            {displayRatioChart}
        </div>
    )
}

export default RatioChartsManager
