import React from 'react';
import { Doughnut } from 'react-chartjs-2';
import 'chartjs-plugin-datalabels';
import { RatioPlotData } from '../../../interfaces/charts.interface';

interface RatioPlotProps {
    data: RatioPlotData;
    heightScreenRatio: number;
    title: string;
}

const colorsPaletteHover = ['#1abc9c', '#2ecc71', '#3498db', '#9b59b6', '#f1c40f', '#e67e22', '#e74c3c', '#34495e'];
const colorsPalette = ['#16a085', '#27ae60', '#2980b9', '#8e44ad', '#f39c12', '#d35400', '#c0392b', '#2c3e50']

function RatioPlot({ heightScreenRatio, data, title }: RatioPlotProps) {
    const plotData = data.data;

    const screenHeight = window.innerHeight;
    const PlotHeight = screenHeight * heightScreenRatio;
    const numOfSlices = plotData.length;

    const colors: string[] = [];
    const colorsHover: string[] = [];
    const values = plotData.map(item => item.value);
    const labels = plotData.map(item => item.label);

    for (let index = 0; index < numOfSlices; index++) {
        const colorIndex = index % colorsPalette.length;
        colors.push(colorsPalette[colorIndex]);
        colorsHover.push(colorsPaletteHover[colorIndex])
    }

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

export default RatioPlot;
