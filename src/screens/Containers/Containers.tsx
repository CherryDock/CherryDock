import React, { useState, useEffect } from 'react';
import ActionButtons from '../../components/Containers/ActionButtons/ActionButtons';
import style from './containers.module.scss';
import ContainerCard from '../../components/Containers/ContainerCard/ContainerCard';
import SelectedContainers from '../../components/Containers/SelectedContainers/SelectedContainers';
import { fetchContainersInfo } from '../../utils/api-fetch';
import { ContainerInfo } from '../../interfaces/data.interface';

function Containers() {

    const [containers, setContainers] = useState<ContainerInfo[]>([])

    useEffect(() => {
        fetchContainersInfo()
            .then(res => {
                const containersInfoWithId = res.map(ctn => {
                    return {
                        ...ctn, selected: false
                    }
                })
                setContainers(containersInfoWithId);
            });

    }, [])

    function toggleContainer(id: string) {
        const updatedContainers = containers.map(container => {
            if (container.Id === id) {
                return {
                    ...container,
                    selected: !container.Selected
                }
            }
            else
                return container;
        });
        setContainers(updatedContainers);
    }

    const displayContainers = containers.map(container => {
        return (
            <div key={container.Id} className={style.containerCard}>
                <ContainerCard
                    key={container.Id}
                    id={container.Id}
                    containerName={container.Name}
                    containerState={container.Status}
                    imageName={container.Image}
                    launchedDate={container.Created}
                    selected={container.Selected!}
                    toggleContainerFunc={toggleContainer}
                />
            </div>
        )
    })

    function resetContainerSelection() {
        const containersWithSelectionReset = containers.map(container => {
            if (container.Selected === true) {
                return {
                    ...container,
                    selected: false
                }
            }
            else return container;
        });
        setContainers(containersWithSelectionReset);
    }

    function displaySelectedContainers() {
        const numSelectedContainers = containers.map(container => container.Selected)
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
