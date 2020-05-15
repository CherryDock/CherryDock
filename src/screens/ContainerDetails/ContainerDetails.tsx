import React, { useState, useEffect } from 'react';
import style from './container-details.module.scss';

interface ContainerDetailsProps {
    match: any;
}

function ContainerDetails() {
    const [containerId, setcontainerId] = useState<number>();

    useEffect(() => {
        //const containerIdFromParams = match
    }, []);

    return (
        <div>
            <h3>Container Details</h3>
        </div>
    )
}

export default ContainerDetails
