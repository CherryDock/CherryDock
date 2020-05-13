import React, { useState, useEffect } from 'react';
import ActionButtons from '../../components/Containers/ActionButtons/ActionButtons';
import style from './containers.module.scss';
import ContainerCard from '../../components/Containers/ContainerCard/ContainerCard';
import containerData from './data';
import SelectedContainers from '../../components/Containers/SelectedContainers/SelectedContainers';

function Containers() {

    interface Container {
        id: number;
        containerName: string;
        containerState: string;
        imageName: string;
        launchedDate: string;
        selected: boolean;
    }

    const [containers, setContainers] = useState<Container[]>([])

    async function getContainersData() {
        const data: Container[] = containerData;
        return data;
    }

    useEffect(() => {
        getContainersData()
            .then(res => {
                setContainers(res);
            });

    }, [])

    function toggleContainer(id: number) {
        const updatedContainers = containers.map(container => {
            if (container.id === id) {
                return {
                    ...container,
                    selected: !container.selected
                }
            }
            else
                return container;
        });
        setContainers(updatedContainers);
    }

    const displayContainers = containers.map(container => {
        return (
            <div className={style.containerCard}>
                <ContainerCard
                    key={container.id}
                    id={container.id}
                    containerName={container.containerName}
                    containerState={container.containerState}
                    imageName={container.imageName}
                    launchedDate={container.launchedDate}
                    selected={container.selected}
                    toggleContainerFunc={toggleContainer}
                />
            </div>
        )
    })

    function resetContainerSelection() {
        const containersWithSelectionReset= containers.map(container => {
            if (container.selected === true) {
                return {
                    ...container,
                    selected: false
                }
            }
            else return container;
        });
        console.log(containersWithSelectionReset);
        setContainers(containersWithSelectionReset);
    }

    function displaySelectedContainers() {
        const numSelectedContainers = containers.map(container => container.selected)
            .filter(isSelected => isSelected === true)
            .length
        if (numSelectedContainers > 0) {
            return (<SelectedContainers
                numContainers={numSelectedContainers} 
                deselectedContainer={resetContainerSelection}
                />)
        }

        else return (<div></div>)
    }

    return (
        <div className={style.container}>
            <h3 className={style.screenName}>Containers</h3>
            <ActionButtons />
            <div className={style.selectedContainers}>
                {displaySelectedContainers()}
            </div>
            <div className={style.containers}>
                {displayContainers}
            </div>
        </div>
    )
}

export default Containers
