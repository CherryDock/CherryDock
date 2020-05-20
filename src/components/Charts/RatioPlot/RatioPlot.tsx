import React from 'react';
import { Doughnut } from 'react-chartjs-2';
import 'chartjs-plugin-datalabels';

interface RatioPlotData {
    id: number;
    label: string;
    value: number;
}

interface RatioPlotProps {
    data: RatioPlotData[];
    heightScreenRatio: number;
    title: string;
}

const colorsPaletteHover = ['#1abc9c', '#2ecc71', '#3498db', '#9b59b6', '#f1c40f', '#e67e22', '#e74c3c', '#34495e'];
const colorsPalette = ['#16a085', '#27ae60', '#2980b9', '#8e44ad', '#f39c12', '#d35400', '#c0392b', '#2c3e50']

function LinePlot({ heightScreenRatio, data, title }: RatioPlotProps) {

    const screenHeight = window.innerHeight;
    const PlotHeight = screenHeight * heightScreenRatio;

    const numOfSlices = data.length;

    const colors = colorsPalette.slice(0, numOfSlices);
    const colorsHover = colorsPaletteHover.slice(0, numOfSlices);
    const values = data.map(item => item.value);
    const labels = data.map(item => item.label);

    const dataConf = {
        labels: labels,
        datasets: [{
            data: values,
            backgroundColor: colors,
            hoverBackgroundColor: colorsHover
        }]
    }

    const optionConf = {
        legend: {
            display: true,
            position: 'bottom'
        },
        title: {
            display: true,
            text: title,
            fontSize: 18
        },
        plugins: {
            datalabels: {
                display: true,
                labels: {
                    title: {
                        font: {
                            weight: 'bold',
                            size: '18'
                        },
                        color: 'white'
                    }
                }
            }
        }
    }


    return (
        <div>
            <Doughnut
                data={dataConf}
                height={PlotHeight}
                options={optionConf}
            />
        </div>
    )
}

export default LinePlot
