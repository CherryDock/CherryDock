import { GlobalContainers, Container } from "../interfaces/data.interface";
import { LinePlotItem } from "../interfaces/charts.interface";

/**
 * Get acquisition time with the following format
 * HH:MM
 */
function makeAcqTime(): string {
  const now = new Date();
  let hour = now.getHours().toString();
  if (hour.length === 1) hour = "0" + hour;
  let minute = now.getMinutes().toString();
  if (minute.length === 1) minute = "0" + minute;
  return hour + ":" + minute;
}

/**
 * Get real-time data of a specific KPI from the Global Containers Data
 * @param fetchedData
 */
function makeGlobalCtnItem(
  fetchedData: GlobalContainers,
  kpiName: string
): LinePlotItem | undefined {
  const kpi = fetchedData.GlobalInfo.find(kpi => kpi.name === kpiName);
  if (kpi !== undefined) {
    //const value = Number(kpi.value.toFixed(2));
    let value = 50 + Math.random() * 3;
    value = Number(value.toFixed(2));
    const label = makeAcqTime();
    return { value, label };
  } else return undefined;
}

/**
 * Get real-time data of a specific KPI from the single Container Data
 * @param fetchedData
 */
function makeSingleCtnItem(
  fetchedData: Container,
  kpiName: string
): LinePlotItem | undefined {
  const kpi = fetchedData.Info.find(kpi => kpi.name === kpiName);
  if (kpi !== undefined) {
    let value = 50 + Math.random() * 3;
    value = Number(value.toFixed(2));
    const label = makeAcqTime();
    return { value, label };
  } else return undefined;
}

export { makeGlobalCtnItem, makeSingleCtnItem };

