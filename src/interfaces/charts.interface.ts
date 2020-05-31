export interface LinePlotProps {
  heightScreenRatio: number;
  data: number[];
  title: string;
  labels: string[];
}

export interface ChartItem {
  id?: number;
  value: number;
  label: string;
}

export interface LinePlotData {
  id: number;
  data: ChartItem[];
  title: string;
}

export interface RatioPlotData {
  id: number;
  title: string;
  data: ChartItem[];
}
