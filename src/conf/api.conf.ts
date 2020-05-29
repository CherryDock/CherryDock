export default {
  apiServer: "http://localhost",
  port: "3000",
  apiToken:
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjY2ODkxNDUsInVzZXIiOiIifQ._nUdbYxWsBI-clzFr16M30D4zyvUu7-SRsAfVisZvLg",
  routes: [
      {
          name: 'GET-TOKEN',
          route: '/token'
      },
    {
      name: "GET-ALL-RT-CTN-STATS",
      route: "/api/monitor/stats",
    },
  ],
};
