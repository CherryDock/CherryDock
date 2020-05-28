import axios from "axios";

const apiServer = "http://0.0.0.0";
const port = "8001";
const apiToken =
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjY2ODkxNDUsInVzZXIiOiIifQ._nUdbYxWsBI-clzFr16M30D4zyvUu7-SRsAfVisZvLg";

/**
 * Fetch data from the API
 * @param apiRoute - API route
 */
async function fetchDataApi<T>(apiRoute: string): Promise<T | null> {
  const apiUrl = apiServer + ":" + port + apiRoute;
  const apiConfig = {
    headers: { Authorization: `Bearer ${apiToken}` },
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

export { fetchDataApi };
