interface Port {
    IP: string;
    PrivatePort: number;
    PublicPort: number;
    Type: string
}

export interface Container {
    Name: string[];
    Status: string;
    Image: string;
    Command: string;
    Created: number;
    Ports: Port[]
}

export interface Containers extends Array<Container>{}

interface StatsInfoValue {
    Value: number;
    Unit: string;
}

interface NetworkInfo {
    In: StatsInfoValue;
    Out: StatsInfoValue;
}

interface StatsInfo {
    CpuUsagePercent: number;
    MemoryUsagePercent: number;
    Memory: StatsInfoValue;
    NetworkInfo: NetworkInfo;
}

export interface RTContainerStats {
    Id: string;
    Info: StatsInfo;
}

export interface RTGlobalContainersStats {
    RunningContainers: number;
    NbCpu: number;
    MemoryLimit: StatsInfoValue;
    MemoryUsagePercent: number;
    CpuUsagePercent: number;
    Containers: RTContainerStats[];
}

export interface HistoricGlobalContainersStats {
    Date: string;
    Stats: RTGlobalContainersStats;
}

export interface HistoricGlobalContainersStatsList extends Array<HistoricGlobalContainersStats>{}