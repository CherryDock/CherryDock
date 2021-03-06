
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
![Go](https://github.com/CherryDock/CherryDock/workflows/Go/badge.svg)
[![Maintainability](https://api.codeclimate.com/v1/badges/05f40aad1870fa973aa9/maintainability)](https://codeclimate.com/github/CherryDock/CherryDock/maintainability)
[![codecov](https://codecov.io/gh/CherryDock/CherryDock/branch/dev/graph/badge.svg)](https://codecov.io/gh/CherryDock/CherryDock)

<p align="center"><img src="https://avatars2.githubusercontent.com/u/64746114?s=200&v=4" /></p>


<h1>Cherry Dock</h1>
Cherry Dock is an open source docker monitoring and management tool written in Golang & React. It allows user to perform basic 
operations on docker containers and analyze various metrics about containers ressources. For the time being,
it only supports standalone docker engine. 
<br><br/>
<img src="app-preview.png" width="90%" />

<h2>Deploy CherryDock</h2>
To deploy cherrydock on a linux host, run le following :

```
docker build -t cherrydock .
```

```
docker run --name cherrydock \
    -v /var/run/docker.sock:/run/docker.sock \
    -p 5000:5000 \
    -i -d cherrydock:latest
```
NB: Api & UI are running in the same container 



