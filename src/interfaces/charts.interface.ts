export interface LinePlotProps {
  heightScreenRatio: number;
  data: number[];
  title: string;
  labels: string[];
}

export interface LinePlotItem {
  value: number;
  label: string;
}

export interface LinePlotData {
  id: number;
  data: LinePlotItem[];
  title: string;
}
