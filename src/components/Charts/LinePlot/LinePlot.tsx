import React from 'react';
import { Line } from 'react-chartjs-2';

interface LinePlotProps {
    heightScreenRatio: number;
    data: number[];
    title: string;
}

function LinePlot({ heightScreenRatio, data, title }: LinePlotProps) {

    const screenHeight = window.innerHeight;
    const PlotHeight = screenHeight * heightScreenRatio;

    const optionConf = {
        legend: {
            display: false
        },
        title: {
            display: true,
            text: title,
            fontSize: 18
        }
    }

    const dataConf = {
        labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
        datasets: [
            {
                label: 'My First dataset',
                fill: false,
                lineTension: 0.1,
                backgroundColor: 'rgba(75,192,192,0.4)',
                borderColor: 'rgba(75,192,192,1)',
                borderCapStyle: 'butt',
                borderDash: [],
                borderDashOffset: 0.0,
                borderJoinStyle: 'miter',
                pointBorderColor: 'rgba(75,192,192,1)',
                pointBackgroundColor: '#fff',
                pointBorderWidth: 1,
                pointHoverRadius: 5,
                pointHoverBackgroundColor: 'rgba(75,192,192,1)',
                pointHoverBorderColor: 'rgba(220,220,220,1)',
                pointHoverBorderWidth: 2,
                pointRadius: 1,
                pointHitRadius: 10,
                data: data
            }
        ]
    };

    return (
        <div>
            <Line
                data={dataConf}
                height={PlotHeight}
                options={optionConf}
            />
        </div>
    )
}

export default LinePlot