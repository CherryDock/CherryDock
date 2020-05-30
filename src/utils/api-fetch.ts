import axios from "axios";
import apiConf from "../conf/api.conf";
import {
  GlobalContainers,
  HistGlobalContainers,
  Container,
} from "../interfaces/data.interface";
import { SingleOrAllCnts } from "../types/data.types";

const globalContainersData = require("../conf/specs-json/rt-global-ctn-stats.json");
const containerData = require("../conf/specs-json/rt-single-ctn-stats.json");
const historicData = require("../conf/specs-json/historic-global-ctn-stats.json");
const containersInfo = require("../conf/specs-json/containers-info.json");

/**
 * Fetch data from the API
 * @param apiRoute - API route
 */
async function fetchDataApi<T>(apiRoute: string): Promise<T | null> {
  const apiUrl = apiRoute;

  const apiConfig = {
    headers: {
      Authorization: `Bearer ${apiConf.apiToken}`,
    },
  };

  try {
    const response = await axios.get(apiUrl, apiConfig);
    const fetchedData: T = response.data;
    return fetchedData;
  } catch (error) {
    console.log(error);
    return null;
  }
}


/**
 * Dummy data simulating fetching the API
 * @param jsonName - Name of the json file
 */
async function fetchGlobalContainers(): Promise<GlobalContainers>{
  return globalContainersData;
}

/**
 * Dummy data simulating fetching the API
 * @param jsonName - Name of the json file
 */
async function fetchSingleContainer(): Promise<Container> {
  return containerData;
}


export { fetchDataApi, fetchGlobalContainers, fetchSingleContainer };
