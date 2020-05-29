export interface Info {
  name: string;
  value: number;
  unit: string;
}

export interface Container {
  Id: string;
  Info: Info[];
}

export interface GlobalContainers {
  GlobalInfo: Info[];
  Containers: Container[];
}

export interface HistGlobalContainers {
  Date: Date;
  GlobalInfo: Info[];
  Containers: Container[];
}

export interface Port {
  IP: string;
  PrivatePort: number;
  PublicPort: number;
  Type: string;
}

export interface ContainerInfo {
  Name: string[];
  Status: string;
  Id: string;
  Image: string;
  Command: string;
  Created: number;
  Ports: Port[];
}
